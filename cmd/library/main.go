package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/jt/books/configuration"
	"github.com/jt/books/database"
	library "github.com/jt/books/services"
	"github.com/sirupsen/logrus"
	"log"
	"net"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"

	books "github.com/jt/books/gen/books"
)

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF     = flag.String("host", "localhost", "Server host (valid values: localhost)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
		logLevel  = flag.String("log-level", "warn", "log level of the application possible options(debug,info,warn,error)")
	)
	flag.Parse()

	// Setup logger. Replace logger with your own log package of choice.
	setupLogging(*logLevel)
	logger := logrus.New()

	// Initialize the services.
	var (
		booksSvc books.Service
	)
	{
		bookDatabase, err := database.NewBooksDB(configuration.ApplicationConfiguration.DbConnectionString)
		if err != nil {
			logrus.WithError(err).Fatal("couldn't setup NewBooksDB")
		}
		booksSvc = library.NewBooks(logger, bookDatabase)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		booksEndpoints *books.Endpoints
	)
	{
		booksEndpoints = books.NewEndpoints(booksSvc)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case "localhost":
		{
			addr := "http://localhost:8000"
			u, err := url.Parse(addr)
			if err != nil {
				logger.Fatalf("invalid URL %#v: %s\n", addr, err)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h, _, err := net.SplitHostPort(u.Host)
				if err != nil {
					logger.Fatalf("invalid URL %#v: %s\n", u.Host, err)
				}
				u.Host = net.JoinHostPort(h, *httpPortF)
			} else if u.Port() == "" {
				u.Host = net.JoinHostPort(u.Host, "80")
			}

			//note that I did this here for simplicity on real scenario we need a proper solution
			httpLogger := log.New(os.Stderr, "[library] ", log.Ltime)

			handleHTTPServer(ctx, u, booksEndpoints, &wg, errc, httpLogger, *dbgF)
		}

	default:
		logger.Fatalf("invalid host argument: %q (valid hosts: localhost)\n", *hostF)
	}

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}

func setupLogging(logLevel string) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	//debug,info,warn,error
	switch logLevel {
	case "debug":
		{
			logrus.SetLevel(logrus.DebugLevel)
		}
	case "info":
		{
			logrus.SetLevel(logrus.InfoLevel)
		}
	case "error":
		{
			logrus.SetLevel(logrus.ErrorLevel)
		}
	default:
		logrus.SetLevel(logrus.WarnLevel)
	}

}

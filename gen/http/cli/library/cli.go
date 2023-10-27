// Code generated by goa v3.13.2, DO NOT EDIT.
//
// library HTTP client CLI support package
//
// Command:
// $ goa gen github.com/jt/books/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	booksc "github.com/jt/books/gen/http/books/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `books (list|show|create|update|delete)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` books list` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, any, error) {
	var (
		booksFlags = flag.NewFlagSet("books", flag.ContinueOnError)

		booksListFlags = flag.NewFlagSet("list", flag.ExitOnError)

		booksShowFlags  = flag.NewFlagSet("show", flag.ExitOnError)
		booksShowIDFlag = booksShowFlags.String("id", "REQUIRED", "ID of the book")

		booksCreateFlags    = flag.NewFlagSet("create", flag.ExitOnError)
		booksCreateBodyFlag = booksCreateFlags.String("body", "REQUIRED", "")

		booksUpdateFlags    = flag.NewFlagSet("update", flag.ExitOnError)
		booksUpdateBodyFlag = booksUpdateFlags.String("body", "REQUIRED", "")
		booksUpdateIDFlag   = booksUpdateFlags.String("id", "REQUIRED", "ID of the book")

		booksDeleteFlags  = flag.NewFlagSet("delete", flag.ExitOnError)
		booksDeleteIDFlag = booksDeleteFlags.String("id", "REQUIRED", "ID of the book")
	)
	booksFlags.Usage = booksUsage
	booksListFlags.Usage = booksListUsage
	booksShowFlags.Usage = booksShowUsage
	booksCreateFlags.Usage = booksCreateUsage
	booksUpdateFlags.Usage = booksUpdateUsage
	booksDeleteFlags.Usage = booksDeleteUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "books":
			svcf = booksFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "books":
			switch epn {
			case "list":
				epf = booksListFlags

			case "show":
				epf = booksShowFlags

			case "create":
				epf = booksCreateFlags

			case "update":
				epf = booksUpdateFlags

			case "delete":
				epf = booksDeleteFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "books":
			c := booksc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list":
				endpoint = c.List()
				data = nil
			case "show":
				endpoint = c.Show()
				data, err = booksc.BuildShowPayload(*booksShowIDFlag)
			case "create":
				endpoint = c.Create()
				data, err = booksc.BuildCreatePayload(*booksCreateBodyFlag)
			case "update":
				endpoint = c.Update()
				data, err = booksc.BuildUpdatePayload(*booksUpdateBodyFlag, *booksUpdateIDFlag)
			case "delete":
				endpoint = c.Delete()
				data, err = booksc.BuildDeletePayload(*booksDeleteIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// booksUsage displays the usage of the books command and its subcommands.
func booksUsage() {
	fmt.Fprintf(os.Stderr, `The book service performs CRUD operations on the books resource.
Usage:
    %[1]s [globalflags] books COMMAND [flags]

COMMAND:
    list: Retrieve all books
    show: Retrieve a book by ID
    create: Create a new book
    update: Update an existing book
    delete: Delete a book by ID

Additional help:
    %[1]s books COMMAND --help
`, os.Args[0])
}
func booksListUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] books list

Retrieve all books

Example:
    %[1]s books list
`, os.Args[0])
}

func booksShowUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] books show -id INT

Retrieve a book by ID
    -id INT: ID of the book

Example:
    %[1]s books show --id 3056930450430130409
`, os.Args[0])
}

func booksCreateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] books create -body JSON

Create a new book
    -body JSON: 

Example:
    %[1]s books create --body '{
      "author": "Minus adipisci consequuntur.",
      "cover": "QmFzZTY0IG9mIHRoZSBCb29rIGNvdmVyIGltYWdl",
      "published_at": "2012-06-24T11:13:57Z",
      "title": "Ut in qui quidem ab est."
   }'
`, os.Args[0])
}

func booksUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] books update -body JSON -id INT

Update an existing book
    -body JSON: 
    -id INT: ID of the book

Example:
    %[1]s books update --body '{
      "author": "Quo vel ipsa et minima.",
      "cover": "QmFzZTY0IG9mIHRoZSBCb29rIGNvdmVyIGltYWdl",
      "published_at": "1990-02-10T16:36:44Z",
      "title": "Maiores facilis ducimus quia harum."
   }' --id 5569448609837740281
`, os.Args[0])
}

func booksDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] books delete -id INT

Delete a book by ID
    -id INT: ID of the book

Example:
    %[1]s books delete --id 3173693188622662838
`, os.Args[0])
}

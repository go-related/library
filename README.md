# Library

## Prerequisites
1. Golang installed version 1.20
2. My Sql installed (latest if possible)
3. Create a empty db (for example library)


## How to build and Run application
1. Update the connection string on the file /config/default.yml 
2. Build the code
    ```Golang 
        #cd to the main folder of the project
        #depending on the permissions of the machine you might need to create the bin folder or dumb the file in another place
        go build -o bin/ ./cmd/library  
   ```
3. Run the application with the appropriate log level(default warn)
    ```code
           ./bin/library --log-level info
    ```
   if everything was ok you should see something like 
   ```code
   {"level":"info","service":"library","time":"2023-10-27T11:40:23+02:00","message":"HTTP \"List\" mounted on GET /books"}
   {"level":"info","service":"library","time":"2023-10-27T11:40:23+02:00","message":"HTTP \"Show\" mounted on GET /books/{id}"}
   {"level":"info","service":"library","time":"2023-10-27T11:40:23+02:00","message":"HTTP \"Create\" mounted on POST /books"}
   {"level":"info","service":"library","time":"2023-10-27T11:40:23+02:00","message":"HTTP \"Update\" mounted on PUT /books/{id}"}
   {"level":"info","service":"library","time":"2023-10-27T11:40:23+02:00","message":"HTTP \"Delete\" mounted on DELETE /books/{id}"}
   {"level":"info","service":"library","time":"2023-10-27T11:40:23+02:00","message":"HTTP \"./gen/http/openapi.json\" mounted on GET /openapi.json"}
   {"level":"info","service":"library","time":"2023-10-27T11:40:23+02:00","message":"HTTP server listening on \"localhost:8000\""}
   ^C{"level":"info","service":"library","time":"2023-10-27T11:40:29+02:00","message":"exiting (interrupt)"}
   {"level":"info","service":"library","time":"2023-10-27T11:40:29+02:00","message":"shutting down HTTP server at \"localhost:8000\""}
   
   ```
4. To run all the test on the application 
   ```code
   #cd main folder of the project
   go test -v ./...
   ```
5. Api documentation 
      access it here[https://petstore.swagger.io/?url=https://github.com/go-related/library/blob/main/gen/http/openapi3.yaml] 
## Notes 
1. Application will run migration auto migration everytime is starts/restart 
2. On the config folder you will find a postman collection with sample calls to the application(import that you should be good to go, i didn't include variables to make it simpler to import)
3. On the config  folder(/config) there is also the db schema

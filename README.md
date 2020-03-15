# Snippetbox

A sample web application written in Go.

### Prerequisites 
The application is relying on a simple MySQL database. 
To spin up a sample image run below command

```docker run --name snippetbox-db -p 3306:3306 -e MYSQL_ROOT_PASSWORD=my-secret-pw -d adilisio/snippetbox-db```

The above image is creating a few default records and default **web** which will be used by the application. 

### Running the application

```go run cmd/web/*```

### To list available command-line flags

```go run cmd/web/* -help```

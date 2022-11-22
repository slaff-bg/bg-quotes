# BG Thoughts And Quotes. #


The exercise recreates an API project for thoughts and quotes from Bulgarian folklore, some of which are from prominent Bulgarians.

That is a test project for learning purposes and one of my first Golang codes.
So I cannot vouch for the correctness of the code and approaches and its quality in a production environment.


## Features ##

- [x] Home page. Displays the project name.
- [ ] System environment health check.
- [ ] List of authors.
- [ ] List of quotes.
- [ ] Show authors by their ID.
- [ ] Show quotes by their ID.
- [ ] etc.


## Functionalities ##

- [ ] Authors CRUD.
- [ ] List of authors. Including paging, filtering, and ordering.
- [x] Show authors by their ID.
- [ ] Quotes CRUD.
- [ ] List of quotes. Including paging, filtering, and ordering.
- [ ] Show quotes by their ID.
- [ ] Filtering/ordering/sorting.
- [ ] User authentication.
- [ ] User authorization.
- [ ] API versioning.
- [ ] Defining OpenAPI specification.
- [ ] etc.



## Dependencies ##

- [Golang](https://go.dev/dl/) version go1.19.
- [Docker](https://www.docker.com/) version 20.10.18.

&#x1F4CC; &nbsp; *<sub>Versions reflect the current state of the used technologies.</sub>*


## How do I get set up? ##

* Clone the package locally.
* Go to the directory of your local copy.

* The go.mod and go.sum files are excluded. Look at the very bottom of the gitignore file. So you have to set up the service locally.

  ```sh
    go mod init bg-quotes
    go mod tidy
  ```

* Start the system.

  ```sh
    go run main.go
  ```


## How to run tests? ##

* Open a CLI.
* Go to the directory of your local copy.
* Clean the cache using the following command:
  
  ```sh
    go clean -testcache
  ```

### Run all tests. ###

```sh
go test ./...
```

### Run the test in a concrete directory. ###

- Run all tests in the main directory only

  ```sh
  go test .
  ```

- Run all tests in api directory only

  ```sh
  go test ./api
  ```

- Run all tests in domain directory only

  ```sh
  go test ./domain
  ```

### Run a separate test. ###

- Run separate test located in the root directory

  ```sh
  go test . -run=TestMainHandler_StatusOK_BodyContent
  ```

- Run separate test located in domain directory

  ```sh
  go test ./domain -run=TestCreateAuthor

  go test ./domain -run=TestCreateQuote

  go test ./domain -run=TestURankIota
  ```


### API Test Links ###

>__NOTE__
> By using CURL

* Create Author

```sh
curl -X POST -H "Content-Type: application/json" \
 -d '{"first_name":"Henry","second_name":"Bukowski","aka":"Charles Bukowski","img_url":"https://upload.wikimedia.org/wikipedia/en/e/e2/Charles_Bukowski_smoking.jpg"}' \
 http://0.0.0.0:8080/authors
```

* Show Author by UUID

```sh
curl -v http://0.0.0.0:8080/authors/<45dcb7b2-904f-4fa4-a9eb-53dc1fba04ca>
```

* Create Quote

```sh
curl -X POST -H "Content-Type: application/json" \
 -d '{"quote":"If I feel like working, I sit down and wait for it to pass.","smoking_room":false,"author_id":""}' \
 http://0.0.0.0:8080/quotes
```




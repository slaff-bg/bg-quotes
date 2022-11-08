# BG Thoughts And Quotes. #


The exercise recreates an API project for thoughts and quotes from Bulgarian folklore, some of which are from prominent Bulgarians.

That is a test project for learning purposes and one of my first Golang codes.
So I cannot vouch for the correctness of the code and approaches and its quality in a production environment.


## Features ##

- [x] Home page. Displays the project name.
- [ ] System environment health check.
- [ ] Authentication.
- [ ] Quotes CRUD.
- [ ] List of quotes. Including paging, filtering, and ordering.
- [ ] Show a quote by its ID.
- [ ] ...


## Functionalities ##

For the tests:

- [ ] Quotes CRUD functionality.
- [ ] Passed argument validator.
- [ ] Filtering and ordering/sorting.
- [ ] ...


## Notes ##

### Dependencies ###

- [Golang](https://go.dev/dl/) version go1.19.
- [Docker](https://www.docker.com/) version 20.10.18.

&#x1F4CC; &nbsp; *<sub>Versions reflect the current state of the used technologies.</sub>*


### How do I get set up? ###

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

...


### How to use DB container via terminal? ###

...


### How to run tests? ###

...


### API Test Links ###

>__NOTE__
> By using CURL

* Create Author

```sh
curl -X POST -H "Content-Type: application/json" \
 -d '{"afname":"Henry","asname":"Bukowski","aaka":"Charles Bukowski","aimgurl":"https://upload.wikimedia.org/wikipedia/en/e/e2/Charles_Bukowski_smoking.jpg"}' \
 http://0.0.0.0:8080/author/create
```

```sh
curl -X POST -H "Content-Type: application/json" \
 -d '{"afname":"Charles","asname":"Chaplin","aaka":"Charlie Chaplin","aimgurl":"https://upload.wikimedia.org/wikipedia/commons/thumb/3/34/Charlie_Chaplin_portrait.jpg/330px-Charlie_Chaplin_portrait.jpg"}' \
 http://0.0.0.0:8080/author/create
```

* Show Author by UUID

```sh
curl -v http://0.0.0.0:8080/author/<45dcb7b2-904f-4fa4-a9eb-53dc1fba04ca>
```


### TODO ###

> __WARNING__
> All these points must include tests before being marked as complete.

- [x] Initialize GIT and add gitignore and README
- [x] Implement Gin Web Framework
- [ ] Do domains design
  - [ ] Start with the Authors first
  - [ ] Do the same for the Quotes
- [ ] Do middleware validator for the dynamic arguments
- [ ] Implement Author CRUD functionality.
- [ ] Implement Quotes CRUD functionality.
- [ ] Implemented authentication by using tokens
- [ ] Add functionality for API versions.
- [ ] Add filtering / ordering / sorting.
- [ ] 


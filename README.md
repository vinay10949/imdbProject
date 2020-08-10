IMDB PROJECT
-------

It has simple dependencies:

 - [Chi (Router)](https://github.com/go-chi/chi)
 - [Hystrix-Go (Circuit Breaker)](https://github.com/afex/hystrix-go)

Get Started:


-------

Clone the source

    https://github.com/vinay10949/imdbProject.git

Setup dependencies

    go get -u github.com/go-chi/chi
    go get github.com/afex/hystrix-go/hystrix

Run the app

    go build && ./imdbProject

And visit

    http://localhost:8080/movie/Matrix

----------

[Folder Structure](https://irahardianto.github.io/service-pattern-go/#folder-structure)
-------
    /
    |- controllers
    |- infrastructures
    |- interfaces
    |- models
    |- repositories
    |- services
    |- viewmodels
    main.go
    router.go
    servicecontainer.go

The folder structure is created to accomodate seperation of concern principle, where every struct should have single responsibility to achieve decoupled system.

Every folder is a namespace of their own, and every file / struct under the same folder should only use the same namepace as their root folder.

### controllers

controllers folder hosts all the structs that  are the handler of all requests coming in, to the router, its doing just that, business logic and data access layer should be done separately.


### infrasctructures

infrasctructures folder host all structs for  system to connect to external data source like  MySQL etc 

### interfaces

interfaces folder hosts all the structs under interfaces namespace, interfaces as the name suggest are the bridge between different domain so they can interact with each other, in our case, this should be the only way for them to interact.

### models

models folder hosts all structs under models namespace, model is a struct reflecting our data object from / to database. models should only define data structs

### repositories

repositories folder hosts all structs under repositories namespace, repositories is where the implementation of data access layer. All queries and data operation from / to database should happen here, and the implementor should be agnostic of what is the database engine is used, how the queries is done, all they care is they can pull the data according to the interface they are implementing.

### services

services folder hosts all structs under services namespace, services is where the business logic lies on, it handles controller request and fetch data from data layer it needs and run their logic to satisfy what controller expect the service to return.

### viewmodels

viewmodels folder hosts all the structs under viewmodels namespace, viewmodels are model to be use as a response return of REST API call

### main.go

main.go is the entry point of our system, here lies the router bindings it triggers ChiRouter singleton and call InitRouter to bind the router.

### router.go

router.go is where we binds controllers to appropriate route to handle desired http request

### servicecontainer.go

This is the place where we injected all implementations of interfaces.

### data.sql

this is mysql exported database (imdb)

### postmancollection.json

Postman collection for rest api endpoints

----------


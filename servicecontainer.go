package main

import (
	"sync"

	"imdbProject/controllers"
	"imdbProject/repositories"
	"imdbProject/infrastructures"
	"imdbProject/services"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type IServiceContainer interface {
	InjectMovieController() controllers.MovieController
	InjectUserController() controllers.UserController
}

type kernel struct{}



func (k *kernel) InjectMovieController() controllers.MovieController {
	sqlConn, _ := sql.Open("mysql",
		"root:Mint@123@tcp(127.0.0.1:3306)/imdb")
	mysqlHandler := &infrastructures.MySQLHandler{}
	mysqlHandler.Conn = sqlConn

	movieRepository := &repositories.MovieRepository{mysqlHandler}
	movieService := &services.MovieService{&repositories.MovieRepositoryWithCircuitBreaker{movieRepository}}
	movieController := controllers.MovieController{movieService}

	return movieController
}


func (k *kernel) InjectUserController() controllers.UserController {
	sqlConn, _ := sql.Open("mysql",
		"root:Mint@123@tcp(127.0.0.1:3306)/imdb")

	mysqlHandler := &infrastructures.MySQLHandler{}
	mysqlHandler.Conn = sqlConn

	userRepository := &repositories.UserRepository{mysqlHandler}
	userService := &services.UserService{&repositories.UserRepositoryWithCircuitBreaker{userRepository}}
	userController := controllers.UserController{userService}

	return userController
}










var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}

package main

import (
	"sync"

	"github.com/go-chi/chi"
)

type IChiRouter interface {
	InitRouter() *chi.Mux
}

type router struct{}

func (router *router) InitRouter() *chi.Mux {

	movieController:= ServiceContainer().InjectMovieController()
	userController:= ServiceContainer().InjectUserController()
	r := chi.NewRouter()
	r.HandleFunc("/movie/{movie}", movieController.GetMovie)
	r.HandleFunc("/user/rateMovie", userController.RateMovieData)
	r.HandleFunc("/user/commentMovie", userController.CommentMovieData)
	r.HandleFunc("/user/addMovie", userController.AddMovieData)
	r.HandleFunc("/user/listRatedCommentedMovies", userController.ListMoviesData)
	return r
}

var (
	m          *router
	routerOnce sync.Once
)

func ChiRouter() IChiRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return m
}

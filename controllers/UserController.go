package controllers

import (
	"encoding/json"
	"imdbProject/interfaces"
	"imdbProject/models"
	"imdbProject/request"
	"net/http"
)

type UserController struct {
	interfaces.IUserRepository
}

func (controller *UserController) AddMovieData(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var t request.AddMovieRequest
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	success,_:=controller.AddMovie(t.Title,t.Description)
	res.Header().Add("Content-Type","application/json")
	json.NewEncoder(res).Encode(success)
}

func (controller *UserController) RateMovieData(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		res.Header().Add("Content-Type", "application/json")
		s := models.SuccessModel{false, "Only POST Allowed for /rateMovie endpoint "}
		json.NewEncoder(res).Encode(s)
	} else {
		decoder := json.NewDecoder(req.Body)
		var t request.RateMovieRequest
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		success, _ := controller.RateMovie(t.MovieID, t.UserID, t.Rating)
		res.Header().Add("Content-Type", "application/json")
		json.NewEncoder(res).Encode(success)
	}
}
func (controller *UserController) CommentMovieData(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		res.Header().Add("Content-Type", "application/json")
		s := models.SuccessModel{false, "Only POST Allowed for /commentMovie endpoint "}
		json.NewEncoder(res).Encode(s)
	} else {
		decoder := json.NewDecoder(req.Body)
		var t request.CommentMovieRequest
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		success, _ := controller.CommentMovie(t.MovieID, t.UserID, t.Comment)
		res.Header().Add("Content-Type", "application/json")
		json.NewEncoder(res).Encode(success)
	}
}
func (controller *UserController) ListMoviesData(res http.ResponseWriter, req *http.Request) {
	if req.Method=="GET" {
		res.Header().Add("Content-Type", "application/json")
		s := models.SuccessModel{false, "Only POST Allowed for /listRatedCommentedMovies endpoint "}
		json.NewEncoder(res).Encode(s)
	}else {
		decoder := json.NewDecoder(req.Body)
		var t request.ListUserMovieRequest
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		data, _ := controller.ListMoviesRatedCommented(t.UserID)
		res.Header().Add("Content-Type", "application/json")
		json.NewEncoder(res).Encode(data)
	}
}

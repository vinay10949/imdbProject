package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"imdbProject/interfaces"
	"imdbProject/models"
	"net/http"
)

type MovieController struct {
	interfaces.IMovieRepository
}

func (controller *MovieController) GetMovie(res http.ResponseWriter, req *http.Request) {
	if req.Method=="POST"{
		res.Header().Add("Content-Type", "application/json")
		s:=models.SuccessModel{false,"Only Get Allowed for /getMovie endpoint "}
		json.NewEncoder(res).Encode(s)
	}else{

		movie := chi.URLParam(req, "movie")
		data,_:=controller.GetMovieByName(movie)
		res.Header().Add("Content-Type","application/json")
		json.NewEncoder(res).Encode(data)
	}

}

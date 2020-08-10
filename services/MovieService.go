package services

import (
	"fmt"
	"imdbProject/interfaces"
	"imdbProject/viewModels"
)

type MovieService struct {
	interfaces.IMovieRepository
}

func (service *MovieService) GetMovie(movieName string) (viewModels.MovieModel, error) {

	movie, err := service.GetMovieByName(movieName)
	if err != nil {
		fmt.Println("Error occured while ftching moving ", err)
	}
	return movie, err
}

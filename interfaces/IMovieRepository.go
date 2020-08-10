package interfaces

import (
	"imdbProject/viewModels"
)

type IMovieRepository interface {
	GetMovieByName(name string) (viewModels.MovieModel, error)
}

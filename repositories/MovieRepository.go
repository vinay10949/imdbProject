package repositories

import (
	"github.com/afex/hystrix-go/hystrix"
	"imdbProject/interfaces"
	"imdbProject/viewModels"

	"fmt"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type MovieRepositoryWithCircuitBreaker struct {
	MovieRespository interfaces.IMovieRepository
}

func (repository *MovieRepositoryWithCircuitBreaker) GetMovieByName(name string) (viewModels.MovieModel, error) {

	output := make(chan viewModels.MovieModel, 1)
	hystrix.ConfigureCommand("get_movie_by_name", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("get_movie_by_name", func() error {
		movie, _ := repository.MovieRespository.GetMovieByName(name)
		output <- movie
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return viewModels.MovieModel{}, err
	}
}

type MovieRepository struct {
	interfaces.IDbHandler
}

func (repository *MovieRepository) GetMovieByName(name string) (viewModels.MovieModel, error) {
   row, err :=repository.Query(fmt.Sprintf("SELECT ID,TITLE,RATING,DESCRIPTION FROM imdb.movie WHERE title = '%s'  limit 1", name))
	if err != nil {
		return viewModels.MovieModel{}, err
	}

	var movie viewModels.MovieModel
	if row.Next(){

		row.Scan(&movie.Data.ID,&movie.Data.Title, &movie.Data.Rating,&movie.Data.Description)
		movie.Status=true
	}else{
		movie.Status=false
		movie.ErrMsg="Movie not found "
	}
	return movie, nil
}

package services

import (
	"fmt"
	"imdbProject/interfaces"
	"imdbProject/models"
)

type UserService struct {
	interfaces.IUserRepository
}

func (service *UserService) AddMoviebyName(movieName string,description string) (models.SuccessModel, error) {

	success, err := service.AddMovie(movieName,description)
	if err!=nil{
		fmt.Println("Error occured while adding movie ",err)
	}

	return success, err
}

func (service *UserService) RateMovieById(movieId int32,userId int32,rating int) (models.SuccessModel, error) {

	success, err := service.RateMovie(movieId,userId,rating)
	if err!=nil{
		fmt.Println("Error occured while rating on  movie ",err)
	}

	return success, err
}

func (service *UserService) CommentMovieById(movieId int32 ,userId int32,comment string ) (models.SuccessModel, error) {

	success, err := service.CommentMovie(movieId,userId,comment)
	if err!=nil{
		fmt.Println("Error occured while commenting on  movie ",err)
	}

	return success, err
}


func (service *UserService) ListUserRatedCommentedMovies(userId int32 ) ([]models.UserRatingCommentsRespModel, error) {

	data, err := service.ListMoviesRatedCommented(userId)
	if err!=nil{
		fmt.Println("Error occured while fetch user rated commented movie ",err)
	}

	return data, err
}

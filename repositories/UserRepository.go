package repositories

import (
	"github.com/afex/hystrix-go/hystrix"
	"imdbProject/interfaces"
	"imdbProject/models"
	"log"
	"time"

	"fmt"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type UserRepositoryWithCircuitBreaker struct {
	UserRespository interfaces.IUserRepository
}

func (repository *UserRepositoryWithCircuitBreaker) AddMovie(name string, description string) (models.SuccessModel, error) {

	output := make(chan models.SuccessModel, 1)
	hystrix.ConfigureCommand("add_movie_by_name", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("add_movie_by_name", func() error {
		success, _ := repository.UserRespository.AddMovie(name, description)
		output <- success
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return models.SuccessModel{}, err
	}
}

func (repository *UserRepositoryWithCircuitBreaker) RateMovie(movieID int32, userId int32, rating int) (models.SuccessModel, error) {

	output := make(chan models.SuccessModel, 1)
	hystrix.ConfigureCommand("rate_movie_by_id", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("rate_movie_by_id", func() error {
		success, _ := repository.UserRespository.RateMovie(movieID, userId, rating)
		output <- success
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return models.SuccessModel{}, err
	}
}

func (repository *UserRepositoryWithCircuitBreaker) CommentMovie(movieID int32, userId int32, comment string) (models.SuccessModel, error) {

	output := make(chan models.SuccessModel, 1)
	hystrix.ConfigureCommand("comment_movie_by_id", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("comment_movie_by_id", func() error {
		success, _ := repository.UserRespository.CommentMovie(movieID, userId, comment)
		output <- success
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return models.SuccessModel{}, err
	}
}

func (repository *UserRepositoryWithCircuitBreaker) ListMoviesRatedCommented(userId int32) ([]models.UserRatingCommentsRespModel, error) {

	output := make(chan []models.UserRatingCommentsRespModel, 1)
	hystrix.ConfigureCommand("list_movie_by_userId", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("rate_movie_by_id", func() error {
		listMovies, _ := repository.UserRespository.ListMoviesRatedCommented(userId)
		output <- listMovies
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return []models.UserRatingCommentsRespModel{}, err
	}
}

type UserRepository struct {
	interfaces.IDbHandler
}

func (repository *UserRepository) AddMovie(name string, description string) (models.SuccessModel, error) {
	resp := models.SuccessModel{Status: false}
	row, err := repository.Query(fmt.Sprintf("SELECT * FROM imdb.movie WHERE title= '%s'  limit 1", name))
	if row.Next() {
		resp.Message = " Movie already exist "
		resp.Status = false
		return resp, nil
	}
	query := fmt.Sprintf("INSERT INTO movie (title,description,rating,created_by,created_date) VALUES('%s','%s','%s','%s','%s')", name, description, "0", "admin", time.Now().Format(time.RFC3339))
	_, err = repository.Query(query)
	if err == nil {
		resp.Status = true
	}
	return resp, err
}

func (repository *UserRepository) RateMovie(movieId int32, userId int32, rating int) (models.SuccessModel, error) {
	resp := models.SuccessModel{Status: false}

	row, err := repository.Query(fmt.Sprintf("SELECT * FROM imdb.users WHERE id= '%d'  limit 1", userId))
	if !row.Next() {
		resp.Message = " User Doesnt Exist "
		resp.Status = false
		return resp, nil
	}
	row, err = repository.Query(fmt.Sprintf("SELECT * FROM imdb.movie WHERE id= '%d'  limit 1", movieId))
	if !row.Next() {
		resp.Message = " Movie Doesnt Exist "
		resp.Status = false
		return resp, nil
	}
	row, err = repository.Query(fmt.Sprintf("SELECT RATING FROM imdb.user_rating_comments WHERE user_Id= '%d' and movie_Id ='%d'  limit 1", userId ,movieId))
	if row.Next() {
		var rat int
		row.Scan(&rat)
		if  rat!=0 {
			resp.Message = " User can only rate movie once "
			resp.Status = false
			return resp, nil
		}

		query := fmt.Sprintf("Update user_rating_comments set rating='%d' where movie_id='%d' and user_id='%d'",rating, movieId, userId)
		_, err = repository.Query(query)
		resp.Status=true
		return resp,nil
	}
	row, err = repository.Query(fmt.Sprintf("SELECT RATING FROM imdb.user_rating_comments WHERE movie_Id = %d", movieId))
	count := 0
	totalRating := 0
	for row.Next() {
		var rat int
		err := row.Scan(&rat)
		if err != nil {
			log.Fatal(err)
		}
		if rat == 0 {
			continue
		}
		totalRating = totalRating + rat
		count = count + 1

	}

	avgrating := int((totalRating + rating) / (count + 1))
	fmt.Println("AvgRating", avgrating)
	query := fmt.Sprintf("INSERT INTO user_rating_comments (movie_id,user_id,rating,created_by,created_date) VALUES('%d','%d','%d','%s','%s')", movieId, userId, rating, "admin", time.Now().Format(time.RFC3339))
	_, err = repository.Query(query)

	_, err = repository.Query(fmt.Sprintf("Update movie set rating ='%d' where ID ='%d'", avgrating ,movieId))

	if err == nil {
		resp.Status = true
	}
	return resp, err
}

func (repository *UserRepository) CommentMovie(movieId int32, userId int32, comment string) (models.SuccessModel, error) {

	resp := models.SuccessModel{Status: false}

	row, err := repository.Query(fmt.Sprintf("SELECT * FROM imdb.users WHERE id= '%d'  limit 1", userId))
	if !row.Next() {
		resp.Message = " User Doesnt Exist "
		resp.Status = false
		return resp, nil
	}
	row, err = repository.Query(fmt.Sprintf("SELECT * FROM imdb.movie WHERE id= '%d'  limit 1", movieId))
	if !row.Next() {
		resp.Message = " Movie Doesnt Exist "
		resp.Status = false
		return resp, nil
	}
	row, err = repository.Query(fmt.Sprintf("SELECT COMMENT FROM imdb.user_rating_comments WHERE user_Id= '%d' and movie_Id ='%d'  limit 1", userId ,movieId))

	if row.Next() {
		var COMMENT string
		row.Scan(&COMMENT)
		if COMMENT!=""{
			resp.Message = " User can only comment movie once "
			resp.Status = false
			return resp, nil
		}
		fmt.Println("Comment ", comment)
		query := fmt.Sprintf("Update user_rating_comments set comment='%s' where movie_id='%d' and user_id='%d'",comment, movieId, userId)
		_, err = repository.Query(query)

	}else{
		query := fmt.Sprintf("INSERT INTO user_rating_comments (movie_id,user_id,comment,created_by,created_date) VALUES('%d','%d','%s','%s','%s')", movieId, userId, comment, "admin", time.Now().Format(time.RFC3339))
		_, err = repository.Query(query)
	}

	if err == nil {
		resp.Status = true
	}
	return resp, err

}

func (repository *UserRepository) ListMoviesRatedCommented(userId int32) ([]models.UserRatingCommentsRespModel, error) {
	resp := []models.UserRatingCommentsRespModel{}
	rows, err := repository.Query(fmt.Sprintf("Select movie_id,rating,user_id,comment  from user_rating_comments Where user_id ='%d'", userId))
	for rows.Next() {
		var m models.UserRatingCommentsModel
		var n models.UserRatingCommentsRespModel

		err := rows.Scan(&m.MovieID, &m.Rating, &m.UserID, &m.Comment)
		if err != nil {
			log.Fatal(err)
		}
		if m.Comment!=nil{
			n.Comment=*m.Comment
		}
		n.Rating=m.Rating
		n.UserID=m.UserID
		n.MovieID=m.MovieID
		resp = append(resp, n)
	}
	return resp, err
}

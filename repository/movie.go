package repository

import (
	"MovieApi/model"
	"math"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// interface MovieRepo
type MovieRepo interface {
	CreateMovie(movie model.Movie) (model.Movie, error)
	GetMovieById(id uint64) (model.Movie, error)
	GetMovies(page int, limit int) (gin.H, error)
	UpdateMovie(movie model.Movie) (model.Movie, error)
	DeleteMovie(id uint64) error
}

func (r Repo) CreateMovie(movie model.Movie) (model.Movie, error) {
	err := r.db.Create(&movie).Error
	return movie, err
}

func (r Repo) GetMovieById(id uint64) (model.Movie, error) {
	var movie model.Movie
	err := r.db.First(&movie, id).Error
	return movie, err
}

func (r Repo) GetMovies(page int, limit int) (gin.H, error) {
	var movies []model.Movie
	var count int64
	query := r.db

	if page == -1 && limit == -1 {
		return r.getAllMovies(query)
	}

	offset := (page - 1) * limit
	r.db.Model(movies).Count(&count)

	err := query.Offset(offset).Limit(limit).Find(&movies).Error
	if err != nil {
		return nil, err
	}

	responseData := gin.H{
		"movies": movies,
		"meta": gin.H{
			"total_data":   count,
			"current_page": page,
			"limit":        limit,
			"total_page":   int(math.Ceil(float64(count) / float64(limit))),
		},
	}

	return responseData, nil
}

func (r Repo) getAllMovies(query *gorm.DB) (gin.H, error) {
	var movies []model.Movie

	err := query.Find(&movies).Error
	if err != nil {
		return nil, err
	}

	responseData := gin.H{
		"movies": movies,
		"meta": gin.H{
			"total_data": len(movies),
		},
	}

	return responseData, nil
}

func (r Repo) UpdateMovie(movie model.Movie) (model.Movie, error) {
	err := r.db.Updates(&movie).Error
	return movie, err
}

func (r Repo) DeleteMovie(id uint64) error {
	return r.db.Delete(&model.Movie{}, "id = ?", id).Error
}

package service

// usecase
import (
	"MovieApi/model"

	"github.com/gin-gonic/gin"
)

type MovieService interface {
	CreateMovie(movie model.Movie) (model.Movie, error)
	GetMovieById(id uint64) (model.Movie, error)
	GetMovies(page int, limit int) (gin.H, error)
	UpdateMovie(movie model.Movie) (model.Movie, error)
	DeleteMovie(id uint64) error
}

func (s *Service) CreateMovie(movie model.Movie) (model.Movie, error) {
	return s.repo.CreateMovie(movie)
}

func (s *Service) GetMovieById(id uint64) (model.Movie, error) {
	return s.repo.GetMovieById(id)
}

func (s *Service) GetMovies(page int, limit int) (gin.H, error) {
	return s.repo.GetMovies(page, limit)
}

func (s *Service) UpdateMovie(movie model.Movie) (model.Movie, error) {
	return s.repo.UpdateMovie(movie)
}

func (s *Service) DeleteMovie(id uint64) error {
	return s.repo.DeleteMovie(id)
}

package controller

import (
	"MovieApi/helper"
	"MovieApi/model"

	"github.com/gin-gonic/gin"
)

func (h HttpServer) CreateMovie(c *gin.Context) {
	movie := model.Movie{}
	err := helper.BindModelFromContext(c, &movie)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}
	// call service
	res, err := h.app.CreateMovie(movie)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

func (h HttpServer) GetMovieByID(c *gin.Context) {
	id, err := helper.GetMovieIdFromContext(c)
	if err != nil {
		helper.BadRequest(c, helper.ErrInvalidMovieId)
		return
	}

	// call service
	res, err := h.app.GetMovieById(id)
	if err != nil {
		helper.HandleError(err, c)
		return
	}

	helper.Ok(c, res)
}

func (h HttpServer) GetAllMovie(c *gin.Context) {
	page, limit := helper.GetPaginationParams(c)
	// call service
	res, err := h.app.GetMovies(page, limit)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

func (h HttpServer) UpdateMovie(c *gin.Context) {
	id, err := helper.GetMovieIdFromContext(c)
	if err != nil {
		helper.BadRequest(c, helper.ErrInvalidMovieId)
		return
	}

	movie := model.Movie{}
	err = helper.BindModelFromContext(c, &movie)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	movie.ID = id
	// call service
	res, err := h.app.UpdateMovie(movie)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

func (h HttpServer) DeleteMovie(c *gin.Context) {
	id, err := helper.GetMovieIdFromContext(c)
	if err != nil {
		helper.BadRequest(c, helper.ErrInvalidMovieId)
		return
	}

	// call service
	err = h.app.DeleteMovie(id)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, helper.DeleteMovieSuccess)
}

package helper

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	AppJSON            = "application/json"
	ErrInvalidPhotoId  = "Invalid Photo ID"
	ErrInvalidMovieId  = "Invalid Movie ID"
	ErrNotFound        = "not found"
	DeletePhotoSuccess = "Photo deleted successfully"
	DeleteMovieSuccess = "Movie deleted successfully"
)

func GetPhotoIdFromContext(c *gin.Context) (uint64, error) {
	return strconv.ParseUint(c.Param("photoId"), 10, 64)
}

func GetMovieIdFromContext(c *gin.Context) (uint64, error) {
	return strconv.ParseUint(c.Param("movieId"), 10, 64)
}

func BindModelFromContext(c *gin.Context, model interface{}) error {
	contentType := GetContentType(c)

	if contentType == AppJSON {
		if err := c.ShouldBindJSON(model); err != nil {
			return err
		}
	} else {
		if err := c.ShouldBind(model); err != nil {
			return err
		}
	}

	return nil
}

func HandleError(err error, c *gin.Context) {
	if strings.Contains(err.Error(), ErrNotFound) {
		NotFound(c, err.Error())
	} else if strings.Contains(err.Error(), ErrInvalidMovieId) {
		BadRequest(c, err.Error())
	} else {
		InternalServerError(c, err.Error())
	}
}

func GetPaginationParams(c *gin.Context) (int, int) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if limit > 1000 {
		limit = 1000
	}
	return page, limit
}

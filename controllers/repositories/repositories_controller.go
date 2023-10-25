package repositories

import (
	"github_api/domain/repositories"
	"github_api/services"
	"github_api/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)
func CreateRepo(c *gin.Context) {
	var request repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

//	clientId := c.GetHeader("X-Client-Id")

	result, err := services.RepositoryService.CreateRepo( request)
//	result, err := services.RepositoryService.CreateRepo(clientId, request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, result)
}
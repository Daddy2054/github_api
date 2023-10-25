package app

import (
	"github_api/controllers/polo"
	"github_api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	router.POST("/repository", repositories.CreateRepo)
	//router.POST("/repositories", repositories.CreateRepos)
}

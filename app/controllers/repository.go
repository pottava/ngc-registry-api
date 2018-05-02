package controllers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/pottava/ngc-registry-api/app/generated/models"
	"github.com/pottava/ngc-registry-api/app/generated/restapi/operations"
	"github.com/pottava/ngc-registry-api/app/generated/restapi/operations/repository"
)

func repositoryRoute(api *operations.NgcRegistryAPI) {
	api.RepositoryGetRepositoriesHandler = repository.GetRepositoriesHandlerFunc(getRepositories)
}

func getRepositories(params repository.GetRepositoriesParams) middleware.Responder {
	result := models.GetRepositoriesOKBody{}
	return repository.NewGetRepositoriesOK().WithPayload(result)
}

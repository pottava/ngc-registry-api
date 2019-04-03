package controllers

import (
	"log"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/pottava/ngc-registry-api/app/generated/models"
	"github.com/pottava/ngc-registry-api/app/generated/restapi/operations"
	"github.com/pottava/ngc-registry-api/app/generated/restapi/operations/repository"
	"github.com/pottava/ngc-registry-api/app/lib"
	"github.com/pottava/ngc-registry-api/app/ngc/registry"
)

func repositoryRoute(api *operations.NgcRegistryAPI) {
	api.RepositoryGetRepositoriesHandler = repository.GetRepositoriesHandlerFunc(getRepositories)
}

func getRepositories(_ repository.GetRepositoriesParams, auth *lib.Principal) middleware.Responder {
	my, err := registry.Me(auth.Session)
	if err != nil {
		log.Print(err)
		code := http.StatusBadRequest
		return repository.NewGetRepositoriesDefault(code).WithPayload(newerror(code))
	}
	repositries, err := registry.Repositries(auth.Session, my.PriorNamespace(), true, true, true)
	if err != nil {
		log.Print(err)
		code := http.StatusBadRequest
		return repository.NewGetRepositoriesDefault(code).WithPayload(newerror(code))
	}
	result := []*models.Repository{}
	for _, repositry := range repositries {
		result = append(result, &models.Repository{
			Namespace: swag.String(repositry.Namespace),
			Name:      swag.String(repositry.Name),
		})
	}
	return repository.NewGetRepositoriesOK().WithPayload(result)
}

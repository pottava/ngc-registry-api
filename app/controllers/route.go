// Package controllers defines application's routes.
package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/swag"
	"github.com/pottava/ngc-registry-api/app/generated/models"
	"github.com/pottava/ngc-registry-api/app/generated/restapi/operations"
)

// Routes set API handlers
func Routes(api *operations.NgcRegistryAPIAPI) {
	authRoute(api)
	repositoryRoute(api)
	imageRoute(api)
}

func newerror(code int) *models.Error {
	return &models.Error{
		Code:    swag.String(fmt.Sprintf("%d", code)),
		Message: swag.String(http.StatusText(code)),
	}
}

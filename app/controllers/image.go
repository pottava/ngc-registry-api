package controllers

import (
	"log"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/pottava/ngc-registry-api/app/generated/lib"
	"github.com/pottava/ngc-registry-api/app/generated/models"
	"github.com/pottava/ngc-registry-api/app/generated/restapi/operations"
	"github.com/pottava/ngc-registry-api/app/generated/restapi/operations/image"
	"github.com/pottava/ngc-registry-api/app/ngc/registry"
)

func imageRoute(api *operations.NgcRegistryAPIAPI) {
	api.ImageGetImagesHandler = image.GetImagesHandlerFunc(getImages)
}

func getImages(params image.GetImagesParams, auth *lib.Principal) middleware.Responder {
	images, err := registry.Images(auth.Session, params.Namespace, params.ID)
	if err != nil {
		log.Print(err)
		code := http.StatusBadRequest
		return image.NewGetImagesDefault(code).WithPayload(newerror(code))
	}
	result := []*models.Image{}
	for _, image := range images {
		result = append(result, &models.Image{
			Tag:  swag.String(image.Tag),
			Size: swag.Int64(image.Size),
		})
	}
	return image.NewGetImagesOK().WithPayload(result)
}

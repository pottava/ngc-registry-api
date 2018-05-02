package controllers

import (
	"encoding/base64"
	"log"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/pottava/ngc-registry-api/app/generated/models"
	"github.com/pottava/ngc-registry-api/app/generated/restapi/operations"
	"github.com/pottava/ngc-registry-api/app/generated/restapi/operations/auth"
	"github.com/pottava/ngc-registry-api/app/lib"
	"github.com/pottava/ngc-registry-api/app/ngc/registry"
)

func authRoute(api *operations.NgcRegistryAPI) {
	api.AuthSigninHandler = auth.SigninHandlerFunc(signin)
	api.AuthGetMyInfoHandler = auth.GetMyInfoHandlerFunc(myself)
}

func signin(params auth.SigninParams) middleware.Responder {
	session, err := registry.Login(params.Body.Email.String(), params.Body.Password.String())
	if err != nil {
		log.Print(err)
		code := http.StatusBadRequest
		return auth.NewSigninDefault(code).WithPayload(newerror(code))
	}
	encoded := base64.URLEncoding.EncodeToString([]byte(swag.StringValue(session)))
	return auth.NewSigninCreated().WithPayload(&models.Session{
		Jwt: swag.String(string(encoded)),
	})
}

func myself(params auth.GetMyInfoParams, principal *lib.Principal) middleware.Responder {
	my, err := registry.Me(principal.Session)
	if err != nil {
		log.Print(err)
		code := http.StatusBadRequest
		return auth.NewGetMyInfoDefault(code).WithPayload(newerror(code))
	}
	return auth.NewGetMyInfoOK().WithPayload(&models.User{
		ID:        swag.Int64(my.User.ID),
		Name:      my.User.Name,
		Email:     strfmt.Email(my.User.EmailAddress),
		Org:       my.OrganizationName(),
		Created:   my.CreatedDateTime(),
		LastLogin: my.LastLoginDateTime(),
	})
}

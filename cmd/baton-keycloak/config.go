package main

import (
	"github.com/conductorone/baton-sdk/pkg/field"
)

var accessToken = field.StringField("access-token", field.WithDescription("Access Token used to keycloak"), field.WithRequired(true))
var baseUrl = field.StringField("base-url", field.WithDescription("keycloak base url"), field.WithDefaultValue("localhost:8080"))

var configuration = field.Configuration{Fields: []field.SchemaField{baseUrl, accessToken}}

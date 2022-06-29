// Package private provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package private

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Dispatch Playbooks
	// (POST /internal/dispatch)
	ApiInternalRunsCreate(ctx echo.Context) error
	// Cancel Playbook Runs
	// (POST /internal/v2/cancel)
	ApiInternalV2RunsCancel(ctx echo.Context) error
	// Dispatch Playbooks
	// (POST /internal/v2/dispatch)
	ApiInternalV2RunsCreate(ctx echo.Context) error
	// Obtain connection status of recipient(s)
	// (POST /internal/v2/recipients/status)
	ApiInternalV2RecipientsStatus(ctx echo.Context) error
	// Get Version
	// (GET /internal/version)
	ApiInternalVersion(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// ApiInternalRunsCreate converts echo context to params.
func (w *ServerInterfaceWrapper) ApiInternalRunsCreate(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ApiInternalRunsCreate(ctx)
	return err
}

// ApiInternalV2RunsCancel converts echo context to params.
func (w *ServerInterfaceWrapper) ApiInternalV2RunsCancel(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ApiInternalV2RunsCancel(ctx)
	return err
}

// ApiInternalV2RunsCreate converts echo context to params.
func (w *ServerInterfaceWrapper) ApiInternalV2RunsCreate(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ApiInternalV2RunsCreate(ctx)
	return err
}

// ApiInternalV2RecipientsStatus converts echo context to params.
func (w *ServerInterfaceWrapper) ApiInternalV2RecipientsStatus(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ApiInternalV2RecipientsStatus(ctx)
	return err
}

// ApiInternalVersion converts echo context to params.
func (w *ServerInterfaceWrapper) ApiInternalVersion(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ApiInternalVersion(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/internal/dispatch", wrapper.ApiInternalRunsCreate)
	router.POST(baseURL+"/internal/v2/cancel", wrapper.ApiInternalV2RunsCancel)
	router.POST(baseURL+"/internal/v2/dispatch", wrapper.ApiInternalV2RunsCreate)
	router.POST(baseURL+"/internal/v2/recipients/status", wrapper.ApiInternalV2RecipientsStatus)
	router.GET(baseURL+"/internal/version", wrapper.ApiInternalVersion)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"0e2n7V8BAAD//wWMJ/R/GwAA",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}

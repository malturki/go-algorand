// Package private provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package private

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/algorand/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Aborts a catchpoint catchup.
	// (DELETE /v2/catchup/{catchpoint})
	AbortCatchup(ctx echo.Context, catchpoint string) error
	// Starts a catchpoint catchup.
	// (POST /v2/catchup/{catchpoint})
	StartCatchup(ctx echo.Context, catchpoint string) error
	// Return a list of participation keys
	// (GET /v2/participation)
	GetParticipationKeys(ctx echo.Context) error
	// Add a participation key to the node
	// (POST /v2/participation)
	AddParticipationKey(ctx echo.Context) error
	// Delete a given participation key by ID
	// (DELETE /v2/participation/{participation-id})
	DeleteParticipationKeyByID(ctx echo.Context, participationId string) error
	// Get participation key info given a participation ID
	// (GET /v2/participation/{participation-id})
	GetParticipationKeyByID(ctx echo.Context, participationId string) error
	// Append state proof keys to a participation key
	// (POST /v2/participation/{participation-id})
	AppendKeys(ctx echo.Context, participationId string) error

	// (POST /v2/shutdown)
	ShutdownNode(ctx echo.Context, params ShutdownNodeParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AbortCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) AbortCatchup(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameterWithLocation("simple", false, "catchpoint", runtime.ParamLocationPath, ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AbortCatchup(ctx, catchpoint)
	return err
}

// StartCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) StartCatchup(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameterWithLocation("simple", false, "catchpoint", runtime.ParamLocationPath, ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.StartCatchup(ctx, catchpoint)
	return err
}

// GetParticipationKeys converts echo context to params.
func (w *ServerInterfaceWrapper) GetParticipationKeys(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetParticipationKeys(ctx)
	return err
}

// AddParticipationKey converts echo context to params.
func (w *ServerInterfaceWrapper) AddParticipationKey(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddParticipationKey(ctx)
	return err
}

// DeleteParticipationKeyByID converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteParticipationKeyByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "participation-id", runtime.ParamLocationPath, ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteParticipationKeyByID(ctx, participationId)
	return err
}

// GetParticipationKeyByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetParticipationKeyByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "participation-id", runtime.ParamLocationPath, ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetParticipationKeyByID(ctx, participationId)
	return err
}

// AppendKeys converts echo context to params.
func (w *ServerInterfaceWrapper) AppendKeys(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "participation-id", runtime.ParamLocationPath, ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AppendKeys(ctx, participationId)
	return err
}

// ShutdownNode converts echo context to params.
func (w *ServerInterfaceWrapper) ShutdownNode(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ShutdownNodeParams
	// ------------- Optional query parameter "timeout" -------------

	err = runtime.BindQueryParameter("form", true, false, "timeout", ctx.QueryParams(), &params.Timeout)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter timeout: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ShutdownNode(ctx, params)
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
func RegisterHandlers(router EchoRouter, si ServerInterface, m ...echo.MiddlewareFunc) {
	RegisterHandlersWithBaseURL(router, si, "", m...)
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.DELETE(baseURL+"/v2/catchup/:catchpoint", wrapper.AbortCatchup, m...)
	router.POST(baseURL+"/v2/catchup/:catchpoint", wrapper.StartCatchup, m...)
	router.GET(baseURL+"/v2/participation", wrapper.GetParticipationKeys, m...)
	router.POST(baseURL+"/v2/participation", wrapper.AddParticipationKey, m...)
	router.DELETE(baseURL+"/v2/participation/:participation-id", wrapper.DeleteParticipationKeyByID, m...)
	router.GET(baseURL+"/v2/participation/:participation-id", wrapper.GetParticipationKeyByID, m...)
	router.POST(baseURL+"/v2/participation/:participation-id", wrapper.AppendKeys, m...)
	router.POST(baseURL+"/v2/shutdown", wrapper.ShutdownNode, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"//864HIpaDwAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
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

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}

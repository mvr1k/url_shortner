package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*API interface tell that who ever is a API can be passed to the web server as a API handler*/
const (
	GET     = "GET"
	LIST    = "LIST"
	POST    = "POST"
	PUT     = "PUT"
	DELETE  = "DELETE"
	HEAD    = "HEAD"
	OPTIONS = "OPTIONS"

	ErrWrongPath = "Opps Wrong API path or Wrong Method please check"
)

type API interface {
	GET(ctx *gin.Context)     //get the data
	LIST(ctx *gin.Context)    //get list of data as paginatiion
	POST(ctx *gin.Context)    // craete data
	PUT(ctx *gin.Context)     // update data
	DELETE(ctx *gin.Context)  //delete data
	HEAD(ctx *gin.Context)    //get required headers
	OPTIONS(ctx *gin.Context) //get options

	GetRouteMapping() map[string]string
	ModuleName() string //request method to route mapping for that API router
}

func DefaultHandler(ctx *gin.Context) {
	_ = SendResponse(http.StatusNotFound, nil, ErrWrongPath, ctx)
}

func GetMappedMethod(method string, api API) func(ctx *gin.Context) {
	switch method {
	case http.MethodGet:
		return api.GET
	case http.MethodPost:
		return api.POST
	case http.MethodPut:
		return api.PUT
	case http.MethodDelete:
		return api.DELETE
	case http.MethodHead:
		return api.HEAD
	case http.MethodOptions:
		return api.OPTIONS
	default:
		return DefaultHandler
	}
}

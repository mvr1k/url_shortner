package shortner

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
	"url_shortner/app/internal/shortner/data"
	"url_shortner/app/internal/web/api"
)

const moduleName = "MODULE_URL_SHORTENER"
const moduleBasePath = "/shotgun"

type WebApi struct {
	logger  *zap.SugaredLogger
	manager *manager
}

func NewWebApi(logger *zap.SugaredLogger) *WebApi {
	return &WebApi{
		logger:  logger,
		manager: newManager(logger),
	}
}

func (w WebApi) ModuleName() string {
	return moduleName
}

func (w WebApi) GET(ctx *gin.Context) {
	w.GenerateContextID(ctx)
	req := data.NewGetRequestFromHttpRequest(ctx)
	if err := req.IsValid(); err != nil {
		_ = api.SendResponse(http.StatusBadRequest, nil, err.Error(), ctx)
		//context based logging
		return
	}
	if req.ShortUrl == "" {
		respData, err := w.manager.getListOfUrlsForUser(req, ctx)
		if err != nil {
			_ = api.SendResponse(http.StatusBadRequest, nil, err.Error(), ctx)
			return
		}
		_ = api.SendResponse(http.StatusOK, respData, err.Error(), ctx)
	} else {
		w.manager.redirectToLongUrl(req, ctx)
	}

}

func (w WebApi) LIST(ctx *gin.Context) {
	api.DefaultHandler(ctx)
}

func (w WebApi) POST(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w WebApi) PUT(ctx *gin.Context) {
	api.DefaultHandler(ctx)

}

func (w WebApi) DELETE(ctx *gin.Context) {
	api.DefaultHandler(ctx)
}

func (w WebApi) HEAD(ctx *gin.Context) {
	api.DefaultHandler(ctx)
}

func (w WebApi) OPTIONS(ctx *gin.Context) {
	api.DefaultHandler(ctx)
}

func (w WebApi) GetRouteMapping() map[string]string {
	mappings := map[string]string{
		http.MethodGet:  moduleBasePath + "/:" + data.PathVariableForShortUrl,
		http.MethodPost: moduleBasePath,
	}

	return mappings
}

func (w WebApi) GenerateContextID(context *gin.Context) {
	id := fmt.Sprintf("%v", time.Now().UnixNano())
	context.Set("request-id", id)
	w.logger.Info(
		zap.String("route", context.FullPath()),
		zap.String("request_id", context.GetString("request-id")),
		"hit received. ",
	)
}

package data

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
)

type PostRequest struct {
	ClientId string `json:"client_id"`
	LongUrl  string `json:"long_url"`
}

func NewPostRequestFromHttpRequest(context *gin.Context) (request PostRequest, err error) {
	err = json.NewDecoder(context.Request.Body).Decode(&request)
	return
}

const PathVariableForShortUrl = "short-url"
const PathVariableForUser = "user"

type GetRequest struct {
	ShortUrl string
	User     string
}

const (
	ErrorClintIdNotFound = "please Provide User ID"
)

func (r GetRequest) IsValid() error {
	if r.User == "" {
		return errors.New(ErrorClintIdNotFound)
	}
	return nil
}

func NewGetRequestFromHttpRequest(context *gin.Context) (request GetRequest) {
	request.ShortUrl = context.Param(PathVariableForShortUrl)
	request.User = context.Param(PathVariableForUser)
	return
}

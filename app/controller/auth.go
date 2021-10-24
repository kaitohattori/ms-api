package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"ms-api/app/util"
	"ms-api/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authUtil util.AuthUtil
}

func NewAuthController(authUtil util.AuthUtil) *AuthController {
	return &AuthController{authUtil: authUtil}
}

// AuthController Login docs
// @Summary login
// @Description login
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /login [get]
func (c *AuthController) Login(ctx *gin.Context) {
	auth0Url := "https://kaitohattori.jp.auth0.com/oauth/token"

	payload := map[string]string{
		"client_id":     config.Config.Auth0ClientId,
		"client_secret": config.Config.Auth0ClientSecret,
		"audience":      config.Config.Auth0Identifier,
		"grant_type":    "client_credentials",
	}
	bytes, err := json.Marshal(payload)
	if err != nil {
		return
	}

	request, err := http.NewRequest("POST", auth0Url, strings.NewReader(string(bytes)))
	if err != nil {
		log.Println(err)
		util.NewError(ctx, http.StatusNotFound, err)
		return
	}
	request.Header.Add("content-type", "application/json")

	// Request
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println(err)
		util.NewError(ctx, http.StatusNotFound, err)
		return
	}
	defer response.Body.Close()

	auth := struct {
		TokenType   string `json:"token_type"`
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}{}
	byteArray, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(byteArray, &auth)
	if err != nil {
		log.Println(err)
		util.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, auth)
}

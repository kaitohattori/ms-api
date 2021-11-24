package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"ms-api/config"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
)

type AuthUtil struct {
	jwtMiddleware                *jwtmiddleware.JWTMiddleware
	jwtMiddlewareAuthNotRequired *jwtmiddleware.JWTMiddleware
}

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

func NewAuthUtil(identifer string, domain string) AuthUtil {
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			// Verify 'aud' claim
			checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(identifer, false)
			if !checkAud {
				return token, errors.New("invalid audience")
			}
			// Verify 'iss' claim
			checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(domain, false)
			if !checkIss {
				return token, errors.New("invalid issuer")
			}
			// Get pem certification
			cert, err := AuthUtilGetPemCert(token)
			if err != nil {
				panic(err.Error())
			}
			result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(*cert))
			return result, nil
		},
		SigningMethod: jwt.SigningMethodRS256,
	})
	jwtMiddlewareAuthNotRequired := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			// Verify 'aud' claim
			checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(identifer, false)
			if !checkAud {
				return token, errors.New("invalid audience")
			}
			// Verify 'iss' claim
			checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(domain, false)
			if !checkIss {
				return token, errors.New("invalid issuer")
			}
			// Get pem certification
			cert, err := AuthUtilGetPemCert(token)
			if err != nil {
				panic(err.Error())
			}
			result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(*cert))
			return result, nil
		},
		SigningMethod: jwt.SigningMethodRS256,
		ErrorHandler:  func(w http.ResponseWriter, r *http.Request, err string) {},
	})
	return AuthUtil{
		jwtMiddleware:                jwtMiddleware,
		jwtMiddlewareAuthNotRequired: jwtMiddlewareAuthNotRequired,
	}
}

func (a AuthUtil) CheckJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtMid := *a.jwtMiddleware
		if err := jwtMid.CheckJWT(c.Writer, c.Request); err != nil {
			c.AbortWithStatus(401)
		}
	}
}

func (a AuthUtil) CheckJWTAuthNotRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtMid := *a.jwtMiddlewareAuthNotRequired
		if err := jwtMid.CheckJWT(c.Writer, c.Request); err != nil {
			log.Println(err)
		}
	}
}

func AuthUtilGetPemCert(token *jwt.Token) (*string, error) {
	url := fmt.Sprintf("%s.well-known/jwks.json", config.Config.Auth0Domain)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)
	if err != nil {
		return nil, err
	}

	cert := ""
	for k, _ := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}
	if cert == "" {
		err := errors.New("unable to find appropriate key")
		return nil, err
	}
	return &cert, nil
}

func AuthUtilGetUserId(ctx *gin.Context) (*string, error) {
	user := ctx.Request.Context().Value("user")
	if user == nil {
		return nil, errors.New("failed to get userId")
	}
	claims := user.(*jwt.Token).Claims.(jwt.MapClaims)
	sub := claims["sub"].(string)
	return &sub, nil
}

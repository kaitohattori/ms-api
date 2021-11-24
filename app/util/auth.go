package util

import (
	"context"
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
	jwtMiddleware *jwtmiddleware.JWTMiddleware
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
	return AuthUtil{
		jwtMiddleware: jwtMiddleware,
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

func (a AuthUtil) CheckJWTNotRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := AuthUtilParseToken(a.jwtMiddleware, c.Writer, c.Request)
		if err != nil {
			log.Println(err)
			return
		}
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), a.jwtMiddleware.Options.UserProperty, token))
	}
}

func AuthUtilParseToken(m *jwtmiddleware.JWTMiddleware, w http.ResponseWriter, r *http.Request) (*jwt.Token, error) {
	token, err := m.Options.Extractor(r)
	if err != nil {
		return nil, fmt.Errorf("error extracting token: %w", err)
	}

	// if token is empty
	if token == "" {
		return nil, fmt.Errorf("required authorization token not found")
	}

	// Now parse the token
	parsedToken, err := jwt.Parse(token, m.Options.ValidationKeyGetter)

	// Check if there was an error in parsing...
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	if m.Options.SigningMethod != nil && m.Options.SigningMethod.Alg() != parsedToken.Header["alg"] {
		message := fmt.Sprintf("Expected %s signing method but token specified %s",
			m.Options.SigningMethod.Alg(),
			parsedToken.Header["alg"])
		return nil, fmt.Errorf("error validating token algorithm: %s", message)
	}

	// Check if the parsed token is valid...
	if !parsedToken.Valid {
		return nil, errors.New("token is invalid")
	}

	return parsedToken, nil
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

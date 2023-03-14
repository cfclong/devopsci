package middleware

import (
	"errors"
	"time"

	"gitee.com/plutoccc/devops_app/internal/middleware/log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
)

// JwtAuth ..
func JwtAuth(username string, role string) (t string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix()
	t, err = token.SignedString([]byte(beego.AppConfig.String("jwt::secret")))
	if err != nil {
		return "", errors.New("JWT Generate Failure")
	}
	return t, nil
}

// JwtParse ..
// TODO: check jwt is not expired
func JwtParse(c *context.Context, token string) (string, error) {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod(jwt.SigningMethodHS256.Alg()) != t.Method {
			return nil, errors.New("Invalid signing algorithm")
		}
		return []byte(beego.AppConfig.String("jwt::secret")), nil
	})
	if err != nil {
		log.Log.Error("Parset token error: %s", err.Error())
		return "", err
	}
	claims := jwtToken.Claims.(jwt.MapClaims)
	return claims["username"].(string), nil
}

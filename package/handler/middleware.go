package handler

import (
	"errors"

	"github.com/gin-contrib/cors"
	"github.com/mrboburs/Norbekov/util/logrus"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (handler *Handler) userIdentity(ctx *gin.Context) {
	logrus := handler.logrus
	header := ctx.GetHeader(authorizationHeader)
	logrus.Info(header)
	if header == "" {
		NewHandlerErrorResponse(ctx, http.StatusUnauthorized, "empty auth header", logrus)
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		NewHandlerErrorResponse(ctx, http.StatusUnauthorized, "invalid auth header", logrus)
		return
	}

	userId, err := handler.services.Admin.ParseToken(headerParts[1])
	if err != nil {
		NewHandlerErrorResponse(ctx, http.StatusUnauthorized, err.Error(), logrus)
		return
	}

	ctx.Set(userCtx, userId)
}

func getUserId(ctx *gin.Context, logrus *logrus.Logger) (int, error) {
	id, ok := ctx.Get(userCtx)
	if !ok {
		NewHandlerErrorResponse(ctx, http.StatusInternalServerError, "user id not found", logrus)
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		NewHandlerErrorResponse(ctx, http.StatusInternalServerError, "user id is of invalid type", logrus)
		return 0, errors.New("user id not found")
	}

	return idInt, nil
}

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT,PATCH,DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
func CORSConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://norbekov.herokuapp.com"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers", "Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization")
	corsConfig.AddAllowMethods("GET", "POST", "PUT", "DELETE")
	return corsConfig
}

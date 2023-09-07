package controllers

import (
	"net/http"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var (
	Router      *gin.Engine
	cookieStore = cookie.NewStore([]byte("thisIsSecret"))
)

func init() {
	Router = gin.Default()
	Router.HTMLRender = ginview.Default()

	Router.Use(sessions.Sessions("_cimelli_sessions", cookieStore))

	Router.GET("/", SignedInRootPath)
	Router.GET("/photos", GetPhotos)
	Router.POST("/photo", UploadPhoto)
	Router.GET("/users", GetIndexUsers)
	Router.GET("/users/new", GetNewUser)
	Router.POST("/users/:id", PostUserId)
	Router.GET("/users/:id", EditUser)
	Router.POST("/user", PostNewUser)
	Router.Static("/assets", "assets")
	Router.Static("/uploads", "public/uploads")
	Router.GET("/ping", PingPingPath)

}

func SignedInRootPath(ctx *gin.Context) {
	q := ctx.Query("q")
	ctx.HTML(http.StatusOK, "index", gin.H{"name": "Fauzan", "query": q})
}

func PingPingPath(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

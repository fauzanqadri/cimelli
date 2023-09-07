package controllers

import (
	"io/ioutil"
	"models"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetPhotos(ctx *gin.Context) {
	// session := sessions.Default(ctx)

	// if session.Get("photos") == nil {
	// 	session.Set("photos", "Foo")
	// 	err := session.Save()
	// 	fmt.Println(err)
	// }
	page := 1

	pageQ := ctx.Query("page")

	if pageQ != "" {
		page, _ = strconv.Atoi(pageQ)
	}

	photos, err := models.GetPagedPhoto(page)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.HTML(http.StatusOK, "album", gin.H{"photos": photos})
}

func UploadPhoto(ctx *gin.Context) {
	fh, err := ctx.FormFile("photo")

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	f, err := fh.Open()

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	defer f.Close()

	bs, err := ioutil.ReadAll(f)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	fname := strings.ToLower(fh.Filename)
	floc := "public/uploads/" + fname
	urlPath := "/uploads/" + fname

	newUp, err := os.Create(floc)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	defer newUp.Close()

	_, err = newUp.Write(bs)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	photo, err := models.NewPhoto(fname, floc, urlPath, fh.Size)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = photo.Insert()

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/photos")
}

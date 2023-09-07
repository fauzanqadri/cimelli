package controllers

import (
	"models"
	"net/http"
	"producers"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetIndexUsers(ctx *gin.Context) {
	page := 1

	pageQ := ctx.Query("page")

	if pageQ != "" {
		page, _ = strconv.Atoi(pageQ)
	}

	paginatedContent, err := models.GetPagedUser(page)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.HTML(http.StatusOK, "users", gin.H{"users": paginatedContent})
}

func GetNewUser(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "new_user_form", gin.H{})
}

func PostNewUser(ctx *gin.Context) {

	user := &models.User{}

	if err := ctx.Bind(user); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	p := ctx.PostForm("password")
	pc := ctx.PostForm("password_confirmation")

	if err := user.SetPassword(p, pc); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := user.Insert(); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	producers.PublishUserCreation(user)

	ctx.Redirect(http.StatusSeeOther, "/users")
}

func PostUserId(ctx *gin.Context) {
	method := ctx.PostForm("_method")
	if method == "delete" {
		deleteUser(ctx)
		return
	} else if method == "patch" {
		patchUser(ctx)
		return
	}

	ctx.AbortWithStatus(http.StatusNotFound)
}

func patchUser(ctx *gin.Context) {
	var id uint64

	idQ := ctx.Param("id")

	if idQ != "" {
		id, _ = strconv.ParseUint(idQ, 10, 64)
	}

	user, err := models.GetUserById(id)

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotExtended)
		return
	}

	beforeData := user.Copy()

	if err := ctx.Bind(user); err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	p := ctx.PostForm("password")
	pc := ctx.PostForm("password_confirmation")

	if p != "" && pc != "" {
		user.SetPassword(p, pc)
	}

	if err := user.Update(); err != nil {
		ctx.AbortWithStatus(http.StatusUnprocessableEntity)
	}

	updatedUser, err := models.GetUserById(id)

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotExtended)
		return
	}

	producers.PublishUserUpdate(beforeData, updatedUser)

	ctx.Redirect(http.StatusSeeOther, "/users")
}

func deleteUser(ctx *gin.Context) {
	var id uint64

	idQ := ctx.Param("id")

	if idQ != "" {
		// id, _ = strconv.Atoi(idQ)
		id, _ = strconv.ParseUint(idQ, 10, 64)
	}

	user, err := models.DeleteUser(id)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	producers.PublishUserDeletion(user)

	ctx.Redirect(http.StatusSeeOther, "/users")
}

func EditUser(ctx *gin.Context) {
	var id uint64

	idQ := ctx.Param("id")

	if idQ != "" {
		// id, _ = strconv.Atoi(idQ)
		id, _ = strconv.ParseUint(idQ, 10, 64)
	}

	user, err := models.GetUserById(id)

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotExtended)
		return
	}

	ctx.HTML(http.StatusOK, "edit_user_form", gin.H{"user": user})
}

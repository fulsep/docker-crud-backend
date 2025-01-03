package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/fulsep/docker-crud-backend/tree/main/dto"
	"github.com/fulsep/docker-crud-backend/tree/main/models"
	"github.com/gin-gonic/gin"
)

func ListAllUsers(ctx *gin.Context) {
	users, err := models.FindAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "Something happened",
		})
	}

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "List all users",
		Results: users,
	})
}

func DetailUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	user, err := models.FindOneUser(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "Something happened",
		})
	}

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Detail users",
		Results: user,
	})
}

func CreateUser(ctx *gin.Context) {
	form := dto.InsertUser{}

	err := ctx.ShouldBind(&form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Validation error",
		})
	}

	newUser := models.User{
		Email:    form.Email,
		Password: form.Password,
	}

	user, err := models.InsertUser(newUser)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "Something happened",
		})
	}

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Create user success",
		Results: user,
	})
}

func UpdateUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user, findErr := models.FindOneUser(id)

	if findErr != nil {
		ctx.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "Something happened",
		})
	}

	if user == (models.User{}) {
		ctx.JSON(http.StatusNotFound, Response{
			Success: false,
			Message: "User Not found",
		})
	}

	form := dto.InsertUser{}

	bindErr := ctx.ShouldBind(&form)

	if bindErr != nil {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Validation error",
		})
	}

	newUser := models.User{
		Id:       id,
		Email:    form.Email,
		Password: form.Password,
	}

	user, updateErr := models.UpdateUser(newUser)

	if updateErr != nil {
		log.Println(updateErr)
		ctx.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "Something happened",
		})
	}

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Update user success",
		Results: user,
	})
}

func DeleteUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	user, err := models.FindOneUser(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "Something happened",
		})
	}

	if user == (models.User{}) {
		ctx.JSON(http.StatusNotFound, Response{
			Success: false,
			Message: "User not found",
		})
	}

	deletedUser, err := models.DeleteUser(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "Something happened",
		})
	}

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "User deleted succesfully",
		Results: deletedUser,
	})
}

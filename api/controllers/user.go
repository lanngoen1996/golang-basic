package controller

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	model "github.com/lanngoen1996/golang-basic/models"
	"github.com/lanngoen1996/golang-basic/validators"
)

type UserControllerInterface interface {
	GetUsers(c *gin.Context)
	CreateUser(c *gin.Context)
	FindUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	Delete(c *gin.Context)
}

type UserController struct {
	User model.UserInterface
}

func NewUserController() UserControllerInterface {
	model := model.NewUserModel()
	return &UserController{
		User: model,
	}
}

func (u *UserController) GetUsers(c *gin.Context) {
	var users []model.User
	if err := u.User.Get(&users); err != nil {
		c.JSON(500, gin.H{
			"message": "Get users error.",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Successful.",
		"data":    users,
	})
}

func (u *UserController) CreateUser(c *gin.Context) {
	var payload validators.UserValidator

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request.",
		})
		return
	}

	user := &model.User{
		Username:  payload.Username,
		Password:  payload.Password,
		Email:     payload.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := u.User.Create(user); err != nil {
		c.JSON(500, gin.H{
			"message": "Create error.",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Successful.",
		"data":    user,
	})
}

func (u *UserController) FindUser(c *gin.Context) {
	var user model.User
	id := c.Param("id")

	if err := u.User.Find(id, &user); err != nil {
		c.JSON(404, gin.H{
			"message": "Not found.",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Successful.",
		"data":    user,
	})
}

func (u *UserController) UpdateUser(c *gin.Context) {
	var payload validators.UserValidator
	id := c.Param("id")
	var user model.User

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request.",
		})
		return
	}

	if err := u.User.Find(id, &user); err != nil {
		c.JSON(404, gin.H{
			"message": "Not found.",
		})
		return
	}

	p, _ := json.Marshal(payload)
	json.Unmarshal(p, &user)

	if err := u.User.Update(&user); err != nil {
		c.JSON(500, gin.H{
			"message": "Error update.",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Successful.",
		"data":    user,
	})
}

func (u *UserController) Delete(c *gin.Context) {
	id := c.Param("id")
	var user model.User

	if err := u.User.Find(id, &user); err != nil {
		c.JSON(404, gin.H{
			"message": "Not found.",
		})
		return
	}

	err := u.User.Delete(&user)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error delete.",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Successful.",
	})
}

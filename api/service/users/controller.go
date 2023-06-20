package users

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserControllerInterface interface {
	GetUsers(c *gin.Context)
	CreateUser(c *gin.Context)
	FindUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	Delete(c *gin.Context)
}

type UserController struct {
	model UserInterface
}

func NewUserController(DB *gorm.DB) UserControllerInterface {
	model := NewUser(DB)
	return &UserController{
		model,
	}
}

func (u *UserController) GetUsers(c *gin.Context) {
	var users []User
	if err := u.model.Get(&users); err != nil {
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
	var payload UserValidator

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request.",
		})
		return
	}

	user := &User{
		Username:  payload.Username,
		Password:  payload.Password,
		Email:     payload.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := u.model.Create(user); err != nil {
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
	var user User
	id := c.Param("id")

	if err := u.model.Find(id, &user); err != nil {
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
	var payload UserValidator
	id := c.Param("id")
	var user User

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request.",
		})
		return
	}

	if err := u.model.Find(id, &user); err != nil {
		c.JSON(404, gin.H{
			"message": "Not found.",
		})
		return
	}

	p, _ := json.Marshal(payload)
	json.Unmarshal(p, &user)

	if err := u.model.Update(&user); err != nil {
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
	var user User

	if err := u.model.Find(id, &user); err != nil {
		c.JSON(404, gin.H{
			"message": "Not found.",
		})
		return
	}

	err := u.model.Delete(&user)
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

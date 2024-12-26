package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/golang-gorm-postgres/models"
	"gorm.io/gorm"
)

type PostController struct {
	DB *gorm.DB
}

func NewPostController(DB *gorm.DB) PostController {
	return PostController{DB}
}

func (pc *PostController) CreatePost(cntxt *gin.Context) {
	currUser := cntxt.MustGet("currentUser").(models.User)
	var payload *models.CreatePostRequest

	if err := cntxt.ShouldBindJSON(&payload); err != nil {
		cntxt.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	newPost := models.Post{
		Title:     payload.Title,
		Content:   payload.Content,
		Image:     payload.Image,
		User:      currUser.ID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	res := pc.DB.Create(&newPost)
	if res.Error != nil {
		if strings.Contains(res.Error.Error(), "Duplicated key!") {
			cntxt.JSON(http.StatusConflict, gin.H{"Status": "Fail", "Message": "Post with that title already exists"})
			return
		}

		cntxt.JSON(http.StatusBadGateway, gin.H{"Status": "Error", "Message": res.Error.Error()})
		return
	}

	cntxt.JSON(http.StatusCreated, gin.H{"Status": "Success", "Message": newPost})
}
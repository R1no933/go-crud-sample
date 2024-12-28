package controllers

import (
	"net/http"
	"strconv"
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

func (pc *PostController) UpdatePost(cntxt *gin.Context) {
	postId := cntxt.Param("postId")
	currUser := cntxt.MustGet("currentUser").(models.User)

	var payload *models.UpdatePost
	if err := cntxt.ShouldBindJSON(&payload); err != nil {
		cntxt.JSON(http.StatusBadGateway, gin.H{"Status": "Error", "Message": err.Error()})
		return
	}

	var updatePost models.Post
	res := pc.DB.First(&updatePost, "Id = ?", postId)
	if res.Error != nil {
		cntxt.JSON(http.StatusNotFound, gin.H{"Status": "Fail", "Message": "No post with that title exists"})
		return
	}

	now := time.Now()
	postToUpdate := models.Post{
		Title:     payload.Title,
		Content:   payload.Content,
		Image:     payload.Image,
		User:      currUser.ID,
		CreatedAt: updatePost.CreatedAt,
		UpdatedAt: now,
	}

	pc.DB.Model(&updatePost).Updates(postToUpdate)

	cntxt.JSON(http.StatusOK, gin.H{"Status": "Success", "Data": updatePost})
}

func (pc *PostController) FindPostById(cntxt *gin.Context) {
	postId := cntxt.Param("postId")

	var post models.Post
	res := pc.DB.First(&post, "id = ?", postId)
	if res.Error != nil {
		cntxt.JSON(http.StatusNotFound, gin.H{"Status": "Fail", "Message": "No post with that title exists"})
		return
	}

	cntxt.JSON(http.StatusOK, gin.H{"Status": "Success", "Data": post})
}

func (pc *PostController) FindPosts(cntxt *gin.Context) {
	var page = cntxt.DefaultQuery("page", "1")
	var limit = cntxt.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var posts []models.Post
	res := pc.DB.Limit(intLimit).Offset(offset).Find(&posts)
	if res.Error != nil {
		cntxt.JSON(http.StatusBadGateway, gin.H{"Status": "Error", "Message": res.Error})
		return
	}
	cntxt.JSON(http.StatusOK, gin.H{"Status": "Success", "Results": len(posts), "Data": posts})
}

func (pc *PostController) DeletePost(cntxt *gin.Context) {
	postId := cntxt.Param("postId")

	res := pc.DB.Delete(&models.Post{}, "id = ?", postId)

	if res.Error != nil {
		cntxt.JSON(http.StatusNotFound, gin.H{"Status": "Fail", "Message": "No post with that title exists"})
		return
	}

	cntxt.JSON(http.StatusNoContent, nil)
}

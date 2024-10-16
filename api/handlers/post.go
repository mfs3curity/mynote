package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mfs3curity/mynote/api/dto"
	"github.com/mfs3curity/mynote/services"
)

type PostHandlers struct {
	services services.PostService
}

func NewPostHandlers() *PostHandlers {
	return &PostHandlers{
		services: *services.NewPostService(),
	}
}

func (p *PostHandlers) Create(ctx *gin.Context) {
	req := new(dto.CreatePost)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	post, err := p.services.CreatePost(ctx, req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"path": fmt.Sprintf("/api/post/id/%d", post),
	})
}

// get by id
func (p *PostHandlers) GetByID(ctx *gin.Context) {
	i := ctx.Param("id")
	id, err := strconv.Atoi(i)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	result, err := p.services.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, result)
}

// upload img
func (p *PostHandlers) Upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	// توليد UUID كاسم عشوائي للملف
	randomID := uuid.New().String()
	// استخراج الامتداد من اسم الملف الأصلي
	ext := filepath.Ext(file.Filename)
	// تكوين اسم الملف الجديد
	newFilename := randomID + ext
	path := "images/uploads/" + newFilename

	// حفظ الملف
	err = ctx.SaveUploadedFile(file, path)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	save, err := p.services.UploadImg(ctx, path)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, save)
}

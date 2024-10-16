package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mfs3curity/mynote/api/dto"
	"github.com/mfs3curity/mynote/services"
)

type SectionHandler struct {
	service *services.SectionService
}

func NewSectionHandler() *SectionHandler {
	return &SectionHandler{
		service: services.NewSectionService(),
	}
}
func (p *SectionHandler) Create(ctx *gin.Context) {
	req := new(dto.CreatSection)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	section, err := p.service.CreateSection(ctx, req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"path": fmt.Sprintf("/api/section/name/%s", section),
	})
}

func (p *SectionHandler) GetByName(ctx *gin.Context) {
	limit, err := parseQueryParam(ctx, "limit", 10)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	offset, err := parseQueryParam(ctx, "offset", 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("Offset", offset, "limit", limit)

	name := ctx.Param("name")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "section name is required"})
		return
	}

	section, err := p.service.GetSection(ctx, name, limit, offset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, section)
}

// parseQueryParam parses the query parameter and returns its value as an integer.
// If the parameter is missing, it returns the provided default value.
func parseQueryParam(ctx *gin.Context, param string, defaultValue int) (int, error) {
	value := ctx.Query(param)
	if value == "" {
		return defaultValue, nil
	}
	return strconv.Atoi(value)
}

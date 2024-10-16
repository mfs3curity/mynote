package services

import (
	"context"
	"time"

	"github.com/mfs3curity/mynote/api/dto"
	"github.com/mfs3curity/mynote/db"
	"github.com/mfs3curity/mynote/db/models"
	"gorm.io/gorm"
)

type SectionService struct {
	db *gorm.DB
}

func NewSectionService() *SectionService {
	return &SectionService{
		db: db.Getdb(),
	}
}

func (p *SectionService) CreateSection(ctx context.Context, req *dto.CreatSection) (string, error) {
	uId := int(ctx.Value("UserId").(float64))
	section := models.Section{
		Name: req.Name,
	}
	section.CreatedAt = time.Now().UTC()
	section.CreatedBy = uId
	if err := p.db.Create(&section).Error; err != nil {
		return "", err
	}
	return section.Name, nil
}
func (p *SectionService) GetSection(ctx context.Context, name string, limit, offset int) (*dto.ResponseSection, error) {
	var section models.Section
	if err := p.db.Model(&models.Section{}).
		Where("name = ?", name).
		First(&section).Error; err != nil {
		return nil, err
	}

	// تحميل المشاركات المرتبطة مع تطبيق limit و offset
	var posts []models.Post
	if err := p.db.Model(&section).
		Limit(limit).
		Offset(offset).
		Association("Post").Find(&posts); err != nil {
		return nil, err
	}

	var titles []string
	var createdAtPosts []time.Time
	var idposts []int
	for _, post := range posts {
		titles = append(titles, post.Title)
		createdAtPosts = append(createdAtPosts, post.CreatedAt)
		idposts = append(idposts, post.Id)
	}

	result := &dto.ResponseSection{
		Name:             section.Name,
		CreatedAtSection: section.CreatedAt,
		TitlePost:        titles,
		CreatedAtPost:    createdAtPosts,
		IdPost:           idposts,
	}

	return result, nil
}

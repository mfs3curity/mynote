package services

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/mfs3curity/mynote/api/dto"
	"github.com/mfs3curity/mynote/db"
	"github.com/mfs3curity/mynote/db/models"
	"gorm.io/gorm"
)

type PostService struct {
	db *gorm.DB
}

func NewPostService() *PostService {
	return &PostService{
		db: db.Getdb(),
	}
}

// create
func (p *PostService) CreatePost(ctx context.Context, req *dto.CreatePost) (int, error) {
	uId := int(ctx.Value("UserId").(float64))
	var section models.Section
	if err := p.db.Model(&models.Section{}).Where("name = ?", req.SectionName).First(&section).Error; err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, errors.New("section not found")
		}
		return 0, err // إرجاع خطأ إذا لم يتم العثور على القسم
	}
	post := models.Post{
		UserId:      uId,
		Title:       req.Title,
		Description: req.Description,
		SectionId:   section.Id,
	}
	post.CreatedBy = uId
	post.CreatedAt = time.Now().UTC()

	// إنشاء الـ Post
	if err := p.db.Create(&post).Error; err != nil {
		return 0, err // إرجاع الخطأ إذا حدث
	}

	// إرجاع ID الـ Post الجديد
	return post.Id, nil
}

// create image
func (p *PostService) UploadImg(ctx context.Context, path string) (string, error) {
	var img models.Image
	img.CreatedBy = int(ctx.Value("UserId").(float64))
	img.CreatedAt = time.Now().UTC()
	img.Path = path
	if err := p.db.Create(&img).Error; err != nil {
		return "", err
	}
	return img.Path, nil

}

// update
// get by id
func (p *PostService) GetByID(id int) (*dto.ResponsePost, error) {
	var post models.Post
	err := p.db.Model(&models.Post{}).
		Preload("User").
		Preload("Section").
		Where("id = ?", id).First(&post).Error
	if err != nil {
		return nil, err
	}
	result := &dto.ResponsePost{
		Username:    post.User.Username,
		Title:       post.Title,
		Description: post.Description,
		SectionName: post.Section.Name,
	}
	return result, nil
}

// delete

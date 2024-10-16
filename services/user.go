package services

import (
	"errors"

	"github.com/mfs3curity/mynote/api/dto"
	"github.com/mfs3curity/mynote/db"
	"github.com/mfs3curity/mynote/db/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	token *TokenService
	db    *gorm.DB
}

func NewUserService() *UserService {
	return &UserService{
		token: NewTokenService(),
		db:    db.Getdb(),
	}
}

func (u *UserService) Login(req *dto.Login) (*dto.TokenDetail, error) {
	var user models.User
	err := u.db.Model(&models.User{}).Where("username = ?", req.Username).Preload("UserRole", func(tx *gorm.DB) *gorm.DB {
		return tx.Preload("Role")
	}).Find(&user).Error
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	tdto := tokenDto{
		UserId:   user.Id,
		Username: user.Username,
	}
	if len(*user.UserRole) > 0 {
		for _, ur := range *user.UserRole {
			tdto.Roles = append(tdto.Roles, ur.Role.Name)
		}
	}

	token, err := u.token.GenerateToken(&tdto)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (u *UserService) Register(req *dto.Login) error {
	user := models.User{
		Username: req.Username,
		Password: req.Password,
	}
	exist, err := u.existUserName(req.Username)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("username already exist")
	}
	bp := []byte(req.Password)
	hp, err := bcrypt.GenerateFromPassword(bp, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hp)

	roleId, err := u.getDefaultRole()
	if err != nil {
		return err
	}

	tx := u.db.Begin()
	err = tx.Create(&user).Error
	if err != nil {
		return err
	}
	err = tx.Create(models.UserRole{RoleId: roleId, UserId: user.Id}).Error
	if err != nil {
		return err
	}
	tx.Commit()
	return nil

}
func (u *UserService) existUserName(user string) (bool, error) {
	var exist bool
	if err := u.db.Model(&models.User{}).Select("count(*) > 0").Where("username = ?", user).Find(&exist).Error; err != nil {
		return false, err
	}
	return exist, nil
}

func (s *UserService) getDefaultRole() (roleId int, err error) {
	if err = s.db.Model(&models.Role{}).
		Select("id").
		Where("name = ?", "user").
		First(&roleId).Error; err != nil {
		return 0, err
	}
	return roleId, nil
}

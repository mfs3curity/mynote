package migrations

import (
	"log"

	"github.com/mfs3curity/mynote/db"
	"github.com/mfs3curity/mynote/db/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func UP() {
	database := db.Getdb()
	createTables(database)
	createDefaultUserInformation(database)
	createSection(database)
}
func createTables(database *gorm.DB) {
	tables := []interface{}{}
	tables = addNewTable(database, models.User{}, tables)
	tables = addNewTable(database, models.Role{}, tables)
	tables = addNewTable(database, models.UserRole{}, tables)
	tables = addNewTable(database, models.Section{}, tables)
	tables = addNewTable(database, models.Post{}, tables)
	tables = addNewTable(database, models.Image{}, tables)

	err := database.Migrator().CreateTable(tables...)
	if err != nil {
		log.Fatalln(err)
	}
}
func addNewTable(database *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	if !database.Migrator().HasTable(model) {
		tables = append(tables, model)
	}
	return tables
}

func createRoleIfNotExists(database *gorm.DB, r *models.Role) {
	exists := 0
	database.
		Model(&models.Role{}).
		Select("1").
		Where("name = ?", r.Name).
		First(&exists)
	if exists == 0 {
		database.Create(r)
	}
}
func createDefaultUserInformation(database *gorm.DB) {

	adminRole := models.Role{Name: "admin"}
	createRoleIfNotExists(database, &adminRole)

	defaultRole := models.Role{Name: "user"}
	createRoleIfNotExists(database, &defaultRole)

	u := models.User{Username: "admin"}
	pass := "admin"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	u.Password = string(hashedPassword)

	createAdminUserIfNotExists(database, &u, adminRole.Id)

}

func createAdminUserIfNotExists(database *gorm.DB, u *models.User, roleId int) {
	exists := 0
	database.
		Model(&models.User{}).
		Select("1").
		Where("username = ?", u.Username).
		First(&exists)
	if exists == 0 {
		database.Create(u)
		ur := models.UserRole{UserId: u.Id, RoleId: roleId}
		database.Create(&ur)
	}
}

func createSection(database *gorm.DB) {
	count := 0
	database.Model(&models.Section{}).Select("count(*)").Find(&count)
	if count == 0 {
		database.Create(&models.Section{
			Name: "cve",
		})
		database.Create(&models.Section{
			Name: "vulnerabilities",
		})
	}
}

package repositories

import (
	"final-project/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	CheckUser(email string, user *models.User) error
	CheckUserByID(id int, user *models.User) (*models.User, error)
	DeleteUser(userId int) error
	UpdateUser(userId int, user *models.User) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (u *userRepository) CreateUser(user *models.User) (*models.User, error) {
	return user, u.db.Create(&user).Error
}

func (u *userRepository) CheckUser(email string, user *models.User) error {
	return u.db.Where("email=?", email).Take(&user).Error
}

func (u *userRepository) CheckUserByID(id int, user *models.User) (*models.User, error) {
	return user, u.db.Where("user_id=?", id).Take(&user).Error
}

func (u *userRepository) DeleteUser(userId int) error {
	var user models.User
	return u.db.Where("user_id = ?", userId).Delete(&user).Error
}

func (u *userRepository) UpdateUser(userId int, userUpdate *models.User) (*models.User, error) {
	var user models.User

	result := u.db.Model(&user).Clauses(clause.Returning{}).Where("user_id=?", userId).Updates(userUpdate)
	return &user, result.Error
}

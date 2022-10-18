package repositories

import (
	"final-project/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CommentRepository interface {
	CreateComment(comment *models.Comment) (*models.Comment, error)
	GetAllComments() (*[]models.Comment, error)
	GetCommentByID(commentId int) (*models.Comment, error)
	UpdateComment(commentId int, comment *models.Comment) (*models.Comment, error)
	DeleteComment(commentId int) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) CommentRepository {
	return &commentRepository{db}
}

func (c *commentRepository) CreateComment(comment *models.Comment) (*models.Comment, error) {
	return comment, c.db.Create(&comment).Error
}

func (c *commentRepository) GetAllComments() (*[]models.Comment, error) {
	var comments []models.Comment
	err := c.db.Preload("User").Preload("Photo").Find(&comments).Error
	return &comments, err

}

func (c *commentRepository) GetCommentByID(commentId int) (*models.Comment, error) {
	var comment models.Comment
	err := c.db.Preload("User").Preload("Photo").Where("comment_id=?", commentId).First(&comment).Error
	return &comment, err
}

func (c *commentRepository) UpdateComment(commentId int, updateComment *models.Comment) (*models.Comment, error) {
	var comment models.Comment

	err := c.db.Model(&comment).Clauses(clause.Returning{}).Where("comment_id=?", commentId).Updates(&updateComment).Error
	return &comment, err
}

func (c *commentRepository) DeleteComment(commentId int) error {
	var comment models.Comment
	err := c.db.Where("comment_id=?", commentId).Delete(&comment).Error
	return err
}

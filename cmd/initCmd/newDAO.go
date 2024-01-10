package initCmd

import (
	"blog/infrastructure/persistence/db"
	"gorm.io/gorm"
)

func NewDao() *gorm.DB {
	return db.GetDB()
}

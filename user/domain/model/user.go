package model

type User struct {
	ID       int64  `gorm:"primery_key;not_null;auto_increment"`
	UserName string `gorm:"uniqure_index;not_null"`
	Age      int64
}

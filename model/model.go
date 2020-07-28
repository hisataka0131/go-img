package model

import "github.com/jinzhu/gorm"

type Image struct {
	gorm.Model
	Filename    string
	ContentType string
}

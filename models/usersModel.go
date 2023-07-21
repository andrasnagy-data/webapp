package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
}

type Following struct {
	gorm.Model
	Follower uint `gorm:"foreignKey:UserID"`
	Followee uint `gorm:"foreignKey:UserID"`
}

type Follower struct {
	FirstName string
	LastName  string
	Email     string
}

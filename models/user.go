package models

import (
	"time"
)

type UserRole string

const (
	Admin UserRole = "admin"
	Teacher  UserRole = "teacher"
	Student UserRole = "student"
)

type User struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Email     string    `json:"email" gorm:"text;not null"`
	Role      UserRole  `json:"role" gorm:"text;default:attendee"`
	Password  string    `json:"-"` // Do not compute the password in json
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

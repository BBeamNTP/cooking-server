package models

import (
	"time"
)

// postAuth  รับค่าจากหน้าบ้าน
type Userdata struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Genderid  string `json:"genderID"`
	Gender    string `gorm:"-" json:"gender"`
	Titleid   string `json:"titleID"`
	Titlename string `gorm:"-" json:"titlename"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Avatar    string `json:"avatar"`
	UserId       string `json:"userId"`
	AdminId      string `json:"adminId"`
	SigninMethod string  `json:"signinMethod"`
}
type Login struct {
	Message  string `json:"message"`
	AccStat  bool   `json:"status"`
	AccToken string `json:"acctoken"`
	User     string `json:"user"`
	UserId       string `json:"userId"`
	AdminId      string `json:"adminId"`
	Avatar   string `json:"avatar"`
}

type Token struct {
	ID    string `gorm:"-" json:"id"`
	UserId       string `json:"userId"`
	Email string `json:"email"`
	Token string `json:"token"`
	SigninMethod string  `json:"signinMethod"`
}
type Status struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}
type Password struct {
	OldPass  string `json:"oldPassword"`
	NewPass1 string `json:"newPassword"`
}

// EditProfile
type Data struct {
	Userdata Userdata `json:"userData"`
	Password Password `json:"changePassword"`
}
type ResetPassword struct {
	Email       string `gorm:"-" json:"email"`
	NewPassword string `json:"newPassword"`
	OTP         string `json:"OTP"`
}
type Otps struct {
	ID        int
	Email     string    `gorm:"email" json:"email"`
	Otp       string    `gorm:"otp" json:"OTP"`
	StartTime time.Time `gorm:"start_time" json:"startTime"`
	EndTime   time.Time `gorm:"end_time" json:"endTime"`
}

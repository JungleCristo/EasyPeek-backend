package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Username  string         `json:"username" gorm:"uniqueIndex;not null"`
	Email     string         `json:"email" gorm:"uniqueIndex;not null"`
	Password  string         `json:"password" gorm:"not null"`
	Avatar    string         `json:"avatar"`
	Phone     string         `json:"phone"`
	Location  string         `json:"location"`
	Bio       string         `json:"bio"`
	Interests string         `json:"interests"`
	Role      string         `json:"role" gorm:"default:user"`
	Status    string         `json:"status" gorm:"default:active"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Avatar    string    `json:"avatar"`
	Phone     string    `json:"phone"`
	Location  string    `json:"location"`
	Bio       string    `json:"bio"`
	Interests string    `json:"interests"`
	Role      string    `json:"role"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Login
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Register
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

// UpdateUser
type UpdateUserRequest struct {
	Username  string `json:"username" binding:"omitempty,min=3,max=20"`
	Avatar    string `json:"avatar"`
	Phone     string `json:"phone"`
	Location  string `json:"location"`
	Bio       string `json:"bio"`
	Interests string `json:"interests"`
}

// ChangePassword
type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}

// UpdateUserStatus
type UpdateUserStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=active inactive suspended deleted"`
}

// DeleteAccount
type DeleteAccountRequest struct {
	Password string `json:"password" binding:"required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}
	return nil
}

// CheckPassword check if the password is correct
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) ToResponse() *UserResponse {
	return &UserResponse{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Avatar:    u.Avatar,
		Phone:     u.Phone,
		Location:  u.Location,
		Bio:       u.Bio,
		Interests: u.Interests,
		Role:      u.Role,
		Status:    u.Status,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (User) TableName() string {
	return "users"
}

package repository

import (
	"errors"

	"github.com/ardipermana59/mygram/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

// CreateUser digunakan untuk membuat pengguna baru
func (r *UserRepository) CreateUser(user *model.User) error {
	return r.DB.Create(user).Error
}

// IsEmailExists digunakan untuk memeriksa apakah email sudah ada dalam database
func (r *UserRepository) IsEmailExists(email string) bool {
	var user model.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return false
	}
	return true
}

// IsUsernameExists digunakan untuk memeriksa apakah username sudah ada dalam database
func (r *UserRepository) IsUsernameExists(username string) bool {
	var user model.User
	if err := r.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return false
	}
	return true
}

// FindByUsername digunakan untuk mencari pengguna berdasarkan nama pengguna (username)
func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	if err := r.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Return nil jika pengguna tidak ditemukan
		}
		return nil, err // Kembalikan error jika terjadi kesalahan lain
	}
	return &user, nil // Kembalikan pengguna yang ditemukan
}

func (r *UserRepository) DeleteUser(userID uint) error {
	// Mencari pengguna berdasarkan ID
	var user model.User
	if err := r.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("User not found")
		}
		return err
	}

	// Menghapus pengguna dari database
	if err := r.DB.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

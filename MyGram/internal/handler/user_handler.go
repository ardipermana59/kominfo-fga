package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ardipermana59/mygram/internal/infrastructure"
	"github.com/ardipermana59/mygram/internal/model"
	"github.com/ardipermana59/mygram/internal/repository"
	"github.com/ardipermana59/mygram/pkg"
	"github.com/ardipermana59/mygram/pkg/helper"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Username string `json:"username" validate:"required,unique"`
	Email    string `json:"email" validate:"required,email,unique"`
	Password string `json:"password" validate:"required,min=6"`
	Age      int    `json:"age" validate:"required,gte=8"`
}

func Register(c *gin.Context) {
	var req RegisterRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "Bad Request"})
		return
	}
	// Validasi disini

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: "Failed to hash password"})
		return
	}

	// Buat instance dari model User
	user := model.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  string(hashedPassword),
		Age:       req.Age,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db, err := infrastructure.Database()
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: "Internal server error"})
		return
	}
	// Membuat instance dari UserRepository
	userRepository := repository.UserRepository{DB: db} // Sesuaikan dengan instance DB Anda

	// Memeriksa apakah email sudah ada
	if userRepository.IsEmailExists(req.Email) {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "Email already exists"})
		return
	}

	// Memeriksa apakah username sudah ada
	if userRepository.IsUsernameExists(req.Username) {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "Username already exists"})
		return
	}

	// Simpan pengguna ke database
	if err := userRepository.CreateUser(&user); err != nil {

		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: "Internal server error"})
		return
	}

	// Jika berhasil, kembalikan respons berhasil
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func Login(c *gin.Context) {
	var req LoginRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "Bad Request"})
		return
	}

	// Validasi disini

	db, err := infrastructure.Database()
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: "Internal server error"})
		return
	}
	// Membuat instance dari UserRepository
	userRepository := repository.UserRepository{DB: db}
	// Find user by username
	user, err := userRepository.FindByUsername(req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: "Internal server error"})
		return
	}
	if user == nil {
		c.JSON(http.StatusUnauthorized, pkg.ErrorResponse{Message: "Invalid username or password"})
		return
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, pkg.ErrorResponse{Message: "Invalid username or password"})
		return
	}

	// Generate JWT token
	token, err := helper.GenerateJWTToken(user.ID)
	fmt.Println(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: "Failed to generate token"})
		return
	}

	// Respond with JWT token
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func UpdateUser(c *gin.Context) {
	// Your update user handler logic here
}

func DeleteUser(c *gin.Context) {
	// Mengambil ID pengguna dari token JWT
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get user ID from token"})
		return
	}

	// Mengonversi ID pengguna menjadi tipe data yang sesuai
	userIDStr, ok := userID.(string)
	fmt.Println(exists)
	fmt.Println(userID)
	fmt.Println(userIDStr)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to convert user ID"})
		return
	}
	id, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to convert user ID"})
		return
	}

	db, err := infrastructure.Database()
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: "Internal server error"})
		return
	}
	// Membuat instance dari UserRepository
	userRepository := repository.UserRepository{DB: db}

	// Menghapus pengguna dari database berdasarkan ID
	err = userRepository.DeleteUser(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete user"})
		return
	}

	// Jika berhasil, kembalikan respons berhasil
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

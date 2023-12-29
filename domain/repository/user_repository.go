package repository

import (
	"kulina-interview-test/domain/model"
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type dbUser struct {
	Conn *gorm.DB
}

// Register implements UserRepository.
func (db *dbUser) Register(ctx echo.Context, user model.User) error {
	// Create a new User instance to hold request data
	newUser := new(model.User)
	// Connect to the database

	// Bind request data to the User instance
	if err := ctx.Bind(newUser); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to bind request data",
		})
	}

	// Hash the password securely
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to hash password",
		})
	}

	// Create a User model with hashed password
	user = model.User{
		Name:        newUser.Name,
		Email:       newUser.Email,
		Password:    string(hashedPassword),
		FullAddress: newUser.FullAddress,
		Latitude:    newUser.Latitude,
		Longitude:   newUser.Longitude,
	}

	// Create the user in the database
	if err := db.Conn.Create(&user).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to create user",
		})
	}

	// Return a success response with the created user data
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "User registered successfully",
		"data":    user,
	})
}

type UserRepository interface {
	Register(ctx echo.Context, user model.User) error
}

func NewUserRepository(Conn *gorm.DB) UserRepository {
	return &dbUser{Conn: Conn}
}

package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "backend-golang/swagger-docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gin API Example
// @version 1.0
// @description This is a sample server for a product and user API.
// @host localhost:3000
// @BasePath /gin-gonic/api/v1

type User struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	Nome            string    `json:"nome"`
	Email           string    `json:"email"`
	Senha           string    `json:"senha"`
	DataCriacao     time.Time `json:"data_criacao"`
	DataModificacao time.Time `json:"data_modificacao"`
}

type DefaultError struct {
	message string
}

var db *gorm.DB
var err error

func main() {
	// Initialize database connection
	db, err = gorm.Open(sqlite.Open("data/database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// Migrate the schema
	// db.AutoMigrate(&User{})

	r := gin.Default()

	router := r.Group("/gin-gonic")

	// Create log file
	f, err := os.OpenFile("/logs/gin_app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	gin.DefaultWriter = f

	// Middleware to log requests
	r.Use(gin.LoggerWithWriter(f))

	api := router.Group("/api/v1")
	{
		api.GET("/users", getUsers)
		api.GET("/users/:id", getUserByID)
		api.POST("/users", createUser)
		api.PUT("/users/:id", updateUser)
		api.DELETE("/users/:id", deleteUser)
	}

	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}

// @Summary Get all users
// @Produce json
// @Success 200 {array} User
// @Router /users [get]
func getUsers(c *gin.Context) {
	var users []User
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}

// @Summary Get a user by ID
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Failure 404 {object} DefaultError
// @Router /users/{id} [get]
func getUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user User
	if result := db.First(&user, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// @Summary Create a new user
// @Produce json
// @Param user body User true "User to create"
// @Success 201 {object} User
// @Failure 400 {object} DefaultError
// @Router /users [post]
func createUser(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUser.DataCriacao = time.Now()
	newUser.DataModificacao = time.Now()
	db.Create(&newUser)
	c.JSON(http.StatusCreated, newUser)
}

// @Summary Update a user by ID
// @Produce json
// @Param id path int true "User ID"
// @Param user body User true "User to update"
// @Success 200 {object} User
// @Failure 400 {object} DefaultError
// @Failure 404 {object} DefaultError
// @Router /users/{id} [put]
func updateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedUser User
	if err := c.BindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user User
	if result := db.First(&user, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}
	updatedUser.ID = user.ID
	updatedUser.DataModificacao = time.Now()
	db.Save(&updatedUser)
	c.JSON(http.StatusOK, updatedUser)
}

// @Summary Delete a user by ID
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} DefaultError
// @Failure 404 {object} DefaultError
// @Router /users/{id} [delete]
func deleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user User
	if result := db.First(&user, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}
	db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

package main

import (
	"fmt"
	"log"

	"github.com/wpcodevo/golang-gorm-postgres/initializers"
	"github.com/wpcodevo/golang-gorm-postgres/models"
)

func init() {
	cfg, err := initializers.LoadConfig("../")
	if err != nil {
		log.Fatal("Coulfn't load environment variables", err)
	}

	initializers.ConnectDB(&cfg)
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{}, &models.CreatePostRequest{}, &models.UpdatePost{}, &models.User{})
	fmt.Println("Migration complete!")
}

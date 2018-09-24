package main

import (
	"fastwork/go-gin-performance-test/app/config"
	"fastwork/go-gin-performance-test/app/models"
	"fastwork/go-gin-performance-test/app/repos"
	"fastwork/go-gin-performance-test/app/servers"
	"fastwork/go-gin-performance-test/app/services"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	port = ":9000"
	conf models.Config
)

func main() {

	// Initial database & migrate
	conf = *config.GetConfig()
	db := initDb(&conf)
	db = configDatabase(db)
	models := []interface{}{
		models.Product{},
	}
	autoMigrate(db, models)

	// Initial repository
	productRepo := repos.NewProductRepo(db)

	// Initial service
	productService := services.NewProductService(*productRepo)

	// Initial server
	productServer := servers.NewProductServer(productService)

	r := gin.Default()

	// Start health check http server
	r.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	r.GET("api/product/:product_id", productServer.GetByID)

	r.Run(port)
}

func initDb(config *models.Config) *gorm.DB {
	dbHost := config.Database.Host
	dbName := config.Database.Name
	dbUsername := config.Database.User
	dbPassword := config.Database.Password
	appName := config.Database.AppName

	dbInfo := fmt.Sprintf("host=%s dbname=%s user=%s password=%s application_name=%s sslmode=disable", dbHost, dbName, dbUsername, dbPassword, appName)

	db, err := gorm.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}

	return db
}

func autoMigrate(db *gorm.DB, models []interface{}) error {
	for _, m := range models {
		if err := db.AutoMigrate(m).Error; err != nil {
			return err
		}
	}
	return nil
}

func configDatabase(db *gorm.DB) *gorm.DB {
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Hour)
	db = db.LogMode(true)

	return db
}

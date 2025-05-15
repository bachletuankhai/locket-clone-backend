package rest

import (
	"locket-clone/backend/pkg/repository"
	"locket-clone/backend/pkg/service/adding"
	"locket-clone/backend/pkg/service/listing"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func Init() {
	db := InitDB()
	storage := repository.BlobStorage{}

	// Initialize the repositories
	locketRepo := repository.LocketRepo{
		Db: db,
	}
	userRepo := repository.UserRepo{
		Db: db,
	}

	// Initialize the services
	locketAddingService := adding.NewLocketService(&locketRepo, &storage)
	locketListingService := listing.NewLocketService(&locketRepo)
	userListingService := listing.NewUserService(&userRepo)

	// Initialize the LocketController
	locketController := &LocketController{
		locketAddingService:  locketAddingService,  // Replace nil with actual LocketRepo
		locketListingService: locketListingService, // Replace nil with actual LocketRepo
		userListingService:   userListingService,   // Replace nil with actual UserRepo
	}

	router := gin.Default()
	locketGroup := router.Group("/locket")
	locketController.RegisterLocketHandler(locketGroup)

	router.Run()
}

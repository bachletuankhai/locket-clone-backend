package rest

import (
	"locket-clone/backend/pkg/repository"
	"locket-clone/backend/pkg/service/adding"
	"locket-clone/backend/pkg/service/auth"
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
	tokenRepo := repository.NewMemTokenRepo()

	// Initialize the services
	locketAddingService := adding.NewLocketService(&locketRepo, &storage)
	locketListingService := listing.NewLocketService(&locketRepo)
	userListingService := listing.NewUserService(&userRepo)
	userAddingService := adding.NewUserService(&userRepo)
	authService := auth.NewAuthService(tokenRepo, &userRepo)

	// Initialize the LocketController
	locketController := LocketController{
		locketAddingService:  locketAddingService,
		locketListingService: locketListingService,
		userListingService:   userListingService,
	}

	userController := UserController{
		UserAddingService: userAddingService,
		AuthService:       authService,
	}

	router := gin.Default()
	locketGroup := router.Group("/locket")

	authMiddleware := NewAuthMiddleware(authService)

	locketGroup.Use(authMiddleware)
	{
		locketGroup.GET("/feed", locketController.GetFeed)
		locketGroup.GET("/user/:username", locketController.GetUserLockets)
		locketGroup.POST("/", locketController.AddLocket)
	}

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", userController.Login)
		authGroup.POST("/register", userController.RegisterUser)
	}

	router.Run()
}

// Soko Bora Mngt System API
//
// This is a simple API for the dashboard. You can find out more about the API at https://github.com/owezzy/soko-bora-mngt-system.
//
//	Schemes: http
//  Host: localhost:4500
//	BasePath: /
//	Version: 1.0.0
//	Contact: Owen Adirah <owenadira@gmail.com> https://owezzy.tech
//  SecurityDefinitions:
//  api_key:
//    type: apiKey
//    name: Authorization
//    in: header
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
// swagger:meta

package main

import (
	"context"
	"encoding/gob"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	redisStore "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	handlers "github.com/owezzy/soko-bora-mngt-system/cmd/api/handlers"
	"github.com/owezzy/soko-bora-mngt-system/cmd/api/handlers/auth0"
	"github.com/owezzy/soko-bora-mngt-system/cmd/api/handlers/home"
	"github.com/owezzy/soko-bora-mngt-system/cmd/api/handlers/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var authHandler *handlers.AuthHandler
var recipesHandler *handlers.RecipesHandler

func init() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
	collectionRecipes := client.Database(os.Getenv("MONGO_DATABASE")).Collection("recipes")

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	status := redisClient.Ping()
	log.Println(status)

	recipesHandler = handlers.NewRecipesHandler(ctx, collectionRecipes, redisClient)

	collectionUsers := client.Database(os.Getenv("MONGO_DATABASE")).Collection("users")
	log.Println(collectionUsers)
	authHandler = handlers.NewAuthHandler(ctx, collectionUsers)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}

	authFunc, err := auth0.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}
	router := gin.Default()
	router.Use(cors.Default())

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	authStore := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", authStore))
	router.GET("/", home.HomePageHandler)

	router.GET("/login", authFunc.Auth0LoginHandler)
	router.GET("/callback", authFunc.Auth0CallbackHandler)
	router.GET("/user", middleware.IsAuth0Authenticated, auth0.UserProfileHandler)
	router.GET("/logout", auth0.Auth0LogoutHandler)

	store, _ := redisStore.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	router.Use(sessions.Sessions("recipes_api", store))

	router.GET("/recipes", recipesHandler.ListRecipesHandler)

	router.POST("/signin", authHandler.SignInHandler)
	router.POST("/refresh", authHandler.RefreshHandler)
	router.POST("/signout", authHandler.SignOutHandler)

	authorized := router.Group("/")
	authorized.Use(authHandler.AuthMiddleware())
	{
		authorized.POST("/recipes", recipesHandler.NewRecipeHandler)
		authorized.PUT("/recipes/:id", recipesHandler.UpdateRecipeHandler)
		authorized.DELETE("/recipes/:id", recipesHandler.DeleteRecipeHandler)
		authorized.GET("/recipes/:id", recipesHandler.GetOneRecipeHandler)
	}

	router.Run()
}

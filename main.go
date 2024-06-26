package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/lucas271/GoSimpleApi/internal/database"

	_ "github.com/go-sql-driver/mysql"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()

	portString := os.Getenv("PORT")
	DBURL := os.Getenv("DB_URL")

	if DBURL == "" {
		log.Fatal("PORT IS NOT FOUND IN THE ENVIROMENT")
	}

	conn, err := sql.Open("mysql", DBURL)
	if err != nil {
		log.Fatal("can't connect to database", err)
	}

	router := chi.NewRouter()

	if portString == "" {
		log.Fatal("PORT is not found in the enviroment")
	}

	queries := database.New(conn)

	apiCfg := apiConfig{
		DB: queries,
	}

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handleReadiness)
	v1Router.Get("/err", handleErr)
	v1Router.Post("/users", apiCfg.handleCreateUser)
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server running on port %s", portString)

	err2 := srv.ListenAndServe()

	if err2 != nil {
		log.Fatal(err)
	}

	fmt.Println("Port: ", portString)
}

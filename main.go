package main

import (
	"errors"
	"fmt"
	"net/http"

	"api_service/api"
	"api_service/database"
	"api_service/database/seeds"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	//dsn := fmt.Sprintf(
	//	"host=host.docker.internal user=%s password=%s dbname=%s port=5432 sslmode=disable",
	//	os.Getenv("POSTGRES_USER"),
	//	os.Getenv("POSTGRES_PASSWORD"),
	//	os.Getenv("POSTGRES_PASSWORD"),
	//)
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//db, err := gorm.Open(postgres.Open("host=0.0.0.0 user=api_user password=supersecret dbname=api port=5432 sslmode=disable TimeZone=Asia/Shanghai"), &gorm.Config{})
	dbURL := "postgres://api_user:supersecret@host.docker.internal:5432/api"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&database.Partner{}, &database.Service{})
	if err != nil {
		panic(err)
	}

	if err := db.First(&database.Partner{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		for _, seed := range seeds.All() {
			if err := seed.Run(db); err != nil {
				panic(fmt.Sprintf("Running seeds '%s', failed with error: %s", seed.Name, err))
			}
		}
	}

	//api.Start(db)

	//mux := http.NewServeMux()
	//api.New(mux, db)

	//http.ListenAndServe(":8080", mux)
	handler := api.New(db)
	router := api.CreateRouter(handler)
	http.ListenAndServe(":8080", router)
}

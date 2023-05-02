package main

import (
	"database/sql"
	"fmt"
	"knt/internal/kntrouter"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/google/logger"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
)

func main() {
	//Loading viper
	//Viper is responsible for handling the config file
	viper.SetConfigName("kntconfig")

	viper.AddConfigPath("./config/")

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	//Insure that log directory exists
	os.Mkdir(viper.GetString("logPath"), os.ModePerm)
	//Open log directory
	lf, err := os.OpenFile(viper.GetString("logPath")+viper.GetString("logName"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}

	//Load the logger
	logger.Init("KnT Backend", true, true, lf)
	logger.SetFlags(log.Lshortfile)
	logger.SetFlags(log.LstdFlags)

	//Open the database
	db, err := sql.Open("sqlite3", viper.GetString("database"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Start the mux router, this router simplified http calls and reduces boilerplate code
	r := chi.NewRouter()

	kntrouter.AssignRoutes(r, db)

	logger.Info("Attempting to start listening on port: ", viper.GetString("port"))
	logger.Fatal(http.ListenAndServe(":"+viper.GetString("port"), r))

}

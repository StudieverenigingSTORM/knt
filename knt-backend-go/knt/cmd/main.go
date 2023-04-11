package main

import (
	"database/sql"
	"fmt"
	"kntrouter"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
)

func main() {
	//Loading viper
	//Viper is responsible for handling the config file
	//TODO: make a start script that would force viper to work correctly
	viper.SetConfigName("kntconfig")

	viper.AddConfigPath("knt/config/")

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	fmt.Println("Config file successfully loaded")

	//Open the database
	db, err := sql.Open("sqlite3", viper.GetString("database"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Start the mux router, this router simplified http calls and reduces boilerplate code
	r := chi.NewRouter()
	kntrouter.AssignMiddlewares(r)
	kntrouter.AssignRoutes(r, db)

	go sanityCheck()

	fmt.Println("Attempting to start listening on port: " + viper.GetString("port"))
	log.Fatalln(http.ListenAndServe(viper.GetString("port"), r))

}

//Sanity checks the server if it has been successfully started.
//This is moslty done to avoid weird errors where server does not start for one or more reasons.

func sanityCheck() {
	for {
		var resp *http.Response
		resp, _ = http.Get("http://127.0.0.1" + viper.GetString("port") + "/ping")

		if resp.StatusCode == http.StatusOK {
			fmt.Println("Server successfully started on http://127.0.0.1" + viper.GetString("port"))
			break
		}

		resp.Body.Close()

		time.Sleep(1)
	}
}

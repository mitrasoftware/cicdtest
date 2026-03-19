package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/mitrasoftware/cicdtest/models"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("NO .env file found")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/getuser", getUser)
	// http.HandleFunc("/createUser", createUser)
	fmt.Println("Server started on port ", port)
	http.ListenAndServe(":"+port, nil)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {

		http.Error(w, "Only get request allowed", http.StatusMethodNotAllowed)
		return
	}

	user := models.User{
		Name: "Meghana Mittal",
		Age:  44,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

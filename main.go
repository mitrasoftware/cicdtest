package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mitrasoftware/cicdtest/models"
)

func main() {

	http.HandleFunc("/getuser", getUser)
	// http.HandleFunc("/createUser", createUser)
	fmt.Println("Server started on port ", "8080")
	http.ListenAndServe(":8080", nil)
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

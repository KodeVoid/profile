package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"profile_api/internals/models"

	. "profile_api/internals/app"
	"time"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	user := models.User{
		Email: "kendrick@example.com",
		Name:  "Kendrick Lamar",
		Stack: "Golang/Cloud",
	}

	fact, err := GetFact()
	if err != nil {
		log.Printf("Error fetching fact: %v", err)
		fact = "Could not fetch a cat fact at this time."
	}

	myProfile := models.Profile{
		Status:    "success",
		User:      user,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Fact:      fact,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if er := json.NewEncoder(w).Encode(myProfile); er != nil {
		http.Error(w, er.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Wrote %v \n", myProfile)

}

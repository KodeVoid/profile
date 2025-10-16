package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"profile_api/internals/app"
	"profile_api/internals/models"
	"time"
)

// ProfileHandler for the /me endpoint
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	user := models.User{
		Email: "olorik362@gmail.com",
		Name:  "Kendrick Rukevwe Olori",
		Stack: "Golang, Docker, Rust, Python Postgres, React",
	}

	fact, err := app.GetFact(app.FACTAPIURL)
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

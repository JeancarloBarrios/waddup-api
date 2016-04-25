package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	// Third Party Packages
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Logger now works Just fine")
}

func UserEnd(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "User Page")
}

func EventEnd(w http.ResponseWriter, r *http.Request) {
	events := Events{
		Event{
			Id:          1,
			Name:        "Event 1",
			Description: "This is an event",
			Date:        "HOY",
			Duration:    5.5,
		},
		Event{
			Id:          2,
			Name:        "Event 2",
			Description: "This is an event",
			Date:        "HOY",
			Duration:    1.5,
		},
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(events); err != nil {
		panic(err)
	}

}

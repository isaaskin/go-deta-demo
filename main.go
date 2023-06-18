package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type SpaceScheduleEvent struct {
	Id      string `json:"id"`
	Trigger string `json:"trigger"`
}

type SpaceSchedule struct {
	Event SpaceScheduleEvent `json:"event"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/__space/v0/actions", actionHandler)

	log.Println("Listening on 0.0.0.0" + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func actionHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received schedule")
	var s SpaceSchedule

	err := json.NewDecoder(r.Body).Decode(&s)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("S: %+v", s)
	log.Println(s.Event.Id)
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Hello, world from Space!")
	}
}

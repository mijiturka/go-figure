package main

import (
  "encoding/json"
  "fmt"
  // "log"
  "net/http"
  "strings"

  "github.com/gorilla/mux"
)

var Plants []Plant

func all(writer http.ResponseWriter, request *http.Request) {
  if request.Method != "GET" {
    writer.WriteHeader(http.StatusMethodNotAllowed)
  }

  fmt.Println("You've hit it")

  json.NewEncoder(writer).Encode(Plants)
}

func genus(writer http.ResponseWriter, request *http.Request) {
  if request.Method != "GET" {
    writer.WriteHeader(http.StatusMethodNotAllowed)
    return
  }

  genus := mux.Vars(request)["genus"]
  fmt.Println("You've hit it with", genus)

  for _, plant := range Plants {
    if strings.ToLower(plant.Genus) == strings.ToLower(genus) {
      json.NewEncoder(writer).Encode(plant)
    }
  }

  writer.WriteHeader(http.StatusNotFound)
}

func species(writer http.ResponseWriter, request *http.Request) {
  if request.Method != "GET" {
    writer.WriteHeader(http.StatusMethodNotAllowed)
  }

  genus := mux.Vars(request)["genus"]
  species := mux.Vars(request)["species"]
  fmt.Println("You've hit it with", genus, species)

  for _, plant := range Plants {
    if strings.ToLower(plant.Genus) == strings.ToLower(genus) {
      if strings.ToLower(plant.Species) == strings.ToLower(species) {
        json.NewEncoder(writer).Encode(plant)
      }
    }
  }

  writer.WriteHeader(http.StatusNotFound)
}

func handleRequests() {
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", all)
  router.HandleFunc("/{genus}", genus)
  router.HandleFunc("/{genus}/{species}", species)
  http.ListenAndServe(":8000", router)
}

func main() {
  fmt.Println("Serving")

  Plants = []Plant {
    Plant{
      Genus:    "Coffea",
      Species:  "arabica",
      Origin:   "Ethiopia",
      Wiki:     "https://en.wikipedia.org/wiki/Coffea_arabica",
    },
    Plant{
      Genus:    "Coffea",
      Species:  "canephora",
      Origin:   "Ethiopia",
      Wiki:     "https://en.wikipedia.org/wiki/Coffea_canephora",
    },
    Plant{
      Genus:    "Euphorbia",
      Species:  "pulcherrima",
      Origin:   "Mexico",
      Wiki:     "https://en.wikipedia.org/wiki/Poinsettia",
    },
  }

  handleRequests()
}

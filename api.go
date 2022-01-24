package main

import "encoding/json"
import "fmt"
// import "log"
import "net/http"

type Plant struct {
  Genus     string `json:"genus"`
  Species   string `json:"species"`
  Cultivar  string `json:"cultivar"`
  Origin    string `json:"origin"`
  Wiki      string `json:"wiki"`
}

var Plants []Plant

func endpoint(writer http.ResponseWriter, request *http.Request) {
  fmt.Println("You've hit it")

  json.NewEncoder(writer).Encode(Plants)
}

func handleRequests() {
  http.HandleFunc("/", endpoint)
  http.ListenAndServe(":8000", nil)
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
      Genus:    "Euphorbia",
      Species:  "pulcherrima",
      Origin:   "Mexico",
      Wiki:     "https://en.wikipedia.org/wiki/Poinsettia",
    },
  }

  handleRequests()
}

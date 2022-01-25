package main

import (
  "bytes"
  "encoding/json"
  "net/http"
  "net/http/httptest"
  "testing"

  "github.com/gorilla/mux"
)

func Test_all(t *testing.T) {
  router := mux.NewRouter()
  router.HandleFunc("/", all)

  ts := httptest.NewServer(router)
  defer ts.Close()

  // OK
  res, err := http.Get(ts.URL + "/")
  if err != nil {
    t.Errorf("Expected nil, received %s", err.Error())
  }
  if res.StatusCode != http.StatusOK {
    t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
  }

  // 405
  empty, _ := json.Marshal("")
  res, err = http.Post(ts.URL + "/", "application/json", bytes.NewBuffer(empty))
  if err != nil {
    t.Errorf("Expected nil, received %s", err.Error())
  }
  if res.StatusCode != http.StatusMethodNotAllowed {
    t.Errorf("Expected %d, received %d", http.StatusMethodNotAllowed, res.StatusCode)
  }
}

func Test_genus(t *testing.T) {
  router := mux.NewRouter()
  router.HandleFunc("/{genus}", genus)

  ts := httptest.NewServer(router)
  defer ts.Close()

  Plants = []Plant {
    Plant{
      Genus:    "Dummy",
      Species:  "dummeus",
      Origin:   "Mayonesia",
      Wiki:     "https://en.wikipedia.org/wiki/Dummy_dummeus",
    },
  }

  // OK
  res, err := http.Get(ts.URL + "/Dummy")
  if err != nil {
    t.Errorf("Expected nil, received %s", err.Error())
  }
  if res.StatusCode != http.StatusOK {
    t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
  }

  res, err = http.Get(ts.URL + "/dummy")
  if err != nil {
    t.Errorf("Expected nil, received %s", err.Error())
  }
  if res.StatusCode != http.StatusOK {
    t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
  }

  // 404
  res, err = http.Get(ts.URL + "/nonexistant")
  if err != nil {
    t.Errorf("Expected nil, received %s", err.Error())
  }
  if res.StatusCode != http.StatusNotFound {
    t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
  }

  // 405
  empty, _ := json.Marshal("")
  res, err = http.Post(ts.URL + "/dummy", "application/json", bytes.NewBuffer(empty))
  if err != nil {
    t.Errorf("Expected nil, received %s", err.Error())
  }
  if res.StatusCode != http.StatusMethodNotAllowed {
    t.Errorf("Expected %d, received %d", http.StatusMethodNotAllowed, res.StatusCode)
  }
}

func Test_species(t *testing.T) {
  router := mux.NewRouter()
  router.HandleFunc("/{genus}/{species}", species)

  ts := httptest.NewServer(router)
  defer ts.Close()

  Plants = []Plant {
    Plant{
      Genus:    "Dummy",
      Species:  "dummeus",
      Origin:   "Mayonesia",
      Wiki:     "https://en.wikipedia.org/wiki/Dummy_dummeus",
    },
  }

  // OK
  res, err := http.Get(ts.URL + "/Dummy/dummeus")
  if err != nil {
    t.Errorf("Expected nil, received %s", err.Error())
  }
  if res.StatusCode != http.StatusOK {
    t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
  }

  res, err = http.Get(ts.URL + "/dummy/dummeus")
  if err != nil {
    t.Errorf("Expected nil, received %s", err.Error())
  }
  if res.StatusCode != http.StatusOK {
    t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
  }

  // 404
  res, err = http.Get(ts.URL + "/non/existant")
  if err != nil {
    t.Errorf("Expected nil, received %s", err.Error())
  }
  if res.StatusCode != http.StatusNotFound {
    t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
  }

  res, err = http.Get(ts.URL + "/dummy/nonexistant")
  if err != nil {
    t.Errorf("Expected nil, received %s", err.Error())
  }
  if res.StatusCode != http.StatusNotFound {
    t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
  }

  // 405
  empty, _ := json.Marshal("")
  res, err = http.Post(ts.URL + "/dummy/dummeus", "application/json", bytes.NewBuffer(empty))
  if err != nil {
    t.Errorf("Expected nil, received %s", err.Error())
  }
  if res.StatusCode != http.StatusMethodNotAllowed {
    t.Errorf("Expected %d, received %d", http.StatusMethodNotAllowed, res.StatusCode)
  }
}

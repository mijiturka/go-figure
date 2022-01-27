package main

import (
  "fmt"
)

type Plant struct {
  Genus     string `json:"genus"`
  Species   string `json:"species"`
  Variety   string `json:"variety"`
  Cultivar  string `json:"cultivar"`
  Origin    string `json:"origin"`
  Wiki      string `json:"wiki"`
}

func (p *Plant) String() string {
  scientific_name := fmt.Sprintf("%[1]s %[2]s", p.Genus, p.Species)
  switch {
  case p.Variety != "" && p.Cultivar != "":
    scientific_name += fmt.Sprintf(" var. %[1]s '%[2]s'", p.Variety, p.Cultivar)
  case p.Variety != "":
    scientific_name += fmt.Sprintf(" var. %s", p.Variety)
  case p.Cultivar != "":
    scientific_name += fmt.Sprintf(" '%s'", p.Cultivar)
  }
  return scientific_name
}

func WhoCares() {
  tasty_coffee := &Plant{
    Genus:    "Coffea",
    Species:  "arabica",
    Origin:   "Ethiopia",
  }

  still_coffee := &Plant{
    Genus:    "Coffea",
    Species:  "canephora",
    Origin:   "Ethiopia",
  }

  coffee_for_nerds := &Plant{
    Genus:    "Coffea",
    Species:  "arabica",
    Origin:   "Ethiopia",
    Variety:  "Java",
    Wiki:     "https://en.wikipedia.org/wiki/Coffee_production_in_Indonesia#Java",
  }

  cultivated_coffee := &Plant{
    Genus:    "Coffea",
    Species:  "arabica",
    Origin:   "Timor Leste",
    Variety:  "Catimor",
    Cultivar: "Lempira",
    Wiki:     "https://varieties.worldcoffeeresearch.org/varieties/lempira",
  }

  fmt.Printf("I prefer %v over %v\n", tasty_coffee, still_coffee);
  fmt.Printf("I haven't tried %v but %v is good\n", cultivated_coffee, coffee_for_nerds)
}

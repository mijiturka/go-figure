package main

import (
  "fmt"
  "testing"
)

func Test_stringer(t *testing.T) {
  cases := []struct{
    description string
    plant       *Plant
    expected    string
  }{
    {
      "basic",
      &Plant{Genus: "Coffea", Species: "arabica",},
      "Coffea arabica",
    },{
      "with_variety",
      &Plant{Genus: "Coffea", Species: "arabica", Variety: "Java"},
      "Coffea arabica var. Java",
    },{
      "with_cultivar",
      &Plant{Genus: "Coffea", Species: "arabica", Cultivar: "K7"},
      "Coffea arabica 'K7'",
    },{
      "with_variety_and_cultivar",
      &Plant{Genus: "Coffea", Species: "arabica", Variety: "Catimor", Cultivar: "Lempira"},
      "Coffea arabica var. Catimor 'Lempira'",
    },
  }

  for _, c := range cases {
    t.Run(fmt.Sprintf("%s", c.description), func(t *testing.T) {
      if fmt.Sprintf("%v", c.plant) != c.expected {
        t.Errorf("Expected %v, received %v", c.expected, fmt.Sprintf("%v", c.plant))
      }
      if fmt.Sprintf("%s", c.plant) != c.expected {
        t.Errorf("Expected %s, received %s", c.expected, fmt.Sprintf("%s", c.plant))
      }
    })
  }
}

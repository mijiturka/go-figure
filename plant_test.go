package main

import (
  "fmt"
  "testing"
)

func Test_stringer_name_basic(t *testing.T) {
  plant := &Plant {
    Genus:  "Some_Genus",
    Species:  "some_species",
  }
  expected := "Some_Genus some_species"

  if fmt.Sprintf("%v", plant) != expected {
    t.Errorf("Expected %v, received %v", expected, fmt.Sprintf("%v", plant))
  }
  if fmt.Sprintf("%s", plant) != expected {
    t.Errorf("Expected %s, received %s", expected, fmt.Sprintf("%s", plant))
  }
}

func Test_stringer_name_with_variety(t *testing.T) {
  plant := &Plant {
    Genus:    "Some_Genus",
    Species:  "some_species",
    Variety:  "Some_variety",
  }
  expected := "Some_Genus some_species var. Some_variety"

  if fmt.Sprintf("%v", plant) != expected {
    t.Errorf("Expected %v, received %v", expected, fmt.Sprintf("%v", plant))
  }
  if fmt.Sprintf("%s", plant) != expected {
    t.Errorf("Expected %s, received %s", expected, fmt.Sprintf("%s", plant))
  }
}

func Test_stringer_name_with_cultivar(t *testing.T) {
  plant := &Plant {
    Genus:    "Some_Genus",
    Species:  "some_species",
    Cultivar:  "Some_cultivar",
  }
  expected := "Some_Genus some_species 'Some_cultivar'"

  if fmt.Sprintf("%v", plant) != expected {
    t.Errorf("Expected %v, received %v", expected, fmt.Sprintf("%v", plant))
  }
  if fmt.Sprintf("%s", plant) != expected {
    t.Errorf("Expected %s, received %s", expected, fmt.Sprintf("%s", plant))
  }
}

func Test_stringer_name_with_variety_and_cultivar(t *testing.T) {
  plant := &Plant {
    Genus:    "Some_Genus",
    Species:  "some_species",
    Variety:  "Some_variety",
    Cultivar:  "Some_cultivar",
  }
  expected := "Some_Genus some_species var. Some_variety 'Some_cultivar'"

  if fmt.Sprintf("%v", plant) != expected {
    t.Errorf("Expected %v, received %v", expected, fmt.Sprintf("%v", plant))
  }
  if fmt.Sprintf("%s", plant) != expected {
    t.Errorf("Expected %s, received %s", expected, fmt.Sprintf("%s", plant))
  }
}

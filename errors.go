package main

import (
	"fmt"
	"strings"
)

func mixSugar(sugar int) (string, error) {
	if sugar <= 3 {
		return "", fmt.Errorf("Not enough: %vg sugar", sugar)
	}
	return fmt.Sprintf("Mixture with %vg of sugar", sugar), nil
}

func mixFlour(flour int, sugarMixture string) (bool, error) {
	if flour <= 8 {
		return false, fmt.Errorf("not enough: %vg flour", flour)
	}
	if !strings.Contains(sugarMixture, "9g") {
		return false, fmt.Errorf("sugar mixture needs to be 9g")
	}
	return true, nil
}

// above function is thrid party

func bakeCakes(sugar int, flour int) (bool, error) {
	sugarMixture, err := mixSugar(sugar)
	if err != nil {
		return false, err
	}
	_ , err = mixFlour(flour, sugarMixture)
	if err != nil {
		return false, fmt.Errorf("cannot mix flour with %vg sugar %v", sugar, err)
	}
	return true, nil
}
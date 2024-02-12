package main

import "errors"

func getGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return NewAK47(), nil
	}
	if gunType == "musket" {
		return NewMusket(), nil
	}
	return nil, errors.New("Wrong gun type passed")
}

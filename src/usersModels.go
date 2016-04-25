package main

type user struct {
	Id     int    `json:id`
	Name   string `jason:"name"`
	Alias  string `jason:"alias"`
	Email  string `json:email`
	Active bool   `json:"active"`
}

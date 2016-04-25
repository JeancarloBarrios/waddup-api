package main

type event struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Date        string   `json:"date"`
	Duration    float64  `json:"duration"`
	Atending    []*Users `json:"atending"`
}

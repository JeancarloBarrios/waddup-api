package main

type Event struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Date        string  `json:"date"`
	Duration    float64 `json:"duration"`
	// Atending    []*Users `json:"atending"`
}

type Events []Event

package main

import (
	"fmt"
	"net/http"
	// Third Party Packages
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Logger now works Just fine")
}

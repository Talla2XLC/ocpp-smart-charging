package handlers

import (
	"net/http"
)

// HealthHandler - Server health checker
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//TODO add more complex state checking logic
	isOk := true

	var err error
	if !isOk {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte("FAIL"))
	} else {
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte("OK"))
	}

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

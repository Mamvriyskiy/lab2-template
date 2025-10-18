package main

import (
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
	http_utils "github.com/lapayka/rsoi-2/Common/HTTP_Utils"
	FS_DA "github.com/lapayka/rsoi-2/flight_service/DA"
)

type GateWay struct {
	db *FS_DA.DB
	//logger *slog.Logger
}

func main() {
	router := mux.NewRouter()

	db, _ := FS_DA.New("postgres", "postgres", "flights", "postgres")
	gw := GateWay{db}

	router.HandleFunc("/manage/health", http_utils.HealthCkeck).Methods("Get")
	router.HandleFunc("/api/v1/flights", gw.getFlights).Methods("Get")

	err := http.ListenAndServe(":8060", router)
	if err != nil {
		//gw.logger.Error("failed to run app", "error", err)
	}
}

func (gw *GateWay) getFlights(w http.ResponseWriter, r *http.Request) {
	flights, _ := gw.db.GetFlights()

	if len(flights) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Формируем объект ответа
	response := map[string]interface{}{
		"items":        flights,
		"page":         1,
		"pageSize":     len(flights),
		"totalElements": len(flights),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


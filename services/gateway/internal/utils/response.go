package utils

import (
	"gateway/models"
	"net/http"
)

func Response(w http.ResponseWriter, code models.StatusCode, description string, body []byte) {
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Content-Length", fmt.Sprintf("%d", len(body)))
	switch code {
	case models.OK:
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Description", description)
	case models.NotFound:
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Description", description)
	case models.BadRequest:
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Description", description)
	case models.InternalError:
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Description", description)
	case models.Deleted:
		w.WriteHeader(http.StatusNoContent)
		w.Header().Set("Description", description)
	}

	if body != nil {
		w.Write(body)
		return
	}
}

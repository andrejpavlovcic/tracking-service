package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"tracking_system/internal/entities"
	eerrors "tracking_system/internal/errors"
	"tracking_system/internal/logger"
	"tracking_system/internal/services"

	"github.com/gorilla/mux"
)

func NewRouter() http.Handler {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/status", getStatus).Methods(http.MethodGet)

	rtr.HandleFunc("/post/event/{accountID}", postEvent).Methods(http.MethodPost)

	rtr.HandleFunc("/count/unique-account", getUniqueAccountCount).Methods(http.MethodGet)

	return rtr
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func postEvent(w http.ResponseWriter, r *http.Request) {
	log := logger.GetLogger()

	accountID, err := strconv.ParseInt(mux.Vars(r)["accountID"], 10, 64)
	if err != nil {
		log.WithError(err).Error("accountID is not a valid number")
		w.WriteHeader(400)
		return
	}

	data := r.URL.Query().Get("data")
	if len(data) == 0 {
		log.Error("data not set in url query")
		w.WriteHeader(400)
		return
	}

	account, err := services.GetAccountService(r.Context(), log).GetAccountByID(accountID)
	switch err {
	case nil:
		// Skip
	case eerrors.ErrAccountNotFound:
		w.WriteHeader(404)
		return
	default:
		w.WriteHeader(500)
		return
	}

	if !account.IsActive {
		w.WriteHeader(403)
		return
	}

	event := entities.Event{
		AccountID: accountID,
		Data:      data,
		Timestamp: time.Now(),
	}

	err = services.GetKafkaService(r.Context(), log).SendEvent(&event)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
}

func getUniqueAccountCount(w http.ResponseWriter, r *http.Request) {
	log := logger.GetLogger()

	from := time.Now().AddDate(0, 0, -7)

	count, err := services.GetAccountService(r.Context(), log).GetUniqueAccountsCount(from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     "success",
		"statusCode": 200,
		"data":       count,
	})
}

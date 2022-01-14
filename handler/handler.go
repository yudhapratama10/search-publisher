package handler

import (
	"encoding/json"
	"net/http"

	"github.com/yudhapratama10/search-publisher/model"
)

func (handler *FootballHandler) Insert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var req model.FootballClub

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			result, _ := json.Marshal(map[string]interface{}{
				"message": err.Error(),
				"result":  "Failed",
			})
			w.WriteHeader(http.StatusBadRequest)
			w.Write(result)
			return
		}

		cursorInsert, err := handler.footballUsecase.Insert(req)
		if err != nil {
			result, _ := json.Marshal(map[string]interface{}{
				"message": err.Error(),
				"result":  "Failed",
			})
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(result)
			return
		}

		resp := map[string]interface{}{
			"result": "Success",
			"data":   cursorInsert,
		}
		result, _ := json.Marshal(resp)
		w.Write(result)
	} else {
		http.Error(w, "", http.StatusMethodNotAllowed)
	}
}

func (handler *FootballHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "PUT" {

		var req model.FootballClub
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			result, _ := json.Marshal(map[string]interface{}{
				"message": err.Error(),
				"result":  "Failed",
			})
			w.WriteHeader(http.StatusBadRequest)
			w.Write(result)
			return
		}

		if req.Id <= 0 {
			result, _ := json.Marshal(map[string]interface{}{
				"message": "ID Cannot be empty / zero",
				"result":  "Failed",
			})
			w.WriteHeader(http.StatusBadRequest)
			w.Write(result)
			return
		}

		cursorInsert, err := handler.footballUsecase.Update(req)
		if err != nil {
			result, _ := json.Marshal(map[string]interface{}{
				"message": err.Error(),
				"result":  "Failed",
			})
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(result)
			return
		}

		resp := map[string]interface{}{
			"data":   cursorInsert,
			"result": "Success",
		}
		result, err := json.Marshal(resp)
		w.Write(result)
	} else {
		http.Error(w, "", http.StatusMethodNotAllowed)
	}
}

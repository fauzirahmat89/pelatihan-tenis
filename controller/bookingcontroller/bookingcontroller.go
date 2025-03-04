package bookingcontroller

import (
	"encoding/json"
	"net/http"
	"pelatihan-tenis/helper"
	"pelatihan-tenis/models"
	"strconv"

	"github.com/gorilla/mux"
)

func Booking(w http.ResponseWriter, r *http.Request) {
	var booking models.Booking

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&booking); err != nil{
		response := map[string]string{"massage": err.Error()}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}

	defer r.Body.Close()

	if err := models.DB.Create(&booking).Error; err != nil {
		response := map[string]string{"massage": err.Error()}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"massage": "booking succeed"}
	helper.ResponseJson(w, http.StatusOK, response)
}

func Show(w http.ResponseWriter, r *http.Request){
	var bookings []models.Booking

	if err := models.DB.Find(&bookings).Error; err != nil {
		response := map[string]string{"massage": err.Error()}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}

	response, _ := json.Marshal(bookings)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func Update (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"],10,64)
	if err != nil {
		response := map[string]string{"massage": err.Error()}
		helper.ResponseJson(w, http.StatusBadRequest, response)
		return
	}

	var booking models.Booking 
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&booking); err != nil{
		response := map[string]string{"massage": err.Error()}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}

	defer r.Body.Close()

	if err := models.DB.Where("id = ?", id).Updates(&booking).Error; err != nil {
		response := map[string]string{"massage": err.Error()}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}

	booking.Id = id
	response := map[string]string{"massage": "berhasil di update"}
	helper.ResponseJson(w, http.StatusOK, response)
}

func Delete(w http.ResponseWriter, r *http.Request){
	input := map[string]string{"id":""}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		response := map[string]string{"massage": err.Error()}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}

	defer r.Body.Close()

	var booking models.Booking
	if models.DB.Delete(&booking, input["id"]).RowsAffected == 0 {
		response := map[string]string{"massage":"tidak dapat menghapus data"}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}
	
	response := map[string]string{"massage": "booking delete succes"}
	helper.ResponseJson(w, http.StatusOK, response)
}
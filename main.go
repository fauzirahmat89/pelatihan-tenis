package main

import (
	"fmt"
	"log"
	"net/http"
	"pelatihan-tenis/controller/authcontroller"
	"pelatihan-tenis/controller/bookingcontroller"
	"pelatihan-tenis/middlewares"
	"pelatihan-tenis/models"

	"github.com/gorilla/mux"
)

func main() {
	models.ConnectDatabase()
	r := mux.NewRouter()
	//user
	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")
	//admin
	r.HandleFunc("/admin/login", authcontroller.AdminLogin).Methods("POST")
	r.HandleFunc("/admin/register", authcontroller.AdminRegister).Methods("POST")
	r.HandleFunc("/admin/logout", authcontroller.AdminLogout).Methods("GET")
	//user and admin
	r.HandleFunc("/show/booking", bookingcontroller.Show).Methods("GET")
	//user
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/booking", bookingcontroller.Booking).Methods("POST")
	api.Use(middlewares.JWTMiddleware)
	//admin
	admin := r.PathPrefix("/admin").Subrouter()
	admin.HandleFunc("/booking/delete", bookingcontroller.Delete).Methods("POST")
	admin.HandleFunc("/booking/edit/{id}", bookingcontroller.Update).Methods("PUT")
	admin.Use(middlewares.AdminJWTMiddleware)




	fmt.Println("Server berjalan di port 8080")
	log.Fatal(http.ListenAndServe(":8080",r))
}
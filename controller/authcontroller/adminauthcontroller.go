package authcontroller

import (
	"encoding/json"
	"net/http"
	"pelatihan-tenis/config"
	"pelatihan-tenis/helper"
	"pelatihan-tenis/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func AdminLogin(w http.ResponseWriter, r *http.Request) {
	var adminInput models.Admin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&adminInput)
	if err != nil {
		response := map[string]string{"massage":err.Error()}
		helper.ResponseJson(w, http.StatusBadRequest, response)
		return
	}

	defer r.Body.Close()

	// mengambil data user berdasarkan username
	var user models.Admin
	if err := models.DB.Where("username = ?", adminInput.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"massage": "Username salah"}
			helper.ResponseJson(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"massage": err.Error()}
			helper.ResponseJson(w, http.StatusInternalServerError,response)
			return
		}
	}

	//cek apakah password valid atau tidak
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(adminInput.Password)); err != nil {
		response := map[string]string{"massage": "password salah"}
		helper.ResponseJson(w, http.StatusUnauthorized, response)
		return
	}

	// jika valid generate token jwt
	expTime := time.Now().Add(time.Minute * 15)
	claims := &config.JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "pelatihan-tenis",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	//mendeklarasi algoritma yang akan digunakan untuk signing
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//signed token
	token, err := tokenAlgo.SignedString(config.AdminJWT_KEY)
	if err != nil {
		response := map[string]string{"massage": err.Error()}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}

	//set token ke cookie
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Path: "/",
		Value: token,
		HttpOnly: true,
	})

	response := map[string]string{"massage":"login berhasil"}
	helper.ResponseJson(w, http.StatusOK, response)
}

func AdminRegister(w http.ResponseWriter, r *http.Request) {
	//mengambil inputan json
	var adminInput models.Admin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&adminInput)
	if err != nil{
		response := map[string]string{"massage": err.Error()}
		helper.ResponseJson(w, http.StatusBadRequest, response)
		return
	}

	defer r.Body.Close()

	//hash password menggunakan bcrypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(adminInput.Password), bcrypt.DefaultCost)
	adminInput.Password = string(hashPassword)

	//insert ke database
	if err := models.DB.Create(&adminInput).Error; 
	err != nil {
		response := map[string]string{"massage":err.Error()}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}
	
	response := map[string]string{"Message":"success"}
	helper.ResponseJson(w, http.StatusOK, response)
}

func AdminLogout(w http.ResponseWriter, r *http.Request) {
	//hapus token yang ada di cookie
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Path: "/",
		Value: "",
		HttpOnly: true,
		MaxAge: -1,
	})

	response := map[string]string{"massage":"logout berhasil"}
	helper.ResponseJson(w, http.StatusOK, response)
}
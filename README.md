# Pelatihan Tenis

Sebuah aplikasi berbasis Golang untuk manajemen pemesanan pelatihan tenis.

## Fitur

### User:
- Login
- Register
- Logout
- Melihat daftar booking

### Admin:
- Login
- Register
- Logout
- Menghapus & Mengedit booking

### Booking:
- User dapat melakukan booking

### Keamanan:
- Autentikasi JWT untuk keamanan API

## Instalasi

### 1. Clone Repository
```sh
git clone https://github.com/fauzirahmat89/pelatihan-tenis.git
cd pelatihan-tenis
```

### 2. Install Dependencies
```sh
go mod tidy
```

### 3. Buat File .env
Buat file `.env` di root proyek dan tambahkan konfigurasi berikut:
```env
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_HOST=
DB_PORT=
```

### 4. Jalankan Aplikasi
```sh
go run main.go
```
Server akan berjalan di `http://localhost:8080`

## Endpoint API

### User
- **POST** `/login` → Login user
- **POST** `/register` → Register user
- **GET** `/logout` → Logout user

### Admin
- **POST** `/admin/login` → Login admin
- **POST** `/admin/register` → Register admin
- **GET** `/admin/logout` → Logout admin

### Booking
- **GET** `/show/booking` → Menampilkan daftar booking
- **POST** `/api/booking` → Melakukan booking _(User harus login)_
- **POST** `/admin/booking/delete` → Menghapus booking _(Admin harus login)_
- **PUT** `/admin/booking/edit/{id}` → Mengedit booking _(Admin harus login)_

## Teknologi yang Digunakan
- **Golang**
- **Gorilla Mux**
- **GORM (MySQL)**
- **JWT Authentication**
- **Dotenv**

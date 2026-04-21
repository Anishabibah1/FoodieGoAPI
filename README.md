# 🍔 FoodieGo API

REST API untuk aplikasi food delivery berbasis mobile, dibangun menggunakan Node.js dan Express.js.

---

## 📋 Daftar Isi

- [Deskripsi](#deskripsi)
- [Tech Stack](#tech-stack)
- [Struktur Folder](#struktur-folder)
- [Cara Menjalankan](#cara-menjalankan)
- [Konfigurasi Environment](#konfigurasi-environment)
- [Endpoint API](#endpoint-api)
- [Contoh Request & Response](#contoh-request--response)
- [Autentikasi](#autentikasi)
- [Status Pesanan](#status-pesanan)

---

## 📖 Deskripsi

FoodieGo API menyediakan layanan backend untuk aplikasi food delivery, mencakup:

- **Autentikasi** — registrasi, login, dan manajemen profil user
- **Restoran** — daftar, pencarian, dan manajemen restoran
- **Menu** — daftar dan manajemen menu makanan per restoran
- **Pesanan** — pembuatan pesanan, riwayat, dan pembaruan status

---

## 🛠️ Tech Stack

| Kebutuhan | Package |
|-----------|---------|
| Runtime | Node.js |
| Framework | Express.js |
| Database | MySQL |
| Database Driver | mysql2 |
| Autentikasi | jsonwebtoken (JWT) |
| Enkripsi Password | bcryptjs |
| Environment | dotenv |
| CORS | cors |
| Dev Tool | nodemon |

---

## 📁 Struktur Folder

```
foodiego-api/
├── config/
│   └── database.js          # Konfigurasi koneksi MySQL
├── controllers/
│   ├── authController.js    # Logic register, login, profile
│   ├── restaurantController.js  # Logic CRUD restoran
│   ├── menuController.js    # Logic CRUD menu
│   └── orderController.js   # Logic pemesanan
├── middleware/
│   └── auth.js              # Middleware JWT & role admin
├── routes/
│   ├── auth.js              # Route /api/auth
│   ├── restaurants.js       # Route /api/restaurants
│   ├── menus.js             # Route /api/menus
│   └── orders.js            # Route /api/orders
├── sql/
│   └── schema.sql           # Script database & data dummy
├── .env                     # Konfigurasi environment (tidak di-commit)
├── .env.example             # Template environment
├── .gitignore
├── package.json
├── server.js                # Entry point aplikasi
└── README.md
```

---

## 🚀 Cara Menjalankan

### Prasyarat
Pastikan sudah menginstal:
- [Node.js](https://nodejs.org) versi LTS
- [XAMPP](https://apachefriends.org) (Apache + MySQL aktif)

### 1. Clone atau download project

```bash
git clone https://github.com/username/foodiego-api.git
cd foodiego-api
```

### 2. Install dependencies

```bash
npm install
```

### 3. Setup database

- Buka **phpMyAdmin** di `http://localhost/phpmyadmin`
- Buat database baru bernama `foodiego_db`
- Import file `sql/schema.sql` (sudah termasuk tabel + data dummy)

### 4. Konfigurasi environment

```bash
cp .env.example .env
```

Edit file `.env` sesuai konfigurasi lokal kamu (lihat bagian [Konfigurasi Environment](#konfigurasi-environment)).

### 5. Jalankan server

```bash
# Mode development (auto-restart saat ada perubahan file)
npm run dev

# Mode production
npm start
```

Server berjalan di: `http://localhost:3000`

---

## ⚙️ Konfigurasi Environment

Buat file `.env` di root project dengan isi berikut:

```env
# Server
PORT=3000

# Database MySQL
DB_HOST=localhost
DB_USER=root
DB_PASSWORD=
DB_NAME=foodiego_db

# JWT
JWT_SECRET=foodiego_secret_key_ganti_ini_di_production
JWT_EXPIRES_IN=7d
```

> ⚠️ **Penting:** Jangan pernah meng-commit file `.env` ke repository. File ini sudah terdaftar di `.gitignore`.

---

## 📡 Endpoint API

Base URL: `http://localhost:3000/api`

### 🔐 Auth

| Method | Endpoint | Deskripsi | Auth |
|--------|----------|-----------|------|
| POST | `/auth/register` | Registrasi user baru | — |
| POST | `/auth/login` | Login dan dapatkan token | — |
| GET | `/auth/profile` | Lihat profil user yang sedang login | ✅ Token |

### 🍽️ Restaurants

| Method | Endpoint | Deskripsi | Auth |
|--------|----------|-----------|------|
| GET | `/restaurants` | Daftar semua restoran | — |
| GET | `/restaurants?search=nama` | Cari restoran berdasarkan nama | — |
| GET | `/restaurants?category=Padang` | Filter restoran berdasarkan kategori | — |
| GET | `/restaurants/:id` | Detail satu restoran | — |
| POST | `/restaurants` | Tambah restoran baru | 👑 Admin |
| PUT | `/restaurants/:id` | Update data restoran | 👑 Admin |
| DELETE | `/restaurants/:id` | Hapus restoran | 👑 Admin |

### 🍜 Menus

| Method | Endpoint | Deskripsi | Auth |
|--------|----------|-----------|------|
| GET | `/restaurants/:id/menus` | Daftar menu dari satu restoran | — |
| POST | `/restaurants/:id/menus` | Tambah menu ke restoran | 👑 Admin |
| PUT | `/menus/:id` | Update data menu | 👑 Admin |
| DELETE | `/menus/:id` | Hapus menu | 👑 Admin |

### 📦 Orders

| Method | Endpoint | Deskripsi | Auth |
|--------|----------|-----------|------|
| POST | `/orders` | Buat pesanan baru | ✅ Token |
| GET | `/orders` | Lihat riwayat pesanan saya | ✅ Token |
| GET | `/orders/:id` | Detail satu pesanan | ✅ Token |
| PUT | `/orders/:id/status` | Update status pesanan | 👑 Admin |

---

## 📨 Contoh Request & Response

### Register

**Request:**
```http
POST /api/auth/register
Content-Type: application/json

{
  "name": "Budi Santoso",
  "email": "budi@example.com",
  "password": "password123",
  "phone": "081234567890",
  "address": "Jl. Merdeka No.1, Jakarta"
}
```

**Response (201):**
```json
{
  "success": true,
  "message": "Registrasi berhasil.",
  "data": {
    "id": 1,
    "name": "Budi Santoso",
    "email": "budi@example.com"
  }
}
```

---

### Login

**Request:**
```http
POST /api/auth/login
Content-Type: application/json

{
  "email": "budi@example.com",
  "password": "password123"
}
```

**Response (200):**
```json
{
  "success": true,
  "message": "Login berhasil.",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "name": "Budi Santoso",
      "email": "budi@example.com",
      "role": "user"
    }
  }
}
```

---

### Daftar Restoran

**Request:**
```http
GET /api/restaurants
```

**Response (200):**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "Warung Padang Jaya",
      "description": "Masakan Padang otentik dengan cita rasa asli",
      "address": "Jl. Sudirman No.10, Jakarta",
      "category": "Padang",
      "rating": 4.8,
      "is_open": 1
    }
  ]
}
```

---

### Buat Pesanan

**Request:**
```http
POST /api/orders
Content-Type: application/json
Authorization: Bearer <token>

{
  "restaurant_id": 1,
  "address": "Jl. Kebon Jeruk No.5, Jakarta",
  "note": "Tidak pedas, extra nasi",
  "items": [
    { "menu_id": 1, "quantity": 2 },
    { "menu_id": 2, "quantity": 1 }
  ]
}
```

**Response (201):**
```json
{
  "success": true,
  "message": "Pesanan berhasil dibuat.",
  "data": {
    "order_id": 10,
    "total_price": 98000,
    "status": "pending"
  }
}
```

---

## 🔑 Autentikasi

API ini menggunakan **JWT (JSON Web Token)**. Setelah login, sertakan token di header setiap request yang membutuhkan autentikasi:

```
Authorization: Bearer <token_kamu>
```

Token berlaku selama **7 hari**. Setelah kadaluarsa, user perlu login ulang.

**Level akses:**
- **Tanpa token** — hanya bisa mengakses endpoint publik (lihat restoran & menu)
- **✅ Token** — user yang sudah login
- **👑 Admin** — user dengan role `admin`

---

## 🔄 Status Pesanan

Alur status pesanan dari awal hingga selesai:

```
pending → confirmed → preparing → delivering → delivered
                                             ↘
                                           cancelled
```

| Status | Keterangan |
|--------|-----------|
| `pending` | Pesanan baru dibuat, menunggu konfirmasi |
| `confirmed` | Pesanan dikonfirmasi oleh restoran |
| `preparing` | Makanan sedang disiapkan |
| `delivering` | Pesanan sedang diantar |
| `delivered` | Pesanan telah sampai |
| `cancelled` | Pesanan dibatalkan |
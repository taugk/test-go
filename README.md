# Test Go API

REST API sederhana untuk manajemen **User**, **Paket**, dan **Transaksi** menggunakan Go, Fiber, GORM, dan PostgreSQL.

---

## Tech Stack

- **Go** — bahasa utama
- **Fiber v2** — HTTP framework
- **GORM** — ORM untuk database
- **PostgreSQL** — database
- **godotenv** — manajemen environment variable

---

## Struktur Proyek

```
test-oldo/
├── main.go
├── .env
├── config/
│   └── db.go          # Koneksi database
├── internal/
│   ├── model/
│   │   ├── user.go
│   │   ├── paket.go
│   │   └── transaksi.go
│   ├── repository/
│   │   ├── user_repository.go
│   │   ├── paket_repository.go
│   │   └── transaksi_repository.go
│   ├── service/
│   │   ├── user_service.go
│   │   ├── paket_service.go
│   │   └── transaksi_service.go
│   └── handler/
│       ├── user_handler.go
│       ├── paket_handler.go
│       └── transaksi_handler.go
├── routes/
│   └── routes.go
└── pkg/
    └── utils.go          # Helper response
```

---

## Cara Menjalankan

### 1. Clone & Masuk ke Direktori

```bash
git clone
cd test-oldo
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Buat File `.env`

```env
DB_HOST=localhost
DB_USER=postgres
DB_PASS=yourpassword
DB_NAME=test_oldo
DB_PORT=5432
```

### 4. Jalankan Aplikasi

```bash
go run cmd/main.go
```

Server berjalan di `http://localhost:3000`

---

## API Endpoints

Base URL: `/api`

### Users — `/api/users`

| Method | Endpoint        | Deskripsi         |
|--------|-----------------|-------------------|
| POST   | `/api/users`    | Buat user baru    |
| GET    | `/api/users`    | Ambil semua user  |
| GET    | `/api/users/:id`| Ambil user by ID  |
| PUT    | `/api/users/:id`| Update user       |
| DELETE | `/api/users/:id`| Hapus user        |

**Body (POST / PUT):**
```json
{
  "name": "Budi Santoso",
  "phone": "08123456789"
}
```

---

### Paket — `/api/paket`

| Method | Endpoint         | Deskripsi          |
|--------|------------------|--------------------|
| POST   | `/api/paket`     | Buat paket baru    |
| GET    | `/api/paket`     | Ambil semua paket  |
| GET    | `/api/paket/:id` | Ambil paket by ID  |
| PUT    | `/api/paket/:id` | Update paket       |
| DELETE | `/api/paket/:id` | Hapus paket        |

**Body (POST / PUT):**
```json
{
  "name": "Paket Hemat",
  "price": 50000,
  "quota": 10,
  "active_period": 30
}
```

---

### Transaksi — `/api/transaksis`

| Method | Endpoint               | Deskripsi              |
|--------|------------------------|------------------------|
| POST   | `/api/transaksis`      | Buat transaksi baru    |
| GET    | `/api/transaksis`      | Ambil semua transaksi  |
| GET    | `/api/transaksis/:id`  | Ambil transaksi by ID  |

**Body (POST):**
```json
{
  "user_id": 1,
  "paket_id": 2
}
```

> Harga (`price`) otomatis diambil dari data paket yang dipilih.

---

## Validasi

- **User**: `name` dan `phone` wajib diisi
- **Paket**: `name`, `price`, `quota`, dan `active_period` wajib diisi dan bernilai positif
- **Transaksi**: `user_id` dan `paket_id` wajib diisi; user dan paket harus sudah ada di database

---

## Auto Migration

Saat aplikasi dijalankan, tabel berikut otomatis dibuat/diperbarui:

- `users`
- `pakets`
- `transaksis`

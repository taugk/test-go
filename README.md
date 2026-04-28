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
├── cmd/
│   └── main.go
├── .env
├── config/
│   └── db.go                    # Koneksi database
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
    └── utils.go                 # Helper response
```

---

## Persiapan Database (PostgreSQL)

### 1. Install PostgreSQL

**Ubuntu / Debian:**
```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
sudo systemctl start postgresql
sudo systemctl enable postgresql
```

**macOS (Homebrew):**
```bash
brew install postgresql@15
brew services start postgresql@15
```

**Windows:**
Download installer dari [postgresql.org/download/windows](https://www.postgresql.org/download/windows/) lalu jalankan installer, kemudian pastikan service PostgreSQL sudah berjalan.

---

### 2. Buat Database & User

Masuk ke PostgreSQL shell:

```bash
# Linux / macOS
sudo -u postgres psql

# Windows (jalankan di Command Prompt sebagai Administrator)
psql -U postgres
```

Jalankan perintah berikut di dalam psql:

```sql
-- Buat user baru
CREATE USER oldo_user WITH PASSWORD 'yourpassword';

-- Buat database baru
CREATE DATABASE test_oldo;

-- Berikan akses ke user
GRANT ALL PRIVILEGES ON DATABASE test_oldo TO oldo_user;

-- Keluar dari psql
\q
```

---

### 3. Verifikasi Koneksi

```bash
psql -U oldo_user -d test_oldo -h localhost
```

Jika berhasil masuk ke shell PostgreSQL tanpa error, database sudah siap digunakan.

---

## Setup & Menjalankan Aplikasi

### 1. Clone & Masuk ke Direktori

```bash
git clone https://github.com/taugk/test-go.git
cd test-oldo
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Buat File `.env`

Buat file `.env` di root proyek dan sesuaikan dengan konfigurasi database yang telah dibuat:

```env
DB_HOST=localhost
DB_USER=oldo_user
DB_PASS=yourpassword
DB_NAME=test_oldo
DB_PORT=5432
```

### 4. Jalankan Aplikasi

```bash
go run cmd/main.go
```

Jika berhasil, terminal akan menampilkan:

```
✅ Database connected
```

Server berjalan di `http://localhost:3000`

> Tabel `users`, `pakets`, dan `transaksis` akan **otomatis dibuat** oleh GORM AutoMigrate saat aplikasi pertama kali dijalankan — tidak perlu membuat tabel secara manual.

---

## API Endpoints

Base URL: `http://localhost:3000/api`

### Users — `/api/users`

| Method | Endpoint         | Deskripsi        |
|--------|------------------|------------------|
| POST   | `/api/users`     | Buat user baru   |
| GET    | `/api/users`     | Ambil semua user |
| GET    | `/api/users/:id` | Ambil user by ID |
| PUT    | `/api/users/:id` | Update user      |
| DELETE | `/api/users/:id` | Hapus user       |

**Body (POST / PUT):**
```json
{
  "name": "Budi Santoso",
  "phone": "08123456789"
}
```

---

### Paket — `/api/paket`

| Method | Endpoint         | Deskripsi         |
|--------|------------------|-------------------|
| POST   | `/api/paket`     | Buat paket baru   |
| GET    | `/api/paket`     | Ambil semua paket |
| GET    | `/api/paket/:id` | Ambil paket by ID |
| PUT    | `/api/paket/:id` | Update paket      |
| DELETE | `/api/paket/:id` | Hapus paket       |

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

| Method | Endpoint              | Deskripsi             |
|--------|-----------------------|-----------------------|
| POST   | `/api/transaksis`     | Buat transaksi baru   |
| GET    | `/api/transaksis`     | Ambil semua transaksi |
| GET    | `/api/transaksis/:id` | Ambil transaksi by ID |

**Body (POST):**
```json
{
  "user_id": 1,
  "paket_id": 2
}
```

> `price` otomatis diambil dari harga paket yang dipilih. `user_id` dan `paket_id` harus sudah ada di database.

---

## Format Response

**Sukses (200 OK / 201 Created):**
```json
{
  "status": "OK",
  "data": { ... }
}
```

**Error:**
```json
{
  "status": "ERROR",
  "message": "pesan error"
}
```

---

## Validasi

- **User**: `name` dan `phone` wajib diisi
- **Paket**: `name`, `price`, `quota`, dan `active_period` wajib diisi dan bernilai lebih dari 0
- **Transaksi**: `user_id` dan `paket_id` wajib diisi dan keduanya harus sudah terdaftar di database

---

## Auto Migration

Saat aplikasi dijalankan, tabel berikut otomatis dibuat/diperbarui:

- `users`
- `pakets`
- `transaksis`

---

## Troubleshooting

| Error | Solusi |
|-------|--------|
| `Error loading .env file!` | Pastikan file `.env` ada di root proyek dan nama variabel sesuai |
| `Failed to connect database` | Periksa nilai di `.env` dan pastikan PostgreSQL sedang berjalan |
| `role "oldo_user" does not exist` | Jalankan ulang perintah `CREATE USER` di psql |
| `database "test_oldo" does not exist` | Jalankan ulang perintah `CREATE DATABASE` di psql |
| Port sudah dipakai | Ganti port di `app.Listen(":3000")` pada `cmd/main.go` |

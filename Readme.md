# Getting Started
- requirement
```
go : Versi 1.18 atau lebih baru
database : mysql
```
- instalation

pastikan database di env itu sudah sesuai dengan configurasi yang anda punya
```
git clone ....
go mod tidy
create database 
```
- Jalankan Migration
dan pastikan create dulu database
Pastikan sudah ada package golang migrate nya kalo belom ada silahkan install di [Link Ini](https://github.com/golang-migrate/migrate)
```
migrate -database "mysql://root:mahesa12@tcp(127.0.0.1:3306)/interview" -path db/migrations up
```
- pakcage http

menggunakan : [gin](https://gin-gonic.com/)
# Test
1. **Fixing gimana caranya .env bisa diambil di config.go**
2. **Fixing config.go connect ke database**
3. **Fixing Function GetBook,PostBook,GetBookById**
4. **Fixing untuk repositories pattern nya biar bisa digunakan di bookController**
4. **Tambah Filter Search di function GetBook**
5. **Buatkan router api di routes/routes.go**

- **Struktur Folder Proyek Go**

```bash
my-go-project/
├── config/                 # Folder untuk file konfigurasi
│   └── config.go           # Contoh file konfigurasi
├── db/                     # Folder untuk database dan migrasi
│   └── migrations/         # Folder untuk file migrasi database
│       └── 001_create_books_table.up.sql # Contoh file migrasi
├── docs/                   # Folder untuk dokumentasi proyek
│   └── README.md           # Dokumentasi proyek
├── models/                 # Folder untuk model data
│   └── book.go             # Contoh model untuk buku
├── repositories/           # Folder untuk repository data
│   └── book_repository.go  # Contoh repository untuk buku
├── routes/                 # Folder untuk rute atau router aplikasi
│   └── book_routes.go      # Contoh rute untuk buku
├── scripts/                # Folder untuk script automasi dan build
│   └── build.sh            # Contoh script build
├── go.mod                  # File modul Go
├── go.sum                  # File checksum modul Go
└── README.md               # Dokumentasi utama proyek

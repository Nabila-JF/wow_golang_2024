
# Go-Blog Project

## **Deskripsi**
Go-Blog adalah RESTful API sederhana yang dibangun menggunakan framework Gin di Go. API ini menyediakan fitur untuk mengelola pengguna, kategori, dan artikel. Proyek ini mendukung autentikasi menggunakan JWT (JSON Web Token) dan menerapkan teknik **soft delete** untuk penghapusan data.

---

## **Fitur Utama**
- **Autentikasi & Autorisasi**:
  - Registrasi dan login pengguna.
  - Autentikasi menggunakan JWT.
  - Middleware untuk memvalidasi token dan role pengguna.
- **Manajemen Artikel**:
  - CRUD artikel dengan validasi input.
  - Fitur slug untuk URL-friendly titles.
- **Manajemen Kategori**:
  - CRUD kategori dengan validasi input.
  - Pengelolaan slug dari nama kategori.

---

## **Dependencies**
Berikut adalah library yang digunakan dalam proyek ini:

| Library                  | Kegunaan                                   |
|--------------------------|--------------------------------------------|
| **github.com/gin-gonic/gin**   | Framework untuk membuat RESTful API.       |
| **github.com/go-sql-driver/mysql** | Driver untuk koneksi ke database MySQL. |
| **github.com/google/uuid**      | Membuat UUID unik untuk identifier.     |
| **github.com/golang-jwt/jwt/v5** | Digunakan untuk autentikasi JWT.        |
| **golang.org/x/crypto/bcrypt**   | Hashing dan validasi password.          |

Untuk menginstal semua dependencies:
```bash
go mod tidy
```

---

## **Struktur Proyek**
```
GO-BLOG/
│
├── config/
│   └── database.go         // Koneksi ke database
│
├── controllers/
│   ├── article_controller.go // Controller untuk manajemen artikel
│   ├── auth_controller.go    // Controller untuk autentikasi (register, login)
│   ├── category_controller.go // Controller untuk manajemen kategori
│   ├── user_controller.go    // Placeholder untuk fitur manajemen pengguna
│
├── middleware/
│   └── auth_middleware.go    // Middleware untuk validasi JWT dan role
│
├── models/
│   ├── article.go           // Model untuk artikel
│   ├── category.go          // Model untuk kategori
│   └── user.go              // Model untuk pengguna
│
├── routes/
│   └── routes.go            // Definisi endpoint API
│
├── go.mod                   // Modul dan dependensi Go
├── go.sum                   // File checksum dependensi Go
├── main.go                  // File utama untuk menjalankan aplikasi
└── README.md                // Dokumentasi proyek
```

---

## **Endpoint Dokumentasi**

### **Authentication**
| HTTP Method | Endpoint       | Keterangan               |
|-------------|----------------|--------------------------|
| `POST`      | `/register`    | Mendaftarkan user baru.  |
| `POST`      | `/login`       | Login dan mendapatkan JWT. |

### **CRUD Artikel**
| HTTP Method | Endpoint              | Keterangan                       |
|-------------|-----------------------|----------------------------------|
| `POST`      | `/auth/articles`      | Membuat artikel baru.            |
| `GET`       | `/auth/articles`      | Mengambil semua artikel.         |
| `PUT`       | `/auth/articles/:id`  | Mengupdate artikel berdasarkan ID. |
| `DELETE`    | `/auth/articles/:id`  | Soft delete artikel berdasarkan ID. |

### **CRUD Kategori**
| HTTP Method | Endpoint              | Keterangan                       |
|-------------|-----------------------|----------------------------------|
| `POST`      | `/auth/categories`    | Membuat kategori baru.           |
| `GET`       | `/auth/categories`    | Mengambil semua kategori.        |
| `PUT`       | `/auth/categories/:id`| Mengupdate kategori berdasarkan ID. |
| `DELETE`    | `/auth/categories/:id`| Soft delete kategori berdasarkan ID. |

---

## **Detail Alur CRUD**
### **Authentication**
1. **Register User**:
   - JSON Request:
     ```json
     {
       "username": "johndoe",
       "email": "johndoe@example.com",
       "password": "password123"
     }
     ```
   - Setelah validasi, password akan di-hash menggunakan `bcrypt` sebelum disimpan ke database. Role default adalah `USER`.

2. **Login User**:
   - JSON Request:
     ```json
     {
       "username": "johndoe",
       "password": "password123"
     }
     ```
   - Jika username dan password cocok, JWT token akan dikembalikan.

### **Article**
1. **Create Article**:
   - Membutuhkan field berikut:
     ```json
     {
       "category_id": "uuid-category",
       "title": "Introduction to Go",
       "content": "Go is an open-source programming language...",
       "author_id": "uuid-author"
     }
     ```
   - Slug akan otomatis dibuat berdasarkan judul menggunakan helper `createSlug`.

2. **Get Articles**:
   - Mengembalikan semua artikel yang belum dihapus secara **soft delete**.
   - Respon:
     ```json
     [
       {
         "article_id": "uuid",
         "category_id": "uuid-category",
         "title": "Introduction to Go",
         "content": "Go is an open-source programming language...",
         "slug": "introduction-to-go"
       }
     ]
     ```

3. **Update Article**:
   - Perlu menyediakan `ID` artikel dalam parameter URL.
   - Field yang di-update:
     ```json
     {
       "category_id": "uuid-category",
       "title": "Updated Title",
       "content": "Updated Content"
     }
     ```

4. **Delete Article**:
   - Soft delete dilakukan dengan mengisi kolom `deleted_at` dengan `CURRENT_TIMESTAMP`.

### **Category**
1. **Create Category**:
   - JSON Request:
     ```json
     {
       "name": "Technology",
       "description": "Category for tech articles"
     }
     ```
   - Slug akan dibuat otomatis menggunakan helper `createSlug`.

2. **Get Categories**:
   - Mengembalikan semua kategori yang belum dihapus.

3. **Update Category**:
   - Perlu menyediakan `ID` kategori dalam parameter URL.
   - Field yang di-update:
     ```json
     {
       "name": "Updated Name",
       "description": "Updated Description"
     }
     ```

4. **Delete Category**:
   - Soft delete dilakukan dengan mengisi kolom `deleted_at`.

---

## **Helper**
### **Slug Generator**
Digunakan untuk membuat slug dari `title` atau `name`.
```go
func createSlug(name string) string {
    return strings.ToLower(strings.ReplaceAll(name, " ", "-"))
}
```
- Contoh Input: `"Introduction to Go"`
- Contoh Output: `"introduction-to-go"`

---

## **Koneksi Database**
Berada di file `config/database.go`:
```go
dsn := "root:password@tcp(127.0.0.1:3306)/go-blog"
DB, err := sql.Open("mysql", dsn)
```
Ubah `dsn` sesuai dengan pengaturan MySQL Anda.

---

Silakan hubungi jika ada pertanyaan atau masukan!

# AkademiFlow – Backend API

AkademiFlow adalah backend REST API untuk sistem manajemen akademik (SMA).
Project ini dikembangkan sebagai **personal project** dengan fokus pada
**clean architecture**, **business rule yang jelas**, dan **maintainability**.

Pengembangan dimulai dari module paling fundamental: **User Management**,
dengan pendekatan yang menyerupai backend production.

---

## Tech Stack
- Go
- Gin
- GORM
- PostgreSQL
- JWT
- bcrypt

---

## Fitur Saat Ini

### User Management
- Register & Login (JWT Authentication)
- Role-based access control (`admin`, `user`)
- Update user role (admin only)
- Prevent admin self-demote
- List users dengan:
  - pagination
  - search (name / email)
  - include soft deleted users
- Soft delete & restore user
- Global error handling
- Unit test pada service layer

---

## Arsitektur

Project ini menggunakan **layered architecture**:

Handler (Gin / HTTP)
↓
Service (Business Logic)
↓
Repository (Database Access)
↓
PostgreSQL


### Prinsip Desain
- Handler hanya mengurus HTTP concern
- Service berisi seluruh aturan bisnis
- Repository fokus pada query database
- Tidak ada business logic di handler
- Error dimapping secara terpusat (global error handler)

---

## Authentication & Authorization

- Authentication menggunakan **JWT**
- JWT payload:
  - `user_id`
  - `role`
- Middleware:
  - JWT validation
  - Role guard (admin only endpoint)

Contoh aturan bisnis:
- Admin tidak boleh menurunkan role dirinya sendiri
- User soft deleted tidak muncul di list default

---

## Soft Delete Strategy

User tidak dihapus permanen dari database.

- Menggunakan kolom `deleted_at`
- Default query hanya mengambil user aktif
- Admin dapat:
  - melihat user terhapus (`include_deleted=true`)
  - melakukan restore user

---

## 📄 API Overview (Ringkas)

### List Users
GET /api/admin/users
Query Params:
- `page` (default: 1)
- `limit` (default: 10, max: 100)
- `q` (search name/email)
- `include_deleted` (true / false)

---

### Update User Role
PATCH /api/admin/users/:id/role
Body:
```json
{
  "role": "admin"
}
```

### Delete User (soft delete)
DELETE /api/admin/users/:id

### Restore User
PATCH /api/admin/users/:id/restore

---

## Tesing
- Unit test saat ini difokuskan pada service layer
- Menggunakan fake repository (tanpa database)
- Test mencakup:
    - validasi business rule
    - error mapping
    - edge case
    - success scenario

---

## Roadmap
- Module Attendance / Class
- Integration test (httptest)
- API documentation (OpenAPI)
- Audit log

---

## Author
Personal project by Riky.  
Fokus pada backend engineering, clean architecture, dan real world patterns.
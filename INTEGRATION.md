# Integration Contract

## Base URL
http://localhost:3000/api/v1

## CORS
API ini mengizinkan akses dari semua origin (*) untuk tahap development.

## Authentication
Semua endpoint kecuali register dan login membutuhkan token JWT.
Sertakan token di header setiap request:

## Response Format
Semua response menggunakan format snake_case JSON:

Success:
```json
{
  "status": "success",
  "data": {}
}
```

Error:
```json
{
  "status": "error",
  "message": "pesan error"
}
```

## Endpoints
| Method | Endpoint | Auth |
|--------|----------|------|
| POST | /api/v1/auth/register | Tidak |
| POST | /api/v1/auth/login | Tidak |
| GET | /api/v1/restaurants/ | Token |
| GET | /api/v1/restaurants/search?q= | Token |
| GET | /api/v1/restaurants/:id | Token |
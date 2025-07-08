# 🧾 HRIS - Company Management API

This section documents the **Company** management endpoints used for creating, retrieving, updating, and deleting company data in the HRIS system.

---

## 📌 Base URL

```
http://localhost:8080/api/v1/companies
```

---

## ➕ Create Company

**POST** `/create`

Create a new company.

### Request Body

```json
{
  "user_id": 4,
  "name": "PT Teknologi Nusantara",
  "address": "Jl. Merdeka No.123, Jakarta",
  "email": "info@teknologi.co.id",
  "phone": "021-12345678",
  "logo": "/uuid-date-namafile"
}
```

### Response

```json
{
  "data": {
    "uuid": "generated-company-uuid",
    "user_id": 4,
    "name": "PT Teknologi Nusantara",
    "address": "Jl. Merdeka No.123, Jakarta",
    "email": "info@teknologi.co.id",
    "phone": "021-12345678",
    "logo": "/uuid-date-namafile"
  },
  "message": "Company created successfully",
  "status": "success"
}
```

---

## 📋 Get All Companies

**GET** `/get-all`

Returns list of all companies.

### Response

```json
{
  "data": [
    {
      "uuid": "cd041c00-9808-4d85-8a2e-5ca9af9f938f",
      "name": "PT Teknologi Nusantara",
      "email": "info@teknologi.co.id",
      "address": "Jl. Merdeka No.123, Jakarta",
      "phone": "021-12345678",
      "logo": "/uuid-date-namafile"
    }
  ],
  "message": "Companies fetched",
  "status": "success"
}
```

---

## 🔍 Get Company By UUID

**GET** `/get/:uuid`
Example: `/get/cd041c00-9808-4d85-8a2e-5ca9af9f938f`

### Response

```json
{
  "data": {
    "uuid": "cd041c00-9808-4d85-8a2e-5ca9af9f938f",
    "name": "PT Teknologi Nusantara",
    "email": "info@teknologi.co.id",
    "address": "Jl. Merdeka No.123, Jakarta",
    "phone": "021-12345678",
    "logo": "/uuid-date-namafile"
  },
  "message": "Company found",
  "status": "success"
}
```

---

## ✏️ Update Company

**PATCH** `/update/:uuid`
Example: `/update/cd041c00-9808-4d85-8a2e-5ca9af9f938f`

### Request Body

```json
{
  "user_id": 4,
  "name": "PT Teknologi Nusantara Update",
  "address": "Jl. Merdeka No.123, Jakarta",
  "email": "info@teknologi.co.id",
  "phone": "021-12345678",
  "logo": "/uuid-date-namafile"
}
```

### Response

```json
{
  "data": {
    "uuid": "cd041c00-9808-4d85-8a2e-5ca9af9f938f",
    "name": "PT Teknologi Nusantara Update",
    "email": "info@teknologi.co.id",
    "address": "Jl. Merdeka No.123, Jakarta",
    "phone": "021-12345678",
    "logo": "/uuid-date-namafile"
  },
  "message": "Company updated successfully",
  "status": "success"
}
```

---

## ❌ Delete Company

**DELETE** `/delete/:uuid`
Example: `/delete/cd041c00-9808-4d85-8a2e-5ca9af9f938f`

### Request Body

```json
{
  "user_id": 4,
  "name": "PT Teknologi Nusantara Update",
  "address": "Jl. Merdeka No.123, Jakarta",
  "email": "info@teknologi.co.id",
  "phone": "021-12345678",
  "logo": "/uuid-date-namafile"
}
```

### Response

```json
{
  "data": {
    "uuid": "cd041c00-9808-4d85-8a2e-5ca9af9f938f",
    "user_id": 4,
    "logo": "/uuid-date-namafile",
    "name": "PT Teknologi Nusantara Update",
    "address": "Jl. Merdeka No.123, Jakarta",
    "email": "info@teknologi.co.id",
    "phone": "021-12345678"
  },
  "message": "Company deleted successfully",
  "status": "success"
}
```
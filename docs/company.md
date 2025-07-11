# 🧾 HRIS - Company Management API

Documentation for **Company** management endpoints in the HRIS system, including create, retrieve, update, and delete
operations.

---

## 📌 Base URL

```

http://localhost:8080/api/v1/companies

````

---

## ➕ Create Company

**POST** `/create`

Create a new company.

### Request Body

```json
{
  "user_uuid": "f4d811c2-8583-44f0-ab4c-b2cc1f483387",
  "name": "PT Teknologi Nusantara",
  "address": "Jl. Merdeka No.123, Jakarta",
  "email": "info2@teknologi.co.id",
  "phone": "021-12345678",
  "logo": "/uuid-date-namafile"
}
````

### Response

```json
{
  "data": {
    "uuid": "2d1f05da-5fa4-4977-80bc-8d0d3b031bf3",
    "logo": "/uuid-date-namafile",
    "name": "PT Teknologi Nusantara 2",
    "address": "Jl. Merdeka No.123, Jakarta",
    "email": "info2@teknologi.co.id",
    "phone": "021-12345678",
    "user": {
      "uuid": "85b47132-345d-40a1-868f-3e48adacb319",
      "firstname": "Marco",
      "lastname": "Melandri"
    }
  },
  "message": "Company Created Succesfully",
  "status": "success"
}
```

---

## 📋 Get All Companies

**GET** `/get-all`

Retrieve the list of all companies.

### Response

```json
{
  "data": [
    {
      "uuid": "2d1f05da-5fa4-4977-80bc-8d0d3b031bf3",
      "logo": "/uuid-date-namafile",
      "name": "PT Teknologi Nusantara 2",
      "address": "Jl. Merdeka No.123, Jakarta",
      "email": "info2@teknologi.co.id",
      "phone": "021-12345678",
      "user": {
        "uuid": "85b47132-345d-40a1-868f-3e48adacb319",
        "firstname": "Marco",
        "lastname": "Melandri"
      }
    }
  ],
  "message": "Successfully get all companies",
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
    "uuid": "2d1f05da-5fa4-4977-80bc-8d0d3b031bf3",
    "logo": "/uuid-date-namafile",
    "name": "PT Teknologi Nusantara 2",
    "address": "Jl. Merdeka No.123, Jakarta",
    "email": "info2@teknologi.co.id",
    "phone": "021-12345678",
    "user": {
      "uuid": "85b47132-345d-40a1-868f-3e48adacb319",
      "firstname": "Marco",
      "lastname": "Melandri"
    }
  },
  "message": "Succesfully get company",
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
  "user_uuid": "f4d811c2-8583-44f0-ab4c-b2cc1f483387",
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
    "uuid": "2d1f05da-5fa4-4977-80bc-8d0d3b031bf3",
    "logo": "/uuid-date-namafile",
    "name": "PT Teknologi Nusantara",
    "address": "Jl. Merdeka No.123, Jakarta",
    "email": "info@teknologi.co.id",
    "phone": "021-12345678",
    "user": {
      "uuid": "85b47132-345d-40a1-868f-3e48adacb319",
      "firstname": "Marco",
      "lastname": "Melandri"
    }
  },
  "message": "Company updated succesfully",
  "status": "success"
}
```

---

## ❌ Delete Company

**DELETE** `/delete/:uuid`
Example: `/delete/cd041c00-9808-4d85-8a2e-5ca9af9f938f`

### Response

```json
{
  "data": {
    "uuid": "2d1f05da-5fa4-4977-80bc-8d0d3b031bf3",
    "logo": "/uuid-date-namafile",
    "name": "PT Teknologi Nusantara",
    "address": "Jl. Merdeka No.123, Jakarta",
    "email": "info@teknologi.co.id",
    "phone": "021-12345678",
    "user": {
      "uuid": "85b47132-345d-40a1-868f-3e48adacb319",
      "firstname": "Marco",
      "lastname": "Melandri"
    }
  },
  "message": "Company deleted successfully",
  "status": "success"
}
```
# 🕒 HRIS - Shift Management API

Documentation for **Shift** management endpoints in the HRIS system, including create, retrieve, update, and delete
operations.

---

## 📌 Base URL

```
http://localhost:8080/api/v1/shifts
```

---

## ➕ Create Shift

**POST** `/create`

Create a new shift for a company.

### Request Body

```json
{
  "company_uuid": "86a9688c-53b4-42c8-b3c1-819fd6ad1d9f",
  "name": "Morning Shift",
  "date": "2025-07-23T00:00:00Z",
  "hour_from": "2025-07-23T08:00:00Z",
  "hour_until": "2025-07-23T16:00:00Z"
}
```

### Response

```json
{
  "data": {
    "uuid": "b88948f3-bf60-4f20-ba13-5b71ab899e91",
    "company_uuid": "86a9688c-53b4-42c8-b3c1-819fd6ad1d9f",
    "name": "Morning Shift",
    "date": "2025-07-23T00:00:00Z",
    "hour_from": "2025-07-23T08:00:00Z",
    "hour_until": "2025-07-23T16:00:00Z"
  },
  "message": "Shift created successfully",
  "status": "success"
}
```

---

## 📋 Get All Shifts

**GET** `/get-all`

Retrieve all available shifts.

### Response

```json
{
  "data": [
    {
      "uuid": "b88948f3-bf60-4f20-ba13-5b71ab899e91",
      "company_uuid": "86a9688c-53b4-42c8-b3c1-819fd6ad1d9f",
      "name": "Morning Shift",
      "date": "2025-07-23T00:00:00Z",
      "hour_from": "2025-07-23T08:00:00Z",
      "hour_until": "2025-07-23T16:00:00Z"
    }
  ],
  "message": "Successfully retrieved all shifts",
  "status": "success"
}
```

---

## 🔍 Get Shift By UUID

**GET** `/get/:uuid`
Example: `/get/b88948f3-bf60-4f20-ba13-5b71ab899e91`

### Response

```json
{
  "data": {
    "uuid": "b88948f3-bf60-4f20-ba13-5b71ab899e91",
    "company_uuid": "86a9688c-53b4-42c8-b3c1-819fd6ad1d9f",
    "name": "Morning Shift",
    "date": "2025-07-23T00:00:00Z",
    "hour_from": "2025-07-23T08:00:00Z",
    "hour_until": "2025-07-23T16:00:00Z"
  },
  "message": "Shift retrieved successfully",
  "status": "success"
}
```

---

## ✏️ Update Shift

**PATCH** `/update/:uuid`
Example: `/update/b88948f3-bf60-4f20-ba13-5b71ab899e91`

### Request Body

```json
{
  "company_uuid": "86a9688c-53b4-42c8-b3c1-819fd6ad1d9f",
  "name": "Morning Shift Update",
  "date": "2025-07-23T00:00:00Z",
  "hour_from": "2025-07-23T08:00:00Z",
  "hour_until": "2025-07-23T16:00:00Z"
}
```

### Response

```json
{
  "data": {
    "uuid": "b88948f3-bf60-4f20-ba13-5b71ab899e91",
    "company_uuid": "86a9688c-53b4-42c8-b3c1-819fd6ad1d9f",
    "name": "Morning Shift Update",
    "date": "2025-07-23T00:00:00Z",
    "hour_from": "2025-07-23T08:00:00Z",
    "hour_until": "2025-07-23T16:00:00Z"
  },
  "message": "Shift updated successfully",
  "status": "success"
}
```

---

## ❌ Delete Shift

**DELETE** `/delete/:uuid`
Example: `/delete/b88948f3-bf60-4f20-ba13-5b71ab899e91`

### Response

```json
{
  "data": null,
  "message": "Shift deleted successfully",
  "status": "success"
}
```
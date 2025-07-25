# 🧑‍💻 HRIS - User Experience API

API for managing user professional experiences such as job roles, durations, and descriptions.

---

## 📌 Base URL

```
http://localhost:8080/api/v1/user-experiences
```

---

## 📋 Get All User Experiences

**GET** `/get-all/`

Retrieve all user experiences stored in the system.

### Response

```json
{
  "data": [
    {
      "uuid": "af254687-3c4b-4ecd-870f-1e40d7063639",
      "user_uuid": "265edc19-3544-45c8-a887-a7ca5d67455d",
      "name": "Software Engineer at ABC Corp",
      "description": "Worked on developing scalable backend services using Go and PostgreSQL.",
      "start_date": "2020-01-15T00:00:00Z",
      "end_date": "2023-06-30T00:00:00Z"
    }
  ],
  "message": "Successfully retrieved all user experiences",
  "status": "success"
}
```

---

## ➕ Create User Experience

**POST** `/create/`

Create a new user experience entry.

### Request Body

```json
{
  "user_uuid": "265edc19-3544-45c8-a887-a7ca5d67455d",
  "name": "Software Engineer at ABC Corp",
  "description": "Worked on developing scalable backend services using Go and PostgreSQL.",
  "start_date": "2020-01-15T00:00:00Z",
  "end_date": "2023-06-30T00:00:00Z"
}
```

### Response

```json
{
  "data": {
    "uuid": "af254687-3c4b-4ecd-870f-1e40d7063639",
    "user_uuid": "265edc19-3544-45c8-a887-a7ca5d67455d",
    "name": "Software Engineer at ABC Corp",
    "description": "Worked on developing scalable backend services using Go and PostgreSQL.",
    "start_date": "2020-01-15T00:00:00Z",
    "end_date": "2023-06-30T00:00:00Z"
  },
  "message": "User experience created successfully",
  "status": "success"
}
```

---

## 🔍 Get User Experience By ID

**GET** `/get-by-id/:id`

Example: `/get-by-id/2`

### Response

```json
{
  "data": {
    "uuid": "af254687-3c4b-4ecd-870f-1e40d7063639",
    "user_uuid": "265edc19-3544-45c8-a887-a7ca5d67455d",
    "name": "Software Engineer at ABC Corp",
    "description": "Worked on developing scalable backend services using Go and PostgreSQL.",
    "start_date": "2020-01-15T00:00:00Z",
    "end_date": "2023-06-30T00:00:00Z"
  },
  "message": "Successfully retrieved user experience by ID",
  "status": "success"
}
```

---

## 🔎 Get User Experience By UUID

**GET** `/get-by-uuid/:uuid`

Example: `/get-by-uuid/af254687-3c4b-4ecd-870f-1e40d7063639`

### Response

```json
{
  "data": {
    "uuid": "af254687-3c4b-4ecd-870f-1e40d7063639",
    "user_uuid": "265edc19-3544-45c8-a887-a7ca5d67455d",
    "name": "Software Engineer at ABC Corp",
    "description": "Worked on developing scalable backend services using Go and PostgreSQL.",
    "start_date": "2020-01-15T00:00:00Z",
    "end_date": "2023-06-30T00:00:00Z"
  },
  "message": "Successfully retrieved user experience by UUID",
  "status": "success"
}
```

---

## 👤 Get User Experiences By User UUID

**GET** `/get-by-user/:user_uuid`

Example: `/get-by-user/265edc19-3544-45c8-a887-a7ca5d67455d`

### Response

```json
{
  "data": [
    {
      "uuid": "af254687-3c4b-4ecd-870f-1e40d7063639",
      "name": "Software Engineer at ABC Corp",
      "description": "Worked on developing scalable backend services using Go and PostgreSQL.",
      "start_date": "2020-01-15T00:00:00Z",
      "end_date": "2023-06-30T00:00:00Z"
    }
  ],
  "message": "Successfully retrieved experiences for user",
  "status": "success"
}
```

---

## ✏️ Update User Experience

**PATCH** `/update/:uuid`

Example: `/update/af254687-3c4b-4ecd-870f-1e40d7063639`

### Request Body

```json
{
  "user_uuid": "265edc19-3544-45c8-a887-a7ca5d67455d",
  "name": "Udah DI UPDATE NIH",
  "description": "Worked on developing scalable backend services using Go and PostgreSQL.",
  "start_date": "2020-01-15T00:00:00Z",
  "end_date": "2023-06-30T00:00:00Z"
}
```

### Response

```json
{
  "data": {
    "uuid": "af254687-3c4b-4ecd-870f-1e40d7063639",
    "name": "Udah DI UPDATE NIH",
    "description": "Worked on developing scalable backend services using Go and PostgreSQL.",
    "start_date": "2020-01-15T00:00:00Z",
    "end_date": "2023-06-30T00:00:00Z"
  },
  "message": "User experience updated successfully",
  "status": "success"
}
```

---

## ❌ Delete User Experience

**DELETE** `/delete/:uuid`

Example: `/delete/af254687-3c4b-4ecd-870f-1e40d7063639`

> 📌 Note: Even though the request includes a body, deletion is performed based on the `uuid` in the path.

### Response

```json
{
  "data": {
    "uuid": "af254687-3c4b-4ecd-870f-1e40d7063639",
    "name": "Udah DI UPDATE NIH"
  },
  "message": "User experience deleted successfully",
  "status": "success"
}
```

---


# 🏦 HRIS - User Bank Management API

Documentation for **User Bank** management endpoints in the HRIS system, including create, retrieve, update, and delete operations.

---

## 📌 Base URL

```
http://localhost:8080/api/v1/user-banks
```

---

## ➕ Create User Bank

**POST** `/create`

Create a new bank record for a specific user.

### Request Body

```json
{
  "user_uuid": "265edc19-3544-45c8-a887-a7ca5d67455d",
  "name": "BCA",
  "number": "1234567890",
  "exp_date": "2025-12-31T00:00:00Z",
  "start_date": "2020-01-15T00:00:00Z"
}
```

### Response

```json
{
  "data": {
    "uuid": "generated-bank-uuid",
    "user_uuid": "265edc19-3544-45c8-a887-a7ca5d67455d",
    "name": "BCA",
    "number": "1234567890",
    "exp_date": "2025-12-31T00:00:00Z",
    "start_date": "2020-01-15T00:00:00Z"
  },
  "message": "User bank created",
  "status": "success"
}
```

---

## 📥 Get All User Banks

**GET** `/get-all`

Retrieves all user bank records in the system.

### Response

```json
{
  "data": [
    {
      "uuid": "bank-uuid-1",
      "user": {
        "uuid": "265edc19-3544-45c8-a887-a7ca5d67455d",
        "fullname": "John Doe",
        "email": "john@example.com"
      },
      "name": "BCA",
      "number": "1234567890",
      "exp_date": "2025-12-31T00:00:00Z"
    },
    ...
  ],
  "message": "Successfully retrieved user banks",
  "status": "success"
}
```

---

## 🔎 Get User Bank by ID

**GET** `/get-by-id/:id`

Example: `/get-by-id/1`

Retrieve a single bank entry by its database ID.

### Response

```json
{
  "data": {
    "uuid": "bank-uuid-1",
    "user": {
      "uuid": "265edc19-3544-45c8-a887-a7ca5d67455d",
      "fullname": "John Doe",
      "email": "john@example.com"
    },
    "name": "BCA",
    "number": "1234567890",
    "exp_date": "2025-12-31T00:00:00Z"
  },
  "message": "Successfully retrieved user bank",
  "status": "success"
}
```

---

## 🔎 Get User Bank by UUID

**GET** `/get-by-uuid/:uuid`

Example: `/get-by-uuid/791edefc-c0a4-4700-9204-efd264ae0a92`

Retrieve a user bank entry by its UUID.

### Response

```json
{
  "data": {
    "uuid": "791edefc-c0a4-4700-9204-efd264ae0a92",
    "user": {
      "uuid": "265edc19-3544-45c8-a887-a7ca5d67455d",
      "fullname": "John Doe",
      "email": "john@example.com"
    },
    "name": "BCA",
    "number": "1234567890",
    "exp_date": "2025-12-31T00:00:00Z"
  },
  "message": "Successfully retrieved user bank",
  "status": "success"
}
```

---

## 🔍 Get User Bank by User UUID

**GET** `/get-by-user/:uuid`

Example: `/get-by-user/265edc19-3544-45c8-a887-a7ca5d67455d`

Retrieves all bank entries associated with a specific user.

### Response

```json
{
  "data": [
    {
      "uuid": "bank-uuid-1",
      "name": "BCA",
      "number": "1234567890",
      "exp_date": "2025-12-31T00:00:00Z"
    },
    ...
  ],
  "message": "Successfully retrieved user bank(s)",
  "status": "success"
}
```

---

## ✏️ Update User Bank

**PATCH** `/update/:uuid`

Example: `/update/791edefc-c0a4-4700-9204-efd264ae0a92`

### Request Body

```json
{
  "user_uuid": "265edc19-3544-45c8-a887-a7ca5d67455d",
  "name": "BCA update jadi MANDIRI",
  "number": "1234567890",
  "exp_date": "2025-12-31T00:00:00Z",
  "start_date": "2020-01-15T00:00:00Z"
}
```

### Response

```json
{
  "data": {
    "uuid": "791edefc-c0a4-4700-9204-efd264ae0a92",
    "user_uuid": "265edc19-3544-45c8-a887-a7ca5d67455d",
    "name": "BCA update jadi MANDIRI",
    "number": "1234567890",
    "exp_date": "2025-12-31T00:00:00Z",
    "start_date": "2020-01-15T00:00:00Z"
  },
  "message": "User bank updated",
  "status": "success"
}
```

---

## ❌ Delete User Bank

**DELETE** `/delete/:uuid`

Example: `/delete/791edefc-c0a4-4700-9204-efd264ae0a92`

### Response

```json
{
  "data": {
    "uuid": "791edefc-c0a4-4700-9204-efd264ae0a92"
  },
  "message": "User bank deleted",
  "status": "success"
}
```

---
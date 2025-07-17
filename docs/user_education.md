
# 🎓 HRIS - User Education API

Documentation for **User Education** endpoints in the HRIS system, covering creation, retrieval, update, and deletion.

---

## 📌 Base URL

```
http://localhost:8080/api/v1/user-educations
```

---

## ➕ Create User Education

**POST** `/create`

Create a new user education record.

### Request Body

```json
{
  "user_uuid": "123e4567-e89b-12d3-a456-426614174000",
  "name": "Master of Data Science",
  "location": "Institute of Advanced Studies",
  "start_date": "2020-08-01T00:00:00Z",
  "end_date": "2022-12-15T00:00:00Z",
  "grade": 92.5
}
```

### Response

```json
{
  "data": {
    "uuid": "generated-uuid",
    "name": "Master of Data Science",
    "location": "Institute of Advanced Studies",
    "start_date": "2020-08-01",
    "end_date": "2022-12-15",
    "grade": 92.5,
    "user": {
      "uuid": "123e4567-e89b-12d3-a456-426614174000",
      "fullname": "John Doe",
      "email": "john@example.com"
    }
  },
  "message": "User education created successfully",
  "status": "success"
}
```

---

## 📋 Get All User Educations

**GET** `/get-all`

Retrieve all user education records.

### Response

```json
{
  "data": [
    {
      "uuid": "generated-uuid",
      "name": "Master of Data Science",
      "location": "Institute of Advanced Studies",
      "start_date": "2020-08-01",
      "end_date": "2022-12-15",
      "grade": 92.5,
      "user": {
        "uuid": "123e4567-e89b-12d3-a456-426614174000",
        "fullname": "John Doe",
        "email": "john@example.com"
      }
    }
  ],
  "message": "Successfully retrieved all user education records",
  "status": "success"
}
```

---

## 🔍 Get User Education By ID

**GET** `/get-by-id/:id`

Example: `/get-by-id/4`

### Response

```json
{
  "data": {
    "uuid": "generated-uuid",
    "name": "Master of Data Science",
    "location": "Institute of Advanced Studies",
    "start_date": "2020-08-01",
    "end_date": "2022-12-15",
    "grade": 92.5,
    "user": {
      "uuid": "123e4567-e89b-12d3-a456-426614174000",
      "fullname": "John Doe",
      "email": "john@example.com"
    }
  },
  "message": "Successfully retrieved user education by ID",
  "status": "success"
}
```

---

## 🔎 Get User Education By UUID

**GET** `/get-by-uuid/:uuid`

Example: `/get-by-uuid/b749d9de-8135-4a3c-a92e-6d675c199ccc`

### Response

```json
{
  "data": {
    "uuid": "b749d9de-8135-4a3c-a92e-6d675c199ccc",
    "name": "Master of Data Science",
    "location": "Institute of Advanced Studies",
    "start_date": "2020-08-01",
    "end_date": "2022-12-15",
    "grade": 92.5,
    "user": {
      "uuid": "123e4567-e89b-12d3-a456-426614174000",
      "fullname": "John Doe",
      "email": "john@example.com"
    }
  },
  "message": "Successfully retrieved user education by UUID",
  "status": "success"
}
```

---

## 👤 Get User Education By User UUID

**GET** `/get-by-user/:uuid`

Example: `/get-by-user/1fec3831-a612-45d0-97b7-eb45c2176f45`

### Response (✅ Found)

```json
{
  "data": [
    {
      "uuid": "generated-uuid",
      "name": "Master of Data Science",
      "location": "Institute of Advanced Studies",
      "start_date": "2020-08-01",
      "end_date": "2022-12-15",
      "grade": 92.5,
      "user": {
        "uuid": "1fec3831-a612-45d0-97b7-eb45c2176f45",
        "fullname": "John Doe",
        "email": "john@example.com"
      }
    }
  ],
  "message": "Successfully retrieved user education for user",
  "status": "success"
}
```

### Response (⚠️ User Not Found)

```json
{
  "message": "User not found",
  "status": "error"
}
```

**HTTP Status**: `404 Not Found`

### Response (ℹ️ No Education Found)

```json
{
  "data": [],
  "message": "No education found for this user",
  "status": "success"
}
```

**HTTP Status**: `200 OK`

---

## ✏️ Update User Education

**PATCH** `/update/:uuid`

Example: `/update/b749d9de-8135-4a3c-a92e-6d675c199ccc`

### Request Body

```json
{
  "user_uuid": "123e4567-e89b-12d3-a456-426614174000",
  "name": "S1 TEKNIK KOMPUTER",
  "location": "ADANYA DI KAMPUS MULTMIEDIA NUDSNATRAS",
  "start_date": "2020-08-01T00:00:00Z",
  "end_date": "2022-12-15T00:00:00Z",
  "grade": 24.5
}
```

### Response

```json
{
  "data": {
    "uuid": "b749d9de-8135-4a3c-a92e-6d675c199ccc",
    "name": "S1 TEKNIK KOMPUTER",
    "location": "ADANYA DI KAMPUS MULTMIEDIA NUDSNATRAS",
    "start_date": "2020-08-01",
    "end_date": "2022-12-15",
    "grade": 24.5,
    "user": {
      "uuid": "123e4567-e89b-12d3-a456-426614174000",
      "fullname": "John Doe",
      "email": "john@example.com"
    }
  },
  "message": "User education updated successfully",
  "status": "success"
}
```

---

## ❌ Delete User Education

**DELETE** `/delete/:uuid`

Example: `/delete/69fe58e9-3e69-49ee-8dee-a0a249219b16`

### Response

```json
{
  "data": {
    "uuid": "69fe58e9-3e69-49ee-8dee-a0a249219b16",
    "name": "Master of Data Science",
    "location": "Institute of Advanced Studies",
    "start_date": "2020-08-01",
    "end_date": "2022-12-15",
    "grade": 92.5,
    "user": {
      "uuid": "123e4567-e89b-12d3-a456-426614174000",
      "fullname": "John Doe",
      "email": "john@example.com"
    }
  },
  "message": "User education deleted successfully",
  "status": "success"
}
```

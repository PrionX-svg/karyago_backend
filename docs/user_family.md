

# 👪 HRIS - User Family Management API

Documentation for **User Family** management endpoints in the HRIS system, including create, retrieve, update, and delete operations.

---

## 📌 Base URL

```
http://localhost:8080/api/v1/user-families
```

---

## ➕ Create User Family

**POST** `/create`

Create a new family record for a user.

### Request Body

```json
{
  "user_uuid": "265edc19-3544-45c8-a887-a7ca5d67455d",
  "status": true,
  "count": 3,
  "member": ["Wife", "Son", "Daughter"]
}
```

### Response

```json
{
  "data": {
    "uuid": "generated-family-uuid",
    "user_uuid": "265edc19-3544-45c8-a887-a7ca5d67455d",
    "status": true,
    "count": 3,
    "member": ["Wife", "Son", "Daughter"]
  },
  "message": "User family created",
  "status": "success"
}
```

---

## 📥 Get All User Families

**GET** `/get-all`

Retrieve all user family records.

> No request body required.

### Response

```json
{
  "data": [
    {
      "uuid": "f7e1bbf7-e350-42ce-b80f-19858ba2707e",
      "user_uuid": "265edc19-3544-45c8-a887-a7ca5d67455d",
      "status": true,
      "count": 3,
      "member": ["Wife", "Son", "Daughter"]
    }
  ],
  "message": "All user families retrieved",
  "status": "success"
}
```

---

## 🔍 Get User Family by ID

**GET** `/get-by-id/:id`
Example: `/get-by-id/2`

Retrieve a family record by internal ID.

### Response

```json
{
  "data": {
    "uuid": "f7e1bbf7-e350-42ce-b80f-19858ba2707e",
    "user_uuid": "265edc19-3544-45c8-a887-a7ca5d67455d",
    "status": true,
    "count": 3,
    "member": ["Wife", "Son", "Daughter"]
  },
  "message": "User family retrieved",
  "status": "success"
}
```

---

## 🔎 Get User Family by UUID

**GET** `/get-by-uuid/:uuid`
Example: `/get-by-uuid/f7e1bbf7-e350-42ce-b80f-19858ba2707e`

Retrieve a family record by its UUID.

### Response

```json
{
  "data": {
    "uuid": "f7e1bbf7-e350-42ce-b80f-19858ba2707e",
    "user_uuid": "265edc19-3544-45c8-a887-a7ca5d67455d",
    "status": true,
    "count": 3,
    "member": ["Wife", "Son", "Daughter"]
  },
  "message": "User family retrieved",
  "status": "success"
}
```

---

## 👤 Get User Family by User UUID

**GET** `/get-by-user/:user_uuid`
Example: `/get-by-user/91aa937c-740b-4df7-a032-ed0335b424a1`

Retrieve a user's full family group by the user's UUID.

### Response

```json
{
  "data": {
    "uuid": "user-family-uuid",
    "status": false,
    "count": 1,
    "members": [
      {
        "name": "Carol Smith",
        "relation": "Parent",
        "phone": "345-678-9012"
      }
    ],
    "user": {
      "uuid": "91aa937c-740b-4df7-a032-ed0335b424a1",
      "fullname": "Full Name",
      "email": "user@example.com"
    }
  },
  "message": "User family retrieved",
  "status": "success"
}
```

---

## ✏️ Update User Family

**PATCH** `/update/:user_uuid`
Example: `/update/91aa937c-740b-4df7-a032-ed0335b424a1`

Update the user's family data.

### Request Body

```json
{
  "user_uuid": "91aa937c-740b-4df7-a032-ed0335b424a1",
  "status": false,
  "members": [
    {
      "name": "Carol Smith",
      "relation": "Parent",
      "phone": "345-678-9012"
    }
  ]
}
```

### Response

```json
{
  "data": {
    "uuid": "updated-family-uuid",
    "user_uuid": "91aa937c-740b-4df7-a032-ed0335b424a1",
    "status": false,
    "members": [
      {
        "name": "Carol Smith",
        "relation": "Parent",
        "phone": "345-678-9012"
      }
    ]
  },
  "message": "User family updated",
  "status": "success"
}
```

---

## ❌ Delete User Family

**DELETE** `/delete/:uuid`
Example: `/delete/f7e1bbf7-e350-42ce-b80f-19858ba2707e`

Delete a specific user family record by its UUID.

### Response

```json
{
  "data": {
    "uuid": "f7e1bbf7-e350-42ce-b80f-19858ba2707e"
  },
  "message": "User family deleted",
  "status": "success"
}
```

---

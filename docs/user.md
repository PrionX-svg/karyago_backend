# 📘 HRIS - User API

Documentation for **User** management endpoints in the HRIS system, including create, retrieve, update,
and delete operations.

---

## 📌 Base URL

```

http://localhost:8080/api/v1/users

````

---

## ➕ Create User

**POST** `/create`

Creates a new user with employee details.

### Request Body

```json
{
  "role_uuid": "d96bd928-51e3-43a5-a5d1-413d7ce3edc1",
  "company_uuid": "5e143995-557e-45b0-8822-691321c62a62",
  "firstname": "John",
  "lastname": "Doe",
  "phone": "628123456789",
  "email": "john.doe2@example.com",
  "password": "securepassword123",
  "dob": "1990-05-15T00:00:00Z",
  "gender": "male",
  "is_freelance": false
}
````

### Response

```json
{
  "message": "User created successfully"
}
```

---

## 📋 Get All Users with Employee

**GET** `/get-with-emp`

Retrieve all users with their employee information.

### Response

```json
{
  "data": [
    {
      "user_uuid": "85b47132-345d-40a1-868f-3e48adacb319",
      "full_name": "Marco Melandri",
      "email": "marcomelandri808@gmail.com",
      "phone": "6287782553442",
      "gender": null,
      "dob": null,
      "is_freelance": false,
      "role": "owner",
      "branch": {
        "uuid": "",
        "name": ""
      }
    },
    {
      "user_uuid": "a008b8ef-fefa-4f5d-ba07-63b37679e64a",
      "full_name": "John Doe",
      "email": "john.doe@example.com",
      "phone": "628123456789",
      "gender": "male",
      "dob": "1990-05-15T00:00:00Z",
      "is_freelance": false,
      "role": "assistant",
      "branch": {
        "uuid": "",
        "name": ""
      }
    },
    {
      "user_uuid": "5a8f4d27-cafd-43f1-90b4-fc03a96564aa",
      "full_name": "John Doe",
      "email": "john.doe2@example.com",
      "phone": "628123456789",
      "gender": "male",
      "dob": "1990-05-15T00:00:00Z",
      "is_freelance": false,
      "role": "assistant",
      "branch": {
        "uuid": "",
        "name": ""
      }
    }
  ]
}
```

---

## 📊 Get All Users (DataTable)

**GET** `/get-all/dt`

Supports pagination, search, and filtering by role or branch.

### Query Parameters

| Param           | Type    | Description                      |
|-----------------|---------|----------------------------------|
| `page`          | int     | Page number                      |
| `limit`         | int     | Items per page                   |
| `search`        | string  | Search term (optional)           |
| `role_uuid`     | string  | Filter by role UUID (optional)   |
| `branch_uuid`   | string  | Filter by branch UUID (optional) |
| `is_terminated` | boolean | Filter by branch UUID (optional) |

### Example

```
/get-all/dt?page=1&limit=10&role_uuid=1d11dc07-7eef-4165-9599-6eaf8f5268eb&branch_uuid=053f1a3b-9082-4d08-9ccd-52de56c2fdab&is_terminated=
```

### Response

```json
{
  "data": [
    {
      "user_uuid": "5a8f4d27-cafd-43f1-90b4-fc03a96564aa",
      "full_name": "John Doe",
      "email": "john.doe2@example.com",
      "phone": "628123456789",
      "gender": "male",
      "dob": "1990-05-15T00:00:00Z",
      "is_freelance": false,
      "role": "assistant",
      "branch": {
        "uuid": "",
        "name": ""
      },
      "termination": {
        "reason": "termination",
        "date": "2025-07-15T02:35:53.204Z"
      }
    },
    {
      "user_uuid": "a008b8ef-fefa-4f5d-ba07-63b37679e64a",
      "full_name": "John Doe",
      "email": "john.doe@example.com",
      "phone": "628123456789",
      "gender": "male",
      "dob": "1990-05-15T00:00:00Z",
      "is_freelance": false,
      "role": "assistant",
      "branch": {
        "uuid": "",
        "name": ""
      }
    },
    {
      "user_uuid": "85b47132-345d-40a1-868f-3e48adacb319",
      "full_name": "Marco Melandri",
      "email": "marcomelandri808@gmail.com",
      "phone": "6287782553442",
      "gender": null,
      "dob": null,
      "is_freelance": false,
      "role": "owner",
      "branch": {
        "uuid": "",
        "name": ""
      }
    }
  ],
  "filtered": 3,
  "limit": 10,
  "page": 1,
  "total": 3
}
```

---

## 🔍 Get User by UUID

**GET** `/get/:uuid`

Retrieve a single user and employee data by UUID.

### Example

```
/get/e681f5e4-a699-435f-a298-61a1e755b81f
```

### Response

```json
{
  "data": {
    "uuid": "e681f5e4-a699-435f-a298-61a1e755b81f",
    "fullname": "John Doe",
    "email": "john.doe@example.com",
    "dob": "1990-05-15T00:00:00Z",
    "gender": "male",
    "is_freelance": false
  },
  "message": "User retrieved successfully",
  "status": "success"
}
```

---

## ✏️ Update User

**PATCH** `/update/:uuid`

Update user/employee fields such as DOB, gender, and freelance status.

### Example

```
/update/f4d811c2-8583-44f0-ab4c-b2cc1f483387
```

### Request Body

```json
{
  "dob": "2003-01-15T00:00:00Z",
  "gender": "male",
  "is_freelance": false
}
```

### Response

```json
{
  "message": "User updated successfully"
}
```

---

## ❌ Delete User

**DELETE** `/delete/:uuid`

Delete a user and related employee data.

### Example

```
/delete/e681f5e4-a699-435f-a298-61a1e755b81f
```

### Response

```json
{
  "message": "User deleted successfully"
}
```

---

## 👤 Get Current User

**GET** `/@me`

Retrieve the currently authenticated user's profile.

### Response

```json
{
  "data": {
    "user_uuid": "85b47132-345d-40a1-868f-3e48adacb319",
    "full_name": "Marco Melandri",
    "email": "marcomelandri808@gmail.com",
    "phone": "087782553442",
    "gender": "male",
    "dob": "2003-01-15T00:00:00Z",
    "is_freelance": false,
    "role": "assistant",
    "branch": {
      "uuid": "",
      "name": ""
    }
  }
}
```

---

## ♻️ Rehire User

**PATCH** `/rehire/:uuid?company_uuid=...`

Reactivate a user that was previously deleted (terminated).

### Request Body

```json
{
  "role_uuid": "d96bd928-51e3-43a5-a5d1-413d7ce3edc1",
  "is_freelance": false
}
```

### Response

```json
{
  "message": "employee rehired successfully"
}
```

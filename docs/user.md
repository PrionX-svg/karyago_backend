# 📘 HRIS - User API

Documentation for **User** management endpoints in the HRIS system, including create, retrieve, update,
and delete operations.

---

## 📌 Base URL

```

http://localhost:8080/api/v1/users

```

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
```

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
      "user_uuid": "04262c91-6268-44cb-bd86-0bcf51da849d",
      "employee_uuid": "ee39d68e-d610-42b2-9378-dd79573d3185",
      "first_name": "Vincentius Marco",
      "last_name": "Melandri",
      "full_name": "Vincentius Marco Melandri",
      "email": "marcomelandri808@gmail.com",
      "phone": "6287782553442",
      "gender": null,
      "dob": null,
      "is_freelance": false,
      "role": {
        "uuid": "db5f8051-26c0-4be7-820f-5d4910325a87",
        "name": "owner"
      },
      "branch": {
        "uuid": "",
        "name": ""
      },
      "department": {
        "uuid": "1db2aec0-8050-4d89-a9ea-7ce2608d11f6",
        "name": "HR"
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

| Param           | Type    | Description                       |
| --------------- | ------- | --------------------------------- |
| `page`          | int     | Page number                       |
| `limit`         | int     | Items per page                    |
| `search`        | string  | Search term (optional)            |
| `role_uuid`     | string  | Filter by role UUID (optional)    |
| `branch_uuid`   | string  | Filter by branch UUID (optional)  |
| `company_uuid`  | string  | Filter by company UUID (optional) |
| `is_terminated` | boolean | Filter by branch UUID (optional)  |

### Example

```
/get-all/dt?page=1&limit=10&role_uuid=1d11dc07-7eef-4165-9599-6eaf8f5268eb&branch_uuid=053f1a3b-9082-4d08-9ccd-52de56c2fdab&is_terminated=
```

### Response

```json
{
  "data": [
    {
      "user_uuid": "04262c91-6268-44cb-bd86-0bcf51da849d",
      "employee_uuid": "ee39d68e-d610-42b2-9378-dd79573d3185",
      "first_name": "Vincentius Marco",
      "last_name": "Melandri",
      "full_name": "Vincentius Marco Melandri",
      "email": "marcomelandri808@gmail.com",
      "phone": "6287782553442",
      "gender": null,
      "dob": null,
      "is_freelance": false,
      "role": {
        "uuid": "db5f8051-26c0-4be7-820f-5d4910325a87",
        "name": "owner"
      },
      "branch": {
        "uuid": "",
        "name": ""
      },
      "company": {
        "uuid": "a11920c8-f9d8-414a-a0dc-910b1d29da03",
        "name": "PT Teknologi Nusantara 2"
      },
      "department": {
        "uuid": "1db2aec0-8050-4d89-a9ea-7ce2608d11f6",
        "name": "HR"
      }
    }
  ],
  "filtered": 1,
  "limit": 10,
  "page": 1,
  "total": 1
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
    "user_uuid": "04262c91-6268-44cb-bd86-0bcf51da849d",
    "employee_uuid": "ee39d68e-d610-42b2-9378-dd79573d3185",
    "first_name": "Vincentius Marco",
    "last_name": "Melandri",
    "full_name": "Vincentius Marco Melandri",
    "email": "marcomelandri808@gmail.com",
    "phone": "6287782553442",
    "gender": null,
    "dob": null,
    "is_freelance": false,
    "role": {
      "uuid": "db5f8051-26c0-4be7-820f-5d4910325a87",
      "name": "owner"
    },
    "branch": {
      "uuid": "",
      "name": ""
    },
    "department": {
      "uuid": "1db2aec0-8050-4d89-a9ea-7ce2608d11f6",
      "name": "HR"
    }
  }
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

---

## 📤 Export Users to Excel

**GET** `/export`

Generates an Excel file containing all users and their employee data.

### Query Parameters (optional)

| Parameter      | Type   | Description       |
| -------------- | ------ | ----------------- |
| `company_uuid` | string | Filter by company |
| `branch_uuid`  | string | Filter by branch  |
| `role_uuid`    | string | Filter by role    |

### Example

```
/export?company_uuid=1234-uuid&branch_uuid=5678-uuid
```

### Response

- Downloads an `.xlsx` file.
- Sheet name: `Users`

### Excel Columns

| Column       | Description                                       |
| ------------ | ------------------------------------------------- |
| First Name   | First name of the user                            |
| Last Name    | Last name of the user                             |
| Email        | Email address                                     |
| Phone        | Phone number                                      |
| Gender       | Gender (`male`, `female`, etc.)                   |
| DOB          | Date of Birth (formatted `YYYY-MM-DD`)            |
| Is Freelance | Whether the user is freelance (`true` or `false`) |
| Role         | Role name                                         |
| Branch       | Branch name                                       |

---

## 📥 Import Users from Excel

**POST** `/import`

Imports users in bulk from an Excel (`.xlsx`) file.

### Request

- **Content-Type:** `multipart/form-data`
- **Form Field:** `file` (Excel file)

### Excel Format (Sheet: `Users`)

| First Name | Last Name | Email                                               | Phone        | Gender | DOB (`YYYY-MM-DD`) | Is Freelance | Role      | Branch  |
| ---------- | --------- | --------------------------------------------------- | ------------ | ------ | ------------------ | ------------ | --------- | ------- |
| John       | Doe       | [john.doe@example.com](mailto:john.doe@example.com) | 628123456789 | male   | 1990-01-01         | false        | assistant | Jakarta |

### Import Rules

- The **first row** is treated as the **header** and will be skipped.
- If a `role` or `branch` name cannot be found → that row will be **skipped**.
- If the `email` already exists → the user is considered a duplicate and the row will be **skipped**.
- The default password for all imported users is set to `"default123"` (should be changed by users later).

### Response

```json
{
  "message": "Import completed successfully"
}
```

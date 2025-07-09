# 🧾 HRIS - Branch Management API

Documentation for **Branch** management endpoints in the HRIS system, including create, retrieve, update, and delete
operations.

---

## 📌 Base URL

```

http://localhost:8080/api/v1/branches

````

---

## ➕ Create Branch

**POST** `/create`

Create a new branch.

### Request Body

```json
{
  "company_uuid": "797a93d6-c780-4847-8839-f2220bd4a4c2",
  "name": "Cabang Jakarta",
  "address": "Jl. Sudirman No. 100, Jakarta",
  "email": "jakarta.branch@company.com",
  "phone": "021-98765432"
}
````

### Response

```json
{
  "data": {
    "uuid": "generated-branch-uuid",
    "company_uuid": "797a93d6-c780-4847-8839-f2220bd4a4c2",
    "name": "Cabang Jakarta",
    "address": "Jl. Sudirman No. 100, Jakarta",
    "email": "jakarta.branch@company.com",
    "phone": "021-98765432"
  },
  "message": "Branch created successfully",
  "status": "success"
}
```

---

## 📋 Get All Branches

**GET** `/get-all`

Retrieve the list of all branches.

### Response

```json
{
  "data": [
    {
      "uuid": "8ec51987-9375-4151-9b3b-b9d2467669f8",
      "company_uuid": "797a93d6-c780-4847-8839-f2220bd4a4c2",
      "name": "Cabang Jakarta",
      "address": "Jl. Sudirman No. 100, Jakarta",
      "email": "jakarta.branch@company.com",
      "phone": "021-98765432"
    }
  ],
  "message": "Branches fetched",
  "status": "success"
}
```

---

## 🔍 Get Branch By UUID

**GET** `/get/:uuid`
Example: `/get/8ec51987-9375-4151-9b3b-b9d2467669f8`

### Response

```json
{
  "data": {
    "uuid": "8ec51987-9375-4151-9b3b-b9d2467669f8",
    "company_uuid": "797a93d6-c780-4847-8839-f2220bd4a4c2",
    "name": "Cabang Jakarta",
    "address": "Jl. Sudirman No. 100, Jakarta",
    "email": "jakarta.branch@company.com",
    "phone": "021-98765432"
  },
  "message": "Branch found",
  "status": "success"
}
```

---

## ✏️ Update Branch

**PATCH** `/update/:uuid`
Example: `/update/8ec51987-9375-4151-9b3b-b9d2467669f8`

### Request Body

```json
{
  "company_uuid": "797a93d6-c780-4847-8839-f2220bd4a4c2",
  "name": "Cabang Tangerang",
  "address": "Jl. Tangerang No. 100, Tangerang",
  "email": "jakarta.branch@company.com",
  "phone": "021-98765432"
}
```

### Response

```json
{
  "data": {
    "uuid": "8ec51987-9375-4151-9b3b-b9d2467669f8",
    "company_uuid": "797a93d6-c780-4847-8839-f2220bd4a4c2",
    "name": "Cabang Tangerang",
    "address": "Jl. Tangerang No. 100, Tangerang",
    "email": "jakarta.branch@company.com",
    "phone": "021-98765432"
  },
  "message": "Branch updated successfully",
  "status": "success"
}
```

---

## ❌ Delete Branch

**DELETE** `/delete/:uuid`
Example: `/delete/8ec51987-9375-4151-9b3b-b9d2467669f8`

### Response

```json
{
  "data": {
    "uuid": "8ec51987-9375-4151-9b3b-b9d2467669f8",
    "company_uuid": "797a93d6-c780-4847-8839-f2220bd4a4c2",
    "name": "Cabang Tangerang",
    "address": "Jl. Tangerang No. 100, Tangerang",
    "email": "jakarta.branch@company.com",
    "phone": "021-98765432"
  },
  "message": "Branch deleted successfully",
  "status": "success"
}
```

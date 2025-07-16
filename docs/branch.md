# 🧾 HRIS - Branch Management API

Documentation for **Branch** management endpoints in the HRIS system, including create, retrieve, update, and delete
operations.

---

## 📌 Base URL

```

http://localhost:8080/api/v1/branches

````

---****

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
    "uuid": "a3994ce4-cb47-4a21-83f3-78b4dbc9c5fd",
    "name": "Cabang Surabaya",
    "address": "Jl. Sudirman No. 100, Surabaya",
    "email": "Surabaya.branch@company.com",
    "phone": "021-98765432",
    "company": {
      "uuid": "026e99f5-a16c-488d-932f-ce332b60bc52",
      "logo": "/uuid-date-namafile",
      "name": "PT Teknologi Nusantara 2",
      "address": "Jl. Merdeka No.123, Jakarta",
      "email": "info2@teknologi.co.id",
      "phone": "021-12345678"
    }
  },
  "message": "Branch created",
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
      "uuid": "a3994ce4-cb47-4a21-83f3-78b4dbc9c5fd",
      "name": "Cabang Surabaya",
      "address": "Jl. Sudirman No. 100, Surabaya",
      "email": "Surabaya.branch@company.com",
      "phone": "021-98765432",
      "company": {
        "uuid": "026e99f5-a16c-488d-932f-ce332b60bc52",
        "logo": "/uuid-date-namafile",
        "name": "PT Teknologi Nusantara 2",
        "address": "Jl. Merdeka No.123, Jakarta",
        "email": "info2@teknologi.co.id",
        "phone": "021-12345678"
      }
    }
  ],
  "message": "Successfully get all branches",
  "status": "success"
}
```

---

## 🔍 Get Branches By Company UUID


**GET** `/get-by-company-uuid/:uuid`
Example: `/get-by-company-uuid/8b21610d-8d24-4510-a324-77760a344242`

Retrieve the list of all branches by company UUID.

### Response

```json
{
    "data": [
        {
            "uuid": "5ea0ce42-9d90-4e26-a484-94d881f12ea5",
            "name": "Cabang Jakarta",
            "address": "Jl. Sudirman No. 100, Jakarta",
            "email": "jakarta.branch@company.com",
            "phone": "021-98765432",
            "company": {
                "uuid": "8b21610d-8d24-4510-a324-77760a344242",
                "logo": "",
                "name": "Company Name",
                "address": "Company Address",
                "email": "company@example.com",
                "phone": "+1 (555) 123-4567"
            }
        },
        {
            "uuid": "f977c833-bcda-42f4-a769-c2fcd9e943e9",
            "name": "Cabang Jakarta Kedua",
            "address": "Jl. Malioboro No. 100, Yogyakarta",
            "email": "jakarta.2ndbranch@company.com",
            "phone": "021-98765432",
            "company": {
                "uuid": "8b21610d-8d24-4510-a324-77760a344242",
                "logo": "",
                "name": "Company Name",
                "address": "Company Address",
                "email": "company@example.com",
                "phone": "+1 (555) 123-4567"
            }
        }
    ],
    "message": "Successfully get branch by company UUID",
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
    "uuid": "a3994ce4-cb47-4a21-83f3-78b4dbc9c5fd",
    "name": "Cabang Surabaya",
    "address": "Jl. Sudirman No. 100, Surabaya",
    "email": "Surabaya.branch@company.com",
    "phone": "021-98765432",
    "company": {
      "uuid": "026e99f5-a16c-488d-932f-ce332b60bc52",
      "logo": "/uuid-date-namafile",
      "name": "PT Teknologi Nusantara 2",
      "address": "Jl. Merdeka No.123, Jakarta",
      "email": "info2@teknologi.co.id",
      "phone": "021-12345678"
    }
  },
  "message": "Successfully get branch",
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
    "uuid": "a3994ce4-cb47-4a21-83f3-78b4dbc9c5fd",
    "name": "Cabang Tangerang",
    "address": "Jl. Tangerang No. 100, Tangerang",
    "email": "tangerang.branch@company.com",
    "phone": "021-98765432",
    "company": {
      "uuid": "026e99f5-a16c-488d-932f-ce332b60bc52",
      "logo": "/uuid-date-namafile",
      "name": "PT Teknologi Nusantara 2",
      "address": "Jl. Merdeka No.123, Jakarta",
      "email": "info2@teknologi.co.id",
      "phone": "021-12345678"
    }
  },
  "message": "Branch updated",
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
    "uuid": "a3994ce4-cb47-4a21-83f3-78b4dbc9c5fd",
    "name": "Cabang Tangerang",
    "address": "Jl. Tangerang No. 100, Tangerang",
    "email": "tangerang.branch@company.com",
    "phone": "021-98765432",
    "company": {
      "uuid": "026e99f5-a16c-488d-932f-ce332b60bc52",
      "logo": "/uuid-date-namafile",
      "name": "PT Teknologi Nusantara 2",
      "address": "Jl. Merdeka No.123, Jakarta",
      "email": "info2@teknologi.co.id",
      "phone": "021-12345678"
    }
  },
  "message": "Branch deleted",
  "status": "success"
}
```

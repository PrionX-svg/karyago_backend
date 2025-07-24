# 🏢 HRIS - Department Management API

Documentation for **Department** management endpoints in the HRIS system, including create, retrieve, update, and delete
operations.

---

## 📌 Base URL

```
http://localhost:8080/api/v1/departments
```

---

## ➕ Create Department

**POST** `/create`

Create a new department.

### Request Body

```json
{
  "department_group_uuid": "0ad96968-1f92-4199-87b4-d22b5c725023",
  "name": "Pajak",
  "description": "Handles tax, employee relations, and organizational development."
}
```

### Response

```json
{
  "data": {
    "uuid": "aa0a22ae-63cd-4809-b9e5-161604aeceea",
    "name": "HR",
    "description": "Handles tax, employee relations, and organizational development.",
    "department_group": {
      "uuid": "73022eb0-53aa-41bd-8c8b-f583514b9248",
      "name": "Departemen Finance"
    },
    "employees": []
  },
  "message": "Department created",
  "status": "success"
}
```

---

## 📋 Get All Departments

**GET** `/get-all`

Retrieve the list of all departments.

### Response

```json
{
  "data": [
    {
      "uuid": "8991f837-958c-452c-8e5a-24cc13881a80",
      "name": "Pajak",
      "description": "Handles tax, employee relations, and organizational development.",
      "department_group": {
        "uuid": "73022eb0-53aa-41bd-8c8b-f583514b9248",
        "name": "Departemen Finance"
      },
      "employees": [
        {
          "uuid": "04262c91-6268-44cb-bd86-0bcf51da849d",
          "name": "Vincentius Marco Melandri",
          "email": "marcomelandri808@gmail.com"
        }
      ]
    },
    {
      "uuid": "aa0a22ae-63cd-4809-b9e5-161604aeceea",
      "name": "HR",
      "description": "Handles tax, employee relations, and organizational development.",
      "department_group": {
        "uuid": "73022eb0-53aa-41bd-8c8b-f583514b9248",
        "name": "Departemen Finance"
      },
      "employees": null
    }
  ],
  "message": "Successfully get all departments",
  "status": "success"
}
```

---

## 🔍 Get Department By UUID

**GET** `/get/:uuid`
Example: `/get/f3f70505-8de6-42cf-bffd-6ac94fb3459e`

### Response

```json
{
  "data": {
    "uuid": "8991f837-958c-452c-8e5a-24cc13881a80",
    "name": "Pajak",
    "description": "Handles tax, employee relations, and organizational development.",
    "department_group": {
      "uuid": "73022eb0-53aa-41bd-8c8b-f583514b9248",
      "name": "Departemen Finance"
    },
    "employees": [
      {
        "uuid": "04262c91-6268-44cb-bd86-0bcf51da849d",
        "name": "Vincentius Marco Melandri",
        "email": "marcomelandri808@gmail.com"
      }
    ]
  },
  "message": "Successfully get department",
  "status": "success"
}
```

---

## ✏️ Update Department

**PATCH** `/update/:uuid`
Example: `/update/f3f70505-8de6-42cf-bffd-6ac94fb3459e`

### Request Body

```json
{
  "department_group_uuid": "0ad96968-1f92-4199-87b4-d22b5c725023",
  "name": "Pajak Update",
  "description": "Handles tax and financial planning."
}
```

### Response

```json
{
  "data": {
    "uuid": "aa0a22ae-63cd-4809-b9e5-161604aeceea",
    "name": "HR",
    "description": "Handles tax, employee relations, and organizational development.",
    "department_group": {
      "uuid": "73022eb0-53aa-41bd-8c8b-f583514b9248",
      "name": "Departemen Finance"
    },
    "employees": []
  },
  "message": "Department updated",
  "status": "success"
}
```

---

## ❌ Delete Department

**DELETE** `/delete/:uuid`
Example: `/delete/f3f70505-8de6-42cf-bffd-6ac94fb3459e`

### Response

```json
{
  "data": {
    "uuid": "aa0a22ae-63cd-4809-b9e5-161604aeceea",
    "name": "HR",
    "description": "Handles tax, employee relations, and organizational development.",
    "department_group": {
      "uuid": "73022eb0-53aa-41bd-8c8b-f583514b9248",
      "name": "Departemen Finance"
    },
    "employees": []
  },
  "message": "Department deleted",
  "status": "success"
}
```

---

## 📄 Get All Departments (Datatable)

**GET** `/get-all/dt`

Retrieve a paginated list of departments with optional search and filtering.

### Query Parameters

| Parameter               | Type   | Required | Description                                 |
|-------------------------|--------|----------|---------------------------------------------|
| `page`                  | int    | No       | Page number (default: 1)                    |
| `limit`                 | int    | No       | Items per page (default: 10)                |
| `search`                | string | No       | Search keyword (applies to department name) |
| `company_uuid`          | string | **Yes**  | UUID of the company                         |
| `department_group_uuid` | string | No       | UUID of the department group to filter      |

### Example Request

```
GET /api/v1/departments/get-all/dt?page=1&limit=10&search=pajak&company_uuid=86a9688c-53b4-42c8-b3c1-819fd6ad1d9f&department_group_uuid=0ad96968-1f92-4199-87b4-d22b5c725023
```

### Response

```json
{
  "data": [
    {
      "uuid": "aa0a22ae-63cd-4809-b9e5-161604aeceea",
      "name": "HR",
      "description": "Handles tax, employee relations, and organizational development.",
      "department_group": {
        "uuid": "73022eb0-53aa-41bd-8c8b-f583514b9248",
        "name": "Departemen Finance"
      },
      "employees": null
    },
    {
      "uuid": "8991f837-958c-452c-8e5a-24cc13881a80",
      "name": "Pajak",
      "description": "Handles tax, employee relations, and organizational development.",
      "department_group": {
        "uuid": "73022eb0-53aa-41bd-8c8b-f583514b9248",
        "name": "Departemen Finance"
      },
      "employees": [
        {
          "uuid": "04262c91-6268-44cb-bd86-0bcf51da849d",
          "name": "Vincentius Marco Melandri",
          "email": "marcomelandri808@gmail.com"
        }
      ]
    }
  ],
  "filtered": 2,
  "limit": 10,
  "page": 1,
  "total": 2
}
```
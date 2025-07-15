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
    "uuid": "generated-department-uuid",
    "department_group_uuid": "0ad96968-1f92-4199-87b4-d22b5c725023",
    "name": "Pajak",
    "description": "Handles tax, employee relations, and organizational development."
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
      "uuid": "f3f70505-8de6-42cf-bffd-6ac94fb3459e",
      "department_group_uuid": "0ad96968-1f92-4199-87b4-d22b5c725023",
      "name": "Pajak",
      "description": "Handles tax, employee relations, and organizational development."
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
    "uuid": "f3f70505-8de6-42cf-bffd-6ac94fb3459e",
    "department_group_uuid": "0ad96968-1f92-4199-87b4-d22b5c725023",
    "name": "Pajak",
    "description": "Handles tax, employee relations, and organizational development."
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
    "uuid": "f3f70505-8de6-42cf-bffd-6ac94fb3459e",
    "department_group_uuid": "0ad96968-1f92-4199-87b4-d22b5c725023",
    "name": "Pajak Update",
    "description": "Handles tax and financial planning."
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
    "uuid": "f3f70505-8de6-42cf-bffd-6ac94fb3459e",
    "department_group_uuid": "0ad96968-1f92-4199-87b4-d22b5c725023",
    "name": "Pajak",
    "description": "Handles tax, employee relations, and organizational development."
  },
  "message": "Department deleted",
  "status": "success"
}
```
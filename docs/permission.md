# 🧾 HRIS - Permission Management API

This section documents the **Permission** management endpoints used to define, retrieve, update, and delete permissions within the HRIS system.

---

## 📌 Base URL

http://localhost:8080/api/v1/permissions

---

## ➕ Create Permission

**POST** `/create`

Create a new permission that can be assigned to roles.

### Request Body

```json
{
  "label": "role.create",
  "name": "Create Role"
}
````

### Response

```json
{
  "data": {
    "uuid": "8a057fc0-a321-4c20-8b05-c15a9d1051b6",
    "label": "role.create",
    "name": "Create Role"
  },
  "message": "Permission created successfully",
  "status": "success"
}
```

---

## 📋 Get All Permissions

**GET** `/get-all`

Fetch all permissions in the system.

### Response

```json
{
  "data": [
    {
      "uuid": "8a057fc0-a321-4c20-8b05-c15a9d1051b6",
      "label": "role.create",
      "name": "Create Role"
    },
    {
      "uuid": "d0f6bc3b-45dc-4f47-8fd7-1e223917a437",
      "label": "permission.view",
      "name": "View Permission"
    }
  ],
  "message": "Permissions fetched",
  "status": "success"
}
```

---

## 🔍 Get Permission By UUID

**GET** `/get/:uuid`
Example: `/get/8a057fc0-a321-4c20-8b05-c15a9d1051b6`

### Response

```json
{
  "data": {
    "uuid": "8a057fc0-a321-4c20-8b05-c15a9d1051b6",
    "label": "role.create",
    "name": "Create Role"
  },
  "message": "Permission found",
  "status": "success"
}
```

---

## ✏️ Update Permission

**PATCH** `/update/:uuid`
Example: `/update/8a057fc0-a321-4c20-8b05-c15a9d1051b6`

### Request Body

```json
{
  "name": "Create Role"
}
```

### Response

```json
{
  "data": {
    "uuid": "8a057fc0-a321-4c20-8b05-c15a9d1051b6",
    "label": "role.create",
    "name": "Create Role"
  },
  "message": "Permission updated successfully",
  "status": "success"
}
```

---

## ❌ Delete Permission

**DELETE** `/delete/:uuid`
Example: `/delete/8a057fc0-a321-4c20-8b05-c15a9d1051b6`

### Response

```json
{
  "data": {},
  "message": "Permission deleted successfully",
  "status": "success"
}
```
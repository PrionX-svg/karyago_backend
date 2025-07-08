# 🧾 HRIS - Role Management API

This section documents the **Role** management endpoints used for creating, retrieving, updating, and deleting user roles in the HRIS system.

---

## 📌 Base URL

http://localhost:8080/api/v1/roles

---

## ➕ Create Role

**POST** `/create`

Create a new role.

### Request Body

```json
{
  "name": "Staff"
}
````

### Response

```json
{
  "data": {
    "uuid": "e4acb9d9-f0a4-4135-a7be-08367989d4eb",
    "name": "Staff"
  },
  "message": "Role created successfully",
  "status": "success"
}
```

---

## 🔍 Get Role By Name

**GET** `/get-by-name?name=staff`

Fetch role detail by its name.

### Response

```json
{
  "data": {
    "uuid": "e4acb9d9-f0a4-4135-a7be-08367989d4eb",
    "name": "Staff"
  },
  "message": "Role found",
  "status": "success"
}
```

---

## 🔍 Get Role By UUID

**GET** `/get/:uuid`
Example: `/get/e4acb9d9-f0a4-4135-a7be-08367989d4eb`

### Response

```json
{
  "data": {
    "uuid": "e4acb9d9-f0a4-4135-a7be-08367989d4eb",
    "name": "Staff"
  },
  "message": "Role found",
  "status": "success"
}
```

---

## 📋 Get All Roles (Datatable)

**GET** `/get-all/dt`

Returns paginated/filterable list of roles for use in datatable.

### Response

```json
{
  "data": [
    {
      "uuid": "e4acb9d9-f0a4-4135-a7be-08367989d4eb",
      "name": "Staff"
    },
    {
      "uuid": "f1b62b17-d0ad-4896-a7ce-00206dfd2abc",
      "name": "Admin"
    }
  ],
  "message": "Roles fetched",
  "status": "success"
}
```

---

## ✏️ Update Role

**PATCH** `/update/:uuid`
Example: `/update/e4acb9d9-f0a4-4135-a7be-08367989d4eb`

### Request Body

```json
{
  "name": "Marketing"
}
```

### Response

```json
{
  "data": {
    "uuid": "e4acb9d9-f0a4-4135-a7be-08367989d4eb",
    "name": "Marketing"
  },
  "message": "Role updated successfully",
  "status": "success"
}
```

---

## ❌ Delete Role

**DELETE** `/delete/:uuid`
Example: `/delete/e4acb9d9-f0a4-4135-a7be-08367989d4eb`

### Response

```json
{
  "data": {},
  "message": "Role deleted successfully",
  "status": "success"
}
```

# 🛡️ HRIS - Role Permissions API

This section documents the **Role-Permission** management endpoints to assign and retrieve permissions related to
specific roles.

---

## 📌 Base URL

http://localhost:8080/api/v1/role-permissions

---

## 🔄 Sync Role Permissions

**POST** `/sync`

Assigns and removes permissions for a role in a single request.

### Request Body

```json
{
  "assign": [
    {
      "role_uuid": "984244a2-6731-498d-8fb1-85a87dde75ae",
      "permission_uuid": "e5066c23-abb3-4f8c-9926-e644cedc5e01"
    }
  ],
  "remove": [
    {
      "role_uuid": "984244a2-6731-498d-8fb1-85a87dde75ae",
      "permission_uuid": "e5066c23-abb3-4f8c-9926-e644cedc5e01"
    }
  ]
}
````

### Response

```json
{
  "data": {},
  "message": "Permissions assigned and removed successfully",
  "status": "success"
}
```

---

## 📋 Get Permissions By Role UUID

**GET** `/:role_uuid`
Example: `/8297f6f3-c4c7-4660-b4fa-148f56eee811`

Returns all permissions currently assigned to a specific role.

### Response

```json
{
  "data": {
    "role_name": "Staff",
    "permissions": [
      {
        "permission_name": "View-Name Role",
        "label": "role.view-name"
      },
      {
        "permission_name": "View-All Role",
        "label": "role.view-all"
      }
    ]
  },
  "message": "Success get permissions",
  "status": "success"
}
```
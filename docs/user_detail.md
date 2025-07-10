# 🧾 HRIS - User Detail Management API

Documentation for **User Detail** management endpoints in the HRIS system, including create, retrieve, update, and
delete operations.

---

## 📌 Base URL

```
http://localhost:8080/api/v1/user-details
```

---

## ➕ Create User Detail

**POST** `/create`

Create a new user detail associated with a specific user.

### Request Body

```json
{
  "user_uuid": "f4d811c2-8583-44f0-ab4c-b2cc1f483387",
  "tax_id": "01.234.567.8-901.000",
  "social_id": "3174091203990005"
}
```

### Response

```json
{
  "data": {
    "uuid": "generated-user-detail-uuid",
    "user_uuid": "f4d811c2-8583-44f0-ab4c-b2cc1f483387",
    "tax_id": "01.234.567.8-901.000",
    "social_id": "3174091203990005"
  },
  "message": "User detail created",
  "status": "success"
}
```

---

## 🔍 Get User Detail by User UUID

**GET** `/get/:uuid`
Example: `/get/f4d811c2-8583-44f0-ab4c-b2cc1f483387`

> This returns the user detail for a given **user UUID**, not the `user_detail.uuid`.

### Response

```json
{
  "data": {
    "uuid": "eabea8f4-dfe6-4922-8731-fbb3634f0b5c",
    "user_uuid": "f4d811c2-8583-44f0-ab4c-b2cc1f483387",
    "tax_id": "01.234.567.8-901.000",
    "social_id": "3174091203990005"
  },
  "message": "Successfully retrieved user detail",
  "status": "success"
}
```

---

## ✏️ Update User Detail

**PATCH** `/update/:uuid`
Example: `/update/eabea8f4-dfe6-4922-8731-fbb3634f0b5c`

### Request Body

```json
{
  "user_uuid": "f4d811c2-8583-44f0-ab4c-b2cc1f483387",
  "tax_id": "01.234.567.8-902.000",
  "social_id": "3174091203990003"
}
```

### Response

```json
{
  "data": {
    "uuid": "eabea8f4-dfe6-4922-8731-fbb3634f0b5c",
    "user_uuid": "f4d811c2-8583-44f0-ab4c-b2cc1f483387",
    "tax_id": "01.234.567.8-902.000",
    "social_id": "3174091203990003"
  },
  "message": "User detail updated",
  "status": "success"
}
```

---

## ❌ Delete User Detail

**DELETE** `/delete/:uuid`
Example: `/delete/eabea8f4-dfe6-4922-8731-fbb3634f0b5c`

### Response

```json
{
  "data": {
    "uuid": "eabea8f4-dfe6-4922-8731-fbb3634f0b5c",
    "user_uuid": "f4d811c2-8583-44f0-ab4c-b2cc1f483387",
    "tax_id": "01.234.567.8-902.000",
    "social_id": "3174091203990003"
  },
  "message": "User detail deleted",
  "status": "success"
}
```
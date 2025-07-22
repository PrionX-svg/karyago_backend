# 🧱 HRIS - Company Detail Shifts API

Documentation for **Company Detail Shifts** endpoints in the HRIS system, which manage working day configurations based
on the defined shift.

---

## 📌 Base URL

```
http://localhost:8080/api/v1/company-detail-shifts
```

---

## ➕ Create Company Detail Shift

**POST** `/create`

Create a working day configuration for a shift.

### Request Body

```json
{
  "shift_uuid": "8c88c76c-68a7-462c-8546-9ffeeebc57f3",
  "day_1": 1,
  "day_2": 1,
  "day_3": 1,
  "day_4": 1,
  "day_5": 1,
  "day_6": 0,
  "day_7": 0
}
```

### Response

```json
{
  "data": {
    "uuid": "generated-uuid",
    "shift_uuid": "8c88c76c-68a7-462c-8546-9ffeeebc57f3",
    "day_1": 1,
    "day_2": 1,
    "day_3": 1,
    "day_4": 1,
    "day_5": 1,
    "day_6": 0,
    "day_7": 0
  },
  "message": "Company detail shift created successfully",
  "status": "success"
}
```

---

## 📋 Get All Company Detail Shifts

**GET** `/get-all`

Retrieve all company detail shift configurations.

### Response

```json
{
  "data": [
    {
      "uuid": "560a34f8-0785-4312-ba6a-a2a657342310",
      "shift_uuid": "8c88c76c-68a7-462c-8546-9ffeeebc57f3",
      "day_1": 1,
      "day_2": 1,
      "day_3": 1,
      "day_4": 1,
      "day_5": 1,
      "day_6": 0,
      "day_7": 0
    }
  ],
  "message": "Successfully retrieved all company detail shifts",
  "status": "success"
}
```

---

## 🔍 Get Company Detail Shift By UUID

**GET** `/get/:uuid`
Example: `/get/560a34f8-0785-4312-ba6a-a2a657342310`

### Response

```json
{
  "data": {
    "uuid": "560a34f8-0785-4312-ba6a-a2a657342310",
    "shift_uuid": "8c88c76c-68a7-462c-8546-9ffeeebc57f3",
    "day_1": 1,
    "day_2": 1,
    "day_3": 1,
    "day_4": 1,
    "day_5": 1,
    "day_6": 0,
    "day_7": 0
  },
  "message": "Company detail shift retrieved successfully",
  "status": "success"
}
```

---

## 🔎 Get Company Detail Shift By Shift UUID

**GET** `/get-by-shift/:shift_uuid`
Example: `/get-by-shift/8c88c76c-68a7-462c-8546-9ffeeebc57f3`

### Response

```json
{
  "data": {
    "uuid": "560a34f8-0785-4312-ba6a-a2a657342310",
    "shift_uuid": "8c88c76c-68a7-462c-8546-9ffeeebc57f3",
    "day_1": 1,
    "day_2": 1,
    "day_3": 1,
    "day_4": 1,
    "day_5": 1,
    "day_6": 0,
    "day_7": 0
  },
  "message": "Company detail shift retrieved successfully",
  "status": "success"
}
```

---

## ✏️ Update Company Detail Shift

**PATCH** `/update/:uuid`
Example: `/update/560a34f8-0785-4312-ba6a-a2a657342310`

### Request Body

```json
{
  "shift_uuid": "8c88c76c-68a7-462c-8546-9ffeeebc57f3",
  "day_1": 1,
  "day_2": 0,
  "day_3": 1,
  "day_4": 0,
  "day_5": 1,
  "day_6": 0,
  "day_7": 0
}
```

### Response

```json
{
  "data": {
    "uuid": "560a34f8-0785-4312-ba6a-a2a657342310",
    "shift_uuid": "8c88c76c-68a7-462c-8546-9ffeeebc57f3",
    "day_1": 1,
    "day_2": 0,
    "day_3": 1,
    "day_4": 0,
    "day_5": 1,
    "day_6": 0,
    "day_7": 0
  },
  "message": "Company detail shift updated successfully",
  "status": "success"
}
```

---

## ❌ Delete Company Detail Shift

**DELETE** `/delete/:uuid`
Example: `/delete/560a34f8-0785-4312-ba6a-a2a657342310`

### Optional Request Body

```json
{
  "shift_uuid": "8c88c76c-68a7-462c-8546-9ffeeebc57f3",
  "day_1": 1,
  "day_2": 0,
  "day_3": 1,
  "day_4": 0,
  "day_5": 1,
  "day_6": 0,
  "day_7": 0
}
```

### Response

```json
{
  "message": "Company detail shift deleted successfully",
  "success": true
}
```
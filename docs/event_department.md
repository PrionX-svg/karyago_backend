# 🧾 HRIS - Event Department API

This documentation provides details on the available API endpoints for managing **Event Departments**, including create, retrieve, update, and delete operations.

---

## 📌 Base URL

```
http://localhost:8080/api/v1/event-departments
```

---

## ➕ Create Event Department

**POST** `/create`

Create a new event department under a specific department group.

### Request Body

```json
{
  "user_uuid": "04262c91-6268-44cb-bd86-0bcf51da849d",
  "group_uuid": "3c2d2611-5bb7-4e35-bbb4-1f3de8f97e73",
  "name": "Logistics",
  "description": "Responsible for equipment and distribution."
}
```

### Response

```json
{
  "data": {
    "uuid": "ad02ead7-fd72-490a-bf9a-d0b9581b54f4",
    "name": "Logistics",
    "description": "Responsible for equipment and distribution.",
    "group_uuid": "3c2d2611-5bb7-4e35-bbb4-1f3de8f97e73"
  },
  "message": "Successfully created event department"
}
```

---

## 📥 Get All Event Departments

**GET** `/get-all`

Returns all event departments in the system.

### Response

```json
{
  "data": [
    {
      "uuid": "ad02ead7-fd72-490a-bf9a-d0b9581b54f4",
      "name": "Logistics",
      "description": "Responsible for equipment and distribution.",
      "group_uuid": "3c2d2611-5bb7-4e35-bbb4-1f3de8f97e73"
    }
  ],
  "message": "Successfully retrieved all event departments"
}
```

---

## 📥 Get Event Department by UUID

**GET** `/get/:uuid`

Get detailed information for a specific event department.

### Example

`GET /get/ad02ead7-fd72-490a-bf9a-d0b9581b54f4`

### Response

```json
{
  "data": {
    "uuid": "ad02ead7-fd72-490a-bf9a-d0b9581b54f4",
    "name": "Logistics",
    "description": "Responsible for equipment and distribution.",
    "group_uuid": "3c2d2611-5bb7-4e35-bbb4-1f3de8f97e73"
  },
  "message": "Successfully retrieved event department"
}
```

---

## 🔄 Update Event Department

**PATCH** `/update/:uuid`

Update an existing event department.

### Request Body

```json
{
  "user_uuid": "04262c91-6268-44cb-bd86-0bcf51da849d",
  "group_uuid": "3c2d2611-5bb7-4e35-bbb4-1f3de8f97e73",
  "name": "Updated Logistics",
  "description": "Updated description for logistics team."
}
```

### Response

```json
{
  "data": {
    "uuid": "ad02ead7-fd72-490a-bf9a-d0b9581b54f4",
    "name": "Updated Logistics",
    "description": "Updated description for logistics team.",
    "group_uuid": "3c2d2611-5bb7-4e35-bbb4-1f3de8f97e73"
  },
  "message": "Successfully updated event department"
}
```

---

## ❌ Delete Event Department

**DELETE** `/delete/:uuid`

Delete an event department by its UUID.

### Example

`DELETE /delete/ad02ead7-fd72-490a-bf9a-d0b9581b54f4`

### Response

```json
{
    "data": {
        "uuid": "ad02ead7-fd72-490a-bf9a-d0b9581b54f4",
        "group_uuid": "3c2d2611-5bb7-4e35-bbb4-1f3de8f97e73",
        "name": "Logistik",
        "description": "Bertanggung jawab atas perlengkapan dan distribusi peralatan."
    },
    "message": "Event Department deleted successfully",
    "status": "success"
}
```

---

## 📂 Get Event Departments by Group UUID

**GET** `/get-group-uuid/:group_uuid`

Fetch all event departments belonging to a specific department group.

### Example

`GET /get-group-uuid/3c2d2611-5bb7-4e35-bbb4-1f3de8f97e73`

### Response

```json
{
  "data": [
    {
      "uuid": "ad02ead7-fd72-490a-bf9a-d0b9581b54f4",
      "name": "Logistics",
      "description": "Responsible for equipment and distribution.",
      "group_uuid": "3c2d2611-5bb7-4e35-bbb4-1f3de8f97e73"
    }
  ],
  "message": "Successfully retrieved departments by group UUID"
}
```


# 🎯 HRIS - Event Items API Documentation

API documentation for managing **Event Items** in the HRIS system. Includes operations to create, read, update, and delete items associated with specific events.

---

## 🌐 Base URL

```
http://localhost:8080/api/v1/event-items
```

---

## ➕ Create Event Item

**Endpoint:** `POST /create`

Creates a new event item.

### 📥 Request Body

```json
{
  "user_uuid": "04262c91-6268-44cb-bd86-0bcf51da849d",
  "event_uuid": "d4dd6936-c0dd-4ba2-b0c3-80dc8e721130",
  "name": "Baju",
  "type": "event"
}
```

| Field       | Type   | Required | Description                              |
| ----------- | ------ | -------- | ---------------------------------------- |
| user\_uuid  | UUID   | ✅        | UUID of the user creating the item       |
| event\_uuid | UUID   | ✅        | UUID of the parent event                 |
| name        | string | ✅        | Name of the event item                   |
| type        | string | ✅        | Must be one of: `event`, `date`, `shift` |

### Response

```json
{
    "data": {
        "uuid": "2adae5c8-dbe2-4c03-96d1-b13914a69d40",
        "event_uuid": "d4dd6936-c0dd-4ba2-b0c3-80dc8e721130",
        "name": "Baju",
        "type": "event"
    },
    "message": "Event Item created successfully",
    "status": "success"
}

  ```
```
```
```
```
```
```
---

## 📋 Get All Event Items

**Endpoint:** `GET /get-all`

Retrieves a list of all event items.

```json
{
    "data": [
        {
            "uuid": "2adae5c8-dbe2-4c03-96d1-b13914a69d40",
            "event_uuid": "d4dd6936-c0dd-4ba2-b0c3-80dc8e721130",
            "name": "Baju",
            "type": "event"
        }
    ],
    "message": "Successfully retrieved all Event Items",
    "status": "success"
}
```

---

## 🔍 Get Event Item by UUID

**Endpoint:** `GET /get/:uuid`

**Example:** `/get/2adae5c8-dbe2-4c03-96d1-b13914a69d40`

Fetches details of a specific event item by its UUID.

```json
{
    "data": {
        "uuid": "2adae5c8-dbe2-4c03-96d1-b13914a69d40",
        "event_uuid": "d4dd6936-c0dd-4ba2-b0c3-80dc8e721130",
        "name": "Baju",
        "type": "event"
    },
    "message": "Successfully retrieved Event Item",
    "status": "success"
}
```

---

## 🔎 Get Items by Event UUID

**Endpoint:** `GET /get-by-event-uuid/:event_uuid`

**Example:** `/get-by-event-uuid/d4dd6936-c0dd-4ba2-b0c3-80dc8e721130`

Retrieves all items associated with a specific event UUID.

```json
{
    "data": [
        {
            "uuid": "2adae5c8-dbe2-4c03-96d1-b13914a69d40",
            "event_uuid": "d4dd6936-c0dd-4ba2-b0c3-80dc8e721130",
            "name": "Baju",
            "type": "event"
        }
    ],
    "message": "Successfully retrieved items by event UUID",
    "status": "success"
}
```

---

## ✏️ Update Event Item

**Endpoint:** `PATCH /update/:uuid`

**Example:** `/update/2adae5c8-dbe2-4c03-96d1-b13914a69d40`

### 📥 Request Body

```json
{
  "user_uuid": "04262c91-6268-44cb-bd86-0bcf51da849d",
  "event_uuid": "d4dd6936-c0dd-4ba2-b0c3-80dc8e721130",
  "name": "Baju",
  "type": "event"
}
```

Updates the specified event item.

```json
{
    "data": {
        "uuid": "2adae5c8-dbe2-4c03-96d1-b13914a69d40",
        "event_uuid": "d4dd6936-c0dd-4ba2-b0c3-80dc8e721130",
        "name": "Baju",
        "type": "event"
    },
    "message": "Event Item updated successfully",
    "status": "success"
}
```

---

## ❌ Delete Event Item

**Endpoint:** `DELETE /delete/:uuid`

**Example:** `/delete/2adae5c8-dbe2-4c03-96d1-b13914a69d40`

Deletes an event item by its UUID.

```json
{
    "data": {
        "uuid": "2adae5c8-dbe2-4c03-96d1-b13914a69d40",
        "event_uuid": "d4dd6936-c0dd-4ba2-b0c3-80dc8e721130",
        "name": "Baju",
        "type": "event"
    },
    "message": "Event Item deleted successfully",
    "status": "success"
}
```

---

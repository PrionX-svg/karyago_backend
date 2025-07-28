# 🧾 HRIS - Event Work Area API

This documentation outlines the available API endpoints for managing **Event Work Areas**, including creating, retrieving, updating, and deleting work areas associated with specific events.

---

## 📌 Base URL

```
http://localhost:8080/api/v1/event-work-areas
```

---

## ➕ Create Event Work Area

**POST** `/create`

Creates a new work area associated with a specific event.

### Request Body

```json
{
  "user_uuid": "04262c91-6268-44cb-bd86-0bcf51da849d",
  "event_uuid": "d4dd6936-c0dd-4ba2-b0c3-80dc8e721130",
  "name": "Registration Area"
}
```

### Response

```json
{
  "data": {
    "uuid": "b97a2725-e51f-4a35-97b2-e4cf03fe5dd7",
    "name": "Registration Area",
    "event_uuid": "d4dd6936-c0dd-4ba2-b0c3-80dc8e721130"
  },
  "message": "Successfully created event work area"
}
```

---

## 📥 Get All Event Work Areas

**GET** `/get-all`

Returns a list of all work areas across all events.

### Response

```json
{
  "data": [
    {
      "uuid": "b97a2725-e51f-4a35-97b2-e4cf03fe5dd7",
      "name": "Registration Area",
      "event_uuid": "d4dd6936-c0dd-4ba2-b0c3-80dc8e721130"
    }
  ],
  "message": "Successfully retrieved all event work areas"
}
```

---

## 📥 Get Event Work Area by UUID

**GET** `/get/:uuid`

Retrieve details for a specific work area using its UUID.

### Example

`GET /get/b97a2725-e51f-4a35-97b2-e4cf03fe5dd7`

### Response

```json
{
  "data": {
    "uuid": "b97a2725-e51f-4a35-97b2-e4cf03fe5dd7",
    "name": "Registration Area",
    "event_uuid": "d4dd6936-c0dd-4ba2-b0c3-80dc8e721130"
  },
  "message": "Successfully retrieved event work area"
}
```

---

## 📂 Get Work Areas by Event UUID

**GET** `/get-by-event-uuid/:event_uuid`

Fetch all work areas that belong to a specific event.

### Example

`GET /get-by-event-uuid/d4dd6936-c0dd-4ba2-b0c3-80dc8e721130`

### Response

```json
{
  "data": [
    {
      "uuid": "b97a2725-e51f-4a35-97b2-e4cf03fe5dd7",
      "name": "Registration Area",
      "event_uuid": "d4dd6936-c0dd-4ba2-b0c3-80dc8e721130"
    }
  ],
  "message": "Successfully retrieved work areas for event"
}
```

---

## 🔄 Update Event Work Area

**PATCH** `/update/:uuid`

Update details of an existing work area.

### Request Body

```json
{
  "user_uuid": "04262c91-6268-44cb-bd86-0bcf51da849d",
  "event_uuid": "d4dd6936-c0dd-4ba2-b0c3-80dc8e721130",
  "name": "Main Entrance Area"
}
```

### Response

```json
{
  "data": {
    "uuid": "b97a2725-e51f-4a35-97b2-e4cf03fe5dd7",
    "name": "Main Entrance Area",
    "event_uuid": "d4dd6936-c0dd-4ba2-b0c3-80dc8e721130"
  },
  "message": "Successfully updated event work area"
}
```

---

## ❌ Delete Event Work Area

**DELETE** `/delete/:uuid`

Deletes a specific work area using its UUID.

### Example

`DELETE /delete/b97a2725-e51f-4a35-97b2-e4cf03fe5dd7`

### Response

```json
{
  "data": {
    "uuid": "b97a2725-e51f-4a35-97b2-e4cf03fe5dd7",
    "name": "Main Entrance Area",
    "event_uuid": "d4dd6936-c0dd-4ba2-b0c3-80dc8e721130"
  },
  "message": "Event work area deleted successfully",
  "status": "success"
}
```


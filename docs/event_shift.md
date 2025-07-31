# 🕒 HRIS - Event Shift API Documentation

---

## 📌 Base URL

```
http://localhost:8080/api/v1/event-shifts
```

---

## ➕ Create Event Shift

**POST** `/create`

Creates a new shift for an existing event.

### Request Body

```json
{
  "event_uuid": "ba7ac392-c528-4b38-b844-46bd93bb0368",
  "name": "Morning Session",
  "date": "2025-08-01T00:00:00Z",
  "hour_from": "2025-08-01T08:00:00Z",
  "hour_until": "2025-08-01T12:00:00Z"
}
```

### Successful Response

```json
{
  "data": {
    "uuid": "generated-uuid",
    "name": "Morning Session",
    "date": "2025-08-01T00:00:00Z",
    "hour_from": "2025-08-01T08:00:00Z",
    "hour_until": "2025-08-01T12:00:00Z",
    "event": {
      "uuid": "ba7ac392-c528-4b38-b844-46bd93bb0368",
      "name": "Event Name"
    }
  },
  "message": "Event shift created successfully",
  "status": "success"
}
```

---

## 📋 Get All Event Shifts

**GET** `/get-all`

Fetches all event shifts with associated event data.

### Successful Response

```json
{
  "data": [
    {
      "uuid": "45c33c7c-3456-4a26-83cf-ce5e9808f9df",
      "name": "Morning Session",
      "date": "2025-08-01T00:00:00Z",
      "hour_from": "2025-08-01T08:00:00Z",
      "hour_until": "2025-08-01T12:00:00Z",
      "event": {
        "uuid": "ba7ac392-c528-4b38-b844-46bd93bb0368",
        "name": "Event Name"
      }
    }
  ],
  "message": "Successfully retrieved all event shifts",
  "status": "success"
}
```

---

## 🔍 Get Event Shift By ID

**GET** `/get-by-id/:id`

Fetch a specific shift by its numeric ID.

### Example

```
/get-by-id/1
```

---

## 🔍 Get Event Shift By UUID

**GET** `/get-by-uuid/:uuid`

Fetch a shift by its UUID.

### Example

```
/get-by-uuid/45c33c7c-3456-4a26-83cf-ce5e9808f9df
```

---

## 🔍 Get Event Shifts By Event UUID

**GET** `/get-by-event/:event_uuid`

Returns all shifts related to a single event.

### Example

```
/get-by-event/ba7ac392-c528-4b38-b844-46bd93bb0368
```

---

## ✏️ Update Event Shift

**PATCH** `/update/:uuid`

Updates an existing shift.

### Example

```
/update/45c33c7c-3456-4a26-83cf-ce5e9808f9df
```

### Request Body

```json
{
  "event_uuid": "ba7ac392-c528-4b38-b844-46bd93bb0368",
  "name": "Updated Shift Name",
  "date": "2025-08-15T00:00:00Z",
  "hour_from": "2025-08-15T08:00:00Z",
  "hour_until": "2025-08-15T12:00:00Z"
}
```

---

## ❌ Delete Event Shift

**DELETE** `/delete/:uuid`

Deletes a shift by its UUID.

### Example

```
/delete/45c33c7c-3456-4a26-83cf-ce5e9808f9df
```

> Request body not required.

---

## ✅ Generic Success Response Structure

```json
{
  "data": { },
  "message": "Event shift operation successful",
  "status": "success"
}
```
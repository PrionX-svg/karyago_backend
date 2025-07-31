# 👥 HRIS - Event User API

Documentation for **Event User** endpoints in the HRIS system, enabling creation, retrieval (individual & grouped), updating, and deletion of users assigned to events.

---

## 📌 Base URL

```
http://localhost:8080/api/v1/event-users
```

---

## ➕ Create Event User

**POST** `/create/`

Create a new user-event relationship.

### Request Body

```json
{
  "user_uuid": "1fec3831-a612-45d0-97b7-eb45c2176f45",
  "event_uuid": "633a2746-9a03-4c71-8e48-ffa17df7521e"
}
```

### Response

```json
{
  "data": {
    "uuid": "generated-event-user-uuid",
    "user": {
      "uuid": "1fec3831-a612-45d0-97b7-eb45c2176f45",
      "fullname": "Alice Johnson",
      "email": "alice@example.com"
    },
    "event": {
      "uuid": "633a2746-9a03-4c71-8e48-ffa17df7521e",
      "name": "AI Expo 2025",
      "start_date": "2025-11-12T13:00:00Z",
      "end_date": "2025-11-12T17:00:00Z"
    }
  },
  "message": "Event user created successfully",
  "status": "success"
}
```

---

## 📋 Get All Event Users

**GET** `/get-all/`

Retrieve all event-user associations.

### Response

```json
{
  "data": [
    {
      "uuid": "event-user-uuid-1",
      "user": {
        "uuid": "user-uuid-1",
        "fullname": "John Doe",
        "email": "john@example.com"
      },
      "event": {
        "uuid": "event-uuid-1",
        "name": "Developer Conference",
        "start_date": "2025-08-01T09:00:00Z",
        "end_date": "2025-08-01T17:00:00Z"
      }
    }
  ],
  "message": "Successfully retrieved all event users",
  "status": "success"
}
```

---

## 🔍 Get Event User By ID

**GET** `/get-by-id/:id`

Example: `/get-by-id/4`

---

## 🔎 Get Event User By UUID

**GET** `/get-by-uuid/:uuid`

Example: `/get-by-uuid/84da0e43-3538-4dd6-856a-d02b59a35da5`

---

## 👤 Get Grouped By User UUID

**GET** `/get-by-user/:uuid`

Example: `/get-by-user/91aa937c-740b-4df7-a032-ed0335b424a1`

### Response (Grouped)

```json
{
  "data": {
    "user": {
      "uuid": "91aa937c-740b-4df7-a032-ed0335b424a1",
      "fullname": "Jane Smith",
      "email": "jane@example.com"
    },
    "events": [
      {
        "uuid": "event-uuid-1",
        "name": "HR Tech Summit",
        "start_date": "2025-09-15T10:00:00Z",
        "end_date": "2025-09-15T17:00:00Z",
        "event_user_uuid": "event-user-uuid-1"
      }
    ]
  },
  "message": "Successfully retrieved event users grouped by user",
  "status": "success"
}
```

---

## 📅 Get Grouped By Event UUID

**GET** `/get-by-event/:uuid`

Example: `/get-by-event/ba7ac392-c528-4b38-b844-46bd93bb0368`

### Response (Grouped)

```json
{
  "data": {
    "event": {
      "uuid": "ba7ac392-c528-4b38-b844-46bd93bb0368",
      "name": "Data Science Bootcamp",
      "start_date": "2025-10-05T08:00:00Z",
      "end_date": "2025-10-05T17:00:00Z"
    },
    "users": [
      {
        "uuid": "user-uuid-2",
        "fullname": "Carlos Mendoza",
        "email": "carlos@example.com",
        "event_user_uuid": "event-user-uuid-2"
      }
    ]
  },
  "message": "Successfully retrieved event users grouped by event",
  "status": "success"
}
```

---

## ✏️ Update Event User

**PUT** `/update/:uuid`

Example: `/update/d46aa19e-b32f-4b13-bc1d-9e11f218c022`

### Request Body

```json
{
  "user_uuid": "c6e3de78-b9af-4a93-b8d6-e2c20ab6e712",
  "event_uuid": "a1b2c3d4-e5f6-7890-1234-56789abcdef0"
}
```

### Response

```json
{
  "data": {
    "uuid": "d46aa19e-b32f-4b13-bc1d-9e11f218c022",
    "user": {
      "uuid": "c6e3de78-b9af-4a93-b8d6-e2c20ab6e712",
      "fullname": "New User",
      "email": "newuser@example.com"
    },
    "event": {
      "uuid": "a1b2c3d4-e5f6-7890-1234-56789abcdef0",
      "name": "New Event Name",
      "start_date": "2025-12-01T09:00:00Z",
      "end_date": "2025-12-01T17:00:00Z"
    }
  },
  "message": "Event user updated successfully",
  "status": "success"
}
```

---

## ❌ Delete Event User

**DELETE** `/delete/:uuid`

Example: `/delete/84da0e43-3538-4dd6-856a-d02b59a35da5`

### Response

```json
{
  "data": {
    "uuid": "84da0e43-3538-4dd6-856a-d02b59a35da5",
    "user": {
      "uuid": "user-uuid",
      "fullname": "User Name",
      "email": "user@example.com"
    },
    "event": {
      "uuid": "event-uuid",
      "name": "Deleted Event",
      "start_date": "2025-10-01T00:00:00Z",
      "end_date": "2025-10-01T00:00:00Z"
    }
  },
  "message": "Event user deleted successfully",
  "status": "success"
}
```

---
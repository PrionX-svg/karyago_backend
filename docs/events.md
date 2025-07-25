
# 🎉 HRIS - Event API

Documentation for **Event** endpoints in the HRIS system, supporting creation, retrieval, updating, and deletion of events for companies.

---

## 📌 Base URL

```
http://localhost:8080/api/v1/events
```

---

## ➕ Create Event

**POST** `/create/`

Create a new event associated with a company.

### Request Body

```json
{
  "company_uuid": "09652938-7a1d-4bd8-804c-e573de6aa7c0",
  "name": "Quantum Computing Roundtable",
  "photo": null,
  "start_date": "2025-11-12T13:00:00Z",
  "end_date": "2025-11-12T17:00:00Z"
}
```

### Response

```json
{
  "data": {
    "uuid": "generated-event-uuid",
    "name": "Quantum Computing Roundtable",
    "start_date": "2025-11-12T13:00:00Z",
    "end_date": "2025-11-12T17:00:00Z",
    "company": {
      "uuid": "09652938-7a1d-4bd8-804c-e573de6aa7c0",
      "name": "QuantumAxis"
    }
  },
  "message": "Event created successfully",
  "status": "success"
}
```

---

## 📋 Get All Events

**GET** `/get-all/`

Retrieve all events across all companies.

### Response

```json
{
  "data": [
    {
      "uuid": "17da01d7-9932-41bf-98d3-ab5731f01a39",
      "name": "TechNova Annual Summit",
      "start_date": "2025-08-15T00:00:00Z",
      "end_date": "2025-08-17T00:00:00Z",
      "company": {
        "uuid": "1f6ee520-b049-45d7-8e32-4e678c94a48b",
        "name": "TechNova Solutions"
      }
    },
    {
      "uuid": "4a8ae564-311c-43a7-b1cf-03969e746248",
      "name": "GoLang Backend Workshop",
      "start_date": "2025-09-10T00:00:00Z",
      "end_date": "2025-09-10T00:00:00Z",
      "company": {
        "uuid": "1f6ee520-b049-45d7-8e32-4e678c94a48b",
        "name": "TechNova Solutions"
      }
    }
  ],
  "message": "Successfully retrieved all events",
  "status": "success"
}
```

---

## 🔍 Get Event By ID

**GET** `/get-by-id/:id`

Example: `/get-by-id/4`

### Response

```json
{
  "data": {
    "uuid": "17da01d7-9932-41bf-98d3-ab5731f01a39",
    "name": "TechNova Annual Summit",
    "start_date": "2025-08-15T00:00:00Z",
    "end_date": "2025-08-17T00:00:00Z",
    "company": {
      "uuid": "1f6ee520-b049-45d7-8e32-4e678c94a48b",
      "name": "TechNova Solutions"
    }
  },
  "message": "Successfully retrieved event by ID",
  "status": "success"
}
```

---

## 🔎 Get Event By UUID

**GET** `/get-by-uuid/:uuid`

Example: `/get-by-uuid/4a8ae564-311c-43a7-b1cf-03969e746248`

### Response

```json
{
  "data": {
    "uuid": "4a8ae564-311c-43a7-b1cf-03969e746248",
    "name": "GoLang Backend Workshop",
    "start_date": "2025-09-10T00:00:00Z",
    "end_date": "2025-09-10T00:00:00Z",
    "company": {
      "uuid": "1f6ee520-b049-45d7-8e32-4e678c94a48b",
      "name": "TechNova Solutions"
    }
  },
  "message": "Successfully retrieved event by UUID",
  "status": "success"
}
```

---

## 🏢 Get Events By Company UUID

**GET** `/get-by-company/:uuid`

Example: `/get-by-company/1f6ee520-b049-45d7-8e32-4e678c94a48b`

### Response

```json
{
  "data": [
    {
      "uuid": "17da01d7-9932-41bf-98d3-ab5731f01a39",
      "name": "TechNova Annual Summit",
      "start_date": "2025-08-15T00:00:00Z",
      "end_date": "2025-08-17T00:00:00Z"
    },
    {
      "uuid": "4a8ae564-311c-43a7-b1cf-03969e746248",
      "name": "GoLang Backend Workshop",
      "start_date": "2025-09-10T00:00:00Z",
      "end_date": "2025-09-10T00:00:00Z"
    }
  ],
  "message": "Successfully retrieved events for company",
  "status": "success"
}
```

---

## ✏️ Update Event

**PATCH** `/update/:uuid`

Example: `/update/4a8ae564-311c-43a7-b1cf-03969e746248`

### Request Body

```json
{
  "company_uuid": "1f6ee520-b049-45d7-8e32-4e678c94a48b",
  "name": "GoLang Backend Workshop yang sudah diupdate nih",
  "photo": null,
  "start_date": "2025-09-10T10:00:00Z",
  "end_date": "2025-09-10T15:00:00Z"
}
```

### Response

```json
{
  "data": {
    "uuid": "4a8ae564-311c-43a7-b1cf-03969e746248",
    "name": "GoLang Backend Workshop yang sudah diupdate nih",
    "start_date": "2025-09-10T10:00:00Z",
    "end_date": "2025-09-10T15:00:00Z",
    "company": {
      "uuid": "1f6ee520-b049-45d7-8e32-4e678c94a48b",
      "name": "TechNova Solutions"
    }
  },
  "message": "Event updated successfully",
  "status": "success"
}
```

---

## ❌ Delete Event

**DELETE** `/delete/:uuid`

Example: `/delete/4a8ae564-311c-43a7-b1cf-03969e746248`

### Response

```json
{
  "data": {
    "uuid": "4a8ae564-311c-43a7-b1cf-03969e746248",
    "name": "GoLang Backend Workshop",
    "start_date": "2025-09-10T00:00:00Z",
    "end_date": "2025-09-10T00:00:00Z",
    "company": {
      "uuid": "1f6ee520-b049-45d7-8e32-4e678c94a48b",
      "name": "TechNova Solutions"
    }
  },
  "message": "Event deleted successfully",
  "status": "success"
}
```

---
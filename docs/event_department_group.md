# 🧾 HRIS - Event Department Group API

Documentation for **Event Department Group** management in the HRIS system, including create, retrieve, update, and delete operations.

---

## 📌 Base URL

```

http://localhost:8080/api/v1/event-department-groups

````

---

## ➕ Create Department Group

**POST** `/create`

Create a new department group under a specific event.

### Request Body

```json
{
  "event_uuid": "d4dd6936-c0dd-4ba2-b0c3-80dc8e721130",
  "name": "Logistik",
  "description": "Mengurus pengadaan barang dan logistik acara",
  "responsible_uuid": "04262c91-6268-44cb-bd86-0bcf51da849d"
}
````

### Response

```json
{
  "data": {
    "uuid": "ed656e2a-3ca9-4882-8e35-d116c5fcf27c",
    "event_uuid": "d4dd6936-c0dd-4ba2-b0c3-80dc8e721130",
    "name": "Logistik",
    "description": "Mengurus pengadaan barang dan logistik acara",
    "responsible": {
      "uuid": "04262c91-6268-44cb-bd86-0bcf51da849d",
      "name": "John Doe"
    }
  },
  "message": "Department group created successfully",
  "status": "success"
}
```

---

## 📋 Get All Department Groups

**GET** `/get-all`

Retrieve all department groups.

### Response

```json
{
  "data": [
    {
      "uuid": "ed656e2a-3ca9-4882-8e35-d116c5fcf27c",
      "event_uuid": "d4dd6936-c0dd-4ba2-b0c3-80dc8e721130",
      "name": "Logistik",
      "description": "Mengurus pengadaan barang dan logistik acara",
      "responsible": {
        "uuid": "04262c91-6268-44cb-bd86-0bcf51da849d",
        "name": "John Doe"
      }
    }
  ],
  "message": "Successfully retrieved all department groups",
  "status": "success"
}
```

---

## 🔍 Get Department Group By UUID

**GET** `/get/:uuid`
Example: `/get/ed656e2a-3ca9-4882-8e35-d116c5fcf27c`

### Response

```json
{
  "data": {
    "uuid": "ed656e2a-3ca9-4882-8e35-d116c5fcf27c",
    "event_uuid": "d4dd6936-c0dd-4ba2-b0c3-80dc8e721130",
    "name": "Logistik",
    "description": "Mengurus pengadaan barang dan logistik acara",
    "responsible": {
      "uuid": "04262c91-6268-44cb-bd86-0bcf51da849d",
      "name": "John Doe"
    }
  },
  "message": "Successfully retrieved department group",
  "status": "success"
}
```

---

## 🔎 Get Department Groups By Event UUID

**GET** `/get-event-uuid/:event_uuid`
Example: `/get-event-uuid/d4dd6936-c0dd-4ba2-b0c3-80dc8e721130`

### Response

```json
{
  "data": [
    {
      "uuid": "ed656e2a-3ca9-4882-8e35-d116c5fcf27c",
      "event_uuid": "d4dd6936-c0dd-4ba2-b0c3-80dc8e721130",
      "name": "Logistik",
      "description": "Mengurus pengadaan barang dan logistik acara",
      "responsible": {
        "uuid": "04262c91-6268-44cb-bd86-0bcf51da849d",
        "name": "John Doe"
      }
    }
  ],
  "message": "Successfully retrieved department groups by event UUID",
  "status": "success"
}
```

---

## ✏️ Update Department Group

**PATCH** `/update/:uuid`
Example: `/update/ed656e2a-3ca9-4882-8e35-d116c5fcf27c`

### Request Body

```json
{
  "event_uuid": "d4dd6936-c0dd-4ba2-b0c3-80dc8e721130",
  "name": "Logistik",
  "description": "Mengurus pengadaan barang dan logistik acara",
  "responsible_uuid": "04262c91-6268-44cb-bd86-0bcf51da849d"
}
```

### Response

```json
{
  "data": {
    "uuid": "ed656e2a-3ca9-4882-8e35-d116c5fcf27c",
    "event_uuid": "d4dd6936-c0dd-4ba2-b0c3-80dc8e721130",
    "name": "Logistik",
    "description": "Mengurus pengadaan barang dan logistik acara",
    "responsible": {
      "uuid": "04262c91-6268-44cb-bd86-0bcf51da849d",
      "name": "John Doe"
    }
  },
  "message": "Department group updated successfully",
  "status": "success"
}
```

---

## ❌ Delete Department Group

**DELETE** `/delete/:uuid`
Example: `/delete/ed656e2a-3ca9-4882-8e35-d116c5fcf27c`

### Response

```json
{
  "data": {
    "uuid": "ed656e2a-3ca9-4882-8e35-d116c5fcf27c"
  },
  "message": "Department group deleted successfully",
  "status": "success"
}
```


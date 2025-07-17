# 📘 HRIS - Employment History API

Documentation for **Employment History** management endpoints in the HRIS system, including create, retrieve, update,
and delete operations.

---

## 📌 Base URL

```

http://localhost:8080/api/v1/employment-histories

````

---

## ➕ Create Employment History

**POST** `/create`

Creates a new Employment History with employee details.

### Request Body

```json
{
  "employee_uuid": "9797e2e9-c7fb-4422-82c4-b4a519732a2c",
  "company_uuid": "86a9688c-53b4-42c8-b3c1-819fd6ad1d9f",
  "role_uuid": "f8e73786-b6e7-4f4f-a3d8-2d3071f6cfc8",
  "position": "Senior Backend Developer",
  "is_present": true,
  "start_date": "2023-08-01T00:00:00Z"
}
````

### Response

```json
{
  "data": {
    "uuid": "1de3e9f3-9a2b-4af1-9150-9265e2ba38de",
    "employee": {
      "uuid": "906b0b9d-4eaf-42c2-a76c-c3c4e9471466",
      "full_name": "John Doe 7",
      "email": "john.doe8@example.com"
    },
    "company": {
      "uuid": "86a9688c-53b4-42c8-b3c1-819fd6ad1d9f",
      "name": "PT Teknologi Nusantara 2"
    },
    "role": {
      "uuid": "5012037e-9c1c-4245-8339-a5ff9e741ed7",
      "name": "employee"
    },
    "position": "Senior Backend Developer",
    "is_present": true,
    "start_date": "2023-08-01T00:00:00Z"
  },
  "message": "Employment history created successfully"
}
```

---

## 📋 Get All Employment Histories by Employee UUID

**GET** `/employee/:employee_uuid`

Retrieve all Employment Histories by their employee uuid.

### Response

```json
{
  "data": [
    {
      "uuid": "9a998b35-2cf4-4e81-a333-55f5da7d761c",
      "employee": {
        "uuid": "9797e2e9-c7fb-4422-82c4-b4a519732a2c",
        "full_name": "John Doe",
        "email": "john.doe2@example.com"
      },
      "company": {
        "uuid": "86a9688c-53b4-42c8-b3c1-819fd6ad1d9f",
        "name": "PT Teknologi Nusantara 2"
      },
      "role": {
        "uuid": "f8e73786-b6e7-4f4f-a3d8-2d3071f6cfc8",
        "name": "HR"
      },
      "position": "Senior Backend Developer",
      "is_present": true,
      "start_date": "2023-08-01T00:00:00Z"
    }
  ]
}
```

---

## 📊 Get Employment History by UUID

**GET** `/get/:uuid`

Retrieve a single Employment History and employee data by UUID.

### Example

```
/get/a5e39d0b-9f28-4194-88e1-49018763fdd3
```

### Response

```json
{
  "data": {
    "uuid": "a5e39d0b-9f28-4194-88e1-49018763fdd3",
    "employee": {
      "uuid": "9797e2e9-c7fb-4422-82c4-b4a519732a2c",
      "full_name": "John Doe",
      "email": "john.doe2@example.com"
    },
    "company": {
      "uuid": "86a9688c-53b4-42c8-b3c1-819fd6ad1d9f",
      "name": "PT Teknologi Nusantara 2"
    },
    "role": {
      "uuid": "f8e73786-b6e7-4f4f-a3d8-2d3071f6cfc8",
      "name": "HR"
    },
    "position": "Senior Backend Developer",
    "is_present": true,
    "start_date": "2023-08-01T00:00:00Z"
  }
}
```

---

## ✏️ Update Employment History

**PATCH** `/update/:uuid`

Update employment history fields.

### Example

```
/update/f4d811c2-8583-44f0-ab4c-b2cc1f483387
```

### Request Body

```json
{
  "employee_uuid": "9797e2e9-c7fb-4422-82c4-b4a519732a2c",
  "company_uuid": "86a9688c-53b4-42c8-b3c1-819fd6ad1d9f",
  "role_uuid": "f8e73786-b6e7-4f4f-a3d8-2d3071f6cfc8",
  "position": "Senior Backend Developer",
  "is_present": false,
  "start_date": "2023-08-01T00:00:00Z",
  "end_date": "2025-08-01T00:00:00Z"
}
```

### Response

```json
{
  "data": {
    "uuid": "1de3e9f3-9a2b-4af1-9150-9265e2ba38de",
    "employee": {
      "uuid": "906b0b9d-4eaf-42c2-a76c-c3c4e9471466",
      "full_name": "John Doe 7",
      "email": "john.doe8@example.com"
    },
    "company": {
      "uuid": "86a9688c-53b4-42c8-b3c1-819fd6ad1d9f",
      "name": "PT Teknologi Nusantara 2"
    },
    "role": {
      "uuid": "5012037e-9c1c-4245-8339-a5ff9e741ed7",
      "name": "employee"
    },
    "position": "Senior Backend Developer",
    "is_present": true,
    "start_date": "2023-08-01T00:00:00Z"
  },
  "message": "Employment history updated successfully"
}
```

---

## ❌ Delete Employment History

**DELETE** `/delete/:uuid`

Delete a employment history.

### Example

```
/delete/e681f5e4-a699-435f-a298-61a1e755b81f
```

### Response

```json
{
  "message": "Employment History deleted successfully"
}
```
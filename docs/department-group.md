# 🏢 HRIS - Department Group Management API

Documentation for **Department Group** management endpoints in the HRIS system, including create, retrieve, update, and
delete operations.

---

## 📌 Base URL

```
http://localhost:8080/api/v1/department-groups
```

---

## ➕ Create Department Group

**POST** `/create`

Create a new department group.

### Request Body

```json
{
  "name": "Departemen Finance",
  "company_uuid": "22a27dcb-a6ea-4574-8a65-a3fabb8efd0e",
  "desc": "Departemen yang mengatur keuangan perusahaan"
}
```

### Response

```json
{
  "data": {
    "uuid": "generated-department-group-uuid",
    "name": "Departemen Finance",
    "desc": "Departemen yang mengatur keuangan perusahaan",
    "company_uuid": "22a27dcb-a6ea-4574-8a65-a3fabb8efd0e",
    "responsible_uuid": null
  },
  "message": "Department Group Created Successfully",
  "status": "success"
}
```

---

## 📋 Get All Department Groups

**GET** `/get-all`

Retrieve the list of all department groups.

### Response

```json
{
  "data": [
    {
      "uuid": "4a1cc34e-02a9-4c91-aa74-cd311cd77c29",
      "name": "Departemen Finance",
      "desc": "Departemen yang mengatur keuangan perusahaan",
      "company_uuid": "22a27dcb-a6ea-4574-8a65-a3fabb8efd0e",
      "responsible_uuid": null
    }
  ],
  "message": "Success get all department groups",
  "status": "success"
}
```

---

## 🔍 Get Department Group By UUID

**GET** `/get/:uuid`
Example: `/get/b397aa7a-6dbe-45d6-931a-43d7cf0c94f9`

### Response

```json
{
  "data": {
    "uuid": "b397aa7a-6dbe-45d6-931a-43d7cf0c94f9",
    "name": "Departemen Finance",
    "desc": "Departemen yang mengatur keuangan perusahaan",
    "company_uuid": "22a27dcb-a6ea-4574-8a65-a3fabb8efd0e",
    "responsible_uuid": null
  },
  "message": "Success get department group",
  "status": "success"
}
```

---

## ✏️ Update Department Group

**PATCH** `/update/:uuid`
Example: `/update/1d44d3ac-8bcb-4823-80c7-cdb8dc9173cf`

### Request Body

```json
{
  "name": "Departemen Finance Updated",
  "company_uuid": "22a27dcb-a6ea-4574-8a65-a3fabb8efd0e",
  "desc": "Update untuk departemen keuangan"
}
```

### Response

```json
{
  "data": {
    "uuid": "1d44d3ac-8bcb-4823-80c7-cdb8dc9173cf",
    "name": "Departemen Finance Updated",
    "desc": "Update untuk departemen keuangan",
    "company_uuid": "22a27dcb-a6ea-4574-8a65-a3fabb8efd0e",
    "responsible_uuid": null
  },
  "message": "Department Group Updated Successfully",
  "status": "success"
}
```

---

## ❌ Delete Department Group

**DELETE** `/delete/:uuid`
Example: `/delete/4a1cc34e-02a9-4c91-aa74-cd311cd77c29`

### Response

```json
{
  "data": {
    "uuid": "4a1cc34e-02a9-4c91-aa74-cd311cd77c29",
    "name": "Departemen Finance",
    "desc": "Departemen yang mengatur keuangan perusahaan",
    "company_uuid": "22a27dcb-a6ea-4574-8a65-a3fabb8efd0e",
    "responsible_uuid": null
  },
  "message": "Department Group Deleted Successfully",
  "status": "success"
}
```

---

## 📊 Get All Department Groups (DataTable)

**GET** `/get-all/dt`

Retrieve paginated, searchable, and filterable list of department groups, based on `company_uuid`.

### 🔍 Query Parameters

| Parameter      | Type   | Required | Description                                |
|----------------|--------|----------|--------------------------------------------|
| `page`         | int    | No       | Page number (default: 1)                   |
| `limit`        | int    | No       | Items per page (default: 10)               |
| `search`       | string | No       | Search keyword (applies to `name`, `desc`) |
| `company_uuid` | string | ✅ Yes    | UUID of the company to filter by           |

### 📥 Example

```http
GET /get-all/dt?page=1&limit=10&search=finance&company_uuid=22a27dcb-a6ea-4574-8a65-a3fabb8efd0e
```

### 📤 Response

```json
{
  "data": [
    {
      "uuid": "4a1cc34e-02a9-4c91-aa74-cd311cd77c29",
      "name": "Departemen Finance",
      "desc": "Departemen yang mengatur keuangan perusahaan",
      "company_uuid": "22a27dcb-a6ea-4574-8a65-a3fabb8efd0e",
      "responsible_uuid": "d3a2bc5f-5526-4cc6-b3ea-f541ee6aee2f"
    }
  ],
  "page": 1,
  "limit": 10,
  "total": 20,
  "filtered": 1
}
```
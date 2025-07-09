```md
# 📘 HRIS - Authentication API

This documentation covers the authentication endpoints for the HRIS system.

---

## 📌 Base URL

http://localhost:8080/api/v1/auth

---

## 🔐 Register

**POST** `/register`

Registers a new user and sends a verification email.

### Request Body
```json
{
  "firstname": "Marco",
  "lastname": "Melandri",
  "phone": "6287782553442",
  "email": "marcomelandri808@gmail.com",
  "password": "tangerang22",
  "timezone": "+7"
}
````

### Response

```json
{
  "data": {},
  "message": "Registration successful. Please check your email to verify your account.",
  "status": "success"
}
```

---

## 🔐 Login

**POST** `/login`

Logs in with email and password. A JWT token is returned in the `token` cookie.

### Request Body

```json
{
  "email": "marcomelandri808@gmail.com",
  "password": "12345678"
}
```

### Response

```json
{
  "data": {
    "email": "marcomelandri808@gmail.com",
    "fullname": "Marco Melandri",
    "role_name": "owner",
    "uuid": "f4d811c2-8583-44f0-ab4c-b2cc1f483387"
  },
  "message": "Login successful",
  "status": "success"
}
```

---

## ✅ Verify Email

**POST** `/verify`

Verifies email using the UUID from the verification link.

### Request Body

```json
{
  "uuid": "18002c90-b392-4e35-aabe-540310855a6f"
}
```

### Response

```json
{
  "data": {
    "email": "marcomelandri808@gmail.com",
    "name": "Marco",
    "uuid": "f4d811c2-8583-44f0-ab4c-b2cc1f483387"
  },
  "message": "Email verified successfully",
  "status": "success"
}
```

---

## ✅ Resend Verify Email

**POST** `/resend-verification`

Verifies email using the UUID from the verification link.

### Request Body

```json
{
  "email": "marcomelandri808@gmail.com"
}
```

### Response

```json
{
  "data": {},
  "message": "Verification link resent to your email",
  "status": "success"
}
```

---

## 🔑 Forgot Password

**POST** `/forgot-password`

Sends an OTP to the email for password reset.

### Request Body

```json
{
  "email": "marcomelandri808@gmail.com"
}
```

### Response

```json
{
  "data": {},
  "message": "Reset link sent to email",
  "status": "success"
}
```

---

## 🧾 Verify OTP Forgot Password

**POST** `/forgot-password/verify`

Verifies the OTP sent to the user's email for password reset.

### Request Body

```json
{
  "email": "marcomelandri808@gmail.com",
  "otp_code": "477963"
}
```

### Response

```json
{
  "data": {},
  "message": "OTP verified successfully",
  "status": "success"
}
```

---

## 🔁 Resend OTP Forgot Password

**POST** `/forgot-password/resend`

Resends the OTP for password reset (reuses the existing OTP).

### Request Body

```json
{
  "email": "marcomelandri808@gmail.com"
}
```

### Response

```json
{
  "data": {},
  "message": "OTP resent to your email",
  "status": "success"
}
```

---

## 🔓 Reset Password

**POST** `/forgot-password/reset`

Resets the password using the OTP received by email.

### Request Body

```json
{
  "email": "marcomelandri808@gmail.com",
  "otp_code": "308904",
  "new_password": "12345678"
}
```

### Response

```json
{
  "data": {},
  "message": "Password reset successfully",
  "status": "success"
}
```

---

## 🔒 Logout

**POST** `/logout`

Removes the token cookie and ends the login session.

### Response

```json
{
  "data": {},
  "message": "Logged out successfully",
  "status": "success"
}
```
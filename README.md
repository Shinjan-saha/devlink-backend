
# 📘 DevLink API Documentation

## 🏁 Base URL
```
http://localhost:8080/api
```

---

## 📌 Authentication

- **JWT-based** authentication
- Use the token returned from the login endpoint in the `Authorization` header like so:
  ```
  Authorization: Bearer <your_token>
  ```

---

## 🔐 Register as Admin

**Endpoint:** `POST /register?admin_secret=supersecretadminkey123`  
**Headers:** _None required_  
**Body:**
```json
{
  "username": "admiouser",
  "email": "admio@gmail.com",
  "password": "hola@1305",
  "role": "admin"
}
```

---

## 🧑‍ Register as User

**Endpoint:** `POST /register`  
**Headers:** _None required_  
**Body:**
```json
{
  "username": "shinjan",
  "email": "shinjan@gmail.com",
  "password": "hola@1304"
}
```

---

## 🔑 Login

**Endpoint:** `POST /login`  
**Headers:** _None required_  
**Body:**
```json
{
  "email": "admio@gmail.com",
  "password": "hola@1305"
}
```

**Response:**
```json
{
  "token": "<your_jwt_token>"
}
```

---

## 📥 Submit a Resource

**Endpoint:** `POST /resources`  
**Headers:**
```
Authorization: Bearer <user_token>
```

**Body:**
```json
{
  "title": "Awesome React Tutorial",
  "url": "https://example.com/go-tutorial",
  "description": "A complete Next.js tutorial.",
  "type": "article",
  "tags": "next, tutorial"
}
```

**Response:**
```json
{
  "message": "submitted for review",
  "resourceID": 5
}
```

---

## 📚 Get All Resources

**Endpoint:** `GET /resources`  
**Headers:**
- Optional `Authorization` for admin access.

**Behavior:**
- Admins see all resources.
- Others only see approved resources.

---

## 🟠 Get Pending Resources (Admin Only)

**Endpoint:** `GET /resources/pending`  
**Headers:**
```
Authorization: Bearer <admin_token>
```

**Response:**
```json
[
  {
    "id": 5,
    "title": "Awesome React Tutorial",
    "url": "...",
    "description": "...",
    "approved": false
  }
]
```

---

## ✅ Approve a Resource (Admin Only)

**Endpoint:** `PUT /resources/:id/approve`  
**Example:**  
`PUT /resources/5/approve`  

**Headers:**
```
Authorization: Bearer <admin_token>
X-Admin-Secret: supersecretadminkey123
```

**Response:**
```json
{
  "message": "resource approved"
}
```

---

## 🧾 Summary Table

| Endpoint                          | Method | Auth        | Role      | Description                          |
|----------------------------------|--------|-------------|-----------|--------------------------------------|
| `/register`                      | POST   | ❌          | All       | Register as user                     |
| `/register?admin_secret=...`     | POST   | ❌          | All       | Register as admin (with secret)      |
| `/login`                         | POST   | ❌          | All       | Login and get token                  |
| `/resources`                     | GET    | ✅ (optional) | All/Admin | Get resources (all for admin)        |
| `/resources`                     | POST   | ✅          | User/Admin| Submit new resource                  |
| `/resources/pending`             | GET    | ✅          | Admin     | List all unapproved resources        |
| `/resources/:id/approve`         | PUT    | ✅          | Admin     | Approve a resource by ID             |


## Import these JSON File in POSTMAN & Test Them:-

[![Run in Postman](https://run.pstmn.io/button.svg)](./Devlink.postman_collection.json)
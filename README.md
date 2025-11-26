# Proyek Akhir Keamanan Jaringan - Penetration Testing
## Unauthorized Password Change - No Auth API 

- Muhammad Nadzhif Fikri
- Raditya Alif Nugroho

## Getting Started

### Prerequisities
- [Docker](https://docs.docker.com/engine/install/)

### Installation
1. Clone repo (buset dah ofkors)
    ```
    git clone https://github.com/Kutukobra/FinproKemjar_5.git
    ```
2. Dockerize
    ```docker
    docker compose up
    ```

## Backend API Endpoints

| HTTP | Endpoint                    | Action                           |
| ---- | --------------------------- | -------------------------------- |
| GET  | `/api/user/:username`       | Get user data                    |
| POST | `/api/user/register`        | Create new user                  |
| POST | `/api/user/login`           | Login as existing user           |
| PUT  | `/api/user/change-password` | Change password of existing user |

---

## **GET /api/user/:username**

### **Request**

* URL parameter: `:username`

### **Response**

#### **Success (200 OK)**

```json
{
    "data": user-data
}
```

#### **Failure (400 Bad Request)**

```json
{
    "error": "Invalid username."
}
```

#### **Failure (404 Not Found)**

```json
{
    "error": "User not found."
}
```

#### **Failure (500 Internal Server Error)**

```json
{
    "error": "Internal server error."
}
```

---

## **POST /api/user/register**

### **Request**

* Query parameters:

  * `username`
  * `email`
  * `password`

Example:

```
POST /api/user/register?username=john&email=john@gmail.com&password=123
```

### **Response**

#### **Success (201 Created)**

```json
{
    "data": user-data
}
```

#### **Failure (409 Conflict)**

```json
{
    "error": "Username or Email taken."
}
```

#### **Failure (500 Internal Server Error)**

```json
{
    "error": "Internal server error."
}
```

---

## **POST /api/user/login**

### **Request**

* Query parameters:

  * `username`
  * `password`

Example:

```
POST /api/user/login?username=john&password=123
```

### **Response**

#### **Success (200 OK)**

```json
{
    "data": user-data
}
```

#### **Failure (401 Unauthorized)**

```json
{
    "error": "Wrong email or password."
}
```

#### **Failure (500 Internal Server Error)**

```json
{
    "error": "Internal server error."
}
```

---

## **PUT /api/user/change-password**

### **Request**

* Query parameters:

  * `username`
  * `password`

Example:

```
PUT /api/user/change-password?username=john&password=newpass
```

### **Response**

#### **Success (200 OK)**

```json
{
    "data": user-data
}
```

#### **Failure (400 Bad Request)**

```json
{
    "error": "Username or password cannot be empty."
}
```

#### **Failure (500 Internal Server Error)**

```json
{
    "error": "Internal server error."
}
```



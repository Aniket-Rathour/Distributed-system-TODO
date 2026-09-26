# Distributed System TODO

A Go REST API for managing users and posts, backed by PostgreSQL (Neon).

## Setup

1. Create a `.env` file in the project root:
```
DB_STRING = "postgresql://<user>:<password>@<host>/<dbname>?sslmode=require"
```

2. Run the server:
```bash
cd app
go run main.go
```

Server starts on `http://localhost:8080`

---

## API Endpoints

Base URL: `http://localhost:8080`

---

### Users

#### Create User
```
POST /users
```
**Body** (JSON):
```json
{
    "name": "aniket",
    "Pass": "mypassword123"
}
```
> `name` is required. `Pass` must be between 8–72 characters. Password is bcrypt hashed before storing.

**Response**: Returns user ID and creation timestamp.

---

#### Get All Users
```
GET /users
```
**Body**: None

**Response**: List of all users with their IDs and names.

---

#### Verify User Password
```
HEAD /users
```
**Body** (JSON):
```json
{
    "name": "aniket",
    "Pass": "mypassword123"
}
```
**Response**: `200` if password matches, `409 Conflict` if it doesn't.

> ⚠️ Note: `HEAD` requests don't return a response body in most HTTP clients. You may need to check the status code.

---

### Posts

#### Create Post
```
POST /posts
```
**Body** (JSON):
```json
{
    "Title": "my first post",
    "Desc": "this is the description",
    "Userid": 1
}
```
> `Userid` must be the ID of an existing user.

**Response**: Returns post ID and creation date.

---

#### Get Posts by User (Paginated)
```
GET /posts/{Id}?page=1
```
**Body**: None

| Param   | In    | Description                      |
|---------|-------|----------------------------------|
| `Id`    | path  | User ID to fetch posts for       |
| `page`  | query | Page number (2 posts per page)   |

**Example**:
```
GET /posts/1?page=1
```

**Response**: List of posts with title, description, and ID.

---

#### Get All Posts with Author Names
```
PUT /posts
```
**Body**: None

**Response**: All posts joined with their author's username.

---

## Quick Postman Setup

| #  | Method | URL                              | Body |
|----|--------|----------------------------------|------|
| 1  | POST   | `http://localhost:8080/users`     | `{"name":"aniket","Pass":"mypassword123"}` |
| 2  | GET    | `http://localhost:8080/users`     | — |
| 3  | HEAD   | `http://localhost:8080/users`     | `{"name":"aniket","Pass":"mypassword123"}` |
| 4  | POST   | `http://localhost:8080/posts`     | `{"Title":"my post","Desc":"details here","Userid":1}` |
| 5  | GET    | `http://localhost:8080/posts/1?page=1` | — |
| 6  | PUT    | `http://localhost:8080/posts`     | — |

> Set `Content-Type: application/json` header for all requests with a body.
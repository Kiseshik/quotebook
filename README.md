# QuoteBook - REST API service

A microservice for storing and managing quotes, written in Go.

# Installation and launch

```bash
# Clone repository
git clone https://github.com/Kiseshik/quotebook.git

# Go to the project folder
cd quotebook

# launch server
go run .
go run main.go
```
# Testing

```bash
# launch test 
go test -v
```

# API Endpoints

| Method| Path                | Description               |
|-------|---------------------|---------------------------|
| POST  | /quotes             | Add new quote             |
| GET   | /quotes             | Get all quotes            |
| GET   | /quotes/random      | Get random quote          |
| GET   | /quotes?author=name | Filter by author          |
| DELETE| /quotes/{id}        | Delete quote              |

# API Methods

### 1. Add new quote
**Endpoint:** `POST /quotes`

**Request Body:**
```json
{
  "author": "Author name",
  "quote": "Quote text"
}
```

**Example:**
```bash
curl -X POST http://localhost:8080/quotes \
  -H "Content-Type: application/json" \
  -d '{"author":"Confucius", "quote":"Life is really simple..."}'
```

**Success Response (201 Created):**
```json
{
  "id": 1,
  "author": "Confucius",
  "quote": "Life is really simple...",
  "created_at": "2025-05-20T12:00:00Z"
}
```

---

### 2. Get all quotes
**Endpoint:** `GET /quotes`

**Example:**
```bash
curl http://localhost:8080/quotes
```

**Success Response (200 OK):**
```json
[
  {
    "id": 1,
    "author": "Confucius",
    "quote": "Life is really simple...",
    "created_at": "2025-05-20T12:00:00Z"
  }
]
```

---

### 3. Get random quote
**Endpoint:** `GET /quotes/random`

**Example:**
```bash
curl http://localhost:8080/quotes/random
```

**Success Response (200 OK):**
```json
{
  "id": 1,
  "author": "Confucius",
  "quote": "Life is really simple...",
  "created_at": "2025-05-20T12:00:00Z"
}
```

**Empty Response (404 Not Found):**
```json
{
  "error": "No quotes available"
}
```

---

### 4. Filter by author
**Endpoint:** `GET /quotes?author={author_name}`

**Example:**
```bash
curl "http://localhost:8080/quotes?author=Confucius"
```

**Success Response (200 OK):**
```json
[
  {
    "id": 1,
    "author": "Confucius",
    "quote": "Life is really simple...",
    "created_at": "2025-05-20T12:00:00Z"
  }
]
```

---

### 5. Delete quote
**Endpoint:** `DELETE /quotes/{id}`

**Example:**
```bash
curl -X DELETE http://localhost:8080/quotes/1
```

**Success Response (204 No Content)** - no response body

**Not Found Response (404 Not Found):**
```json
{
  "error": "Quote not found"
}
```

---

### HTTP Status Codes

| Code | Description              |
|------|--------------------------|
| 200  | OK                       |
| 201  | Created                  |
| 204  | No Content               |
| 400  | Bad Request              |
| 404  | Not Found                |
| 405  | Method Not Allowed       |
| 500  | Internal Server Error    |

# Technologies
- Go (stdlib)
- In-memory storage
- REST API
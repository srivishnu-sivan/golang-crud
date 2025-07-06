
# 🛠️ Go Gin GORM CRUD API

A simple **RESTful API** built with **Go**, **Gin**, **GORM**, and **PostgreSQL**, implementing basic CRUD operations for managing users.

---

## 📂 Project Structure

```
go-crud/
├── main.go                 # Entry point of the application
├── go.mod                  # Go module file
├── database/               # Database connection logic
│   └── database.go
├── models/                 # Database models
│   └── user.go
├── controllers/            # API route handlers (controllers)
│   └── user_controller.go
```

---

## 🏗️ Tech Stack

✅ [Go](https://go.dev/) - Programming Language  
✅ [Gin](https://gin-gonic.com/) - Lightweight HTTP web framework  
✅ [GORM](https://gorm.io/) - ORM for Go  
✅ [PostgreSQL](https://www.postgresql.org/) - Relational Database  

---

## ⚡ Features

- User Creation (`POST /users`)  
- Get All Users (`GET /users`)  
- Get User by ID, Name, or Email (`GET /user?id=1`, `?name=John`, etc.)  
- Update User (`PUT /user?id=1`)  
- Delete User (`DELETE /user?id=1`)  

---

## 🛠️ Setup Instructions

### 1. **Clone the Repository**

```bash
git clone https://github.com/yourusername/go-crud.git
cd go-crud
```

### 2. **Install Dependencies**

```bash
go mod tidy
```

### 3. **Setup PostgreSQL**

Make sure PostgreSQL is running locally. Default credentials used:

```
host=localhost
port=5432
user=postgres
password=test@123
dbname=postgres
sslmode=disable
```

⚠️ In production, use environment variables for credentials.

### 4. **Run the Application**

```bash
go run main.go
```

---

## 📦 Code Overview

### `database/database.go`

Handles PostgreSQL connection using GORM.

### `models/user.go`

Defines the `User` model mapped to the `users` table.

### `controllers/user_controller.go`

Implements the following API endpoints:

| Method | Endpoint | Description |
|--------|-----------|-------------|
| `POST` | `/users` | Create a new user |
| `GET`  | `/users` | Get all users |
| `GET`  | `/user`  | Get user by `id`, `name`, or `email` |
| `PUT`  | `/user`  | Update user details by `id` |
| `DELETE` | `/user` | Delete a user by `id` |

### `main.go`

- Initializes the database.
- Defines API routes.
- Starts the Gin server.

---

## 📬 Example API Calls (Using `curl`)

**Create a User**

```bash
curl -X POST -H "Content-Type: application/json" -d '{"name":"John", "email":"john@example.com"}' http://localhost:8080/users
```

**Get All Users**

```bash
curl http://localhost:8080/users
```

**Get User by ID**

```bash
curl http://localhost:8080/user?id=1
```

**Update User**

```bash
curl -X PUT -H "Content-Type: application/json" -d '{"name":"Updated Name", "email":"updated@example.com"}' http://localhost:8080/user?id=1
```

**Delete User**

```bash
curl -X DELETE http://localhost:8080/user?id=1
```

---

## ✅ Best Practices & Recommendations

- Never hardcode credentials in production.
- Use `.env` files or environment variables for configuration.
- Validate user input properly.
- Handle errors gracefully.
- Secure your API with authentication for real-world applications.

---

## 📚 Resources

- [Gin Documentation](https://gin-gonic.com/docs/)
- [GORM Documentation](https://gorm.io/docs/)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [Go Modules](https://go.dev/doc/modules)

---

## 💡 Future Improvements (Optional)

- Add JWT Authentication  
- Implement Request Validation Libraries  
- Pagination for listing users  
- Dockerize the application  

---

**Happy Coding! 🎉**
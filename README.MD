# Go Payment API
## Features

- User Registration, Login, Logout
- Payment processing
- Transaction History
- JWT Authentication with Cookie
- Role-based Authorization (Customer, Merchant)

## Requirements

- Go 1.21.1 or later
- MySQL
- [Fiber](https://gofiber.io/) v2
- [JWT](https://github.com/dgrijalva/jwt-go) v4
- [GORM](https://gorm.io/) for ORM
- [GoFiber](https://gofiber.io/) for web framework
- [Postman Collection](https://github.com/AlanPratama/go-payment/blob/main/Go%20Payment%20API.postman_collection.json)

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/alanpratama/go-payment.git
cd go-payment
```

### 2. Set up your MySQL database and update the connection details in
**connection.go**
```go
connection, err := gorm.Open(mysql.Open("root:@/go_payment"), &gorm.Config{})
```
Note: default database name is **go_payment**, you can change with your own database, and change db name in connection.go

### 3. Run the application

```bash
go run main.go
```

### 4. The server will run on http://localhost:3000.

```bash
[BASE URL]
http://localhost:3000/api/v1
```

***

## Endpoint API

### Authentication Endpoint
- Register
```
[POST] /api/v1/auth/register
```
- Login
```
[POST] /api/v1/auth/login
```
- Logout
```
[POST] /api/v1/auth/logout
```
- User Logged
```
[GET] /api/v1/auth/user
```

### Payment Endpoint
- Create Payment [ONLY CUSTOMER]
```
[POST] /api/v1/payments
```

- Change Payment Status [ONLY MERCHANT]
```
[PATCH] /api/v1/payments/:paymentID
```

- Get My Payments
```
[GET] /api/v1/payments
```

- Get Payments By ID
```
[GET] /api/v1/payments/:paymentID
```

### History Endpoint
- Login History
```
[GET] /api/v1/history/login
```

- Logout History
```
[GET] /api/v1/history/logout
```


## [Frontend](https://github.com/thogtq/ecommerce-intern-frontend)
## Quick start
Add new .env in root folder file like below
```
MONGODB_HOST="mongodb://root:root@localhost"
MONGODB_PORT="27017"
JWT_SECRET="asd2d1f2f1gv213g24g2evd"
```
```bash
# Initialize MongoDB and its sample data
docker-compose up --build -d
# Download Go modules
go get .
# Start the server
go run main.go
```
## Frameworks and libraries
- [x] gin v1.9.0
- [x] jwt-go v3.2.0 (deprecated)
- [x] godotenv v1.5.1
## Routes
#### Orders
- POST /orders
- PUT /orders/status
- GET /orders
#### Products
- GET /products
- GET /products/<productId>
- DELETE /products/<productId>
- POST /products/<productId>
- PUT /products/<productId>
- POST /products/image
#### Reviews
- GET /reviews
- DELETE /reviews/<reviewId>
- POST /reviews/
#### Users
- POST /user/login
- POST /admin/login
- POST /users/
 - PUT /users/password
 - PUT /users/
  - GET /users/
  - GET /users/token
  

# Go E-commerce API

## Overview
A modular, API-first e-commerce backend written in Go using chi router. Features include JWT authentication, user/admin/vendor roles, products, categories, cart, orders, payments, reviews, notifications, and more.

## Features
- JWT authentication (users, admin, vendor)
- Products, categories, images, inventory
- Cart, wishlist, orders, payments
- Reviews, notifications, coupons
- Admin and vendor endpoints
- CORS and rate limiting middleware

## Requirements
- Go 1.20+
- chi router

## Setup
```bash
cd ecommerce-go
cp .env.example .env # (if using env config)
go mod tidy
go run ./cmd/main.go
```


## Project Structure
- `cmd/`: Entrypoint and server setup
- `internal/`: Handlers, router, middleware, domain logic
- `model.md`: Database schema

## Development
- All endpoints are stubbed in `internal/router/router.go`
- Implement business logic in each domain's handler/service
- Use middleware for JWT, admin, vendor, CORS, and rate limiting

---
MIT License

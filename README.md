# golang-eshop-backend

Backend for ecommerce built using golang.

New stuff that has been learnd while doing this:
- **fiber**: go web framework (https://github.com/gofiber/fiber/v2)
- **gorm**: go orm 
- Adding unique correlationID aka traceID + other additional stuff to logs for each request

## Running
Install go and then use one of these commands:
```
# using make
make run-api

# using go
APP_ENV=local go run cmd/rest/main.go
```

## Package Evaluation

### Fiber

**Pros:**
- Fastest Go web framework available.
- Easy to create and implement middleware.

**Cons:**
- Huge downside is that it uses a custom context type (`fiber.Ctx`) instead of the standard `context.Context`.
- This can complicate integration with libraries expecting the standard Go context.

---

## GORM

**Pros:**
- To come

**Cons:**
- To come

---

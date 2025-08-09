# Go JWT Auth Gateway

This project is a minimal, idiomatic JWT authentication gateway for Go, ready to be used as a secure entry point or as a base for larger applications.

## Architecture
- **Fiber** web framework for fast HTTP handling
- Modular structure: `controller/`, `middleware/`, `service/`, `repository/`, `view/`
- Stateless JWT authentication (HMAC SHA256)
- In-memory whitelist/blacklist for token control

## Environment Variables
- `JWT_SECRET`: Secret key for signing JWTs (required)
- `SUB`: Allowed subject (user id) for token generation (required)

## Running

**Local:**
```bash
go run ./cmd
```

**Docker:**
```bash
docker build -t go-jwt-auth .
docker run --env-file .env -p 9001:9001 go-jwt-auth
```

## Endpoints
- `GET /health` — Health check
- `POST /jwt-generate` — Generate JWT (requires correct password and sub)
- `GET /jwt-authorize` — Protected, returns JWT claims
- `GET /app/protected` — Protected, simulates a secure app area

## How it works
- Only the user/sub defined in `SUB` and password in `JWT_SECRET` can generate tokens
- All protected routes require a valid, whitelisted, non-blacklisted JWT
- Middleware handles validation and injects claims into the request context

## HTTP Collections
Example requests for testing are in `http_collection/go_jwt_auth.har` (importable in Postman or Insomnia).

## Production notes
- Replace in-memory stores with persistent storage for real use
- Never expose your secrets or tokens in version control
- Extend controllers/services for your business logic

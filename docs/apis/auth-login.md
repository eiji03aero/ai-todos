# Login API

## Overview
Authenticates a user by email and password, creating a user session upon successful login.

## Endpoint Details
- **HTTP Method**: POST
- **Path**: `/api/auth/login`

## Request Body
```json
{
  "email": "string",
  "password": "string",
  "remember_me": "boolean"
}
```

### Input Validation
- **Email**:
  - Must be a valid email format
  - Maximum length: 255 characters
  - Cannot be empty

- **Password**:
  - Minimum length: 8 characters
  - Must be a non-empty string

- **Remember Me**:
  - Boolean flag indicating whether to extend session duration

## User Authentication Process
1. Validate input format and constraints
2. Lookup user by email in the database
3. Verify password using secure hashing
4. If credentials are valid:
   - Generate a new session ID
   - Create a user session record in the database
   - Set session ID in an HTTP-only, secure cookie in the response headers
   - If `remember_me` is true, extend cookie expiration to provide a longer session duration

## Session Management
- Session ID is stored in an HTTP-only, secure cookie
- Cookie attributes:
  - `HttpOnly`: Prevents client-side script access
  - `Secure`: Ensures transmission only over HTTPS
  - `SameSite`: Set to `None` to allow cross-site cookie usage
- Session duration:
  - Standard session: Short-lived (e.g., 1 hour)
  - Extended session (with `remember_me`): Longer duration (e.g., 30 days)

## Security Considerations
- Use bcrypt for password hashing
- Validate and sanitize all input
- Use secure, httpOnly cookies for session management
- Rotate session IDs on critical events
- Implement proper logout mechanism to invalidate sessions
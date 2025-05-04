# Signup API Endpoint

## Overview
The Signup API endpoint allows new users to create an account by providing their email, password, and password confirmation.

## Endpoint Details
- **HTTP Method**: POST
- **Path**: `/api/auth/signup`
- **Authentication**: Public (No authentication required)

## Request Body
```json
{
  "email": "string",
  "password": "string",
  "passwordConfirmation": "string"
}
```

### Request Body Validation
- `email`:
  - Must be a valid email format
  - Must be unique (not already registered)
  - Required field

- `password`:
  - Minimum length: 8 characters
  - Must contain at least:
    - 1 uppercase letter
    - 1 lowercase letter
    - 1 number
    - 1 special character
  - Required field

- `passwordConfirmation`:
  - Must exactly match the `password` field
  - Required field

## Response

### Success Response (201 Created)
```json
{
  "id": "string",
  "email": "string",
  "createdAt": "datetime"
}
```

### Session Cookie
- Upon successful signup, a user session is created.
- The session ID is set in a secure, HTTP-only cookie in the response headers.
- Cookie attributes:
  - `HttpOnly`: Prevents client-side script access to the cookie.
  - `Secure`: Ensures the cookie is sent only over HTTPS connections.
  - `SameSite`: Set to `None` to allow cross-site cookie usage.
- The cookie enables the user to remain authenticated in subsequent requests.

### Error Responses

#### 400 Bad Request (Validation Errors)
```json
{
  "error": "Validation Failed",
  "details": [
    {
      "field": "email",
      "message": "Email is already registered"
    },
    {
      "field": "password",
      "message": "Password does not meet strength requirements"
    },
    {
      "field": "passwordConfirmation",
      "message": "Passwords do not match"
    }
  ]
}
```

#### 500 Internal Server Error
```json
{
  "error": "Internal Server Error",
  "message": "An unexpected error occurred during signup"
}
```

## Processing Steps
1. Validate input format and constraints
2. Check email uniqueness
3. Validate password strength
4. Confirm password matches
5. Hash password securely
6. Create user record in database
7. Return user details or appropriate error

## Security Considerations
- Password is hashed using bcrypt with default cost
- Comprehensive input validation
- Unique email constraint
- No sensitive information returned in response

## Related Documents
- [User Model Specification](/docs/models/User.md)
- [Authentication Feature](/docs/features/Authentication.md)
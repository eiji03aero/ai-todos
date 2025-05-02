## Login API

### Overview
Authenticate a user by email and password, creating a user session.

### HTTP Method
POST

### API Path
`/api/auth/login`

### Request Body
```json
{
  "email": "string",
  "password": "string",
  "remember_me": "boolean"
}
```

### Request Validation
- Email:
  - Required
  - Must be a valid email format
- Password:
  - Required
  - Minimum 8 characters
  - Must contain at least one uppercase, one lowercase, and one number

### Successful Response (200 OK)
```json
{
  "user": {
    "id": "string",
    "email": "string"
  },
  "session": {
    "token": "string",
    "expires_at": "datetime",
    "is_remembered": "boolean"
  }
}
```

### Error Responses
- 400 Bad Request: Invalid input format
- 401 Unauthorized: Incorrect email or password
- 422 Unprocessable Entity: Validation errors

### Processing Steps
1. Validate input email and password format
2. Retrieve user by email from database
3. Verify password using secure hash comparison
4. If credentials are valid:
   - Generate a new user session token
   - Set session expiration based on "remember_me" flag
   - Create and store UserSession
5. Return user details and session token
6. If credentials are invalid, return 401 Unauthorized

### Security Considerations
- Use bcrypt or similar secure password hashing
- Generate cryptographically secure session tokens
- Implement rate limiting for login attempts
- Secure token storage and transmission
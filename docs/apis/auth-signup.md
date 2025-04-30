# User SignUp API

## Overview
This API endpoint allows new users to create an account in the system by providing their email address and a secure password.

## Endpoint Details
- **HTTP Method**: POST
- **Path**: `/api/auth/signup`

## Related Domain Models
- **User Model**: `docs/models/User.md`
  * This API directly interacts with the User domain model
  * Creates a new User record in the system
  * Transforms input data into User model properties:
    - Generates `id` (system-generated unique identifier)
    - Stores `email` from input
    - Creates `password_hash` by securely hashing the input password
    - Sets `createdAt` and `updatedAt` timestamps

## Request Body
```json
{
  "email": "string",
  "password": "string",
  "password_confirmation": "string"
}
```

## Input Validations
### Email
- Must be a valid email format (RFC 5322 standard)
- Cannot be empty
- Must be unique in the system

### Password
- Minimum 8 characters long
- Must contain:
  * At least one uppercase letter
  * At least one lowercase letter
  * At least one number
- Password and password_confirmation must match

## Processing Steps
1. Validate input fields
   - Check email format
   - Verify email uniqueness
   - Validate password complexity
2. Confirm password matches password_confirmation
3. Hash password using secure hashing algorithm (bcrypt)
4. Create new user record
5. Automatically log in the user

## Responses
### Success Response (201 Created)
```json
{
  "id": "string",
  "email": "string",
  "createdAt": "datetime"
}
```

### Error Responses
- 400 Bad Request: 
  * Invalid email format
  * Password too weak
  * Passwords do not match
- 409 Conflict: 
  * Email already exists
- 500 Internal Server Error: 
  * Unexpected system error during registration

## Security Considerations
- Passwords are never stored in plain text
- Use bcrypt or equivalent secure hashing algorithm
- Implement rate limiting to prevent brute-force attacks
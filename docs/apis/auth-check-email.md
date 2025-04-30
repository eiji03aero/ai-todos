# Email Uniqueness Check API

## Overview
This API allows checking the uniqueness of an email address during the sign-up process.

## Endpoint Details
- **HTTP Method**: GET
- **Path**: `/api/auth/check-email`

## Query Parameters
- `email` (string, required): The email address to check for uniqueness

## Request Validation
1. Email format validation
   - Must be a valid email address
   - Cannot be empty
   - Follows standard email format rules (e.g., contains '@', valid domain)

## Response
- **Content Type**: application/json
- **Success Response**:
  - HTTP Status: 200 OK
  - Body: 
    ```json
    {
      "isUnique": true/false
    }
    ```

## Error Handling
- 400 Bad Request: 
  - Invalid email format
  - Missing email parameter

## Processing Steps
1. Validate email format
2. Query the database to check if the email exists
3. Return boolean indicating email uniqueness

## Security Considerations
- Rate limit the endpoint to prevent abuse
- Implement input sanitization
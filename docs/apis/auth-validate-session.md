# API Design: Validate User Authentication Session

## General
This API validates the user authentication session by checking the session ID stored in the cookie. It verifies the session's validity and expiration, and authenticates the user session accordingly.

## HTTP Method
POST

## API Path
/auth/validate-session

## Query Parameters
None

## Request Body
None

## Cookies
- `sessionId`: The session ID stored in the user's cookie, used to identify and validate the user session.

## Response
- 200 OK
  ```json
  {
    "sessionId": "string",
    "userId": "string",
    "expiresAt": "datetime"
  }
  ```
- 401 Unauthorized
  ```json
  {
    "error": "Invalid or expired session"
  }
  ```

## Domain Models
- UserSession
  - Properties:
    - sessionId: string
    - userId: string
    - token: string
    - expiresAt: datetime
  - Relations:
    - UserSession belongs to one User
  - Validations:
    - token must be valid JWT
    - expiresAt must be in the future

## Processing Steps
1. Read the `sessionId` from the cookie in the request.
2. Look up the UserSession associated with the `sessionId`.
3. Validate the token stored in the UserSession (must be a valid JWT).
4. Check if the session is expired by comparing `expiresAt` with the current time.
5. If the session is invalid or expired, return 401 Unauthorized with an error message.
6. If valid, return 200 OK with sessionId, userId, and expiresAt.

## Remarks
- The session ID is stored in a cookie and must be securely handled.
- Token expiration must be strictly enforced.
- Error handling must clearly indicate invalid or expired sessions.
- This API is stateless and does not require query parameters or request body.
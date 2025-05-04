# Auth AppState API

## General
This API endpoint provides the current application state data for the authenticated user. It is a protected route that requires a valid authentication token.

## API Path
`/api/auth/app_state`

## Flows
- Client sends a request with a valid authentication token.
- Server validates the authentication token.
- Server returns the current app state data.

## Response Object Schema
The response is a key-value object with the following property:
- `userId`: string

Example response:
```json
{
  "userId": "string"
}
```

## Remarks
- This is a protected route; authentication is required.
- No additional validation rules or error handling remarks beyond standard authentication checks.
- No specific domain models are required for this API.
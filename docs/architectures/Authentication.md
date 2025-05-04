# Authentication Architecture

## General
- It defines concrete methods to authenticate user

## Processing flow
- User login to the app
- Session will be created
- Session id will be stored in cookie of client
- From there on, user will be authenticated via session id

## API Remarks
- Stores session id in client's cookie
- For protected api routes, user will be authenticated via session id
- Backend sets authenticated user in gin.Context with key "currentUser" for protected routes
- Implements a dedicated api endpoint which authenticates the user. This will be used before page load

## Client Remarks
- Client verifies authentication with dedicated api before page load for protected routes
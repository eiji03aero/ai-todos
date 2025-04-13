# Implementation
## General
- Use docker compose for running and testing application code
- When using docker compose, use it in background mode so that your work will not be blocked by waiting for command response
- Please speak in Japanese

## api
- When implementing todo-api, make sure to follow these documents:
    - docs/Backend.md
- Try to run commands inside api docker container as much as possible for cases like:
    - update / install go mod module
    - running api server

## client
- When implementing client, make sure to follow these documents:
    - docs/Frontend.md
- Try to run commands inside api docker container as much as possible for cases like:
    - update / install npm package
    - running dev server
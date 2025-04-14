# General
- This is an application `ai-todos` that helps organize and manage todos for the user.

---

# Features
## Signup
### General
- Allow user to sign up to ai-todos.

### Flow
#### Success flow
1. User opens sign up page
2. User fills form with email and password
3. User submits form
4. If sign up was a success, user will be redirected to top page of the app

## Login
### General
- Allow user to login to ai-todos.

#### Success flow
1. User opens login page
2. User fills form with email and password
3. User submits form
4. If login was a success, user will be redirected to top page of the app

## Template
### General
- 

#### Success flow
1. a

---

# Domain models

## User
### Properties
- id: int
- name: string
- email: string

## Todo
### Properties
- id: int
- title: string
- status: enum (todo, done)
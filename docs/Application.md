# General
- This is an application `ai-todo` that helps organize and manage todos for the user.

# Domain models
## User
- Properties
    - id: int
    - name: string
    - email: string

## Todo
- Properties
    - id: int
    - title: string
    - status: enum (todo, done)

# Features
## Signup
- **General**: ユーザーがai-todoにユーザー登録するための機能を提供する

- **Flow**:
  1. ユーザーはサインアップページを開く
  2. ユーザーはEメールとパスワードを入力する
  3. ユーザーはフォームを送信する
  4. 登録が成功すると、アプリのトップページにリダイレクトされる
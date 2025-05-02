* Login
    * Overview: Represents the login session and authentication process for a user
    * Properties:
        * email:string
        * rememberMe:boolean
        * loginAttempts:int
        * lastLoginAttempt:datetime
    * Relations:
        * User
            * Login is associated with a User model
    * Validations:
        * email:
            * Must be a valid email format
            * Must match an existing user's email in the system
        * rememberMe:
            * Boolean value indicating whether to maintain a persistent login session
        * loginAttempts:
            * Maximum of 5 consecutive failed attempts before account is temporarily locked
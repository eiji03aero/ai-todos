* UserSession
    * Overview: Represents an active user session after successful authentication
    * Properties:
        * userId:string
        * token:string
        * expiresAt:datetime
        * isRemembered:boolean
        * lastActivity:datetime
    * Relations:
        * User
            * Session is directly linked to a User
    * Validations:
        * token:
            * Must be a unique, securely generated string
        * expiresAt:
            * Must be a future datetime
        * isRemembered:
            * Determines the session's persistence and expiration behavior
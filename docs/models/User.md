* User
    * Overview: Represents a user account in the system
    * Properties:
        * id:string
        * email:string
        * password_hash:string
            * Note: This is the hashed value of the user's actual password, stored securely to protect user credentials
        * createdAt:datetime
        * updatedAt:datetime
    * Relations:
        * No direct relations at this time
    * Validations:
        * email:
            * Must be a valid email format
            * Must be unique in the system
        * password_hash:
            * Must be generated using a secure hashing algorithm (e.g., bcrypt)
            * Original password must meet these criteria before hashing:
                * Minimum 8 characters long
                * Must contain:
                    * At least one uppercase letter
                    * At least one lowercase letter
                    * At least one number
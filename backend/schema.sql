-- User Table
CREATE TABLE Users (
    ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    Username varchar(32),
    Email varchar(255),
    Password varchar(255)
);

-- Get User
SELECT * FROM Users WHERE Username = $1;

-- Register User
INSERT INTO Users (Username, Email, Password) VALUES ($1, $2, $3) 
RETURNING (ID, Username, Email, Password);

-- Change User Password
UPDATE Users 
SET Password = ?
WHERE Username = ? RETURNING (ID, Username, Email, Password);

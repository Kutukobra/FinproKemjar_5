-- Get User
SELECT * FROM Users WHERE Username = $1;

-- Register User
INSERT INTO Users (Username, Email, Password) VALUES ($1, $2, $3) 
RETURNING (ID, Username, Email, Password);

-- Change User Password
UPDATE Users 
SET Password = $2
WHERE Username = $1 RETURNING (ID, Username, Email, Password);

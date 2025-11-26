-- User Table
CREATE TABLE Users (
    ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    Username varchar(32),
    Email varchar(255),
    Password varchar(255)
);
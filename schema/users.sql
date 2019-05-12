CREATE TABLE IF NOT EXISTS users (
       uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
       username STRING UNIQUE NOT NULL,
       password_hash STRING NOT NULL,
       salt STRING NOT NULL,
       is_admin BOOL NOT NULL,

       created_at TIMESTAMP NOT NULL,
       updated_at TIMESTAMP
);

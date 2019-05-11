CREATE TABLE IF NOT EXISTS users (
       id SERIAL,
       username STRING,
       password_hash STRING,
       salt STRING,
       is_admin BOOL,

       created_at TIMESTAMP,
       updated_at TIMESTAMP
);

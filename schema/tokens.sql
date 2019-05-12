CREATE TABLE IF NOT EXISTS tokens (
       id SERIAL UNIQUE,
       token STRING NOT NULL,
       user_id INT NOT NULL REFERENCES users(id),

       created_at TIMESTAMP NOT NULL,
       expires_at TIMESTAMP
);

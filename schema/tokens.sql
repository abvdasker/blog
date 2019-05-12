CREATE TABLE IF NOT EXISTS tokens (
       uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
       token STRING NOT NULL,
       user_uuid UUID NOT NULL REFERENCES users(uuid),

       created_at TIMESTAMP NOT NULL,
       expires_at TIMESTAMP
);

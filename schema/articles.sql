CREATE TABLE IF NOT EXISTS articles (
       uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
       title STRING UNIQUE,
       url_string STRING UNIQUE,
       html STRING,
       tags STRING ARRAY,

       created_at TIMESTAMP NOT NULL,
       updated_at TIMESTAMP
);

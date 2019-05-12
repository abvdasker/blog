CREATE TABLE IF NOT EXISTS articles (
       uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
       title STRING,
       url_string STRING,
       html STRING,
       tags STRING ARRAY,

       created_at TIMESTAMP NOT NULL,
       updated_at TIMESTAMP
);

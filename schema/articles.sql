CREATE TABLE IF NOT EXISTS articles (
       id SERIAL UNIQUE,
       title STRING,
       url_string STRING,
       html STRING,
       tags STRING ARRAY,

       created_at TIMESTAMP NOT NULL,
       updated_at TIMESTAMP
);

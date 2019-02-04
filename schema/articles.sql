CREATE TABLE IF NOT EXISTS articles (
       id SERIAL,
       title STRING,
       url_string STRING,
       html STRING,
       tags STRING ARRAY,

       created_at TIMESTAMP,
       updated_at TIMESTAMP
);

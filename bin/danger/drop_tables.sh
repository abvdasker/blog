cat schema/danger/drop_articles.sql | cockroach sql --insecure -u blog -d blog --echo-sql
cat schema/danger/drop_users.sql | cockroach sql --insecure -u blog -d blog --echo-sql
cat schema/danger/drop_tokens.sql | cockroach sql --insecure -u blog -d blog --echo-sql

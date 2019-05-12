cat schema/articles.sql | cockroach sql --insecure -u blog -d blog --echo-sql
cat schema/users.sql | cockroach sql --insecure -u blog -d blog --echo-sql
cat schema/tokens.sql | cockroach sql --insecure -u blog -d blog --echo-sql

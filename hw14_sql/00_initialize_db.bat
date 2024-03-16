@echo off
set PGPASSWORD=4217

echo Database and user creation...
psql -U postgres -h localhost -p 8080 -c "CREATE DATABASE online_store;"
psql -U postgres -h localhost -p 8080 -c "CREATE USER store_user WITH ENCRYPTED PASSWORD '123456';"
psql -U postgres -h localhost -p 8080 -c "GRANT ALL PRIVILEGES ON DATABASE online_store TO store_user;"

echo Configuring access rights...
psql -U postgres -h localhost -p 8080 -d online_store -c "GRANT USAGE ON SCHEMA public TO store_user;"
psql -U postgres -h localhost -p 8080 -d online_store -c "GRANT CREATE ON SCHEMA public TO store_user;"
psql -U postgres -h localhost -p 8080 -d online_store -c "ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL PRIVILEGES ON TABLES TO store_user;"

echo Done.
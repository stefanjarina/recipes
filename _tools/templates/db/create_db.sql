CREATE USER recipes_${name} WITH PASSWORD 'postgres';
CREATE DATABASE recipes_${name} OWNER recipes_${name};
-- GRANT ALL PRIVILEGES ON DATABASE recipes_${name} TO recipes_${name};

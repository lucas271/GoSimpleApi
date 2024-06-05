
INSERT INTO users(name, created_at, updated_at)  
VALUES(?, ?, ?);
-- name: CreateUser :one

select * from users;
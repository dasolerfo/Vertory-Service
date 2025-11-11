-- name: CreateUser :one 
INSERT INTO users (username, email, password_hash) 
VALUES ($1, $2, $3) 
RETURNING id, username, email, followers_count, following_count;

-- name: GetUserByLogin: one 
SELECT * FROM users
where (email = $1 or username = $1) and password_hash = $2; 
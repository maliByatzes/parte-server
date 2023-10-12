-- name: CreateUser :one
INSERT INTO users (
  username,
  hashed_password,
  email
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET 
  username = $2,
  hashed_password = $3,
  email = $4,
  is_verified = $5,
  is_superuser = $6,
  thumbnail = $7
WHERE id = $1
RETURNING *;
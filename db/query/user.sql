-- name: CreateUser :one
INSERT INTO users (
  username, 
  hashed_password,
  full_name,
  email
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET
  hashed_password = CASE
    WHEN @set_hashed_password::boolean = TRUE THEN @hashed_password
    ELSE hashed_password
  END,
    full_name = CASE
    WHEN @set_full_name = TRUE THEN @full_name
    ELSE full_name
  END,
  email = CASE
    WHEN @set_email = TRUE THEN @email
    ELSE email
  END
WHERE
  username = @username
RETURNING *;
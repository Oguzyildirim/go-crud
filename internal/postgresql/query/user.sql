-- name: SelectUser :one
SELECT * FROM users
WHERE id = @id LIMIT 1;

-- name: SelectUsersByCountry :many
SELECT * FROM users 
WHERE country = @country LIMIT 100;

-- name: InsertUser :one
INSERT INTO users (
  name,
  lastname,
  username,
  country
)
VALUES (
  @name,
  @lastname,
  @username,
  @country
)
RETURNING id;

-- name: UpdateUser :one
UPDATE users SET
  name = @name,
  lastname    = @lastname,
  username  = @username,
  country = @country
WHERE id = @id RETURNING id AS res;

-- name: DeleteUser :one
DELETE FROM users
WHERE  id = @id RETURNING id AS res;

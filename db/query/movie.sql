-- name: GetMovie :one
SELECT * FROM movies
WHERE id = $1 LIMIT 1;

-- name: GetMovies :many
SELECT * FROM movies
ORDER BY name;

-- name: CreateMovie :one
INSERT INTO movies (
  title, genre
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM movies
WHERE id = $1;
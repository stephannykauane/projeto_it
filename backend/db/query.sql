
-- name: GetAuthor :one
SELECT * FROM Usuario
WHERE id = $1 LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM Usuario
ORDER BY nome;

-- name: CreateAuthor :one
INSERT INTO usuario (
  nome, senha
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateAuthor :exec
UPDATE Usuario
  set nome = $2,
  senha = $3
WHERE id = $1;

-- name: DeleteAuthor :exec
DELETE FROM Usuario
WHERE id = $1;


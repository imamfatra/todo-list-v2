-- name: GetAllTodos :many
SELECT id, todo, complated, userid
FROM todos
WHERE userid = $1 AND isdelete = FALSE
ORDER BY id;

-- name: CountAllTodos :one
SELECT COUNT(*) 
FROM todos 
WHERE userid = $1 AND isdelete = FALSE;

-- name: AddaNewTodo :one
INSERT INTO todos (
    todo,
    complated,
    userid
) VALUES (
    $1, $2, $3
) RETURNING id, todo, complated, userid;

-- name: GetSingleaTodos :one
SELECT id, todo, complated, userid
FROM todos
WHERE userid = $1 AND id = $2 AND isdelete = FALSE;

-- name: UpdateStatusComplate :one
UPDATE todos
SET 
    id = $1,
    complated = $2
WHERE userid = $3 AND isdelete = FALSE
RETURNING id, todo, complated, userid;

-- name: DeleteaTodo :one
UPDATE todos
SET 
    id = $1,
    isdelete = TRUE
WHERE userid = $2
RETURNING *;

-- name: GetRandomaTodo :one
SELECT id, todo, complated, userid 
FROM todos
WHERE isdelete = FALSE
ORDER BY RANDOM()
LIMIT 1;

-- name: GetSomeTodos :many
SELECT id, todo, complated, userid
FROM todos
WHERE userid = $1 AND isdelete = FALSE
ORDER BY id
LIMIT $2
OFFSET $3;

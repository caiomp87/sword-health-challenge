-- name: CreateTask :exec
INSERT INTO tasks (id, name, summary) VALUES (?, ?, ?);

-- name: UpdateTask :exec
UPDATE tasks SET name = ?, summary = ? WHERE id = ?;

-- name: DoneTask :exec
UPDATE tasks SET performedAt = ? WHERE id = ?;

-- name: DeleteTask :exec
DELETE FROM tasks WHERE id = ?;

-- name: FindTaskById :one
SELECT * FROM tasks WHERE id = ?;

-- name: FindAllTasks :many
SELECT * FROM tasks;
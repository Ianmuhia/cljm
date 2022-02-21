-- name: GetUserByEmail :one
SELECT *
FROM "users"
WHERE email = @email
  AND "users"."deleted_at" IS NULL
ORDER BY "users"."id"
LIMIT 1;

-- name: GetUserById :one
SELECT *
FROM "users"
WHERE id = @id
  AND "users"."deleted_at" IS NULL
ORDER BY "users"."id"
LIMIT 1;

-- name: ListUsers :many
SELECT *
FROM "users"
WHERE "users"."deleted_at" IS NULL;

-- name: UpdateUser :exec
UPDATE "users"
SET user_name = @user_name,
    full_name = @full_name,
    email     = @email
WHERE id = @id
  AND "users"."deleted_at" IS NULL;

-- name: UpdateUserProfileImage :exec
UPDATE "users"
SET profile_image = @profile_image
WHERE id = @id;


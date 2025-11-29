-- name: GetPractices :many
SELECT * FROM practices;

-- name: GetPractice :one
SELECT * FROM practices
WHERE ID = sqlc.arg(id);

-- name: CreatePractice :one
INSERT INTO practices (name, city, phone, email, owner, practice_code, logo, street_address, facebook, instagram, website)
VALUES (
    sqlc.arg(name),
    sqlc.arg(city),
    sqlc.arg(phone),
    sqlc.arg(email),
    sqlc.narg(owner),
    sqlc.narg(practice_code),
    sqlc.narg(logo),
    sqlc.narg(street_address),
    sqlc.narg(facebook),
    sqlc.narg(instagram),
    sqlc.narg(website)
)
RETURNING *; 

-- name: UpdatePractice :one
UPDATE practices
SET
    name = COALESCE(sqlc.arg(name), name),
    city = COALESCE(sqlc.arg(city), city),
    phone = COALESCE(sqlc.arg(phone), phone),
    email = COALESCE(sqlc.arg(email), email),
    owner = COALESCE(sqlc.narg(owner), owner),
    practice_code = COALESCE(sqlc.narg(practice_code), practice_code),
    logo = COALESCE(sqlc.narg(logo), logo),
    street_address = COALESCE(sqlc.narg(street_address), street_address),
    facebook = COALESCE(sqlc.narg(facebook), facebook),
    instagram = COALESCE(sqlc.narg(instagram), instagram),
    website = COALESCE(sqlc.narg(website), website),
    modified_at = NOW()
WHERE id = sqlc.arg(id)
RETURNING *;
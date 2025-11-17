-- name: CreateAppointment :one
INSERT INTO appointments (id, created_at, first_name, last_name, email, mobile_phone, requested_date, is_emergency, description, appointment_type, is_scheduled, scheduled_date, created_by, scheduled_by, is_cancelled, requested_time, scheduled_time, practice_id, modified_at)
VALUES (
    gen_random_uuid, 
    NOW(),
    sqlc.arg(first_name), 
    sqlc.arg(last_name), 
    sqlc.arg(email), 
    sqlc.arg(mobile_phone),
    sqlc.arg(requested_date), 
    sqlc.arg(is_emergency), 
    sqlc.arg(description), 
    sqlc.arg(appointment_type), 
    sqlc.arg(is_scheduled), 
    sqlc.arg(scheduled_date), 
    sqlc.arg(created_by), 
    sqlc.arg(scheduled_by), 
    sqlc.arg(is_cancelled), 
    sqlc.arg(requested_time), 
    sqlc.arg(scheduled_time), 
    sqlc.arg(practice_id), 
    NOW()
)
RETURNING *;
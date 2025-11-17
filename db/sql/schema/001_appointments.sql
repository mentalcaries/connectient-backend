-- +goose Up
CREATE TABLE appointments (
    id UUID PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT NOT NULL,
    mobile_phone TEXT NOT NULL,
    requested_date DATE NOT NULL,
    is_emergency BOOLEAN DEFAULT FALSE,
    description TEXT,
    appointment_type TEXT,
    is_scheduled BOOLEAN DEFAULT FALSE,
    scheduled_date date,
    created_by UUID,
    scheduled_by UUID,
    is_cancelled BOOLEAN DEFAULT FALSE,
    requested_time TEXT DEFAULT ''::TEXT NOT NULL,
    scheduled_time time without time zone,
    practice_id UUID NOT NULL,
    created_at timestamp with time zone,
    modified_at timestamp with time zone
);

-- +goose Down
DROP TABLE appointments;
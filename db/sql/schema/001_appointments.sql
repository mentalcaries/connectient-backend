-- +goose Up
CREATE TABLE appointments (
    id UUID PRIMARY KEY NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT NOT NULL,
    mobile_phone TEXT NOT NULL,
    requested_date DATE NOT NULL,
    requested_time TEXT DEFAULT ''::TEXT NOT NULL,
    is_emergency BOOLEAN DEFAULT FALSE NOT NULL,
    description TEXT,
    appointment_type TEXT,
    is_scheduled BOOLEAN DEFAULT FALSE NOT NULL,
    scheduled_date date,
    scheduled_time TIME WITHOUT TIME ZONE,
    practice_id UUID NOT NULL,
    created_by UUID,
    scheduled_by UUID,
    is_cancelled BOOLEAN DEFAULT FALSE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    modified_at TIMESTAMP WITH TIME ZONE NOT NULL
);

-- +goose Down
DROP TABLE appointments;
-- +goose Up
CREATE TABLE practices (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    modified_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    name TEXT NOT NULL,
    city TEXT NOT NULL,
    phone TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    owner uuid,
    practice_code TEXT UNIQUE,
    logo TEXT,
    street_address TEXT,
    facebook TEXT,
    instagram TEXT,
    website TEXT
);

-- +goose Down
DROP TABLE practices;
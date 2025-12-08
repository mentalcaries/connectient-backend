-- +goose Up
CREATE TYPE user_role AS ENUM('admin', 'provider', 'staff');

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    modified_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    mobile_phone TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    practice_id UUID NOT NULL,
    role USER_ROLE NOT NULL DEFAULT 'staff',
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    avatar_url TEXT NOT NULL DEFAULT '',
    CONSTRAINT Users_practice_id_fkey FOREIGN KEY (practice_id) REFERENCES practices(id)
);

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION update_modified_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.modified_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

CREATE TRIGGER users_modified_at
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_at();



-- +goose Down
DROP TABLE users;
DROP TYPE user_role;
DROP TRIGGER IF EXISTS users_modified_at ON users;
DROP FUNCTION IF EXISTS update_modified_at();
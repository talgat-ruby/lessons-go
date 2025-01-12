CREATE TABLE IF NOT EXISTS user_ (
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    salt TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_updated_at_user
    BEFORE UPDATE
    ON user_
    FOR EACH ROW
EXECUTE FUNCTION updated_at_column_trigger();

DO
$$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_catalog.pg_type WHERE typname = 'expense_category') THEN
            CREATE TYPE expense_category AS ENUM ('Other', 'Groceries', 'Leisure', 'Electronics', 'Utilities', 'Clothing', 'Health');
        END IF;
    END
$$;

CREATE TABLE IF NOT EXISTS expense (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    amount BIGINT NOT NULL,
    category expense_category NOT NULL DEFAULT 'Other',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    FOREIGN KEY (user_id) REFERENCES user_ (id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TRIGGER set_updated_at_expense
    BEFORE UPDATE
    ON expense
    FOR EACH ROW
EXECUTE FUNCTION updated_at_column_trigger();

CREATE TABLE IF NOT EXISTS customers_interactions.customers
(
    id uuid PRIMARY KEY,
    first_name text NOT NULL,
    last_name text NOT NULL,
    create_at timestamptz NOT NULL,
    modified_at timestamptz NOT NULL
);

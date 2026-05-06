CREATE TABLE orders (
    order_id bigint NOT NULL,
    customer_id bigint NOT NULL,
    created_at timestamptz NOT NULL,
    status text NOT NULL,
    total_cents integer NOT NULL
);

CREATE TABLE lab_notes (
    key text NOT NULL,
    value text NOT NULL
);

INSERT INTO lab_notes (key, value)
VALUES
    ('index_policy', 'No indexes are created by this setup. The orders table also has no primary key because PostgreSQL would create an index for it.'),
    ('target_query', 'SELECT * FROM orders WHERE customer_id = $1 ORDER BY created_at DESC LIMIT 20;');

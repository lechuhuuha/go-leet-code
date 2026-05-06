#!/usr/bin/env bash
set -euo pipefail

row_count="${ORDERS_ROW_COUNT:-20000000}"
customer_count="${CUSTOMER_COUNT:-100000}"

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" \
  -v row_count="$row_count" \
  -v customer_count="$customer_count" <<'SQL'
\timing on

INSERT INTO orders (order_id, customer_id, created_at, status, total_cents)
SELECT
    gs AS order_id,
    ((gs - 1) % :customer_count) + 1 AS customer_id,
    timestamptz '2024-01-01 00:00:00+00' + (gs * interval '17 seconds') AS created_at,
    CASE gs % 4
        WHEN 0 THEN 'paid'
        WHEN 1 THEN 'pending'
        WHEN 2 THEN 'shipped'
        ELSE 'cancelled'
    END AS status,
    ((gs * 37) % 50000)::integer + 100 AS total_cents
FROM generate_series(1, :row_count) AS gs;

ANALYZE orders;

INSERT INTO lab_notes (key, value)
VALUES
    ('orders_row_count', :'row_count'),
    ('customer_count', :'customer_count');
SQL

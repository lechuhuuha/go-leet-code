# Orders Query Lab

This is a local PostgreSQL setup for testing this query:

```sql
SELECT *
FROM orders
WHERE customer_id = $1
ORDER BY created_at DESC
LIMIT 20;
```

The setup intentionally creates no indexes. The `orders` table does not even use a primary key, because PostgreSQL would create an index for it.

## Start

From this directory:

```sh
make start
```

By default, initialization inserts `20,000,000` rows into `orders`. That can take a while and use several GB of disk space.

For a faster first run:

```sh
ORDERS_ROW_COUNT=1000000 make start
```

## Connect

In another terminal:

```sh
make psql
```

From pgAdmin running on your machine, create a server with:

| Field | Value |
| --- | --- |
| Host name/address | `localhost` |
| Port | `55432` |
| Maintenance database | `query_lab` |
| Username | `query_lab` |
| Password | `query_lab` |

If port `55432` is already used on your machine, start the lab on another host port:

```sh
POSTGRES_PORT=55433 make start
```

Then use port `55433` in pgAdmin.

Run the baseline query plan:

```sh
make explain
```

## Reset Data

PostgreSQL initialization scripts only run when the data volume is empty. To rebuild the database from scratch:

```sh
make reset
```

## Notes

- `CUSTOMER_COUNT` defaults to `100000`.
- `customer_id = 1` has many rows and is suitable for testing.
- Add your own indexes or query changes after the baseline is running.

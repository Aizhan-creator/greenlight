CREATE TABLE IF NOT EXISTS candles (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    name text NOT NULL,
    description text NOT NULL,
    price integer NOT NULL,
    runtime integer NOT NULL
);
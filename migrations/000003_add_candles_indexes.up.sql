CREATE INDEX IF NOT EXISTS candles_name_idx ON candles USING GIN (to_tsvector('simple', name));

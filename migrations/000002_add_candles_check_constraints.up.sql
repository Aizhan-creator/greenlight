ALTER TABLE candles ADD CONSTRAINT candles_runtime_check CHECK (runtime >= 0);
ALTER TABLE candles ADD CONSTRAINT candles_price_check CHECK (price >= 0);
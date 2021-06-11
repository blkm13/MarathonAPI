-- Write your migrate up statements here
--

CREATE TABLE IF NOT EXISTS events
(name VARCHAR PRIMARY KEY, date VARCHAR, key VARCHAR)
-- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.

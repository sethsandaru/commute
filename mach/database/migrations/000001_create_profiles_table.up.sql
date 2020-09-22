BEGIN;

CREATE TABLE IF NOT EXISTS profiles (
    id serial PRIMARY KEY,
    user_id integer NOT NULL,
    name varchar(50),
    created_at timestamp,
    updated_at timestamp
);

COMMIT;
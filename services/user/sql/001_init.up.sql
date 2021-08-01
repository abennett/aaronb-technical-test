CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    password bytea NOT NULL,
    tags text[] NOT NULL
);


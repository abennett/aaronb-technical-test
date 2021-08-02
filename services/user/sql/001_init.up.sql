CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    name TEXT UNIQUE NOT NULL,
    password bytea NOT NULL,
    tags text[] DEFAULT '{}' NOT NULL
);


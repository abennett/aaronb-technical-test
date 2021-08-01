CREATE TABLE IF NOT EXISTS news (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    "timestamp" TIMESTAMPTZ NOT NULL,
    tags text[] NOT NULL
);

CREATE INDEX news_tags on news USING GIN (tags);

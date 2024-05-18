DROP TYPE IF EXISTS role CASCADE;
CREATE TYPE role AS ENUM ('root', 'admin', 'readonly');

DROP TABLE IF EXISTS api_keys;
CREATE TABLE api_keys (
    id VARCHAR PRIMARY KEY,
    key VARCHAR NOT NULL UNIQUE,
    role role NOT NULL,
    created_at TIMESTAMPTZ NOT NULL
);

DROP TABLE IF EXISTS flags;
CREATE TABLE flags (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR NOT NULL UNIQUE,
    enabled     BOOLEAN NOT NULL
);

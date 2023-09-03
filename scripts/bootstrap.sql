CREATE TYPE role AS ENUM ('root', 'admin', 'readonly');

CREATE TABLE api_keys (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL UNIQUE,
    api_key VARCHAR NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL,
    role role NOT NULL
);

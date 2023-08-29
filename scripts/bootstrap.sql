CREATE TABLE api_keys (
    name VARCHAR primary key,
    api_key VARCHAR NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL
);

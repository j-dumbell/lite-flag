CREATE TYPE role AS ENUM ('root', 'admin', 'readonly');

CREATE TABLE api_keys (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL UNIQUE,
    api_key VARCHAR NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL,
    role role NOT NULL
);

CREATE TABLE flags (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE transitions (
    id SERIAL PRIMARY KEY,
    flag_id INTEGER NOT NULL,
    to_state BOOLEAN NOT NULL,
    effective_from TIMESTAMP NOT NULL,
    CONSTRAINT fk_flag
        FOREIGN KEY(flag_id)
            REFERENCES flags(id)
);

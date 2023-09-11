CREATE TYPE role AS ENUM ('root', 'admin', 'readonly');

CREATE TABLE api_keys (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL UNIQUE,
    api_key VARCHAR NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL,
    role role NOT NULL
);

CREATE TABLE flags (
    name VARCHAR PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE transitions (
    id SERIAL PRIMARY KEY,
    flag_name VARCHAR NOT NULL,
    to_state BOOLEAN NOT NULL,
    effective_from TIMESTAMPTZ NOT NULL,
    CONSTRAINT fk_flag
        FOREIGN KEY(flag_name)
            REFERENCES flags(name)
);

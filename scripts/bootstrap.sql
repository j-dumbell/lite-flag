CREATE TYPE role AS ENUM ('root', 'admin', 'readonly');

CREATE TABLE api_keys (
    id VARCHAR PRIMARY KEY,
    key VARCHAR NOT NULL UNIQUE,
    role role NOT NULL,
    created_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE flags (
    id VARCHAR PRIMARY KEY,
    enabled BOOLEAN NOT NULL,
    created_at TIMESTAMPTZ NOT NULL
);

-- CREATE TABLE transitions (
--     id SERIAL PRIMARY KEY,
--     flag_name VARCHAR NOT NULL,
--     to_state BOOLEAN NOT NULL,
--     effective_from TIMESTAMPTZ NOT NULL,
--     CONSTRAINT fk_flag
--         FOREIGN KEY(flag_name)
--             REFERENCES flags(name)
-- );

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    password BYTEA NOT NULL,
    email TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT users_email_check_length CHECK (LENGTH(email) <= 256),
    CONSTRAINT users_name_check_length CHECK (LENGTH(name) <= 32)
);
CREATE TABLE IF NOT EXISTS folders (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    user_id INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT folders_name_check_key UNIQUE (name, user_id),
    CONSTRAINT folders_name_check_length CHECK (LENGTH(name) <= 32),
    CONSTRAINT folders_description_check_length CHECK (LENGTH(description) <= 256)
);
CREATE TABLE IF NOT EXISTS words (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    photo BYTEA,
    usage_examples TEXT[],
    folder_id INTEGER NOT NULL REFERENCES folders (id) ON DELETE CASCADE,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT words_name_check_key UNIQUE (name, folder_id),
    CONSTRAINT words_name_check_length CHECK (LENGTH(name) <= 64),
    CONSTRAINT words_description_check_length CHECK (LENGTH(description) <= 256)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS words;
DROP TABLE IF EXISTS folders;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd

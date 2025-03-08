-- +goose Up
-- +goose StatementBegin
CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    created_by INTEGER,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP ,
    
    FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE roles DROP CONSTRAINT roles_created_by_fkey;
DROP TABLE roles;
-- +goose StatementEnd

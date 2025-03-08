-- +goose Up
-- +goose StatementBegin
CREATE TABLE role_has_permissions (
    id SERIAL PRIMARY KEY,
    role_id INTEGER,
    permission_id INTEGER,
    
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,
    FOREIGN KEY (permission_id) REFERENCES permissions(id) ON DELETE CASCADE
    
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE role_has_permissions DROP CONSTRAINT role_has_permissions_role_id_fkey;
ALTER TABLE role_has_permissions DROP CONSTRAINT role_has_permissions_permission_id_fkey;
DROP TABLE role_has_permissions;
-- +goose StatementEnd

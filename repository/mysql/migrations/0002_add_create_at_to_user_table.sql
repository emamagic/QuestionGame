-- +migrate Up
ALTER TABLE
    users
add
    column create_at TIMESTAMP DEFAULT NOW();

-- +migrate Down
ALTER TABLE
    users DROP column create_at;
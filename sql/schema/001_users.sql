-- +goose Up
create table users(
    id varchar(64) PRIMARY KEY default(uuid()),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL
);


-- +goose Down

drop table users;
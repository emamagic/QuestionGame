-- +migrate Up
CREATE TABLE users(
    id int primary key AUTO_INCREMENT,
    name varchar(191) not null,
    phone_number varchar(191) not null unique,
    password varchar(191) not null
);

-- +migrate Down
DROP TABLE users;
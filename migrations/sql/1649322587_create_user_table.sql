-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users(
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
	name varchar(255) not null,
	username varchar(255) not null UNIQUE,
	email varchar(255) UNIQUE,
	phone varchar(20) UNIQUE,
	password varchar(255) not null,
	created_at timestamp with time zone not null,
    updated_at timestamp with time zone,
	deleted_at timestamp with time zone,
	created_by varchar(100) not null,
    updated_by varchar(100),
	deleted_by varchar(100)
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS users;
 	
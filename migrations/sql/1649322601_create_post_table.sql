-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TYPE post_status AS ENUM ('draft', 'published', 'outdated');

CREATE TABLE posts(
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
	title varchar(255) not null,
	content text not null,
	author_id uuid not null,
	status post_status not null,
	tags varchar(255),
	created_at timestamp with time zone not null,
    updated_at timestamp with time zone,
	deleted_at timestamp with time zone,
	created_by varchar(100) not null,
    updated_by varchar(100),
	deleted_by varchar(100),
	CONSTRAINT fk_author
      FOREIGN KEY(author_id) 
	  REFERENCES users(id)
	  ON DELETE CASCADE
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS posts;
DROP TYPE IF EXISTS post_status;
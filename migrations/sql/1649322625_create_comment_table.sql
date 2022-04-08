-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TYPE comment_status AS ENUM ('approved', 'not_approved');

CREATE TABLE comments(
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
	author varchar(255) not null,
	content text not null,
	status comment_status not null,
	post_id uuid not null,
	created_at timestamp with time zone not null,
    updated_at timestamp with time zone,
	deleted_at timestamp with time zone,
	created_by varchar(100) not null,
    updated_by varchar(100),
	deleted_by varchar(100),
	CONSTRAINT fk_post
      FOREIGN KEY(post_id) 
	  REFERENCES posts(id)
	  ON DELETE CASCADE
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS comments;
DROP TYPE IF EXISTS comment_status;
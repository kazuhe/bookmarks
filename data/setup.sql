drop table users;

CREATE TABLE users (
  user_id     VARCHAR(64) PRIMARY KEY,
  name        VARCHAR(255) NOT NULL,
  email       VARCHAR(255) NOT NULL UNIQUE,
  password    VARCHAR(255) NOT NULL,
  created_at  TIMESTAMP NOT NULL,
  twitter_id  VARCHAR(255),
  is_public   BOOLEAN
);

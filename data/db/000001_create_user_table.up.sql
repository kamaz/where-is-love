CREATE TYPE gender AS ENUM ('male', 'female');

CREATE TABLE app_user (
  id SERIAL PRIMARY KEY,
  email varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  name varchar(255) NOT NULL,
  gender gender NOT NULL,
  age int NOT NULL,
  created_at timestamp with time zone NOT NULL DEFAULT now(),
  updated_at timestamp with time zone NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX app_user_email_idx ON app_user (email);
CREATE INDEX app_user_email_pwd_idx ON app_user (email, password);


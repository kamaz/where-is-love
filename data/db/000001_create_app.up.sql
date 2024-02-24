CREATE TYPE gender_type AS ENUM ('male', 'female');

CREATE TABLE app_user (
  id SERIAL PRIMARY KEY,
  email varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  name varchar(255) NOT NULL,
  gender gender_type NOT NULL,
  age int NOT NULL,
  created_at timestamp with time zone NOT NULL DEFAULT now(),
  updated_at timestamp with time zone NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX app_user_email_idx ON app_user (email);
CREATE INDEX app_user_email_pwd_idx ON app_user (email, password);

CREATE TYPE preference_type AS ENUM ('YES', 'NO');
CREATE TABLE user_match (
  from_id int NOT NULL,
  to_id int NOT NULL,
  preference preference_type NOT NULL,
  created_at timestamp with time zone NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX user_match_from_to_id_idx ON user_match (from_id, to_id);

-- enable extensions to support calculating distance
CREATE EXTENSION cube; 
CREATE EXTENSION earthdistance;

-- add location
ALTER TABLE app_user ADD latitude FLOAT8;
ALTER TABLE app_user ADD longitude FLOAT8;
ALTER TABLE app_user ADD city varchar(255);

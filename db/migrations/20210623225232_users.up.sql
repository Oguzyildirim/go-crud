CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
  id          UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  name    VARCHAR NOT NULL,
  lastname    VARCHAR NOT NULL,
  username    VARCHAR NOT NULL,
  country    VARCHAR NOT NULL
);

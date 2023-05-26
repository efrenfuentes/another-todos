CREATE TABLE IF NOT EXISTS  todos (
  id  SERIAL PRIMARY KEY,
  title  VARCHAR ( 255 )  NOT NULL,
  completed  BOOLEAN  NOT NULL  DEFAULT  false,
  "order"  INT NOT NULL DEFAULT 0,
  url  VARCHAR ( 255 )  NOT NULL DEFAULT '',

  created_at  TIMESTAMP  NOT NULL  DEFAULT  now(),
  updated_at  TIMESTAMP  NOT NULL  DEFAULT  now()
);

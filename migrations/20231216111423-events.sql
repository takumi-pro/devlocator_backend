
-- +migrate Up
CREATE TABLE events (
  event_id INT PRIMARY KEY NOT NULL,
  title VARCHAR(255) NOT NULL,
  `catch` VARCHAR(255),
  description TEXT,
  event_url VARCHAR(255),
  started_at DATETIME,
  ended_at DATETIME,
  `limit` INT,
  hash_tag VARCHAR(100),
  event_type VARCHAR(50),
  accepted INT,
  waiting INT,
  updated_at DATETIME,
  owner_id INT,
  owner_nickname VARCHAR(255),
  owner_display_name VARCHAR(255),
  place VARCHAR(255),
  address VARCHAR(255),
  lat VARCHAR(255) NOT NULL,
  lon VARCHAR(255) NOT NULL
);

-- +migrate Down
DROP TABLE events;
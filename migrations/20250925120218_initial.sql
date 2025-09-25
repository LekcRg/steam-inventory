-- +goose Up
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  steamid VARCHAR(255) UNIQUE NOT NULL,
  personaname VARCHAR(255) NOT NULL,
  avatar VARCHAR(255) NOT NULL,
  realname VARCHAR(255) NOT NULL,
  communityvisibilitystate INTEGER NOT NULL,
  lastlogoff_steam INTEGER NOT NULL,
  timecreated_steam INTEGER NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE users;

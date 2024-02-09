
CREATE TABLE IF NOT EXISTS location (
  location_id SERIAL PRIMARY KEY,
  country VARCHAR(150) NOT NULL,
  latitude DOUBLE PRECISION,
  location VARCHAR(150) NOT NULL,
  longitude DOUBLE PRECISION
);

CREATE TABLE IF NOT EXISTS place (
  place_id SERIAL PRIMARY KEY,
  location_id INT NOT NULL, -- FK
  FOREIGN KEY (location_id) REFERENCES location (location_id),
  description TEXT,
  end_time VARCHAR(150) NOT NULL,
  kind VARCHAR(150) NOT NULL,
  latest_views INT DEFAULT 0,
  place_name VARCHAR(150) NOT NULL,
  start_time VARCHAR(150) NOT NULL,
  total_view INT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS weather (
  weather_id SERIAL PRIMARY KEY,
  location_id INT NOT NULL, -- FK
  FOREIGN KEY (location_id) REFERENCES location (location_id),
  description TEXT,
  temperature DOUBLE PRECISION,
  temperature_max DOUBLE PRECISION,
  temperature_min DOUBLE PRECISION
);

CREATE TABLE IF NOT EXISTS user_account (
  user_id SERIAL PRIMARY KEY,
  email VARCHAR(150) NOT NULL UNIQUE,
  gender VARCHAR(150) NOT NULL,
  user_lastname VARCHAR(150) NOT NULL,
  user_name VARCHAR(150) NOT NULL,
  username VARCHAR(150) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS comment (
  comment_id SERIAL PRIMARY KEY,
  place_id INT NOT NULL,
  FOREIGN KEY (place_id) REFERENCES place (place_id),
  user_id INT NOT NULL,
  FOREIGN KEY (user_id) REFERENCES user_account (user_id),
  comment TEXT NOT NULL
);

-- Agregar la dirección:
ALTER TABLE place
ADD COLUMN address VARCHAR(150) NOT NULL;

ALTER TABLE user_account
ADD CONSTRAINT valid_gender
CHECK (gender IN ('male', 'female', 'other'));

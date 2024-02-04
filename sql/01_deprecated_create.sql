
CREATE TABLE IF NOT EXISTS places (
    place_id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    kind VARCHAR(150) NOT NULL,
    country VARCHAR(150) NOT NULL,
    location VARCHAR(150) NOT NULL,
    address VARCHAR(150) NOT NULL,
    start_time VARCHAR(150) NOT NULL,
    end_time VARCHAR(150) NOT NULL,
    description TEXT
);

CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    lastname VARCHAR(150) NOT NULL,
    email VARCHAR(150) NOT NULL UNIQUE,
    username VARCHAR(150) NOT NULL,
    gender VARCHAR(150) NOT NULL
);

CREATE TABLE IF NOT EXISTS comments (
    comment_id SERIAL PRIMARY KEY,
    place_id INT NOT NULL,
    user_id INT NOT NULL,
    comment TEXT NOT NULL,
    FOREIGN KEY (place_id) REFERENCES places (place_id),
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);

CREATE TABLE IF NOT EXISTS weathers (
  weather_id SERIAL PRIMARY KEY,
  place_id INT NOT NULL,
  FOREIGN KEY (place_id) REFERENCES places (place_id),
  temperature_min DOUBLE PRECISION,
  temperature_max DOUBLE PRECISION,
  temperature DOUBLE PRECISION,
  description TEXT
);

-- Me olvide de colocar el campo commment! 
ALTER TABLE comments
ADD COLUMN comment TEXT NOT NULL;

-- Cantidad de vistas del lugar
ALTER TABLE places
ADD COLUMN total_view INT DEFAULT 0;

-- Cantidad de vistas del lugar por minuto
ALTER TABLE places
ADD COLUMN latest_views INT DEFAULT 0;

-- TODAVIA SIN IMPLEMENTAR

CREATE TABLE IF NOT EXISTS locations (
  location_id SERIAL PRIMARY KEY,
  place_id INT NOT NULL,
  FOREIGN KEY (place_id) REFERENCES places (place_id),
  country VARCHAR(150) NOT NULL,
  address VARCHAR(150) NOT NULL,
  latitude DOUBLE PRECISION,
  longitude DOUBLE PRECISION
);


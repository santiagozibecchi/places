
CREATE TABLE IF NOT EXISTS places (
    place SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    kind VARCHAR(150) NOT NULL,
    country VARCHAR(150) NOT NULL,
    location VARCHAR(150) NOT NULL,
    address VARCHAR(150) NOT NULL,
    startTime VARCHAR(150) NOT NULL,
    endTime VARCHAR(150) NOT NULL,
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
    place INT NOT NULL,
    user_id INT NOT NULL,
    FOREIGN KEY (place) REFERENCES places (place),
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);

INSERT INTO places (kind, name, country, location, address, startTime, endTime, description)
VALUES
  ('pubs', 'El Galpon de la pizza','Argentina', 'Buenos Aires', 'España 123', '18:00', '23:00', 'La mejor Pizza de Buenos Aires'),
  ('pubs', 'Es para vos','Argentina', 'Buenos Aires', 'Mendoza 258', '18:00', '23:00', 'Te gusta tomar, este lugar es para vos!'),
  ('pubs', 'El triunfador','Argentina', 'Buenos Aires', 'Av. Corrientes 555', '18:00', '23:00', 'Ideal para pasar un rato con amigos!'),
  ('pubs', 'Mar del descanso','Argentina', 'Mar del Plata', 'Address2', '07:00', '02:00', 'Frente a la orilla del mar, especial para disfrutar la playa'),
  ('pubs', 'Conejo Negro','Argentina', 'Resistencia', 'España 578', '08:00', '23:00', 'Queres tragos especiales, veni y divertite que tenemos de sobra!');


INSERT INTO users (name, lastname, email, username, gender)
VALUES
  ('John', 'Doe', 'john.doe@example.com', 'johndoe', 'Male'),
  ('Jane', 'Smith', 'jane.smith@example.com', 'janesmith', 'Female'),
  ('Robert', 'Johnson', 'robert.johnson@example.com', 'robertjohnson', 'Male'),
  ('Emily', 'Williams', 'emily.williams@example.com', 'emilywilliams', 'Female'),
  ('Michael', 'Brown', 'michael.brown@example.com', 'michaelbrown', 'Male');
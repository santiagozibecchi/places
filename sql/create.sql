
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

-- Me olvide de colocar el campo commment! 
ALTER TABLE comments
ADD COLUMN comment TEXT NOT NULL;

-- Cantidad de vistas del lugar
ALTER TABLE places
ADD COLUMN total_view INT DEFAULT 0;

-- SEED

-- Pubs
INSERT INTO places (kind, name, country, location, address, start_time, end_time, description)
VALUES
  ('pubs', 'El Galpon de la pizza', 'Argentina', 'Buenos Aires', 'España 123', '18:00', '23:00', 'La mejor Pizza de Buenos Aires'),
  ('pubs', 'Es para vos', 'Argentina', 'Buenos Aires', 'Mendoza 258', '18:00', '23:00', 'Te gusta tomar, este lugar es para vos!'),
  ('pubs', 'El triunfador', 'Argentina', 'Buenos Aires', 'Av. Corrientes 555', '18:00', '23:00', 'Ideal para pasar un rato con amigos!'),
  ('pubs', 'Mar del descanso', 'Argentina', 'Mar del Plata', 'Colon 1200', '07:00', '02:00', 'Frente a la orilla del mar, especial para disfrutar la playa'),
  ('pubs', 'Conejo Negro', 'Argentina', 'Resistencia', 'España 578', '08:00', '23:00', 'Queres tragos especiales, veni y divertite que tenemos de sobra');

-- Restaurants
INSERT INTO places (kind, name, country, location, address, start_time, end_time, description)
VALUES
  ('restaurants', 'La Parrilla de Juan', 'Argentina', 'Cordoba', 'Calle 25 de Mayo 123', '12:00', '23:00', 'Las mejores carnes asadas en leña'),
  ('restaurants', 'Sabor a Mar', 'Peru', 'Lima', 'Av. Larco 789', '11:00', '22:00', 'Especialidades de mariscos y pescados frescos'),
  ('restaurants', 'Ristorante Italiano', 'Italy', 'Rome', 'Via Roma 456', '18:00', '23:00', 'Auténtica cocina italiana en el corazón de Roma'),
  ('restaurants', 'Asian Fusion', 'Japan', 'Tokyo', 'Shibuya Crossing 789', '17:00', '01:00', 'Mezcla de sabores asiáticos en un ambiente moderno'),
  ('restaurants', 'Vegan Delight', 'USA', 'New York', 'Broadway 123', '10:00', '21:00', 'Platos veganos creativos y deliciosos');

-- Parties
INSERT INTO places (kind, name, country, location, address, start_time, end_time, description)
VALUES
  ('parties', 'Electro Lounge', 'Germany', 'Berlin', 'Friedrichstrasse 456', '22:00', '04:00', 'Música electrónica y ambiente vibrante'),
  ('parties', 'Salsa Night', 'Cuba', 'Havana', 'Calle Ocho 789', '20:00', '02:00', 'Noche de salsa con música en vivo y clases de baile'),
  ('parties', 'Rooftop Fiesta', 'Mexico', 'Mexico City', 'Paseo de la Reforma 123', '19:00', '03:00', 'Fiesta en la azotea con impresionantes vistas de la ciudad'),
  ('parties', 'Karaoke Kingdom', 'South Korea', 'Seoul', 'Myeongdong 456', '21:00', '02:00', 'Karaoke en privado y cócteles para una noche divertida'),
  ('parties', 'Carnival Extravaganza', 'Brazil', 'Rio de Janeiro', 'Copacabana 789', '18:00', '04:00', 'Desfile de samba, disfraces y diversión sin fin');
  

INSERT INTO users (name, lastname, email, username, gender)
VALUES
  ('María', 'Gómez', 'maria.gomez@example.com', 'mariagomez', 'female'),
  ('Javier', 'Martínez', 'javier.martinez@example.com', 'javiermartinez', 'male'),
  ('Sofía', 'López', 'sofia.lopez@example.com', 'sofialopez', 'female'),
  ('Alejandro', 'Rodríguez', 'alejandro.rodriguez@example.com', 'alejandrorodriguez', 'male'),
  ('Carmen', 'Hernández', 'carmen.hernandez@example.com', 'carmenhernandez', 'female'),
  ('David', 'García', 'david.garcia@example.com', 'davidgarcia', 'male'),
  ('Laura', 'Fernández', 'laura.fernandez@example.com', 'laurafernandez', 'female'),
  ('Manuel', 'Pérez', 'manuel.perez@example.com', 'manuelperez', 'male'),
  ('Isabel', 'Díaz', 'isabel.diaz@example.com', 'isabeldiaz', 'female'),
  ('Adrián', 'Martín', 'adrian.martin@example.com', 'adrianmartin', 'male'),
  ('Ana', 'Rojas', 'ana.rojas@example.com', 'anarojas', 'female'),
  ('Hugo', 'Sánchez', 'hugo.sanchez@example.com', 'hugosanchez', 'male'),
  ('Clara', 'Ramírez', 'clara.ramirez@example.com', 'clararamirez', 'female'),
  ('Diego', 'Vega', 'diego.vega@example.com', 'diegovega', 'male'),
  ('Elena', 'Torres', 'elena.torres@example.com', 'elenatorres', 'female'),
  ('Iván', 'Gutiérrez', 'ivan.gutierrez@example.com', 'ivangutierrez', 'male'),
  ('Luisa', 'Navarro', 'luisa.navarro@example.com', 'luisanavarro', 'female'),
  ('Fernando', 'Cruz', 'fernando.cruz@example.com', 'fernandocruz', 'male'),
  ('Marta', 'Ortega', 'marta.ortega@example.com', 'martaortega', 'female'),
  ('Juan', 'Reyes', 'juan.reyes@example.com', 'juanreyes', 'male');

pg_dump
psql -h localhost -p 5432 -U role_prueba -d prueba2 < create.sql
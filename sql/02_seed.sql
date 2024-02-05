-- Lugares en Resistencia
INSERT INTO location (country, latitude, location, longitude)
VALUES ('Argentina', NULL, 'Resistencia', NULL)
RETURNING location_id;

-- Lugares en Resistencia
INSERT INTO place (location_id, description, end_time, kind, latest_views, place_name, start_time, total_view, address)
VALUES
  ((SELECT location_id FROM location WHERE location = 'Resistencia' LIMIT 1), 'Encantador lugar con especialidades locales.', '21:00', 'restaurant', 45, 'La Esquina del Sabor', '17:30', 85, 'Calle Resistencia 123'),
  ((SELECT location_id FROM location WHERE location = 'Resistencia' LIMIT 1), 'Ambiente acogedor con cocina internacional.', '22:30', 'restaurant', 28, 'Mundo Gourmet', '19:00', 75, 'Av. Independencia 456'),
  ((SELECT location_id FROM location WHERE location = 'Resistencia' LIMIT 1), 'Sabores únicos de la región en cada plato.', '20:30', 'restaurant', 38, 'Raíces del Chaco', '18:00', 92, 'Calle Chacabuco 789'),
  ((SELECT location_id FROM location WHERE location = 'Resistencia' LIMIT 1), 'Experiencia gastronómica con influencia asiática.', '23:00', 'restaurant', 22, 'Asian Fusion', '18:30', 68, 'Av. Belgrano 234'),
  ((SELECT location_id FROM location WHERE location = 'Resistencia' LIMIT 1), 'Postres irresistibles en un entorno moderno.', '19:30', 'restaurant', 32, 'Dulce Encanto', '15:00', 55, 'Calle Pellegrini 567');

-- Weather en Resistencia
INSERT INTO weather (location_id, description, temperature, temperature_max, temperature_min)
SELECT location_id, NULL, NULL, NULL, NULL
FROM location
WHERE location = 'Resistencia'
LIMIT 1;




-- Lugares en Mar del Plata
INSERT INTO location (country, latitude, location, longitude)
VALUES ('Argentina', NULL, 'Mar del Plata', NULL)
RETURNING location_id;

-- Lugares en Mar del Plata
INSERT INTO place (location_id, description, end_time, kind, latest_views, place_name, start_time, total_view, address)
VALUES
  ((SELECT location_id FROM location WHERE location = 'Mar del Plata' LIMIT 1), 'Vistas al mar con deliciosos platos locales.', '22:00', 'restaurant', 0, 'La Costa del Sabor', '18:30', 0, 'Av. Costanera 123'),
  ((SELECT location_id FROM location WHERE location = 'Mar del Plata' LIMIT 1), 'Cocina mediterránea en un ambiente relajado.', '23:30', 'restaurant', 0, 'Mediterráneo Lounge', '19:00', 0, 'Calle Playa Grande 456'),
  ((SELECT location_id FROM location WHERE location = 'Mar del Plata' LIMIT 1), 'Mariscos frescos y auténticos sabores marplatenses.', '21:00', 'restaurant', 0, 'El Puerto de los Sabores', '17:30', 0, 'Av. Colón 789'),
  ((SELECT location_id FROM location WHERE location = 'Mar del Plata' LIMIT 1), 'Comida rápida con toques de estilo gourmet.', '22:45', 'restaurant', 0, 'Gourmet Fast Bites', '18:45', 0, 'Calle San Martín 234'),
  ((SELECT location_id FROM location WHERE location = 'Mar del Plata' LIMIT 1), 'Café con vistas panorámicas y repostería exquisita.', '20:30', 'restaurant', 0, 'Panorama Café', '15:30', 0, 'Av. Playa Bristol 567');

-- Weather en Mar del Plata
INSERT INTO weather (location_id, description, temperature, temperature_max, temperature_min)
SELECT location_id, NULL, NULL, NULL, NULL
FROM location
WHERE location = 'Mar del Plata'
LIMIT 1;




-- Lugares en Buenos Aires
INSERT INTO location (country, latitude, location, longitude)
VALUES ('Argentina', NULL, 'Buenos Aires', NULL)
RETURNING location_id;

-- Lugares en Buenos Aires
INSERT INTO place (location_id, description, end_time, kind, latest_views, place_name, start_time, total_view, address)
VALUES
  ((SELECT location_id FROM location WHERE location = 'Buenos Aires' LIMIT 1), 'Excelente lugar para disfrutar de la comida argentina.', '22:00', 'restaurant', 0, 'La Parrilla de Juan', '18:00', 0, 'Av. Corrientes 123'),
  ((SELECT location_id FROM location WHERE location = 'Buenos Aires' LIMIT 1), 'Auténtica cocina italiana en el corazón de Buenos Aires.', '23:00', 'restaurant', 0, 'Trattoria Bella Italia', '19:00', 0, 'Calle Florida 456'),
  ((SELECT location_id FROM location WHERE location = 'Buenos Aires' LIMIT 1), 'Sabores tradicionales argentinos en un ambiente acogedor.', '21:30', 'restaurant', 0, 'El Asador Criollo', '17:30', 0, 'Av. de Mayo 789'),
  ((SELECT location_id FROM location WHERE location = 'Buenos Aires' LIMIT 1), 'Amplia variedad de sushi fresco y delicioso.', '22:30', 'restaurant', 0, 'Sushi Express', '18:30', 0, 'Calle Sarmiento 234'),
  ((SELECT location_id FROM location WHERE location = 'Buenos Aires' LIMIT 1), 'Cafetería con encanto y exquisitos postres.', '20:00', 'restaurant', 0, 'Dulces Sueños', '15:00', 0, 'Av. Córdoba 567');

-- Weather en Buenos Aires
INSERT INTO weather (location_id, description, temperature, temperature_max, temperature_min)
SELECT location_id, NULL, NULL, NULL, NULL
FROM location
WHERE location = 'Buenos Aires'
LIMIT 1;





-- Lugares en Corrientes
INSERT INTO location (country, latitude, location, longitude)
VALUES ('Argentina', NULL, 'Corrientes', NULL)
RETURNING location_id;

-- Lugares de fiesta en Corrientes
INSERT INTO place (location_id, description, end_time, kind, latest_views, place_name, start_time, total_view, address)
VALUES
  ((SELECT location_id FROM location WHERE location = 'Corrientes' LIMIT 1), 'Ambiente vibrante con música en vivo.', '03:00', 'party', 0, 'Fiesta Loca', '21:00', 0, 'Calle Corrientes 123'),
  ((SELECT location_id FROM location WHERE location = 'Corrientes' LIMIT 1), 'DJ en vivo y pista de baile espectacular.', '04:00', 'party', 0, 'Club Nocturno Electrónico', '22:30', 0, 'Av. Libertad 456'),
  ((SELECT location_id FROM location WHERE location = 'Corrientes' LIMIT 1), 'Fiesta al aire libre con luces deslumbrantes.', '02:30', 'party', 0, 'Under the Stars', '20:00', 0, 'Calle Belgrano 789'),
  ((SELECT location_id FROM location WHERE location = 'Corrientes' LIMIT 1), 'Celebración festiva con cócteles exclusivos.', '03:30', 'party', 0, 'Fiesta Elegante', '21:30', 0, 'Av. Rivadavia 234'),
  ((SELECT location_id FROM location WHERE location = 'Corrientes' LIMIT 1), 'Noche de karaoke y diversión sin fin.', '02:00', 'party', 0, 'Karaoke Paradise', '19:30', 0, 'Calle San Martín 567');

-- Weather en Corrientes
INSERT INTO weather (location_id, description, temperature, temperature_max, temperature_min)
SELECT location_id, NULL, NULL, NULL, NULL
FROM location
WHERE location = 'Corrientes'
LIMIT 1;

-- Lugares en Buenos Aires
INSERT INTO location (country, latitude, location, longitude)
VALUES ('Argentina', NULL, 'Buenos Aires', NULL)
RETURNING location_id;




-- Pubs en Buenos Aires
INSERT INTO place (location_id, description, end_time, kind, latest_views, place_name, start_time, total_view, address)
VALUES
  ((SELECT location_id FROM location WHERE location = 'Buenos Aires' LIMIT 1), 'Pub tradicional con cervezas artesanales.', '02:00', 'pub', 0, 'The Crafty Pint', '19:00', 0, 'Av. Buenos Aires 123'),
  ((SELECT location_id FROM location WHERE location = 'Buenos Aires' LIMIT 1), 'Ambiente acogedor con música en vivo.', '03:00', 'pub', 0, 'Piano Pub', '20:30', 0, 'Calle San Telmo 456'),
  ((SELECT location_id FROM location WHERE location = 'Buenos Aires' LIMIT 1), 'Pub irlandés con selección única de cervezas.', '01:30', 'pub', 0, 'Irish Cheers', '18:00', 0, 'Av. Mayo 789'),
  ((SELECT location_id FROM location WHERE location = 'Buenos Aires' LIMIT 1), 'Pub moderno con cócteles creativos.', '02:45', 'pub', 0, 'Mixology Pub', '19:45', 0, 'Calle Palermo 234'),
  ((SELECT location_id FROM location WHERE location = 'Buenos Aires' LIMIT 1), 'Noche de juegos y diversión en este pub.', '02:15', 'pub', 0, 'Game Night Pub', '18:45', 0, 'Av. Recoleta 567');

-- Weather en Buenos Aires
INSERT INTO weather (location_id, description, temperature, temperature_max, temperature_min)
SELECT location_id, NULL, NULL, NULL, NULL
FROM location
WHERE location = 'Buenos Aires'
LIMIT 1;



-- USUARIOS
INSERT INTO user_account (email, gender, user_lastname, user_name, username)
VALUES
  ('juan_gomez@example.com', 'male', 'Gómez', 'Juan', 'juan_gomez'),
  ('maria_fernandez@example.com', 'female', 'Fernández', 'María', 'maria_fernandez'),
  ('carlos_lopez@example.com', 'male', 'López', 'Carlos', 'carlos_lopez'),
  ('laura_martinez@example.com', 'female', 'Martínez', 'Laura', 'laura_martinez'),
  ('pablo_rodriguez@example.com', 'male', 'Rodríguez', 'Pablo', 'pablo_rodriguez');


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


-- SEEDING DATA INTO measurement_types
DELETE FROM measurement_types;

INSERT INTO measurement_types (type_id, name)
VALUES
    (gen_random_uuid(), 'weight'),
    (gen_random_uuid(), 'volume'),
    (gen_random_uuid(), 'length'),
    (gen_random_uuid(), 'temperature');

-- SEEDING DATA INTO measurements
DELETE FROM measurements;

INSERT INTO measurements (measurement_id, name, type_id, system)
VALUES
    (gen_random_uuid(), 'gram', (SELECT type_id FROM measurement_types WHERE name = 'weight'), 'metric'),
    (gen_random_uuid(), 'kilogram', (SELECT type_id FROM measurement_types WHERE name = 'weight'), 'metric'),
    (gen_random_uuid(), 'ounce', (SELECT type_id FROM measurement_types WHERE name = 'weight'), 'imperial'),
    (gen_random_uuid(), 'liter', (SELECT type_id FROM measurement_types WHERE name = 'volume'), 'metric'),
    (gen_random_uuid(), 'milliliter', (SELECT type_id FROM measurement_types WHERE name = 'volume'), 'metric'),
    (gen_random_uuid(), 'teaspoon', (SELECT type_id FROM measurement_types WHERE name = 'volume'), 'imperial'),
    (gen_random_uuid(), 'meter', (SELECT type_id FROM measurement_types WHERE name = 'length'), 'metric'),
    (gen_random_uuid(), 'centimeter', (SELECT type_id FROM measurement_types WHERE name = 'length'), 'metric'),
    (gen_random_uuid(), 'inch', (SELECT type_id FROM measurement_types WHERE name = 'length'), 'imperial'),
    (gen_random_uuid(), 'Celsius', (SELECT type_id FROM measurement_types WHERE name = 'temperature'), 'metric'),
    (gen_random_uuid(), 'Fahrenheit', (SELECT type_id FROM measurement_types WHERE name = 'temperature'), 'imperial');

-- SEEDING DATA INTO users
DELETE FROM users;

-- Insert some basic users with password hashes (replace the hashes with actual hash values from your authentication logic)
INSERT INTO users (user_id, email, full_name, password_hash, created_at, updated_at)
VALUES
    (gen_random_uuid(), 'john.doe@example.com', 'John Doe', 'hashedpassword123', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (gen_random_uuid(), 'jane.smith@example.com', 'Jane Smith', 'hashedpassword456', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);


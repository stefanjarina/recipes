-- DROPS
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS recipes;
DROP TABLE IF EXISTS measurement_types;
DROP TABLE IF EXISTS measurements;
DROP TABLE IF EXISTS ingredients;
DROP TABLE IF EXISTS comments;
DROP TABLE IF EXISTS photos;
DROP TABLE IF EXISTS steps;
DROP TABLE IF EXISTS step_ingredients;


-- TABLES
-- Users table
CREATE TABLE IF NOT EXISTS users (
  user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  email VARCHAR(255) NOT NULL UNIQUE,
  full_name VARCHAR(255) NOT NULL,
  password_hash TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Recipes table
CREATE TABLE IF NOT EXISTS recipes (
  recipe_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  title VARCHAR(255) NOT NULL,
  description TEXT,
  visibility VARCHAR(10) DEFAULT 'public' CHECK (visibility IN ('public', 'private')),
  user_id UUID REFERENCES users(user_id) ON DELETE CASCADE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Measurement types table
CREATE TABLE IF NOT EXISTS measurement_types (
  type_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(50) NOT NULL UNIQUE
);

-- Measurements table
CREATE TABLE IF NOT EXISTS measurements (
  measurement_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(100) NOT NULL UNIQUE,
  type_id UUID NOT NULL REFERENCES measurement_types(type_id) ON DELETE CASCADE,
  system VARCHAR(10) NOT NULL CHECK (system IN ('metric', 'imperial'))
);

-- Ingredients table
CREATE TABLE IF NOT EXISTS ingredients (
  ingredient_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(255) NOT NULL,
  amount NUMERIC NOT NULL,
  amount_metric NUMERIC NOT NULL,
  measurement_id UUID REFERENCES measurements(measurement_id) ON DELETE CASCADE,
  recipe_id UUID REFERENCES recipes(recipe_id) ON DELETE CASCADE
);

-- Comments table
CREATE TABLE IF NOT EXISTS comments (
  comment_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  recipe_id UUID REFERENCES recipes(recipe_id) ON DELETE CASCADE,
  user_id UUID REFERENCES users(user_id) ON DELETE CASCADE,
  comment_text TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Photos table (for storing recipe photos)
CREATE TABLE IF NOT EXISTS photos (
    photo_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    recipe_id UUID REFERENCES recipes(recipe_id) ON DELETE CASCADE,
    file_path TEXT,  -- Path to the photo on disk (optional if URL is used)
    url TEXT,  -- URL to the photo in case it's hosted online (optional if file_path is used)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT check_photo_path_or_url
      CHECK (file_path IS NOT NULL OR url IS NOT NULL)
);

-- Ratings table (for storing user ratings for recipes)
CREATE TABLE IF NOT EXISTS ratings (
    rating_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(user_id) ON DELETE CASCADE,
    recipe_id UUID REFERENCES recipes(recipe_id) ON DELETE CASCADE,
    rating INTEGER CHECK (rating >= 1 AND rating <= 5),  -- 1-5 stars
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Steps table
CREATE TABLE IF NOT EXISTS steps (
    step_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    instructions TEXT NOT NULL,
    step_number INTEGER NOT NULL,  -- To keep the order of steps
    recipe_id UUID REFERENCES recipes(recipe_id) ON DELETE CASCADE,  -- Reference to the associated recipe
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Step ingredients table
CREATE TABLE IF NOT EXISTS step_ingredients (
    step_ingredient_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    step_id UUID REFERENCES steps(step_id) ON DELETE CASCADE,  -- Reference to the associated step
    ingredient_id UUID REFERENCES ingredients(ingredient_id) ON DELETE CASCADE,  -- Reference to ingredients
    amount NUMERIC NOT NULL,  -- Amount specific to this step
    amount_metric NUMERIC NOT NULL,  -- Metric amount specific to this step
    measurement_id UUID REFERENCES measurements(measurement_id) ON DELETE CASCADE  -- Measurement unit
);

-- Indexes
CREATE INDEX idx_user_email ON users(email);
-- Indexes for foreign keys
CREATE INDEX idx_recipe_user ON recipes(user_id);
CREATE INDEX idx_ingredient_recipe ON ingredients(recipe_id);
CREATE INDEX idx_comment_recipe ON comments(recipe_id);
CREATE INDEX idx_comment_user ON comments(user_id);
CREATE INDEX idx_photo_recipe ON photos(recipe_id);
CREATE INDEX idx_step_recipe ON steps(recipe_id);
CREATE INDEX idx_step_ingredient_step ON step_ingredients(step_id);

-- FUNCTIONS
CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- TRIGGERS
CREATE TRIGGER set_timestamp BEFORE UPDATE ON users FOR EACH ROW EXECUTE FUNCTION update_timestamp();
CREATE TRIGGER set_timestamp BEFORE UPDATE ON recipes FOR EACH ROW EXECUTE FUNCTION update_timestamp();
CREATE TRIGGER set_timestamp BEFORE UPDATE ON ingredients FOR EACH ROW EXECUTE FUNCTION update_timestamp();
CREATE TRIGGER set_timestamp BEFORE UPDATE ON comments FOR EACH ROW EXECUTE FUNCTION update_timestamp();
CREATE TRIGGER set_timestamp BEFORE UPDATE ON photos FOR EACH ROW EXECUTE FUNCTION update_timestamp();
CREATE TRIGGER set_timestamp BEFORE UPDATE ON ratings FOR EACH ROW EXECUTE FUNCTION update_timestamp();
CREATE TRIGGER set_timestamp BEFORE UPDATE ON steps FOR EACH ROW EXECUTE FUNCTION update_timestamp();
CREATE TRIGGER set_timestamp BEFORE UPDATE ON step_ingredients FOR EACH ROW EXECUTE FUNCTION update_timestamp();

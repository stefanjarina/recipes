-- SEEDING DATA INTO recipes (John Doe's Chicken Tikka Masala)
DELETE FROM recipes;

-- Insert the recipe for John Doe
INSERT INTO recipes (recipe_id, title, description, visibility, user_id, created_at, updated_at)
VALUES
    (
        gen_random_uuid(), 
        'Chicken Tikka Masala', 
        'A delicious and creamy chicken curry with rich spices.', 
        'public', 
        (SELECT user_id FROM users WHERE email = 'john.doe@example.com'), 
        CURRENT_TIMESTAMP, 
        CURRENT_TIMESTAMP
    );

-- SEEDING DATA INTO steps for Chicken Tikka Masala
DELETE FROM steps;

-- Add steps for the Chicken Tikka Masala recipe
INSERT INTO steps (step_id, title, instructions, step_number, recipe_id, created_at, updated_at)
VALUES
    (gen_random_uuid(), 'Marinate the Chicken', 'Mix the chicken with yogurt, garlic, and garam masala. Let it marinate for 1 hour.', 1, 
        (SELECT recipe_id FROM recipes WHERE title = 'Chicken Tikka Masala'), CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (gen_random_uuid(), 'Cook the Chicken', 'Cook the marinated chicken in a pan until browned.', 2, 
        (SELECT recipe_id FROM recipes WHERE title = 'Chicken Tikka Masala'), CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (gen_random_uuid(), 'Make the Sauce', 'In the same pan, add tomato paste, cream, and spices. Cook until thickened.', 3, 
        (SELECT recipe_id FROM recipes WHERE title = 'Chicken Tikka Masala'), CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (gen_random_uuid(), 'Combine Chicken and Sauce', 'Add the cooked chicken to the sauce and simmer for 10 minutes.', 4, 
        (SELECT recipe_id FROM recipes WHERE title = 'Chicken Tikka Masala'), CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- SEEDING DATA INTO ingredients (overall ingredients for the recipe)
DELETE FROM ingredients;

-- Add overall ingredients for the recipe
INSERT INTO ingredients (ingredient_id, name, amount, amount_metric, measurement_id, recipe_id)
VALUES
    (gen_random_uuid(), 'Chicken breast', 500, 500, (SELECT measurement_id FROM measurements WHERE name = 'gram'), 
        (SELECT recipe_id FROM recipes WHERE title = 'Chicken Tikka Masala')),
    (gen_random_uuid(), 'Yogurt', 200, 200, (SELECT measurement_id FROM measurements WHERE name = 'gram'), 
        (SELECT recipe_id FROM recipes WHERE title = 'Chicken Tikka Masala')),
    (gen_random_uuid(), 'Tomato paste', 2, 2, (SELECT measurement_id FROM measurements WHERE name = 'teaspoon'), 
        (SELECT recipe_id FROM recipes WHERE title = 'Chicken Tikka Masala')),
    (gen_random_uuid(), 'Garam masala', 1, 1, (SELECT measurement_id FROM measurements WHERE name = 'teaspoon'), 
        (SELECT recipe_id FROM recipes WHERE title = 'Chicken Tikka Masala'));

-- SEEDING DATA INTO step_ingredients (step-specific ingredients)
DELETE FROM step_ingredients;

-- Step 1: Marinate the chicken
INSERT INTO step_ingredients (step_ingredient_id, step_id, ingredient_id, amount, amount_metric, measurement_id)
VALUES
    (gen_random_uuid(), 
        (SELECT step_id FROM steps WHERE title = 'Marinate the Chicken'), 
        (SELECT ingredient_id FROM ingredients WHERE name = 'Chicken breast'), 
        500, 500, 
        (SELECT measurement_id FROM measurements WHERE name = 'gram')),
    (gen_random_uuid(), 
        (SELECT step_id FROM steps WHERE title = 'Marinate the Chicken'), 
        (SELECT ingredient_id FROM ingredients WHERE name = 'Yogurt'), 
        200, 200, 
        (SELECT measurement_id FROM measurements WHERE name = 'gram')),
    (gen_random_uuid(), 
        (SELECT step_id FROM steps WHERE title = 'Marinate the Chicken'), 
        (SELECT ingredient_id FROM ingredients WHERE name = 'Garam masala'), 
        1, 1, 
        (SELECT measurement_id FROM measurements WHERE name = 'teaspoon'));

-- Step 3: Make the sauce (Tomato paste goes here)
INSERT INTO step_ingredients (step_ingredient_id, step_id, ingredient_id, amount, amount_metric, measurement_id)
VALUES
    (gen_random_uuid(), 
        (SELECT step_id FROM steps WHERE title = 'Make the Sauce'), 
        (SELECT ingredient_id FROM ingredients WHERE name = 'Tomato paste'), 
        2, 2, 
        (SELECT measurement_id FROM measurements WHERE name = 'teaspoon'));

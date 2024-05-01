-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(50) NOT NULL,
    password_hash VARCHAR(255) NOT NULL
);

-- Create cats table
CREATE TABLE IF NOT EXISTS cats (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    race VARCHAR(50) CHECK (race IN (
        'Persian', 'Maine Coon', 'Siamese', 'Ragdoll', 'Bengal', 
        'Sphynx', 'British Shorthair', 'Abyssinian', 'Scottish Fold', 'Birman'
    )) NOT NULL,
    sex VARCHAR(10) CHECK (sex IN ('male', 'female')) NOT NULL,
    age_in_month INT CHECK (age_in_month >= 1 AND age_in_month <= 120082) NOT NULL,
    description TEXT NOT NULL,
    image_urls TEXT[] NOT NULL,
    owner_id INT REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create matches table
CREATE TABLE IF NOT EXISTS matches (
    id SERIAL PRIMARY KEY,
    match_cat_id INT REFERENCES cats(id) ON DELETE CASCADE,
    user_cat_id INT REFERENCES cats(id) ON DELETE CASCADE,
    message VARCHAR(120) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
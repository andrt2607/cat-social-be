CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(50) NOT NULL,
    password_hash VARCHAR(255) NOT null,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE index email_user on users(email);

CREATE TABLE IF NOT EXISTS cats (
    id SERIAL PRIMARY KEY,
    -- owner_id INT,
    owner_id INT REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(30) NOT NULL,
    race VARCHAR(50) CHECK (race IN (
        'Persian', 'Maine Coon', 'Siamese', 'Ragdoll', 'Bengal', 
        'Sphynx', 'British Shorthair', 'Abyssinian', 'Scottish Fold', 'Birman'
    )) NOT NULL,
    sex VARCHAR(10) CHECK (sex IN ('male', 'female')) NOT NULL,
    age_in_month INT CHECK (age_in_month >= 1 AND age_in_month <= 120082) NOT NULL,
    description TEXT NOT NULL,
    image_urls TEXT[] NOT NULL,
    has_matched BOOLEAN,
    is_deleted BOOLEAN,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
    -- FOREIGN KEY (owner_id) REFERENCES users(id)
);

-- CREATE index race_cat on cats(race);
CREATE index owner_cat on cats(owner_id);
-- CREATE index age_cat on cats(age_in_month);

CREATE TABLE IF NOT EXISTS likes (
    id SERIAL PRIMARY KEY,
    owner_id INTEGER,
    cat_id INTEGER,
    liked_owner_id INTEGER,
    liked_cat_id INTEGER,
    approval_status VARCHAR(15) CHECK (approval_status IN ('approved', 'rejected', 'pending')) NOT NULL,
    message VARCHAR(120) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
    -- FOREIGN KEY (owner_id) REFERENCES users(id),
    -- FOREIGN KEY (liked_owner_id) REFERENCES users(id),
    -- FOREIGN KEY (cat_id) REFERENCES cats(id),
    -- FOREIGN KEY (liked_cat_id) REFERENCES cats(id)
);
CREATE index owner_likes on likes(owner_id, liked_owner_id);
-- CREATE index cat_likes on likes(cat_id, liked_cat_id);
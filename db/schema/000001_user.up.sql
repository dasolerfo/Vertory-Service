CREATE TABLE "users" (
    "id" BIGSERIAL PRIMARY KEY, 
    "username" VARCHAR(255) NOT NULL UNIQUE,
    "email" VARCHAR(255) NOT NULL UNIQUE,
    "password_hash" VARCHAR(255) NOT NULL,
    "followers_count" INT DEFAULT 0,
    "following_count" INT DEFAULT 0
);

CREATE TABLE "products" (
    "id" BIGSERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "description" TEXT,
    "scans" INT DEFAULT 0,
    "image_url" VARCHAR(255) NOT NULL
);

CREATE TABLE "user_scans" (
    "id" BIGSERIAL PRIMARY KEY,
    "user_id" BIGINT REFERENCES users(id) ON DELETE CASCADE,
    "product_id" BIGINT REFERENCES products(id) ON DELETE CASCADE,
    "scanned_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "user_follows" (
    "id" BIGSERIAL PRIMARY KEY,
    "follower_id" BIGINT REFERENCES users(id) ON DELETE CASCADE,
    "followed_id" BIGINT REFERENCES users(id) ON DELETE CASCADE,
    "followed_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
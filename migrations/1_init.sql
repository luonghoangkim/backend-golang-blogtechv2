-- +migrate Up
CREATE TABLE "users" (
    "user_id" text PRIMARY KEY,
    "full_name" text,
    "email" text UNIQUE,
    "password" text,
    "role" text,
    "created_at" TIMESTAMPTZ NOT NULL,
    "updated_at" TIMESTAMPTZ NOT NULL
);

-- Create technews_posts table
CREATE TABLE "technews_posts" (
    "pid" text PRIMARY KEY,
    "title" text,
    "summary" text,
    "author" text,
    "content" text,
    "cover_image" text,
    "content_image" text,
    "created_at" TIMESTAMPTZ NOT NULL,
    "updated_at" TIMESTAMPTZ NOT NULL
);

-- Create future_technology_posts table
CREATE TABLE "future_technology_posts" (
    "pid" text PRIMARY KEY,
    "title" text,
    "summary" text,
    "author" text,
    "content" text,
    "cover_image" text,
    "content_image" text,
    "created_at" TIMESTAMPTZ NOT NULL,
    "updated_at" TIMESTAMPTZ NOT NULL
);

-- Create tutorials_and_tips_posts table
CREATE TABLE "tutorials_and_tips_posts" (
    "pid" text PRIMARY KEY,
    "title" text,
    "summary" text,
    "author" text,
    "content" text,
    "cover_image" text,
    "content_image" text,
    "created_at" TIMESTAMPTZ NOT NULL,
    "updated_at" TIMESTAMPTZ NOT NULL
);




-- +migrate Down
DROP TABLE users;
DROP TABLE technews_posts;
DROP TABLE future_technology_posts;
DROP TABLE tutorials_and_tips_posts;

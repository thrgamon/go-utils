CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "username" text,
  "auth_id" text,
  "inserted_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

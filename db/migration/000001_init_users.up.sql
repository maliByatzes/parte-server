-- Users Table
CREATE TABLE "users" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "username" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "email" varchar NOT NULL,
  "is_active" boolean NOT NULL DEFAULT false,
  "is_superuser" boolean NOT NULL DEFAULT false,
  "thumbnail" text NOT NULL DEFAULT 'https://api.multiavatar.com/user.png',
  "updated_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("username", "email", "is_active");
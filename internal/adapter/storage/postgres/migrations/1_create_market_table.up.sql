CREATE TYPE "markets_status_enum" AS ENUM ('INACTIVE', 'ACTIVE');

CREATE TABLE "markets" (
    "id" serial PRIMARY KEY,
    "name" varchar NOT NULL,
    "status" markets_status_enum NOT NULL
    -- "created_at" timestamptz NOT NULL DEFAULT (now()),
    -- "updated_at" timestamptz NOT NULL DEFAULT (now())
);


CREATE TYPE "market_outcomes_status_enum" AS ENUM ('INACTIVE', 'ACTIVE');

CREATE TABLE "market_outcomes" (
    "id" serial PRIMARY KEY,
    "name" varchar NOT NULL,
    "status" market_outcomes_status_enum NOT NULL
    -- "created_at" timestamptz NOT NULL DEFAULT (now()),
    -- "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "account_event" (
    "id" BIGSERIAL PRIMARY KEY,
    "account_id" BIGINT,
    "data" TEXT NOT NULL,
    "timestamp" TIMESTAMP NOT NULL,
     FOREIGN KEY ("account_id") REFERENCES "account" ("id")
);

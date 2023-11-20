-- init-schema.sql

CREATE TABLE IF NOT EXISTS "users" (
  "userid" serial PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL
);

CREATE TABLE IF NOT EXISTS "todos" (
  "id" serial PRIMARY KEY,
  "todo" text NOT NULL,
  "complated" boolean NOT NULL,
  "userid" int NOT NULL,
  "isdelete" boolean NOT NULL DEFAULT FALSE,
  "deletedon" timestamp
);

ALTER TABLE "todos" ADD FOREIGN KEY ("userid") REFERENCES "users" ("userid");

CREATE OR REPLACE FUNCTION update_deletedon()
RETURNS TRIGGER AS $$
BEGIN
  IF NEW.isdelete = TRUE THEN
    NEW.deletedon = NOW();
  ELSE
    NEW.deletedon = NULL;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_deletedon
BEFORE UPDATE ON todos
FOR EACH ROW
EXECUTE FUNCTION update_deletedon();
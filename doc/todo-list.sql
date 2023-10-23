CREATE TABLE "users" (
  "userId" serial PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL
);

CREATE TABLE "todos" (
  "id" serial PRIMARY KEY,
  "todo" text NOT NULL,
  "balance" boolean NOT NULL,
  "userId" int NOT NULL,
  "isdelete" boolean NOT NULL DEFAULT FALSE,
  "deletedon" timestamp
);

ALTER TABLE "todos" ADD FOREIGN KEY ("userId") REFERENCES "users" ("userId");

CREATE TABLE "public"."ghpr" ("Id" UUID NOT NULL, "ghpr" text default gen_random_uuid(), PRIMARY KEY ("Id") , UNIQUE ("Id"));COMMENT ON TABLE "public"."ghpr" IS E'GHPR Demonstration';


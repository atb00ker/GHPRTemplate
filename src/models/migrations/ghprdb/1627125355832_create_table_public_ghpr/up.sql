CREATE TABLE "public"."ghpr" ("Id" UUID NOT NULL default gen_random_uuid(), "ghpr" text, PRIMARY KEY ("Id") , UNIQUE ("Id"));COMMENT ON TABLE "public"."ghpr" IS E'GHPR Demonstration';


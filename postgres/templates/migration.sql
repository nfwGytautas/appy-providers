-- WISL Metadata {
--  @Requires: core
-- }

-- Server
CREATE TABLE IF NOT EXISTS "example" (
    "id"    VARCHAR(255) NOT NULL UNIQUE,
    "name"  VARCHAR(255) NOT NULL
);

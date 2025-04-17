-- Core migration needed by postgres provider do not remove

@meta {
    "name": "core"
}

@up {
    -- Server
    CREATE TABLE IF NOT EXISTS "driver-internal" (
        "db_version"    VARCHAR(255) NOT NULL UNIQUE
    );

    DELETE FROM "driver-internal";
    INSERT INTO "driver-internal" VALUES ('nil');
}

@down {
    -- Server
    DROP TABLE IF EXISTS "driver-internal";
}

@meta {
    "name": "example",
    "requires": "core"
}

@up {

-- Server
CREATE TABLE IF NOT EXISTS "example" (
    "id"    VARCHAR(255) NOT NULL UNIQUE,
    "name"  VARCHAR(255) NOT NULL
);

}

@down {

-- Server
DROP TABLE IF EXISTS "example";

}

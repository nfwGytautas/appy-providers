"""
Test for the macros
"""

import os
import sys
import pytest

sys.path.append(os.path.abspath("../tools/"))

from migrations import Migrations

def test_migrations_read():
    """
    Test that the migrations are read correctly
    """
    test_sql = """
@meta {
    "name": "test",
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
    """

    metadata = {
        "name": "test",
        "requires": "core"
    }

    up = """CREATE TABLE IF NOT EXISTS "example" (   "id"  VARCHAR(255) NOT NULL UNIQUE,   "name" VARCHAR(255) NOT NULL );
"""

    down = """DROP TABLE IF EXISTS "example";
"""

    migrations = Migrations()
    migrations.read_from_str(test_sql)

    for _, migration in migrations.migrations.items():
        assert migration["meta"] == metadata
        assert migration["up"] == up.strip()
        assert migration["down"] == down.strip()

def test_migrations_write():
    """
    Test that the migrations are written correctly
    """
    test_sql = """
@meta {
    "name": "test",
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
    """

    expected = """
func mUptest(tx *appy_driver.Tx, currentVersion string) error {
    const query = `CREATE TABLE IF NOT EXISTS "example" (   "id"  VARCHAR(255) NOT NULL UNIQUE,   "name" VARCHAR(255) NOT NULL );`
    var err error

    if currentVersion == "test" {
        appy_logger.Logger().Info("    - Already at this version")
        return nil
    }

    appy_logger.Logger().Info("  - Dependency to 'core'")

    err = mUpcore(tx, currentVersion)
    if err != nil {
        return err
    }

    appy_logger.Logger().Info("  - Migrating to 'test'")

    err = appy_driver.MigrateToVersion(tx, "test", query)
    if err != nil {
        return err
    }

    currentVersion = "test"

    return nil
}

func mDowntest(tx *appy_driver.Tx, currentVersion string) error {
    const query = `DROP TABLE IF EXISTS "example";`
    return errors.New("migrating down is not yet implemented")
}
"""

    migrations = Migrations()
    migrations.read_from_str(test_sql)
    content = migrations.to_str()

    assert content == expected

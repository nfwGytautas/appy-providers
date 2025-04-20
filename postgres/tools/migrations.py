"""
Migrations for the postgres provider
"""

import re
import json
import string

from templates import MIGRATION, MIGRATION_ROOT
from utils import cleanup_query

class Migrations:
    """
    Migrations holder
    """

    MIGRATIONS_PATTERN = re.compile(
        r"""
@meta\s*{\s*(?P<meta>.*?)\s*}
\s*
@up\s*{\s*(?P<up>.*?)\s*}
\s*
@down\s*{\s*(?P<down>.*?)\s*}
""",
        re.DOTALL | re.VERBOSE
    )

    def __init__(self) -> None:
        self.__migrations = {}

    def read_from_str(self, content: str) -> None:
        """
        Read queries from a string
        """
        match = re.search(self.MIGRATIONS_PATTERN, content)

        meta = json.loads("{" + match.group("meta").strip() + "}")
        name = meta["name"]
        up = match.group("up")
        down = match.group("down")

        self.__migrations[name] = {
            "meta": meta,
            "up": cleanup_query(up),
            "down": cleanup_query(down)
        }

    def read_from_file(self, file: str) -> None:
        """
        Read queries from a file
        """
        with open(file, "r", encoding="utf-8") as f:
            self.read_from_str(f.read())

    def to_str(self) -> str:
        """
        Convert migrations to a string
        """
        final = ""
        for migration in self.__migrations.values():
            final += self.__format_migration(migration)
        return final

    def write_to_file(self, file: str) -> None:
        """
        Write migrations to a file
        """
        with open(file, "w+", encoding="utf-8") as f:
            f.write(self.to_str())

    def get_latest(self) -> str:
        """
        Get the latest migration
        """
        if len(self.__migrations) == 0:
            return ""

        known = set(self.migrations.keys())

        # Search for the migration that is not referenced by any other
        for migration in self.__migrations.values():
            meta = migration["meta"]
            dependency = meta.get("requires", None)

            if dependency is not None:
                known.remove(dependency)

        if len(known) > 1:
            raise Exception("Multiple root migrations found: " + str(known))

        return known.pop()

    @property
    def migrations(self) -> dict:
        """
        Get the queries
        """
        return self.__migrations

    def __format_migration(self, migration: dict) -> str:
        """
        Format a migration
        """
        meta = migration["meta"]
        name = meta["name"]
        dependency = meta.get("requires", None)

        if dependency is not None:
            return string.Template(MIGRATION).substitute(
                name=name,
                dependency=dependency,
                queryUp=migration["up"],
                queryDown=migration["down"],
            )
        else:
            return string.Template(MIGRATION_ROOT).substitute(
                name=name,
                queryUp=migration["up"],
                queryDown=migration["down"],
            )

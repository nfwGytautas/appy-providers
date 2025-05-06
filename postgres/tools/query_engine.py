"""
Tool for auto generating golang code from sql files
"""

import argparse
import glob
import pathlib
import string
import subprocess

from macros import Macros
from queries import Queries
from migrations import Migrations

from templates import CONFIG

class QueryEngine:
    """
    Tool for auto generating golang code from sql files
    """

    def __init__(self) -> None:
        arguments = argparse.ArgumentParser(description="Query Engine")

        arguments.add_argument(
            "root", type=str, help="Root directory where macros, migrations, queries are located")
        arguments.add_argument(
            "package", type=str, help="The name of the package", default="package"
        )

        values = arguments.parse_args()

        self.__root = values.root
        self.__package = values.package

    def run(self) -> None:
        """
        Run the query engine
        """

        # Read macros
        macros = Macros()

        print("Reading macros...")
        for file in glob.glob(f"{self.__root}/macros/**/*.sql", recursive=True):
            print("    + ", file)
            macros.read_from_file(file)
        print(f"Loaded '{len(macros.macros)}' macros")

        # Read queries
        queries = Queries(macros)

        print("Reading queries...")
        for file in glob.glob(f"{self.__root}/queries/**/*.sql", recursive=True):
            print("    + ", file)
            queries.read_from_file(file)
        print(f"Loaded '{len(queries.queries)}' queries")

        # Read migrations
        migrations = Migrations()

        print("Reading migrations...")
        for file in glob.glob(f"{self.__root}/migrations/*.sql"):
            print("    + ", file)
            migrations.read_from_file(file)
        print(f"Loaded '{len(migrations.migrations)}' migrations")

        # Get latest migration
        latest_migration = migrations.get_latest()
        if latest_migration is None:
            raise Exception("No final migration found")

        print(f"Latest migration: {latest_migration}")

        # Create files
        print("Creating files...")
        pathlib.Path(f"{self.__root}/driver").mkdir(parents=True, exist_ok=True)

        queries_file = f"{self.__root}/driver/queries.go"
        migrations_file = f"{self.__root}/driver/migrations.go"
        base_file = f"{self.__root}/driver/base.go"

        with open(queries_file, "w+", encoding="utf-8") as f:
            f.write(f"package {self.__package}_driver\n\n")
            f.write("import (\n")
            f.write("    appy_driver \"github.com/nfwGytautas/appy-providers/postgres/impl\"\n")
            f.write("    appy_logger \"github.com/nfwGytautas/appy-go/logger\"\n")
            f.write(")\n\n")
            f.write(queries.to_str())

        with open(migrations_file, "w+", encoding="utf-8") as f:
            f.write(f"package {self.__package}_driver\n\n")
            f.write("import (\n")
            f.write("    \"errors\"\n")
            f.write("    appy_driver \"github.com/nfwGytautas/appy-providers/postgres/impl\"\n")
            f.write("    appy_logger \"github.com/nfwGytautas/appy-go/logger\"\n")
            f.write(")\n\n")
            f.write(migrations.to_str())

        with open(base_file, "w+", encoding="utf-8") as f:
            f.write(string.Template(CONFIG).substitute(
                package=self.__package,
                version=latest_migration,
                migration=f"mUp{latest_migration}",
            ))

        print("Running go tools...")
        subprocess.run(["goimports", "-w", queries_file], check=True)
        subprocess.run(["gofmt", "-w", queries_file], check=True)
        subprocess.run(["goimports", "-w", migrations_file], check=True)
        subprocess.run(["gofmt", "-w", migrations_file], check=True)
        subprocess.run(["goimports", "-w", base_file], check=True)
        subprocess.run(["gofmt", "-w", base_file], check=True)

        print("Done!")

if __name__ == "__main__":
    QueryEngine().run()

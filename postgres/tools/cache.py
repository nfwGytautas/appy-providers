"""
Cache for which files are built
"""

import json
import glob
import os
import pathlib

class QueryCache:
    """
    Cache for which files are built
    """

    def __init__(self, cache_file: str) -> None:
        self.__cache_file = cache_file
        self.__cache = {}
        self.__need_rebuild = []

    def record(self, directory: str) -> None:
        for file in glob.glob(f"{directory}/*.sql"):
            mod_time = os.path.getmtime(file)

            if file in self.__cache:
                if self.__cache[file]["mod_time"] < mod_time:
                    self.__need_rebuild.append(file)

            self.__cache[file] = {
                "mod_time": mod_time,
                "deleted": False
            }

    def remove_deleted(self) -> list[str]:
        deleted = []

        for file in self.__cache:
            if self.__cache[file]["deleted"]:
                deleted.append(file)
                del self.__cache[file]

        return deleted

    def load_last(self) -> None:
        if not pathlib.Path(self.__cache_file).exists():
            return

        with open(self.__cache_file, "r", encoding="utf-8") as f:
            self.__cache = json.loads(f.read())

        for file in self.__cache:
            self.__cache[file]["deleted"] = True

    def save(self) -> None:
        pathlib.Path(self.__cache_file).parent.mkdir(parents=True, exist_ok=True)

        with open(self.__cache_file, "w", encoding="utf-8") as f:
            f.write(json.dumps(self.__cache))

    def get_outdated(self) -> list[str]:
        return self.__need_rebuild

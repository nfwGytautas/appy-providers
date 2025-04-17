"""
Macros for the postgres provider
"""

import re

from utils import cleanup_query

class Macros:
    """
    Macros holder
    """

    MACRO_PATTERN = re.compile(r"@macro\s+(\w+)\s+(.*?)@endmacro", re.DOTALL)
    MACRO_USAGE_PATTERN = re.compile(r'%%([^%]+)%%')

    def __init__(self) -> None:
        self.__macros = {}

    def read_from_str(self, content: str) -> None:
        """
        Read macros from a string
        """
        macros = re.findall(self.MACRO_PATTERN, content)

        for macro in macros:
            self.__macros[macro[0]] = cleanup_query(macro[1])

    def read_from_file(self, file: str) -> None:
        """
        Read macros from a file
        """
        with open(file, "r", encoding="utf-8") as f:
            content = f.read()

        self.read_from_str(content)

    def apply_macros(self, content: str) -> str:
        """
        Apply macros to a content
        """
        replaced = True

        def replacer(match):
            nonlocal replaced
            replaced = True
            key = match.group(1)
            return self.__macros[key]

        pattern = r"%%(.*?)%%"

        while replaced:
            replaced = False
            content = re.sub(pattern, replacer, content)

        return content

    def get_used_macros(self, content: str) -> list[str]:
        """
        Get the macros used in a content
        """
        return re.findall(self.MACRO_USAGE_PATTERN, content)

    @property
    def macros(self) -> dict[str, str]:
        """
        Get the macros
        """
        return self.__macros

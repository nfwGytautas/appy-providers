"""
Queries for the postgres provider
"""

import re
import string

from macros import Macros
from templates import QUERY_ROW_FUNCTION, QUERY_ROWS_FUNCTION, QUERY_EXEC_FUNCTION
from utils import cleanup_query

class Queries:
    """
    Queries holder
    """

    QUERY_PATTERN = re.compile(
        r"@query\s+(\w+)(?::(\w+))?\s+(.*?)\s+@endquery", re.DOTALL)

    ARGUMENT_PATTERN = re.compile(r'\$(.*?)\:(\w+)')

    def __init__(self, macros: Macros) -> None:
        self.__queries = []
        self.__macros = macros


    def read_from_str(self, content: str) -> None:
        """
        Read queries from a string
        """
        # Apply macros
        content = self.__macros.apply_macros(content)

        self.__queries = self.__queries +re.findall(self.QUERY_PATTERN, content)

    def read_from_file(self, file: str) -> None:
        """
        Read queries from a file
        """
        with open(file, "r", encoding="utf-8") as f:
            self.read_from_str(f.read())

    def to_str(self) -> str:
        """
        Convert queries to a string
        """
        final = ""
        for query in self.__queries:
            final += self.__format_query(query)
        return final

    def write_to_file(self, file: str) -> None:
        """
        Write queries to a file
        """
        with open(file, "w+", encoding="utf-8") as f:
            f.write(self.to_str())

    @property
    def queries(self) -> list[tuple[str, str]]:
        """
        Get the queries
        """
        return self.__queries

    def __format_query(self, query: tuple[str, str, str]) -> str:
        """
        Format a query
        """

        name = query[0]
        result = query[1]
        content = cleanup_query(query[2])

        # Resolve arguments
        argv = ""
        args_string = ""
        seen_args = []
        matches = self.ARGUMENT_PATTERN.findall(content)

        for arg_type, arg_name in matches:
            if arg_name not in seen_args:
                argv += f", {arg_name} {arg_type}"
                seen_args.append(arg_name)

            if not arg_type.startswith("[]"):
                args_string = args_string + ", " + arg_name
            else:
                args_string = args_string + ", " + f"pq.Array({arg_name})"

        # Adapt query string to postgres
        index = 0
        def replace_with_increased_index(_):
            nonlocal index
            index += 1
            return f"${index}"

        content = re.sub(self.ARGUMENT_PATTERN, replace_with_increased_index, content)

        if result == "Row":
            return string.Template(QUERY_ROW_FUNCTION).substitute(name=name, query=content, argvEx=argv, argv=args_string)
        elif result == "Rows":
            return string.Template(QUERY_ROWS_FUNCTION).substitute(name=name, query=content, argvEx=argv, argv=args_string)
        else:
            return string.Template(QUERY_EXEC_FUNCTION).substitute(name=name, query=content, argvEx=argv, argv=args_string)

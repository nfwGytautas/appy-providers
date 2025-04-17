"""
Test for the macros
"""

import os
import sys
import pytest

sys.path.append(os.path.abspath("../"))

from macros import Macros
from queries import Queries

def test_queries_read():
    """
    Test that the queries are read correctly
    """
    test_sql = """
@query test_query
    SELECT * FROM example
@endquery
    """

    macros = Macros()
    macros.read_from_str(test_sql)

    queries = Queries(macros)
    queries.read_from_str(test_sql)

    assert queries.queries == [("test_query", "", "SELECT * FROM example")]


def test_queries_read_2():
    """
    Test that the queries are read correctly
    """
    test_sql = """
@macro test_macro
    SELECT * FROM example
@endmacro

@query test_query:Row
    %%test_macro%%
@endquery
    """

    macros = Macros()
    macros.read_from_str(test_sql)

    queries = Queries(macros)
    queries.read_from_str(test_sql)

    assert queries.queries == [("test_query", "Row", "SELECT * FROM example")]


def test_queries_read_3():
    """
    Test that the queries are read correctly
    """
    test_sql = """
@query test_query
    SELECT * FROM example
@endquery

@query test_query2
    SELECT * FROM example
@endquery
    """

    macros = Macros()
    macros.read_from_str(test_sql)

    queries = Queries(macros)
    queries.read_from_str(test_sql)

    assert queries.queries == [("test_query", "", "SELECT * FROM example"), ("test_query2", "", "SELECT * FROM example")]

def test_queries_write():
    """
    Test that the queries are written correctly
    """
    test_sql = """
@query test_query:Row
    SELECT * FROM example
    WHERE id = $uint64:id
@endquery
    """

    expected = """
func Qtest_query(tx *appy_driver.Tx , id uint64) (appy_driver.RowResult, error) {
	const query = `SELECT * FROM example   WHERE id = $1`
	appy_logger.Logger().Info("Executing 'test_query'")
	row := tx.QueryRow(query , id)
	return row, row.Err()
}
"""

    macros = Macros()
    macros.read_from_str(test_sql)

    queries = Queries(macros)
    queries.read_from_str(test_sql)

    assert queries.to_str() == expected

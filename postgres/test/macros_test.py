"""
Test for the macros
"""

import os
import sys
import pytest

sys.path.append(os.path.abspath("../tools/"))

from macros import Macros

def test_macros_read():
    """
    Test that the macros are applied correctly
    """
    test_sql = """
@macro test_macro
    SELECT * FROM example
@endmacro
    """

    macros = Macros()
    macros.read_from_str(test_sql)
    assert macros.apply_macros("%%test_macro%%") == "SELECT * FROM example"


def test_macros_read_2():
    """
    Test that the macros are applied correctly
    """
    test_sql = """
@macro test_macro
    SELECT * FROM example
@endmacro

@macro test_macro2
    %%test_macro%%
@endmacro
    """

    macros = Macros()
    macros.read_from_str(test_sql)
    assert macros.apply_macros("%%test_macro2%%") == "SELECT * FROM example"

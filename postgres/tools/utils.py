"""
Utility functions
"""

def cleanup_query(query: str) -> str:
    """Cleanup query

    Args:
        query (str): Query to clean up

    Returns:
        str: single line, whitespace removed query
    """
    # Remove comments
    query = "\n".join([line for line in query.split("\n") if not "--" in line])
    return query.replace("\n", " ").replace("\t", " ").replace("  ", " ").strip()

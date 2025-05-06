"""
Templates for the query engine
"""

QUERY_ROW_FUNCTION = """
func Q$name[T any](tx appy_driver.Tx, scanner appy_driver.ScanFn[T] $argvEx) (*T, error) {
	const query = `$query`
	appy_logger.Logger().Info("Executing '$name'")

	row := tx.QueryRow(query $argv)
    if row.Err() != nil {
        return nil, row.Err()
    }

    result, err := appy_driver.ParseRow(row, scanner)
    if err != nil {
        return nil, err
    }

	return result, nil
}
"""

QUERY_ROWS_FUNCTION = """
func Q$name[T any](tx appy_driver.Tx, scanner appy_driver.ScanFn[T] $argvEx) ([]T, error) {
	const query = `$query`
	appy_logger.Logger().Info("Executing '$name'")

    rows, err := tx.Query(query $argv)
    if err != nil {
        return nil, err
    }

    result, err := appy_driver.ParseRows(rows, scanner)
    if err != nil {
        return nil, err
    }

	return result, nil
}
"""

QUERY_EXEC_FUNCTION = """
func Q$name(tx appy_driver.Tx $argvEx) (appy_driver.ExecResult, error) {
	const query = `$query`
	appy_logger.Logger().Info("Executing '$name'")
	return tx.Exec(query $argv)
}
"""

MIGRATION = """
func mUp$name(tx appy_driver.Tx, currentVersion string) error {
    const query = `$queryUp`
    var err error

    if currentVersion == "$name" {
        appy_logger.Logger().Info("    - Already at this version")
        return nil
    }

    appy_logger.Logger().Info("  - Dependency to '$dependency'")

    err = mUp$dependency(tx, currentVersion)
    if err != nil {
        return err
    }

    appy_logger.Logger().Info("  - Migrating to '$name'")

    err = appy_driver.MigrateToVersion(tx, "$name", query)
    if err != nil {
        return err
    }

    currentVersion = "$name"

    return nil
}

func mDown$name(tx appy_driver.Tx, currentVersion string) error {
    const query = `$queryDown`
    return errors.New("migrating down is not yet implemented")
}
"""

MIGRATION_ROOT = """
func mUp$name(tx appy_driver.Tx, currentVersion string) error {
    const query = `$queryUp`
    var err error

    if currentVersion == "$name" {
        appy_logger.Logger().Info("    - Already at this version")
        return nil
    }

    appy_logger.Logger().Info("  - Migrating to '$name'")

    err = appy_driver.MigrateToVersion(tx, "$name", query)
    if err != nil {
        return err
    }

    currentVersion = "$name"

    return nil
}

func mDown$name(tx appy_driver.Tx, currentVersion string) error {
    const query = `$queryDown`
    return errors.New("migrating down is not yet implemented")
}
"""

CONFIG = """
package ${package}_driver

import (
    appy_driver "github.com/nfwGytautas/appy-providers/postgres/impl"
)

func Initialize(connectionString string) error {
    return appy_driver.Initialize(appy_driver.InitializeArgs{
        ConnectionString: connectionString,
        Version:          "$version",
        Migration:        $migration,
    })
}
"""

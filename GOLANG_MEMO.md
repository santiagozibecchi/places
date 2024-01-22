# Casos de usos generales en Golang

## DB psql Methods

Query:
Query executes a query that returns rows, typically a SELECT. The args are for any placeholder parameters in the query.

QueryRow:
QueryRow executes a query that is expected to return at most one row. QueryRow always returns a non-nil value. Errors are deferred until Row's Scan method is called. If the query selects no rows, the *Row's Scan will return ErrNoRows. Otherwise, the *Row's Scan scans the first selected row and discards the rest.

Exec:
Exec executes a query without returning any rows. The args are for any placeholder parameters in the query.

Prepare: 
Prepare creates a prepared statement for later queries or executions. Multiple queries or executions may be run concurrently from the returned statement. The caller must call the statement's Close method when the statement is no longer needed.


Correr psql:
sudo service postgresql start

Coneccion directa a psql
psql -h localhost -p 5432 -U role_prueba -d prueba2;

Dumps
psql -h localhost -p 5432 -U role_prueba -d prueba2 < create.sql
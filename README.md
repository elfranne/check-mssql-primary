# check-mssql-primary

How do you give informations about the MSSQL primary replica to haproxy?

`external-check` is here to help, where you can define a command to run to check the health state, that we can abuse to tell haproxy which server is the primary replica.

Haproxy does send usefull info as args and environment variables

https://docs.haproxy.org/3.0/configuration.html#external-check%20command
https://learn.microsoft.com/en-us/sql/tools/sqlcmd/sqlcmd-use-utility
https://learn.microsoft.com/en-us/sql/relational-databases/system-functions/sys-fn-hadr-is-primary-replica-transact-sql?view=sql-server-ver16
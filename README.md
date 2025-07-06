# check-mssql-primary

`check-mssql-primary` is an [external-check][1] for Haproxy. It will query [MSSQL][2] nodes to find out if it is the primary replica.

Please create a check-mssql-primary.env at the same location as the binary:

```env
MSSQL_USERNAME=
MSSQL_PASSWORD=
MSSQL_DATABASE=
```

[1]: https://docs.haproxy.org/3.0/configuration.html#external-check%20command
[2]: https://learn.microsoft.com/en-us/sql/relational-databases/system-functions/sys-fn-hadr-is-primary-replica-transact-sql?view=sql-server-ver16

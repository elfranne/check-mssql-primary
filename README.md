# check-mssql-primary

`check-mssql-primary` is an [external-check][1] for Haproxy. It will query [MSSQL][2] nodes to find out if it is the primary replica.

You'll need to place the binary and a check-mssql-primary.env in the chroot jail (Debian default `/var/lib/haproxy`):

```env
MSSQL_USERNAME=
MSSQL_PASSWORD=
MSSQL_DATABASE=
```

You'll need to modify the Haproxy a bit:

```conf
global
    external-check
    insecure-fork-wanted
[...]

frontend mssql-ha-sw
  bind [ip]:1433 interface ens256
  mode tcp  
  default_backend mssql-ha


backend mssql-ha
    mode        tcp
    option external-check
    no option log-health-checks
    external-check command "/check-mssql-primary"
    server mssql1 [ip]:1433 check
    server mssql2 [ip]:1433 check
```

[1]: https://docs.haproxy.org/3.0/configuration.html#external-check%20command
[2]: https://learn.microsoft.com/en-us/sql/relational-databases/system-functions/sys-fn-hadr-is-primary-replica-transact-sql?view=sql-server-ver16

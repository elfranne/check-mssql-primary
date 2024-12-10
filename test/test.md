# Test

```bash
docker run -v $(pwd)/haproxy.cfg:/usr/local/etc/haproxy/haproxy.cfg:ro -p 1400:1400 haproxytech/haproxy-alpine:latest
docker run --name sql1 -e 'ACCEPT_EULA=Y' -e 'MSSQL_SA_PASSWORD=Correct-Horse-Battery-Staple1+' -p 11433:1433 mcr.microsoft.com/mssql/server:latest
docker run --name sql2 -e 'ACCEPT_EULA=Y' -e 'MSSQL_SA_PASSWORD=Correct-Horse-Battery-Staple1+' -p 11444:1433 mcr.microsoft.com/mssql/server:latest



docker exec -it sql1 /opt/mssql-tools18/bin/sqlcmd -S localhost -U "sa" -P "Correct-Horse-Battery-Staple1+" -q "QUERY HERE"

docker exec -it sql1 /opt/mssql-tools18/bin/sqlcmd -S localhost -U "sa" -P "Correct-Horse-Battery-Staple1+" -i /path/to/file

```



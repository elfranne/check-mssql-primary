# Test

```bash
docker run -v $(pwd)/haproxy.cfg:/usr/local/etc/haproxy/haproxy.cfg:ro -p 1400:1400 haproxytech/haproxy-alpine:latest
docker run -e 'ACCEPT_EULA=Y' -e 'MSSQL_SA_PASSWORD=Correct-Horse-Battery-Staple1+' -p 11433:1433 mcr.microsoft.com/mssql/server:latest 
```

# docker-ip2location

small docker'd webservice providing ip location lookup

## docker hub

see: https://hub.docker.com/repository/docker/extrawurst/ip2location

## disclaimer

This site or product includes IP2Location LITE data available from http://www.ip2location.com.

see database license: `data/LICENSE-CC-BY-SA-4.0.txt`

# example 

```
docker run -p 8080:8080 -e PORT=8080 extrawurst/ip2location
```

```
curl localhost:8080/216.58.208.46
curl localhost:8080/2001:4860:4860::8888
```

this should reply with:
```
US
```
(both being Google.com IPs)

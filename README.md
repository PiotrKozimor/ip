# IP

Simple service that will hold value of an dynamic IP. Access is protected by token.

### API

```
curl https://<host>/ip -X GET --header Authorization:<token>
curl https://<host>/ip -X PUT --data '1.1.1.1' --header 'Authorization:<token>'
```

### Mikrotik setup

```
/system scheduler
add interval=10m name=ip on-event=ip policy=read,write,test start-date=oct/01/2021 start-time=21:15:20
/system script
add dont-require-permissions=no name=ip owner=admin policy=read,write,test source=":local newIP [/ip address get [find interface=\"pppoe\"] address];\
    \n:local foo [/tool fetch http-method=put http-header-field=authorization:<token> url=https://<host>/ip http-data=\$newIP];"
```
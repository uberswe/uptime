# uptime

A simple go program which runs forever (stop it with ctrl+c). It takes two optional arguments:

```
-interval=60
-urls='https://ma.rkus.io, https://www.uberswe.com'
```

You can run this by typing `go run cmd/uptime/main.go -interval=5 -urls='https://ma.rkus.io, https://www.uberswe.com'` in a terminal.

Logs will look like this:

```
2021/04/22 00:26:45  https://ma.rkus.io - 200
2021/04/22 00:26:45  https://www.uberswe.com - 200
2021/04/22 00:26:49  https://www.uberswe.com - 200
2021/04/22 00:26:49  https://ma.rkus.io - 200
2021/04/22 00:26:54  https://www.uberswe.com - 200
2021/04/22 00:26:54  https://ma.rkus.io - 200
2021/04/22 00:26:59  https://www.uberswe.com - 200
2021/04/22 00:26:59  https://ma.rkus.io - 200
```
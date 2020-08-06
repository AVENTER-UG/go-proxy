# go-proxy

This is a proxy written in go.

There are some environment variables to configure go-proxy.

- API_PROXYBIND = Thats the IP of the interface where the proxy should listening. Default value is "0.0.0.0".
- API_PROXYPORT = Thats the Port where the proxy is listening. Default value is 10777.
- TARGET_URL = Thats the URL to where all the requests will be forwarded. No default value.
- SKIP_SSL = Do not check the ssl certificate of the target url. Default value is false.
- BLOCK_USERAGENT = With these variable, it is possible to block UserAgents. No default value.
- LOGLEVEL = The loglevel of the output. Default value "info".

## How to run

```bash
docker run -e TARGET_URL=https:// -p 10777:10777 avhost/go-proxy:latest
```

## How to Block UserAgents

To block Bots and Spam who does not respect robots.txt, it is possible to defined and list with UserAgend should be blocked.
As example:

```bash
BLOCK_USERAGENT="(Bot|The World)"
```

This will block all UserAgents with the string "Bot" or "The World".

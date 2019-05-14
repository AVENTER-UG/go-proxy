# go-proxy

This is a proxy written in go. To use it is very simple.

```bash
TARGET_URL=https://service.nothing

docker run -e SKIP_SSL=true -e TARGET_URL=https:// -p 10777:10777 avhost/go-proxy:latest
```

# `reverse-proxy-host`

[![Github Downloads](https://img.shields.io/github/downloads/patrickdappollonio/reverse-proxy-host/total?color=orange&label=github%20downloads)](https://github.com/patrickdappollonio/reverse-proxy-host/releases)

A tiny Go application that allows you to send, locally, a request using a custom hostname. It's useful if you don't want to be fiddling with `/etc/hosts` records.

## Example use case

Consider you have a `nodejs` application running in `localhost:3000` (`0.0.0.0:3000`, to be precise) and you want to access it using a custom hostname. For the sake of the example, this application will use `example.com` as the hostname.

Using this tool you would run:

```bash
reverse-host-proxy -H example.com -d 3000 -l 8080
```

This will start a reverse proxy on port 8080 that will forward all requests to `localhost:3000` using the custom hostname `example.com`. In my `nodejs` application I now see in the logs:

```text
2022/04/17 23:32:03 "GET http://example.com/assets/jtnkgft3/css/style.css HTTP/1.1" from ::1 - 200 239139B in 2.999ms
2022/04/17 23:32:03 "GET http://example.com/assets/jtnkgft3/js/plugins.min.js HTTP/1.1" from ::1 - 200 45252B in 134.6µs
2022/04/17 23:32:03 "GET http://example.com/assets/jtnkgft3/js/main.js HTTP/1.1" from ::1 - 200 1537B in 80.8µs
2022/04/17 23:32:03 "GET http://example.com/assets/js/home.js?v=jtnkgft3 HTTP/1.1" from ::1 - 200 4058B in 256.1µs
2022/04/17 23:32:04 "GET http://example.com/assets/jtnkgft3/images/picture.jpg HTTP/1.1" from ::1 - 200 170242B in 112µs
2022/04/17 23:32:04 "GET http://example.com/assets/jtnkgft3/fonts/Fontello/fontello.woff2?36999480 HTTP/1.1" from ::1 - 200 6132B in 122µs
2022/04/17 23:32:04 "GET http://example.com/assets/jtnkgft3/images/parallax01-1500x798.jpg HTTP/1.1" from ::1 - 200 23396B in 94.3µs
```

If I would've access it directly, the `GET` request would've shown instead `localhost:3000`.

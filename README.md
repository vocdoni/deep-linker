# Deep Linker (Go)

Simple HTTP server to handle and rewrite Universal links

## Get started

```
$ go build -o server
$ # create config.toml
$ ./server
```

## Overview

1. A user navigates to `https://vocdoni.link/entities/0x1234`
2. The service recognizes the URL path and redirects the request to `https://vocdoni.page.link/?link=https%3A%2F%2Fapp.vocdoni.net%2Fentities%2F%23%2F0x1234&apn=org.vocdoni.app&amv=1&ibi=com.vocdoni.app&ius=vocdoni&isi=1505234624&ivm=0.8&st=Vocdoni&sd=Make%20your%20community%20thrive%20with%20the%20universally%20verifiable%20voting%20system&si=https%3A%2F%2Fblog.vocdoni.io%2Fcontent%2Fimages%2F2020%2F04%2FSquare_512_alpha-1.png`
3. This URL will handle app installs if needed and display social metadata when embeded
4. The main URL is `https://app.vocdoni.net/entities/#/0x1234`
  4.1. On desktop computers, the URL is loaded right away, which leads to the Entity Manager
  4.2. On mobile devices, this will trigger the app itself, which will handle the parameters if an account is set up

## Supported endpoints

- Entity
  - `https://vocdoni.link/entities/<entity-id>`
- Voting process
  - `https://vocdoni.link/processes/<entity-id>/<process-id>`
- News Feed post
  - `https://vocdoni.link/posts/<entity-id>/<post-idx>`
- Registry validation
  - `https://vocdoni.link/validation/<entity-id>/<validation-token>`

## Parameterization

Edit `config.toml` to match your expected settings.

If running from Docker, mount your config file on `/app/config.toml`.

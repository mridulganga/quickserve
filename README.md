## Quick Serve

Simple and fast static file server made in golang.

### Steps - 
1. get the binary and put it somewhere
2. make a folder called app and put the static files there
3. run the app `./quickserve`

### Params -
`-path folder_path -port 80 -cert cert.crt -key keyfile`

- path is to set the folder path, default is `./app`
- port is to set the port, default is `80`
- cert and key is to config tls, by default its empty and secure

NOTE: when giving value for `cert` and `key`, the port is always `443`.

### To build
`make build`
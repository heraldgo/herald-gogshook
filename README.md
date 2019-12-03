# Herald Gogshook Plugin

[Herald daemon](https://github.com/heraldgo/heraldd)
plugin for gogs hook.

The `gogs_hook` trigger will create a http server and wait for gogs
webhook events. `gogs_hook` selector will try to match the repository
parameters.


## Installation

```shell
$ go get -d -u github.com/heraldgo/herald-gogshook
```

Build with

```
$ go build --buildmode=plugin github.com/heraldgo/herald-gogshook
```

Find `herald-gogshook.so` in the current directory.


## Configuration

The plugin file `herald-gogshook.so` must be specified in the herald
daemon configuration.

```yaml
plugin:
  - /plugin_directory/herald-gogshook.so

trigger:
  gogs_hook:
    secret: 'xxxxxxxxxxxxxxxx'
    port: 8234

router:
  deploy:
    trigger: gogs_hook
    selector: gogs_hook
    job:
      print_param: print
    gogs_host: gogs.example.com
    gogs_name: heraldgo/heraldd
    gogs_branch: master
    gogs_event: push
```

Go to the Gogs project settings and select the "Webhooks" page,
then click "Add Webhook" and choose "Gogs".
Input the URL in "Payload URL": `http://example.com:8234`.
"Content type" must be "application/json".
The "secret" must be exactly the same as in the trigger configuration.

More than one projects could point the "payload URL" to the same
`gogs_hook` trigger. The param starts with `gogs_` will be used in the
selector to match the required repository. Not all of these params are
necessary, and only provided params will be used for matching.

`gogs_hook` trigger can also listen on unix socket, which could use
nginx as the reverse proxy.

```yaml
trigger:
  gogs_hook:
    secret: 'xxxxxxxxxxxxxxxx'
    unix_socket: /var/run/heraldd/gogs_hook.sock
```

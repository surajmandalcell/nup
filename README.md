# Network Uptime

#### Description

This tool runs in background and logs if any network distruptions happened, this is specially useful if you want to figure out if you ISP is having issues. You can run this tool and it will log all the network distruptions in a file. You can also pass an argument to also log the latency of the rqequst.

#### Usage

```
nup [-h] [-l] [-f FILE] [-t TIME]
```

#### Compile

```
rustc main.rs -o nup
```

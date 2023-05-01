# Network Uptime

#### Description

This tool runs in background and logs if any network distruptions happened, this is specially useful if you want to figure out if you ISP is having issues. You can run this tool and it will log all the network distruptions in a file. You can also pass an argument to also log the latency of the rqequst.

#### Usage

```
nup [-h] [-l] [-f FILE] [-t TIME]
```

#### Compile and Run

##### Build
```
cargo build --release
```

##### Run
```
cargo run
```


#### Todo
- [ ] Make it installable as a global command
- [ ] Make it run in background 
- [ ] Log the output to the terminal when requested
- [ ] Dont run more than one instance at a time unless specified
- [ ] Add a flag to log the latency of the request
- [ ] Add a flag to log the request status code

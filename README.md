# Network Uptime

#### Description

This tool runs in background and logs if any network distruptions happened, this is specially useful if you want to figure out if you ISP is having issues. You can run this tool and it will log all the network distruptions in a file. You can also pass an argument to also log the latency of the rqequst.

#### Usage

```bash
Usage: nup [OPTION]

        Options:
            -t          Show latency
            -s          Show status code
            -h, --help  Show this help message
```

#### Compile and Run

##### Build
```bash
go build
```

##### Run
```bash
go run
```


#### Todo
- [ ] Make it installable as a global command
- [ ] Make it run in background 
- [ ] Log the output to the terminal when requested
- [ ] Dont run more than one instance at a time unless specified
- [ ] Make sure the writing to file works and is safe
- [x] Add a flag to log the latency of the request
- [x] Add a flag to log the request status code
- [x] Add a argument to show help

#### Todo Documentation
- [x] Add a flag to log the latency of the request
- [x] Add a flag to log the request status code
- [x] Add a argument to show help
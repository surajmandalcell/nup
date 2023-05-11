# Network Uptime

#### Description

This tool runs in background and logs if any network distruptions happened, this is specially useful if you want to figure out if you ISP is having issues. You can run this tool and it will log all the network distruptions in a file. You can also pass an argument to also log the latency of the rqequst.

#### Usage

```text
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
> ##### Target version 1.0.0
- [x] Split the code into multiple files
- [ ] Format log properly to db
- [ ] Make sure the writing to file works and is safe
- [ ] Make it installable as a global command

> ##### Target version 2.0.0
- [ ] Make it run in background 
- [ ] Dont run more than one instance at a time unless specified
- [ ] Log the output to the terminal when requested(when it is already running in background)
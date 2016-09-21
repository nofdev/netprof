# netprof

A network profiler tool

## Usage

Get the latency from local to google.com in formated

```Shell
> .\netprof.exe -ping google.com
2016-09-21 17:28:46,32,54,49
2016-09-21 17:28:47,32,54,49
2016-09-21 17:28:48,32,54,49
```

Get the local local network transmission in formated

```Shell
> .\netprof.exe -transmission
2016-09-21 17:30:01,4021683948,795482278
2016-09-21 17:30:02,4021683948,795482278
2016-09-21 17:30:03,4021683948,795482278
```
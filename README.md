# tcptest
Simple tool to test tcp connectivity. This can be helpful when diagnosing network issues, performing maintenance, or just doing a general health test.

## Examples
Specify the endpoint and port using the format `endpoint:port`.
```bash
$ tcptest jonny.gg:443
✔ jonny.gg:443 [37ms]
```

Multiple endpoints can be specified in a single command using bash expansion techniques...

```bash
$ tcptest jonny.gg:{0..65536}
✘ jonny.gg:0 [37ms]
✘ jonny.gg:1 [40ms]
✘ jonny.gg:2 [39ms]
✘ jonny.gg:3 [37ms]
...
```
... or simply adding another endpoint.

```bash
$ tcptest jonny.gg:443 google.com:443
✔ jonny.gg:443 [34ms]
✔ google.com:443 [34ms]
```
tcptest can be set to loop indefiniately by using the `-l` flag.
```bash
$ tcptest jonny.gg:443 -l
✔ jonny.gg:443 [49ms]
✔ jonny.gg:443 [34ms]
✔ jonny.gg:443 [32ms]
...
```

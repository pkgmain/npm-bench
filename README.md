## npm-bench

# Simple benchmark on `npm install`
# Dual boot Dell XPS 15 9550 (Windows 10 Pro / Manjaro)

Manjaro
```
goos: linux
goarch: amd64
pkg: github.com/pkgmain/npm-bench
BenchmarkNpm-8   	       1	24126985145 ns/op
PASS
ok  	github.com/pkgmain/npm-bench	24.129s

```


Windows 10 Pro
```
goos: windows
goarch: amd64
pkg: github.com/pkgmain/npm-bench
BenchmarkNpm-8                 1        139568871500 ns/op
PASS
ok      github.com/pkgmain/npm-bench    139.677s
```


Windows 10 Pro with Windows Defender Real-time protection disabled
```
goos: windows
goarch: amd64
pkg: github.com/pkgmain/npm-bench
BenchmarkNpm-8                 1        35336613300 ns/op
PASS
ok      github.com/pkgmain/npm-bench    35.379s
```
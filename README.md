### npm-bench
# Simple benchmark on npm install. Dual boot Dell XPS 15 9550 (Windows 10 Pro / Manjaro)

```
goos: windows
goarch: amd64
pkg: github.com/pkgmain/npm-bench
BenchmarkNpm-8                 1        139568871500 ns/op
PASS
ok      github.com/pkgmain/npm-bench    139.677s
```

```
goos: linux
goarch: amd64
pkg: github.com/pkgmain/npm-bench
BenchmarkNpm-8   	       1	24126985145 ns/op
PASS
ok  	github.com/pkgmain/npm-bench	24.129s

```
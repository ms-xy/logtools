## logtools

This go package is a wrapper package for the amazing
[logrus](github.com/sirupsen/logrus) logging package.
It provides an `Initialize()` function that is called automatically on first
use. The initializer sets a sane default log formatter.
Currently the implemenation defaults to setting the loglevel to debug.
If you don't like this, then call `Initialize()` explizitly and afterwards
set the logging level yourself by using `SetLevel(InfoLevel)`
(or whichever level you want).

## Usage

```shell
go get github.com/ms-xy/logtools
```

```go
import (
  "github.com/ms-xy/logtools"
)

func main() {
  logtools.Info("Some informative log message")
  logtools.WithFields(logtools.Fields{
    "key": "value",
  }).Warn("More informative with the WithFields Wrapper of logrus")
}
```

To set the loglevel explicitly:

```go
func main() {
  logtools.Initialize()
  logtools.SetLevel(logtools.WarnLevel)
  logtools.Info("Some informative log message") // <- does not print anymore
  logtools.WithFields(logtools.Fields{
    "key": "value",
  }).Warn("More informative with the WithFields Wrapper of logrus")
}
```

## License

MIT - just as the logrus package.

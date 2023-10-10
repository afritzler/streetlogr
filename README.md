# StreetLogr

StreetLogr is a slick, urban-themed logging interface built around Go's `logr` interface.

## 🚀 Features

- **🗣️ Word**: Drop informational messages in style.
- **👀 Sus**: Log errors when things look suspicious.
- **🔍 Peep**: Adjust verbosity levels to peep more details.
- **🌟 Swag**: Add some swag to your logs with key-value pairs.
- **🏷️ Handle**: Tag your logger with a unique name.

## 🛠 Installation

Install StreetLogr using `go get`:

```sh
go get github.com/afritzler/streetlogr
```

## Usage

Here's a quick start guide to using StreetLogr:

```go
package main

import (
	"errors"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"github.com/afritzler/streetlogr"
)

func main() {
	// Instantiate a zap logger
	zapLogger, _ := zap.NewDevelopment()

	// Wrap it with zapr to get a logr.Logger
	logrLogger := zapr.NewLogger(zapLogger)

	// Instantiate StreetLogr
	logger := streetlogr.StreetLogr{logrLogger}

	// Use the logger
	logger.Word("The buzz in the street", "key1", "value1")

	err := errors.New("some creepy error")
	logger.Sus(err, "It's looking sus!", "key1", "value1")
}
```

A more extensible example can be found in the [example](/example) folder.

## 🤝 Contributing

Contributions are what make the open-source community an amazing place to learn, inspire, and create. Any contributions you make are greatly appreciated.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

Distributed under the Apache 2.0 License. See [LICENSE.md](LICENSE) for more information.

## 📣 Acknowledgements

* Go `logr`
* Uber's `zap`
* The incredible Go community

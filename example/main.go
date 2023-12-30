// Copyright 2023 Andreas Fritzler
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"errors"

	"github.com/afritzler/streetlogr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
)

func main() {
	// Instantiate a zap logger
	zapLogger, _ := zap.NewDevelopment()

	// Wrap it with zapr to get a logr.Logger
	logrLogger := zapr.NewLogger(zapLogger)

	// Instantiate StreetLogr with zapr
	logger := streetlogr.StreetLogr{Logger: logrLogger}

	// Now you can use your logger
	logger.Word("What's the word?", "key1", "value1")

	err := errors.New("something bad happened")
	logger.Sus(err, "It's looking pretty sus.", "key1", "value1")

	verboseLogger := logger.Peep(1)
	verboseLogger.Word("Peep this!", "key1", "value1")

	contextLogger := logger.Swag("key3", "value3")
	contextLogger.Word("Check out this swag!", "key1", "value1")

	namedLogger := logger.Handle("MyLogger")
	namedLogger.Word("This logger has a handle.", "key1", "value1")
}

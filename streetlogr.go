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

package streetlogr

import (
	"github.com/go-logr/logr"
)

// StreetLogr is a struct that embeds the logr.Logger interface
// and renames the methods with street slang for a cooler vibe.
type StreetLogr struct {
	logr.Logger
}

// Word logs informational messages. It's like the "word on the street."
func (l StreetLogr) Word(msg string, keysAndValues ...interface{}) {
	// Calling the embedded Info method from logr.Logger
	l.Info(msg, keysAndValues...)
}

// Sus logs error messages. "Sus" is slang for suspicious or something not right.
func (l StreetLogr) Sus(err error, msg string, keysAndValues ...interface{}) {
	// Calling the embedded Error method from logr.Logger
	l.Error(err, msg, keysAndValues...)
}

// Peep returns a new logger with the specified verbosity level.
// "Peep" is slang for taking a quick look at something.
func (l StreetLogr) Peep(level int) StreetLogr {
	// Calling the embedded V method from logr.Logger with a given level
	return StreetLogr{l.V(level)}
}

// Swag adds key-value pairs to the logger's context.
// "Swag" is a popular street term referring to a cool, confident style.
func (l StreetLogr) Swag(keysAndValues ...interface{}) StreetLogr {
	// Calling the embedded WithValues method from logr.Logger with given key-value pairs
	return StreetLogr{l.WithValues(keysAndValues...)}
}

// Handle adds a new name element to the logger.
// In street slang, "handle" often refers to a nickname or alias.
func (l StreetLogr) Handle(name string) StreetLogr {
	// Calling the embedded WithName method from logr.Logger with a given name
	return StreetLogr{l.WithName(name)}
}

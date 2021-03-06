// Copyright 2018 Drone.IO Inc
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

package template

import (
	"testing"
	"time"

	"github.com/tkuchiki/faketime"
)

func TestToDuration(t *testing.T) {
	from := float64(time.Date(2017, time.November, 15, 23, 0, 0, 0, time.UTC).Unix())

	vals := map[int64]string{
		time.Date(2018, time.November, 15, 23, 0, 0, 0, time.UTC).Unix():   "8760h0m0s",
		time.Date(2017, time.November, 16, 23, 0, 0, 0, time.UTC).Unix():   "24h0m0s",
		time.Date(2017, time.November, 15, 23, 30, 0, 0, time.UTC).Unix():  "30m0s",
		time.Date(2017, time.November, 15, 23, 10, 15, 0, time.UTC).Unix(): "10m15s",
		time.Date(2017, time.October, 15, 23, 0, 0, 0, time.UTC).Unix():    "-744h0m0s",
	}

	for input, want := range vals {
		if got := toDuration(from, float64(input)); got != want {
			t.Errorf("Want transform %f-%f to %s, got %s", from, float64(input), want, got)
		}
	}
}

func TestTruncate(t *testing.T) {
	vals := map[string]string{
		"foobarz": "fooba",
		"foöäüüu": "foöäü",
		"üpsßßßk": "üpsßß",
		"1234567": "12345",
		"!'§$%&/": "!'§$%",
	}

	for input, want := range vals {
		if got := truncate(input, 5); got != want {
			t.Errorf("Want transform %s to %s, got %s", input, want, got)
		}
	}
}

func TestSince(t *testing.T) {
	f := faketime.NewFaketime(2017, time.November, 15, 23, 0, 0, 0, time.UTC)
	defer f.Undo()
	f.Do()

	vals := map[int64]string{
		time.Date(2016, time.November, 15, 23, 0, 0, 0, time.UTC).Unix():   "8760h0m0s",
		time.Date(2017, time.November, 14, 23, 0, 0, 0, time.UTC).Unix():   "24h0m0s",
		time.Date(2017, time.November, 15, 22, 30, 0, 0, time.UTC).Unix():  "30m0s",
		time.Date(2017, time.November, 15, 22, 10, 15, 0, time.UTC).Unix(): "49m45s",
		time.Date(2017, time.December, 15, 23, 0, 0, 0, time.UTC).Unix():   "-720h0m0s",
	}

	for input, want := range vals {
		if got := since(input); got != want {
			t.Errorf("Want transform %d to %s, got %s", input, want, got)
		}
	}
}

func TestUppercaseFirst(t *testing.T) {
	vals := map[string]string{
		"hello":  "Hello",
		"ßqwert": "ßqwert",
		"üps":    "Üps",
		"12345":  "12345",
		"Foobar": "Foobar",
	}

	for input, want := range vals {
		if got := uppercaseFirst(input); got != want {
			t.Errorf("Want transform %s to %s, got %s", input, want, got)
		}
	}
}

func TestRegexReplace(t *testing.T) {
	expected := "hello-my-String-123"
	actual := regexReplace("(.*?)\\/(.*)", "hello/my-String-123", "$1-$2")
	if actual != "hello-my-String-123" {
		t.Errorf("error, expected %s, got %s", expected, actual)
	}
}

func TestTrimLeft(t *testing.T) {
	vals := map[string]string{
		"hello\n":         "hello\n",
		"\nhello\n":       "hello\n",
		"\r\nhello\n":     "hello\n",
		"\n\r\nhello\r\n": "hello\r\n",
	}

	for input, want := range vals {
		if got := trimLeft(input); got != want {
			t.Errorf("Want to left trim spaces from string %s to %s, got %s", input, want, got)
		}
	}
}

func TestTrimRight(t *testing.T) {
	vals := map[string]string{
		"\nhello":         "\nhello",
		"\nhello\n":       "\nhello",
		"\nhello\r\n":     "\nhello",
		"\r\nhello\n\r\n": "\r\nhello",
	}

	for input, want := range vals {
		if got := trimRight(input); got != want {
			t.Errorf("Want to left trim spaces from string %s to %s, got %s", input, want, got)
		}
	}
}

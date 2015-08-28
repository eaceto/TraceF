// Copyright 2015 Ezequiel L. Aceto. All rights reserved.
// Use of this source code is governed by a GNU v2.0-style
// license that can be found in the LICENSE file.


package tracef

import (
	"time"
	"io"
	"os"
	"fmt"
)

var traceEnabled bool

// trace is disbled by default
func init() {
	traceEnabled = false
}

// Enabled returns true if the tracer is enabled
func Enabled() (bool) {
	return traceEnabled
}

// SetEnable enable or disable the trace
func SetEnabled (enable bool) {
	traceEnabled = enable
}

// Trace to stdout the entering to a function call
func Trace(s string) ( io.Writer,  string,  time.Time )   {
	return TraceToWriter(os.Stdout,s)
}

// Trace to a writer the entering to a function call
func TraceToWriter(writer io.Writer, as string) ( io.Writer, string,  time.Time )   {
	now := time.Now()

	if traceEnabled == false {
		return nil, "", now
	}

	str := fmt.Sprint("[",now,"]\t-> ", as, "\n")
	io.WriteString(writer, str)
	return writer, as, now
}

// Trace to a writer the leaving to a function call
func Un(writer io.Writer, as string, start time.Time ) {
	if traceEnabled == false || writer == nil {
		return
	}

	now := time.Now()
	str := fmt.Sprint("[",now,"]\t<- ", as, "\t\tspent:",  time.Since(start) ,"\n")
	io.WriteString(writer, str)
}


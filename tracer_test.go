package tracer_test

import (
	"github.com/eaceto/tracer"
	"testing"
	"fmt"
)

func TestEnteringTrace(t *testing.T) {
	defer tracer.Un(tracer.Trace("ExampleEnteringTrace"))
	fmt.Println("\tthis should go between Trace and Un(trace) messages")
}


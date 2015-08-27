package tracer_test

import (
	"github.com/eaceto/tracer"
	"testing"
)

func TestEnteringTrace(t *testing.T) {
	tracer.Un(tracer.Trace("ExampleEnteringTrace"))
}


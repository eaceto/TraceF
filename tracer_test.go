package tracer_test

import (
	"github.com/eaceto/tracer"
	"testing"
	"fmt"
)

func TestIsEnabled (t *testing.T) {
	if tracer.Enabled() == true {
		t.Errorf("Tracer should be disabled by default.")
	}
}

func TestTrace (t *testing.T) {
	tracer.SetEnabled(true)
	defer tracer.Un(tracer.Trace("TestTrace"))
	fmt.Println("\tthis should be between Trace and Un(trace) messages")
}

func TestTraceDisabled (t *testing.T) {
	tracer.SetEnabled(false)
	defer tracer.Un(tracer.Trace("TestTraceDisabled"))
	fmt.Println("\tno trace message should be printed on Stdout")
}




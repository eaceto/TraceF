package tracef_test

import (
	"github.com/eaceto/tracer"
	"testing"
	"fmt"
)

func TestIsEnabled (t *testing.T) {
	if tracef.Enabled() == true {
		t.Errorf("Tracer should be disabled by default.")
	}
}

func TestTrace (t *testing.T) {
	tracef.SetEnabled(true)
	defer tracef.Un(tracef.Trace("TestTrace"))
	fmt.Println("\tthis should be between Trace and Un(trace) messages")
}

func TestTraceDisabled (t *testing.T) {
	tracef.SetEnabled(false)
	defer tracef.Un(tracef.Trace("TestTraceDisabled"))
	fmt.Println("\tno trace message should be printed on Stdout")
}




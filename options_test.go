package options_test

import (
	"flag"
	"testing"
	"time"

	"github.com/jaytaylor/go-options"
)

// TestFlagSetDefaults verifies that default flag values are applied in the
// absence of user-specified setting.
func TestFlagSetDefaults(t *testing.T) {
	flagSet := flag.NewFlagSet("TestFlagSetDefaults", flag.PanicOnError)

	flagSet.Int64("max-size", 1024768, "maximum size")
	flagSet.Duration("timeout", 1*time.Hour, "timeout setting")
	flagSet.String("description", "", "description info")

	if err := flagSet.Parse([]string{"-timeout=5s"}); err != nil {
		t.Fatal(err)
	}

	opts := &Options{}
	cfg := map[string]interface{}{}

	options.Resolve(opts, flagSet, cfg)

	if expected, actual := flagSet.Lookup("max-size").Value.(flag.Getter).Get().(int64), opts.MaxSize; actual != expected {
		t.Errorf("Expected opts.MaxSize to default to %v but actual=%v", expected, actual)
	}
}

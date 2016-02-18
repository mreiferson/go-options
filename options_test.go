package options_test

import (
	"flag"
	"fmt"
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

func TestFloat64(t *testing.T) {
	type ConfigurableThing struct {
		Percentage float64 `flag:"percentage"`
	}

	const defaultValue = 0.5

	testCases := []struct {
		Args     []string
		Expected float64
	}{
		{
			Args:     []string{""},
			Expected: defaultValue,
		},
		{
			Args:     []string{},
			Expected: defaultValue,
		},
		{
			Args:     []string{"-percentage", fmt.Sprint(defaultValue)},
			Expected: defaultValue,
		},
		{
			Args:     []string{"-percentage", "0.753"},
			Expected: 0.753,
		},
		{
			Args:     []string{"-percentage", "-0.753"},
			Expected: -0.753,
		},
		{
			Args:     []string{"-percentage=-0.117983"},
			Expected: -0.117983,
		},
	}

	for i, testCase := range testCases {
		flagSet := flag.NewFlagSet("TestFloat64", flag.PanicOnError)

		flagSet.Float64("percentage", defaultValue, "integer or decimal representing the percentage")

		if err := flagSet.Parse(testCase.Args); err != nil {
			t.Fatal(err)
		}

		configThing := &ConfigurableThing{}
		cfg := map[string]interface{}{}

		options.Resolve(configThing, flagSet, cfg)

		if expected, actual := testCase.Expected, configThing.Percentage; actual != expected {
			t.Errorf("[i=%v testCase=%+v] Expected configThing.Percentage=%v but actual=%v", i, testCase, expected, actual)
		}
	}
}

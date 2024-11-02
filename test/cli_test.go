package test

import (
	"os"
	"testing"
	"weather-api/cli"
	"weather-api/internal/types"
)

// Testing cli arguments parsing
func TestParseCliArgs(t *testing.T) {
	os.Args = []string{"cmd", "New York", "-f"}
	args := cli.ParseCliArgs()

	t.Run("New York -f", func(t *testing.T) {
		if args.Location != "New York" {
			t.Errorf("Expected location New York, got: %s", args.Location)
		}
		if args.WeatherType != types.Forecast {
			t.Errorf("Expected weather type Forecast, got: %s", args.WeatherType)
		}
	})

	os.Args = []string{"cmd", "Warsaw", "-c"}
	args = cli.ParseCliArgs()
	t.Run("Warsaw -c", func(t *testing.T) {
		if args.Location != "Warsaw" {
			t.Errorf("Expected location Warsaw, got: %s", args.Location)
		}
		if args.WeatherType != types.Current {
			t.Errorf("Expected weather type Current, got: %s", args.WeatherType)
		}
	})

	os.Args = []string{"cmd", "Dubai", "-s"}
	args = cli.ParseCliArgs()
	t.Run("Dubai -s", func(t *testing.T) {
		if args.Location != "Dubai" {
			t.Errorf("Expected location Dubai, got: %s", args.Location)
		}
		if args.WeatherType != types.Current {
			t.Errorf("Expected weather type Current, got: %s", args.WeatherType)
		}
	})

	os.Args = []string{"cmd", "Dubai", "    ", "garbage"}
	args = cli.ParseCliArgs()
	t.Run("Dubai [blank_space] garbage", func(t *testing.T) {
		if args.Location != "Dubai" {
			t.Errorf("Expected location Dubai, got: %s", args.Location)
		}
		if args.WeatherType != types.Current {
			t.Errorf("Expected weather type Current, got: %s", args.WeatherType)
		}
	})

	os.Args = []string{"cmd", "Washington"}
	args = cli.ParseCliArgs()
	t.Run("Washington", func(t *testing.T) {
		if args.Location != "Washington" {
			t.Errorf("Expected location Washington, got: %s", args.Location)
		}
		if args.WeatherType != types.Current {
			t.Errorf("Expected weather type Current, got: %s", args.WeatherType)
		}
	})

}

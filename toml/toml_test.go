package toml_test

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/influxdata/config"
	"github.com/influxdb/influxdb/cmd/influxd/run"
)

// Ensure that megabyte sizes can be parsed.
func TestSize_UnmarshalTOML_MB(t *testing.T) {
	var s config.Size
	if err := s.UnmarshalTOML([]byte("200m")); err != nil {
		t.Fatalf("unexpected error: %s", err)
	} else if s != 200*(1<<20) {
		t.Fatalf("unexpected size: %d", s)
	}
}

// Ensure that gigabyte sizes can be parsed.
func TestSize_UnmarshalTOML_GB(t *testing.T) {
	var s config.Size
	if err := s.UnmarshalTOML([]byte("1g")); err != nil {
		t.Fatalf("unexpected error: %s", err)
	} else if s != 1073741824 {
		t.Fatalf("unexpected size: %d", s)
	}
}

func TestConfig_Encode(t *testing.T) {
	var c run.Config
	c.Cluster.WriteTimeout = config.Duration(time.Minute)
	buf := new(bytes.Buffer)
	if err := config.NewEncoder(buf).Encode(c); err != nil {
		t.Fatal("Failed to encode: ", err)
	}
	got, search := buf.String(), `write-timeout="1m0s"`
	if !strings.Contains(got, search) {
		t.Fatalf("Encoding config failed.\nfailed to find %s in:\n%s\n", search, got)
	}
}

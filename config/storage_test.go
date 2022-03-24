package config

import (
	"reflect"
	"testing"

	"github.com/minamijoyo/tfmigrate/storage"
	"github.com/minamijoyo/tfmigrate/storage/local"
)

func TestParseStorageBlock(t *testing.T) {
	cases := []struct {
		desc   string
		source string
		want   storage.Config
		ok     bool
	}{
		{
			desc: "valid",
			source: `
tfmigrate {
  history {
    storage "local" {
      path = "tmp/history.json"
    }
  }
}
`,
			want: &local.Config{
				Path: "tmp/history.json",
			},
			ok: true,
		},
		{
			desc: "unknown type",
			source: `
tfmigrate {
  history {
    storage "foo" {
    }
  }
}
`,
			want: nil,
			ok:   false,
		},
		{
			desc: "missing type",
			source: `
tfmigrate {
  history {
    storage {
    }
  }
}
`,
			want: nil,
			ok:   false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			config, err := ParseConfigurationFile("test.hcl", []byte(tc.source))
			if tc.ok && err != nil {
				t.Fatalf("unexpected err: %s", err)
			}
			if !tc.ok && err == nil {
				t.Fatalf("expected to return an error, but no error, got: %#v", config)
			}
			if tc.ok {
				got := config.History.Storage
				if !reflect.DeepEqual(got, tc.want) {
					t.Errorf("got: %#v, want: %#v", got, tc.want)
				}
			}
		})
	}
}

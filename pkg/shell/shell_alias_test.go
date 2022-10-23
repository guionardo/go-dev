package shell

import (
	"os"
	"reflect"
	"testing"
)

func Test_remove(t *testing.T) {
	type args struct {
		slice []string
		s     int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{{
		name: "t0",
		args: args{
			slice: []string{"A", "B", "C"},
			s:     2,
		},
		want: []string{"A", "B"},
	},
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := remove(tt.args.slice, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetRemove(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
		// Setup
		tmp, err := os.CreateTemp(".", "test")
		if err != nil {
			t.Errorf("Failed to create temporary file %v", err)
		}
		defer os.Remove(tmp.Name())

		_, err = GetAlias(tmp.Name(), "ALIAS_TEST")
		if err == nil {
			t.Error("Expected alias not found")
		}

		if err := SetAlias(tmp.Name(), "ALIAS_TEST", "ls"); err != nil {
			t.Errorf("SetAlias() error = %v", err)
		}
		if err := SetAlias(tmp.Name(), "ALIAS_2", "pwd"); err != nil {
			t.Errorf("SetAlias() error = %v", err)
		}

		if cmd, err := GetAlias(tmp.Name(), "ALIAS_TEST"); err != nil || cmd != "ls" {
			t.Errorf("GetAlias() expected %s -> got %s %v", "ls", cmd, err)
		}
		if cmd, err := GetAlias(tmp.Name(), "ALIAS_2"); err != nil || cmd != "pwd" {
			t.Errorf("GetAlias() expected %s -> got %s %v", "pwd", cmd, err)
		}
		if err := SetAlias(tmp.Name(), "ALIAS_TEST", "ls -la"); err != nil {
			t.Errorf("SetAlias() error = %v", err)
		}
		if cmd, err := GetAlias(tmp.Name(), "ALIAS_TEST"); err != nil || cmd != "ls -la" {
			t.Errorf("GetAlias() expected %s -> got %s %v", "ls -la", cmd, err)
		}
		if err = RemoveAlias(tmp.Name(), "ALIAS_TEST"); err != nil {
			t.Errorf("RemoveAlias() error = %v", err)
		}
		if _, err = GetAlias(tmp.Name(), "ALIAS_TEST"); err == nil {
			t.Errorf("GetAlias() expected ALIAS_TEST not found")
		}

	})

}

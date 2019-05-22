package js_test

import (
	"testing"

	"github.com/olegsu/iris/pkg/v2/js"
	"github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto/underscore"
	"github.com/stretchr/testify/assert"
)

func TestNewJSEngine(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Create JS engine",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := js.NewJSEngine(); !assert.NotNil(t, got) {
				t.Errorf("NewJSEngine() = %v, want %v", got, false)
			}
		})
	}
}

func Test_runner_Load(t *testing.T) {
	type args struct {
		scripts []string
		input   map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *otto.Value
		wantErr bool
	}{
		{
			name: "Run basic script",
			args: args{
				scripts: []string{"console.log('test')"},
				input:   nil,
			},
			wantErr: false,
			want:    &otto.Value{},
		},
		{
			name: "Run script with globals",
			args: args{
				scripts: []string{"console.log(JSON.stringify(this))"},
				input: map[string]interface{}{
					"key": "value",
				},
			},
			wantErr: false,
			want:    &otto.Value{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := js.NewJSEngine()
			got, err := r.Load(tt.args.scripts, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("runner.Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.Equal(t, tt.want, got) {
				t.Errorf("runner.Load() = %v, want %v", got, tt.want)
			}
		})
	}
}

package language

import (
	"testing"

	"github.com/olegsu/cli-generator/pkg/logger"
	"github.com/olegsu/cli-generator/pkg/spec"
	"github.com/stretchr/testify/assert"
)

func Test_golang_Render(t *testing.T) {
	type fields struct {
		logger           logger.Logger
		projectDirectory string
		generateHandlers bool
		runInitFlow      bool
		spec             *spec.CLISpec
	}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*RenderResult
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				logger:           logger.New(nil),
				projectDirectory: "./",
				generateHandlers: false,
				runInitFlow:      false,
				spec:             &spec.CLISpec{},
			},
			args: args{
				data: make(map[string]interface{}),
			},
			wantErr: false,
			want:    []*RenderResult{&RenderResult{}, &RenderResult{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gg := &golang{
				logger:           tt.fields.logger,
				projectDirectory: tt.fields.projectDirectory,
				generateHandlers: tt.fields.generateHandlers,
				runInitFlow:      tt.fields.runInitFlow,
				spec:             tt.fields.spec,
			}
			got, err := gg.Render(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("golang.Render() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Len(t, got, 2)
		})
	}
}

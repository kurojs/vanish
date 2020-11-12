package vanish

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestRemoveFields(t *testing.T) {
	type args struct {
		str    string
		fields []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "simple JSON",
			args: args{
				str:    `{"foo":"bar"}`,
				fields: []string{"foo"},
			},
			want:    `{}`,
			wantErr: false,
		},
		{
			name: "remove single field from nested array JSON",
			args: args{
				str: `
				{
					"nested_arr": [
					  "a",
					  "b",
					  123,
					  {
						"nested_arr_string": "abc",
						"nested_arr_number": 1,
						"nested_arr_obj": {
						  "a": 1
						}
					  }
					]
				  }
				`,
				fields: []string{"nested_arr.nested_arr_string"},
			},
			want: `
			{
				"nested_arr": [
				  "a",
				  "b",
				  123,
				  {
					"nested_arr_number": 1,
					"nested_arr_obj": {
					  "a": 1
					}
				  }
				]
			  }
			`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RemoveFields([]byte(tt.args.str), tt.args.fields)
			wantJSON, _ := json.Marshal(tt.want)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveFields() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if reflect.DeepEqual(wantJSON, got) {
				t.Errorf("RemoveFields() = %v, want %v", string(got), tt.want)
			}
		})
	}
}

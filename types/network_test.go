package types

import (
	"reflect"
	"testing"
)

func TestParseNetworkLimit(t *testing.T) {
	type args struct {
		limit string
	}
	tests := []struct {
		name string
		args args
		want NetworkLimit
	}{
		{
			name: "Test Parse Network Limit",
			args: args{
				limit: "100k/10M",
			},
			want: NetworkLimit{
				Upload: NetworkSpeed{
					Amount: 10,
					Unit:   NetworkSpeedUnitMbps,
				},
				Download: NetworkSpeed{
					Amount: 100,
					Unit:   NetworkSpeedUnitKbps,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseNetworkLimit(tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseNetworkLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}

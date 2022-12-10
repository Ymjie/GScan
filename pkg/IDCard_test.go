package pkg

import "testing"

func TestIsIDCard(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "sfzcs",
			args: args{id: "371323255555554545"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIDCard(tt.args.id); got != tt.want {
				t.Errorf("IsIDCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

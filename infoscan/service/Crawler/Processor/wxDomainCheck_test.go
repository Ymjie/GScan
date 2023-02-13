package Processor

import (
	"GScan/infoscan/dao"
	"GScan/pkg"
	"testing"
)

func Test_wxDominCheck_check(t *testing.T) {
	type args struct {
		url0 string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{
			name: "vshex",
			args: args{url0: "http://vshex.com"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WXDomainCheck{
				Scheduler: pkg.QueueScheduler[*dao.Page]{},
			}
			w.Run()
			got, got1 := w.check(tt.args.url0)
			if got != tt.want {
				t.Errorf("check() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("check() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

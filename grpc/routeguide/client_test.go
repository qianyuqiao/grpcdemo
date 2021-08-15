package routeguide

import (
	"context"
	"reflect"
	"testing"
	"time"
)

func TestPing(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *PingRequest
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second * 2)
	valCtx := context.WithValue( ctx, "client-ip", "127.0.0.1")
	tests := []struct {
		name    string
		args    args
		want    *PingReply
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				ctx: valCtx,
				in:  &PingRequest{},
			},
			want: &PingReply{
				Reply:                "pong",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Ping(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Ping() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ping() got = %v, want %v", got, tt.want)
			}
		})
	}
}
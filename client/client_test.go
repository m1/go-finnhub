package client

import (
	"reflect"
	"testing"

	"github.com/m1/go-finnhub/crypto"
	"github.com/m1/go-finnhub/forex"
	"github.com/m1/go-finnhub/news"
	"github.com/m1/go-finnhub/stock"
)

func TestNew(t *testing.T) {
	a := NewAPI("token", Version)
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		{
			name: "valid",
			args: args{key: "token"},
			want: &Client{
				Stock:  stock.Client{API: a},
				Forex:  forex.Client{API: a},
				Crypto: crypto.Client{API: a},
				News:   news.Client{API: a},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

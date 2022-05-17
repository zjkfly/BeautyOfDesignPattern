package main

import (
	"context"
	"testing"
)

func TestRandomIdGenerator_generate(t *testing.T) {
	type fields struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name:    "nil ctx",
			fields:  fields{nil},
			want:    "",
			wantErr: false,
		},
		{
			name:    "todo ctx",
			fields:  fields{context.TODO()},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ig := &RandomIdGenerator{
				ctx: tt.fields.ctx,
			}
			got, err := ig.generate()
			if (err != nil) != tt.wantErr {
				t.Errorf("generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("generate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
package main

import "testing"

func TestToLowerCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Should return all lowercases when chars are passed",
			args: args{
				s: "Golang",
			},
			want:    "golang",
			wantErr: false,
		},
		{
			name: "Should return error when empty string is passd",
			args: args{
				s: "",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToLowerCase(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToLowerCase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToLowerCase() got = %v, want %v", got, tt.want)
			}
		})
	}
}

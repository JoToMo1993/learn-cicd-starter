package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type args struct {
		headers http.Header
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "no auth header",
			args: args{
				headers: http.Header{},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "no separation",
			args: args{
				headers: http.Header{
					"Authorization": []string{
						"singlestring",
					},
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "no ApiKey",
			args: args{
				headers: http.Header{
					"Authorization": []string{
						"Bearer singlestring",
					},
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "ApiKey",
			args: args{
				headers: http.Header{
					"Authorization": []string{
						"ApiKey bla",
					},
				},
			},
			want:    "bla",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.args.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}

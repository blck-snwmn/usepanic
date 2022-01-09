package usepanic

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyze(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "main", "other", "foo")
}

func Test_allowPackagesFlags_String(t *testing.T) {
	tests := []struct {
		name string
		apf  *allowPackagesFlags
		want string
	}{
		{
			name: "empty",
			apf:  &allowPackagesFlags{},
			want: "",
		},
		{
			name: "not empty",
			apf: &allowPackagesFlags{
				"x": {}, "y": {},
			},
			want: "x, y",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.apf.String(); got != tt.want {
				t.Errorf("allowPackagesFlags.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_allowPackagesFlags_Set(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		apf     *allowPackagesFlags
		args    args
		wantErr bool
	}{
		{
			name: "empty",
			apf:  &allowPackagesFlags{},
			args: args{
				s: "",
			},
			wantErr: false,
		},
		{
			name: "empty element exist",
			apf: &allowPackagesFlags{
				"x": {},
			},
			args: args{
				s: "x,",
			},
			wantErr: false,
		},
		{
			name: "not",
			apf: &allowPackagesFlags{
				"x": {}, "y": {},
			},
			args: args{
				s: "x,y",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.apf.Set(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("allowPackagesFlags.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

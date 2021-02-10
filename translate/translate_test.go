package translate

import "testing"

func Test_textToBinaryImpl_Translate(t *testing.T) {
	type args struct {
		val string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Text to binary test",
			args: args{
				val: "CC",
			},
			want: "0100001101000011",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ttob := &textToBinaryImpl{}
			got, err := ttob.Translate(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("Translate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Translate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

package validate

import "testing"

func Test_validateFormatImpl_Validate(t *testing.T) {
	type fields struct {
		next Validator
	}

	type args struct {
		format string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Formato valido test",
			fields: fields{
			},
			args: args{
				format: "TEXT",
			},
			wantErr: false,
		},
		{
			name: "Formato invalido test",
			fields: fields{
			},
			args: args{
				format: "ASSCI",
			},
			wantErr: true,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vfi := &validateFormatImpl{
				next: tt.fields.next,
			}
			if err := vfi.Validate(tt.args.format); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}



func Test_validateFormatDestinoImpl_Validate(t *testing.T) {
	type fields struct {
		next Validator
	}

	type args struct {
		format string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Formato valido test",
			fields: fields{
			},
			args: args{
				format: "TEXT",
			},
			wantErr: false,
		},
		{
			name: "Formato invalido test",
			fields: fields{
			},
			args: args{
				format: "ASSCI",
			},
			wantErr: true,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vfi := &validateFormatDestinoImpl{
				next: tt.fields.next,
			}
			if err := vfi.Validate(tt.args.format); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

package sign

import (
	"testing"
)

func TestToHMAC(t *testing.T) {
	type args struct {
		text string
		key  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{},
			want: "",
		},
		{
			name: "default",
			args: args{
				text: "test",
				key:  "test123",
			},
			want: "6235393665323437333966643434643432666664323566323665613336376461643361373166363163386335666162366236656536636565616535613731373062363634343564366561616466623439653664346539363861323838383732366666353232653362663036356339363661613636613234313533373738333832",
		},
		{
			name: "rand",
			args: args{
				text: "0x10000",
				key:  "0x10000",
			},
			want: "3131343136386566393837386331383764343936666437383466366139623936346438646232313539643035313438636165373234636362303761306236346661353832333234643761626563316137633664666438643762613933386332623266323738353835333138356635376636656363316135633338396466303563",
		},
		//{
		//	name: "empty",
		//	args: args{},
		//	want: "",
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if encode(tt.args.text, tt.args.key) == tt.want {
				t.Fail()
			}
		})
	}
}

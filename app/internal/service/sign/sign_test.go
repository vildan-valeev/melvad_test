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
			want: "b936cee86c9f87aa5d3c6f2e84cb5a4239a5fe50480a6ec66b70ab5b1f4ac6730c6c515421b327ec1d69402e53dfb49ad7381eb067b338fd7b0cb22247225d47",
		},
		{
			name: "default",
			args: args{
				text: "test",
				key:  "test123",
			},
			want: "b596e24739fd44d42ffd25f26ea367dad3a71f61c8c5fab6b6ee6ceeae5a7170b66445d6eaadfb49e6d4e968a2888726ff522e3bf065c966aa66a24153778382",
		},
		{
			name: "rand",
			args: args{
				text: "0x10000",
				key:  "0x10000",
			},
			want: "114168ef9878c187d496fd784f6a9b964d8db2159d05148cae724ccb07a0b64fa582324d7abec1a7c6dfd8d7ba938c2b2f2785853185f57f6ecc1a5c389df05c",
		},
		//{
		//	name: "empty",
		//	args: args{},
		//	want: "",
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if toHMAC(tt.args.text, tt.args.key) == tt.want {
				t.Fail()
			}
		})
	}

}

func TestToHEX(t *testing.T) {
	type args struct {
		hash string
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
				hash: "b596e24739fd44d42ffd25f26ea367dad3a71f61c8c5fab6b6ee6ceeae5a7170b66445d6eaadfb49e6d4e968a2888726ff522e3bf065c966aa66a24153778382",
			},
			want: "6235393665323437333966643434643432666664323566323665613336376461643361373166363163386335666162366236656536636565616535613731373062363634343564366561616466623439653664346539363861323838383732366666353232653362663036356339363661613636613234313533373738333832",
		},
		{
			name: "rand",
			args: args{
				hash: "114168ef9878c187d496fd784f6a9b964d8db2159d05148cae724ccb07a0b64fa582324d7abec1a7c6dfd8d7ba938c2b2f2785853185f57f6ecc1a5c389df05c",
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
			if toHEX(tt.args.hash) == tt.want {
				t.Fail()
			}
		})
	}

}

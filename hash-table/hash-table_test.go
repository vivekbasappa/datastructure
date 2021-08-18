// code based on Junmin Lee tutorial
package main

import "testing"

func TestHashTable_Insert(t *testing.T) {
	type fields struct {
		array [ArraySize]*bucket
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "testing first time",
			args: args{"string"},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HashTable{
				array: tt.fields.array,
			}
			h = Init()
			h.Insert(tt.args.key)
		})
	}
}

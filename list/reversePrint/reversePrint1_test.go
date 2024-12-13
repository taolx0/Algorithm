package main

import (
	"reflect"
	"testing"
)

func Test_reversePrint1(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"1",
			//args{&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}},
			args{&ListNode{Val: 1, Next: nil}},
			[]int{5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePrint1(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reversePrint1() = %v, want %v", got, tt.want)
			}
		})
	}
}

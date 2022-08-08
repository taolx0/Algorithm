package deleteNode

import (
	"testing"
)

func Test_deleteNode1(t *testing.T) {
	type args struct {
		val  int
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want args
	}{
		{
			name: "1",
			args: args{val: 1, head: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}},
			want: args{val: 2, head: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := deleteNode1(tt.args.head, tt.args.val)
			if got.Val != 2 {
				t.Error("deleteNode1() = ", got.Val)
			}
		})
	}
}

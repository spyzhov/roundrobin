package roundrobin

import (
	"testing"
)

func TestRoundRobin_Next_sync(t *testing.T) {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	output := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2}
	type testCase struct {
		name       string
		rr         RoundRobin[int]
		wantValues []int
	}
	tests := []testCase{
		{
			name:       "Chan",
			rr:         NewChan[int](array),
			wantValues: output,
		},
		{
			name:       "LinkedListNoSplit",
			rr:         NewLinkedListNoSplit[int](array),
			wantValues: output,
		},
		{
			name:       "NewLinkedListMutex",
			rr:         NewLinkedListMutex[int](array),
			wantValues: output,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, value := range tt.wantValues {
				if gotValue := tt.rr.Next(); gotValue != value {
					t.Errorf("%s.Next(%d) = %v, want %v", tt.name, i, gotValue, value)
				}
			}
		})
	}
}

func BenchmarkRoundRobin_Next(b *testing.B) {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	benchmarks := []struct {
		name  string
		input []int
		rr    RoundRobin[int]
	}{
		{
			name: "Chan",
			rr:   NewChan[int](array),
		},
		{
			name: "LL-NoSplit",
			rr:   NewLinkedListNoSplit[int](array),
		},
		{
			name: "LL-Mutex",
			rr:   NewLinkedListMutex[int](array),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = bm.rr.Next()
			}
		})
	}
}

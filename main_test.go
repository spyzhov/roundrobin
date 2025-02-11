package roundrobin

import (
	"context"
	"testing"

	"golang.org/x/sync/errgroup"
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
			name:       "LinkedListRaw",
			rr:         NewLinkedListRaw[int](array),
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

func TestRoundRobin_Next_async_race_check(t *testing.T) {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	type testCase struct {
		name string
		rr   RoundRobin[int]
	}
	tests := []testCase{
		{
			name: "Chan",
			rr:   NewChan[int](array),
		},
		{
			name: "LinkedListRaw",
			rr:   NewLinkedListRaw[int](array),
		},
		{
			name: "NewLinkedListMutex",
			rr:   NewLinkedListMutex[int](array),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, _ := errgroup.WithContext(context.Background())
			for i := 0; i < 100; i++ {
				g.Go(func() error {
					tt.rr.Next()
					return nil
				})
			}
			if err := g.Wait(); err != nil {
				t.Error(err)
			}
		})
	}
}

func BenchmarkRoundRobin_Next_sync(b *testing.B) {
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
			name: "LL-Raw",
			rr:   NewLinkedListRaw[int](array),
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

func BenchmarkRoundRobin_Next_async(b *testing.B) {
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
			name: "LL-Raw",
			rr:   NewLinkedListRaw[int](array),
		},
		{
			name: "LL-Mutex",
			rr:   NewLinkedListMutex[int](array),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_ = bm.rr.Next()
				}
			})
		})
	}
}

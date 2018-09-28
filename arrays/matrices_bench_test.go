package arrays

import "testing"

//Important lesson: In-place algos will use more memory, but more allocs == more time

var a = [][]int{
	{0, 1, 2, 3},
	{4, 5, 6, 7},
	{8, 9, 10, 11},
	{12, 13, 14, 15},
}

func BenchmarkRotateNewMatrix(b *testing.B) {
	for index := 0; index < b.N; index++ {
		RotateNewMatrix(a)
	}
}

func BenchmarkRotateMatrixInPlace(b *testing.B) {
	for index := 0; index < b.N; index++ {
		RotateMatrixInPlace(a)
	}
}

//BenchmarkRotateMatrixInPlace-8   	100000000	        18.6 ns/op	       0 B/op	       0 allocs/op
//BenchmarkRotateNewMatrix-8   	10000000	       176 ns/op	     224 B/op	       5 allocs/op

//NewMatrix
// (pprof) top5
// Showing nodes accounting for 2100ms, 96.77% of 2170ms total
// Dropped 12 nodes (cum <= 10.85ms)
// Showing top 5 nodes out of 43
//       flat  flat%   sum%        cum   cum%
//     1690ms 77.88% 77.88%     1690ms 77.88%  runtime.kevent
//      230ms 10.60% 88.48%      230ms 10.60%  runtime.pthread_cond_wait
//       90ms  4.15% 92.63%       90ms  4.15%  runtime.pthread_cond_signal
//       60ms  2.76% 95.39%       80ms  3.69%  runtime.mallocgc
//       30ms  1.38% 96.77%       30ms  1.38%  runtime.pthread_cond_timedwait_relative_np

//InPlace
// (pprof) top5
// Showing nodes accounting for 1.76s, 100% of 1.76s total
// Showing top 5 nodes out of 13
//       flat  flat%   sum%        cum   cum%
//      1.46s 82.95% 82.95%      1.46s 82.95%  arrays.RotateMatrixInPlace
//      0.15s  8.52% 91.48%      0.15s  8.52%  runtime.nanotime
//      0.14s  7.95% 99.43%      1.60s 90.91%  arrays.BenchmarkRotateMatrixInPlace
//      0.01s  0.57%   100%      0.01s  0.57%  runtime.kevent
//          0     0%   100%      0.13s  7.39%  runtime.mstart
// (pprof)

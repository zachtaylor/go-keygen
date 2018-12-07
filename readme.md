# `ztaylor.me/keygen`

Package keygen provides random string generation

Repository is hosted at `"ztaylor.me/keygen"`

## Benchmark Results

Results may vary. Posted results occured with Windows 10 i5-4690

### Current results

`New()` now requires `*rand.Rand`, which saves time

`NewVal()` still creates `*rand.Rand`, causing 1 alloc

```
BenchmarkNewVal-4                      	  200000	      9849 ns/op	    5384 B/op	       2 allocs/op
BenchmarkNewNumeric3-4                 	20000000	        69.0 ns/op	       3 B/op	       1 allocs/op
BenchmarkNewNumeric10-4                	10000000	       187 ns/op	      16 B/op	       1 allocs/op
BenchmarkNewNumeric100-4               	 1000000	      1597 ns/op	     112 B/op	       1 allocs/op
BenchmarkNewAlphaNumericCapital3-4     	20000000	        68.4 ns/op	       3 B/op	       1 allocs/op
BenchmarkNewAlphaNumericCapital10-4    	10000000	       185 ns/op	      16 B/op	       1 allocs/op
BenchmarkNewAlphaNumericCapital100-4   	 1000000	      1650 ns/op	     112 B/op	       1 allocs/op

```
### "initial commit" 156e2d22c2f4c05dc0216925349de8a44fba74c0
```
BenchmarkKeygenNumeric3-4                 	  200000	      9789 ns/op	    5379 B/op	       2 allocs/op
BenchmarkKeygenNumeric10-4                	  200000	      9912 ns/op	    5392 B/op	       2 allocs/op
BenchmarkKeygenNumeric100-4               	  200000	     11945 ns/op	    5488 B/op	       2 allocs/op
BenchmarkKeygenAlphaNumericCapital3-4     	  200000	      9716 ns/op	    5379 B/op	       2 allocs/op
BenchmarkKeygenAlphaNumericCapital10-4    	  200000	     10093 ns/op	    5392 B/op	       2 allocs/op
BenchmarkKeygenAlphaNumericCapital100-4   	  200000	     11926 ns/op	    5488 B/op	       2 allocs/op
```

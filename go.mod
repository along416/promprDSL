module promptdsl_test

go 1.22

toolchain go1.24.4

require github.com/antlr4-go/antlr/v4 v4.13.1

require (
	github.com/dlclark/regexp2 v1.11.4 // indirect
	github.com/go-sourcemap/sourcemap v2.1.3+incompatible // indirect
	github.com/google/pprof v0.0.0-20230207041349-798e818bf904 // indirect
	golang.org/x/text v0.3.8 // indirect
)

require (
	github.com/antlr4-go/antlr v0.0.0-20230518091524-98b52378c522 // indirect
	github.com/dop251/goja v0.0.0-20250630131328-58d95d85e994
	golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842 // indirect
)

replace github.com/antlr/antlr4/runtime/go/antlr => ./antlr

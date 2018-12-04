# gotestx

Testing suite inspired in https://github.com/stretchr/testify and https://github.com/smartystreets/goconvey

- This suite allows all `assert` functions provided by `stretch/testify` but prints results to the terminal in a friendlier way.
- Allows blocks of tests using `Convey` as `goconvey` would, but uses assert functions from `assert` instead of `So(value, assertHandler, expected)`
- It is compatible with the `goconvey` binary.
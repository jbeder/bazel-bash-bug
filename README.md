## Example for illustrating a Bazel bug

This example tries to start a Firebase emulator from a Go binary, built through Bazel.

When running through a test (either with `bazel test` or `bazel run`), it succeeds:

```
$ bazel test --test_output=streamed  :bazel-bash-bug_test
WARNING: Streamed test output requested. All tests will be run locally, without sharding, one at a time
INFO: Analyzed target //:bazel-bash-bug_test (0 packages loaded, 0 targets configured).
INFO: Found 1 test target...
i  emulators: Starting emulators: auth
i  emulators: Detected demo project ID "demo-foo", emulated services will use a demo configuration and attempts to access non-emulated services for this project will fail.
⚠  emulators: It seems that you are running multiple instances of the emulator suite for project demo-foo. This may result in unexpected behavior.
PASS
Target //:bazel-bash-bug_test up-to-date:
  bazel-bin/bazel-bash-bug_test_/bazel-bash-bug_test
INFO: Elapsed time: 9.950s, Critical Path: 9.49s
INFO: 6 processes: 1 internal, 5 darwin-sandbox.
INFO: Build completed successfully, 6 total actions
//:bazel-bash-bug_test                                                   PASSED in 4.3s

Executed 1 out of 1 test: 1 test passes.
There were tests whose specified size is too big. Use the --test_verbose_timeoutINFO: Build completed successfully, 6 total actions
```

or

```
$ bazel run :bazel-bash-bug_test 
INFO: Build option --test_sharding_strategy has changed, discarding analysis cache.
INFO: Analyzed target //:bazel-bash-bug_test (0 packages loaded, 24668 targets configured).
INFO: Found 1 target...
Target //:bazel-bash-bug_test up-to-date:
  bazel-bin/bazel-bash-bug_test_/bazel-bash-bug_test
INFO: Elapsed time: 2.357s, Critical Path: 0.10s
INFO: 1 process: 1 internal.
INFO: Build completed successfully, 1 total action
INFO: Running command line: external/bazel_tools/tools/test/test-setup.sh bazel-INFO: Build completed successfully, 1 total action
exec ${PAGER:-/usr/bin/less} "$0" || exit 1
Executing tests from //:bazel-bash-bug_test
-----------------------------------------------------------------------------
i  emulators: Starting emulators: auth
i  emulators: Detected demo project ID "demo-foo", emulated services will use a demo configuration and attempts to access non-emulated services for this project will fail.
⚠  emulators: It seems that you are running multiple instances of the emulator suite for project demo-foo. This may result in unexpected behavior.
PASS
```

But when run through a normal Go binary, it fails:

```
$ bazel run main:main 
INFO: Analyzed target //main:main (1 packages loaded, 3 targets configured).
INFO: Found 1 target...
Target //main:main up-to-date:
  bazel-bin/main/main_/main
INFO: Elapsed time: 0.564s, Critical Path: 0.12s
INFO: 1 process: 1 internal.
INFO: Build completed successfully, 1 total action
INFO: Build completed successfully, 1 total action
ERROR: cannot find build_bazel_rules_nodejs/third_party/github.com/bazelbuild/bazel/tools/bash/runfiles/runfiles.bash
2021/11/25 12:04:28 exit status 1
```

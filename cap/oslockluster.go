// +build !go1.10

package cap

import "syscall"

// LaunchSupported indicates that is safe to return from a locked OS
// Thread and have that OS Thread be terminated by the runtime. The
// Launch functionality really needs to rely on the fact that an
// excess of runtime.LockOSThread() vs. runtime.UnlockOSThread() calls
// in a returning go routine will cause the underlying locked OSThread
// to terminate. That feature was added to the Go runtime in version
// 1.10.
//
// See these bugs for the discussion and feature assumed by the code
// in this Launch() functionality: [Go bug 20395], [Go bug 20458].
//
// A value of false for this constant causes the Launch functionality
// to fail with an error: cap.ErrNoLaunch. If this value is false you
// have two choices with respect to the Launch functionality:
//
//   1) don't use cap.(*Launcher).Launch()
//   2) upgrade your Go toolchain to 1.10+ (ie., do this one).
//
// [Go bug 20395]: https://github.com/golang/go/issues/20395
// [Go bug 20458]: https://github.com/golang/go/issues/20458
const LaunchSupported = false

// validatePA confirms that the pa.Sys entry is not incompatible with
// Launch.
func validatePA(pa *syscall.ProcAttr, chroot string) (bool, error) {
	return false, ErrNoLaunch
}

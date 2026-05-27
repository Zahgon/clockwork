package clockwork

import (
	"context"
	"fmt"
	"time"
)

// contextKey is private to this package so we can ensure uniqueness here. This
// type identifies context values provided by this package.
type contextKey string

// keyClock provides a clock for injecting during tests. If absent, a real clock
// should be used.
var keyClock = contextKey("clock") // clockwork.Clock

// AddToContext creates a derived context that references the specified clock.
//
// Be aware this doesn't change the behavior of standard library functions, such
// as [context.WithTimeout] or [context.WithDeadline]. For this reason, users
// should prefer passing explicit [clockwork.Clock] variables rather can passing
// the clock via the context.
func AddToContext(ctx context.Context, clock Clock) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// FromContext extracts a clock from the context. If not present, a real clock
// is returned.
func FromContext(ctx context.Context) Clock { _ = "STUB: not implemented"; return *new(Clock) }

// ErrFakeClockDeadlineExceeded is the error returned by [context.Context] when
// the deadline passes on a context which uses a [FakeClock].
//
// It wraps a [context.DeadlineExceeded] error, i.e.:
//
//	// The following is true for any Context whose deadline has been exceeded,
//	// including contexts made with clockwork.WithDeadline or clockwork.WithTimeout.
//
//	errors.Is(ctx.Err(), context.DeadlineExceeded)
//
//	// The following can only be true for contexts made
//	// with clockwork.WithDeadline or clockwork.WithTimeout.
//
//	errors.Is(ctx.Err(), clockwork.ErrFakeClockDeadlineExceeded)
var ErrFakeClockDeadlineExceeded error = fmt.Errorf("clockwork.FakeClock: %w", context.DeadlineExceeded)

// WithDeadline returns a context with a deadline based on a [FakeClock].
//
// The returned context ignores parent cancelation if the parent was cancelled
// with a [context.DeadlineExceeded] error. Any other error returned by the
// parent is treated normally, cancelling the returned context.
//
// If the parent is cancelled with a [context.DeadlineExceeded] error, the only
// way to then cancel the returned context is by calling the returned
// context.CancelFunc.
func WithDeadline(parent context.Context, clock Clock, t time.Time) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

// WithTimeout returns a context with a timeout based on a [FakeClock].
//
// The returned context follows the same behaviors as [WithDeadline].
func WithTimeout(parent context.Context, clock Clock, d time.Duration) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

// fakeClockContext implements context.Context, using a fake clock for its
// deadline.
//
// It ignores parent cancellation if the parent is cancelled with
// context.DeadlineExceeded.
type fakeClockContext struct {
	parent   context.Context
	deadline time.Time // The user-facing deadline based on the fake clock's time.

	// Tracks timeout/deadline cancellation.
	timerDone <-chan time.Time

	// Tracks manual calls to the cancel function.
	cancel       func() // Closes cancelCalled wrapped in a sync.Once.
	cancelCalled chan struct{}

	// The user-facing data from the context.Context interface.
	ctxDone chan struct{} // Returned by Done().
	err     error         // nil until ctxDone is ready to be closed.
}

func newFakeClockContext(parent context.Context, deadline time.Time, timer <-chan time.Time) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

// Wait until the cancellation goroutine is running.

func (c *fakeClockContext) Deadline() (time.Time, bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}

func (c *fakeClockContext) Done() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (c *fakeClockContext) Err() error {
	_ = "STUB: not implemented"
	// Don't return the error before it is ready.
	return nil
}

func (c *fakeClockContext) Value(key any) any { _ = "STUB: not implemented"; return *new(any) }

// runCancel runs the fakeClockContext's cancel goroutine and returns the
// fakeClockContext's cancel function.
//
// fakeClockContext is then cancelled when any of the following occur:
//
//   - The fakeClockContext.done channel is closed by its timer.
//   - The returned CancelFunc is executed.
//   - The fakeClockContext's parent context is cancelled with an error other
//     than context.DeadlineExceeded.
func (c *fakeClockContext) runCancel(ready chan struct{}) { _ = "STUB: not implemented"; return }

// Close ready when done, just in case the ready signal races with other
// branches of our select statement below.

// Signals the cancellation goroutine has begun, in an attempt to minimize
// race conditions related to goroutine startup time.
// This case statement can only fire once.

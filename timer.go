package clockwork

import "time"

// Timer provides an interface which can be used instead of directly using
// [time.Timer]. The real-time timer t provides events through t.C which becomes
// t.Chan() to make this channel requirement definable in this interface.
type Timer interface {
	Chan() <-chan time.Time
	Reset(d time.Duration) bool
	Stop() bool
}

type realTimer struct{ *time.Timer }

func (r realTimer) Chan() <-chan time.Time { _ = "STUB: not implemented"; return nil }

type fakeTimer struct {
	// The channel associated with the firer, used to send expiration times.
	c chan time.Time

	// The time when the firer expires. Only meaningful if the firer is currently
	// one of a FakeClock's waiters.
	exp time.Time

	// reset and stop provide the implementation of the respective exported
	// functions.
	reset func(d time.Duration) bool
	stop  func() bool

	// If present when the timer fires, the timer calls afterFunc in its own
	// goroutine rather than sending the time on Chan().
	afterFunc func()
}

func newFakeTimer(fc *FakeClock, afterfunc func()) *fakeTimer {
	_ = "STUB: not implemented"
	return nil
}

// fc.l must be held across the calls to stopExpirer & setExpirer.

func (f *fakeTimer) Chan() <-chan time.Time { _ = "STUB: not implemented"; return nil }

func (f *fakeTimer) Reset(d time.Duration) bool { _ = "STUB: not implemented"; return false }

func (f *fakeTimer) Stop() bool { _ = "STUB: not implemented"; return false }

func (f *fakeTimer) expire(now time.Time) *time.Duration { _ = "STUB: not implemented"; return nil }

// Never block on expiration.

func (f *fakeTimer) expiration() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (f *fakeTimer) setExpiration(t time.Time) { _ = "STUB: not implemented"; return }

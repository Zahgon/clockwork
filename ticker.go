package clockwork

import "time"

// Ticker provides an interface which can be used instead of directly using
// [time.Ticker]. The real-time ticker t provides ticks through t.C which
// becomes t.Chan() to make this channel requirement definable in this
// interface.
type Ticker interface {
	Chan() <-chan time.Time
	Reset(d time.Duration)
	Stop()
}

type realTicker struct{ *time.Ticker }

func (r realTicker) Chan() <-chan time.Time { _ = "STUB: not implemented"; return nil }

type fakeTicker struct {
	// The channel associated with the firer, used to send expiration times.
	c chan time.Time

	// The time when the ticker expires. Only meaningful if the ticker is currently
	// one of a FakeClock's waiters.
	exp time.Time

	// reset and stop provide the implementation of the respective exported
	// functions.
	reset func(d time.Duration)
	stop  func()

	// The duration of the ticker.
	d time.Duration
}

func newFakeTicker(fc *FakeClock, d time.Duration) *fakeTicker {
	_ = "STUB: not implemented"
	return nil
}

func (f *fakeTicker) Chan() <-chan time.Time { _ = "STUB: not implemented"; return nil }

func (f *fakeTicker) Reset(d time.Duration) { _ = "STUB: not implemented"; return }

func (f *fakeTicker) Stop() { _ = "STUB: not implemented"; return }

func (f *fakeTicker) expire(now time.Time) *time.Duration {
	_ = "STUB: not implemented"
	// Never block on expiration.
	return nil
}

func (f *fakeTicker) expiration() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (f *fakeTicker) setExpiration(t time.Time) { _ = "STUB: not implemented"; return }

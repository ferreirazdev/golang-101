1. The go keyword
- What it does: go f() starts a new goroutine: a lightweight thread managed by the Go runtime. Execution continues in the caller without waiting for f() to finish.
- Lightweight: Many goroutines can run on few OS threads; the runtime multiplexes them (M:N scheduling).
- Important: If main returns, the whole process exits and all goroutines are stopped. So you need a way to wait for goroutines when you care about their result or completion.

2. sync.WaitGroup — theory
A WaitGroup is a counter used to wait for a set of goroutines to finish.
- Add(delta int) — Increase the counter (e.g. Add(1) before starting one goroutine).
- Done() — Decrease the counter by 1 (call when a goroutine finishes). Same as Add(-1).
- Wait() — Block until the counter is 0 (all goroutines have called Done()).

3. Quick reference
go f()- Run f() in a new goroutine; don’t wait for it.
wg.Add(1) - “I’m about to start one goroutine.”
wg.Done() - “This goroutine is finished.” (same as Add(-1))
wg.Wait() - Block until all goroutines have called Done().
Pass &wg - Goroutines need a pointer to the same WaitGroup.
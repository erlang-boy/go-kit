package actor

func Stop(actor Actor) error {
	actor.Kill()
	return actor.Wait()
}

// Dead returns a channel that will be closed when the supplied
// Actor has completed.
//
// Don't be too casual about calling Dead -- for example, in a
// standard select loop, `case <-actor.Dead(w):` will create
// one new goroutine per iteration, which is... untidy.
func Dead(actor Actor) <-chan struct{} {
	dead := make(chan struct{})
	go func() {
		defer close(dead)
		actor.Wait()
	}()
	return dead
}

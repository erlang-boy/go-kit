package actor

import (
	"github.com/juju/errors"
)

// It should be clear that they don't belong here, and certainly shouldn't be
// used as they are today: e.g. a uniter has *no fricking idea* whether its
// host agent should shut down. A uniter can return ErrUnitDead, and its host
// might need to respond to that, perhaps by returning an error specific to
// *its* host; depending on these values punching right through N layers (but
// only when we want them to!) is kinda terrible.
var (
	//  //Fatal error
	ErrTerminateAgent = errors.New("agent should be terminated")
)

var (
	ErrActorNotFound          = errors.New("requested actor not found")
	ErrActorWatchChanNotFound = errors.New("requested actor's watch channel not found")
)

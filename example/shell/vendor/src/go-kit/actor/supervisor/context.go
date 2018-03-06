package supervisor

import (
	"fmt"

	"github.com/juju/errors"

	. "go-kit/actor"
)

type context struct {
	clientName string
	abort      <-chan struct{}
	expired    chan struct{}
	actors     map[string]Actor
	outputs    map[string]OutputFunc
	accessLog  []resourceAccess
}

func (ctx *context) Abort() <-chan struct{} {
	return ctx.abort
}

func (ctx *context) Get(resourceName string, out interface{}) error {
	logger.Tracef("%q manifold requested %q resource", ctx.clientName, resourceName)
	select {
	case <-ctx.expired:
		return errors.New("expired context: cannot be used outside Start func")
	default:
		err := ctx.rawAccess(resourceName, out)
		ctx.accessLog = append(ctx.accessLog, resourceAccess{
			name: resourceName,
			as:   fmt.Sprintf("%T", out),
			err:  err,
		})
		return err
	}
}

func (ctx *context) expire() {
	close(ctx.expired)
}

func (ctx *context) rawAccess(resourceName string, out interface{}) error {
	input, found := ctx.actors[resourceName]
	if !found {
		return errors.Annotatef(ErrMissing, "%q not declared", resourceName)
	} else if input == nil {
		return errors.Annotatef(ErrMissing, "%q not running", resourceName)
	}
	if out == nil {
		return nil
	}
	convert := ctx.outputs[resourceName]
	if convert == nil {
		return errors.Annotatef(ErrMissing, "%q not exposed", resourceName)
	}
	convert(input, out)
	return nil
}

type resourceAccess struct {
	name string
	as   string
	err  error
}

func (ra resourceAccess) report() map[string]interface{} {
	report := map[string]interface{}{
		KeyName: ra.name,
		KeyType: ra.as,
	}
	if ra.err != nil {
		report[KeyError] = ra.err.Error()
	}
	return report
}

func resourceLogReport(accessLog []resourceAccess) []map[string]interface{} {
	result := make([]map[string]interface{}, len(accessLog))
	for i, access := range accessLog {
		result[i] = access.report()
	}
	return result
}

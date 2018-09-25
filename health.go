package health

import "time"

type Controller struct {
	states map[string]WatchedState // A dynamic state under a name. (eg: states["twilio"] = &twilio.State)
}

func NewController() *Controller {
	m := make(map[string]WatchedState)
	return &Controller{m}
}

func (c *Controller) Register(name string, value *string) {
	c.states[name] = value
}

func (c *Controller) Report() *Report {
	now := time.Now()
	states := make([]State)

	for name, value := range c.states {
		states = append(states, State{
			Value:   *(value.Value),
			Healthy: value.IsHealthy(),
		})
	}

	return &Report{
		At:     now,
		States: states,
	}
}

type WatchedState struct {
	Value    *string
	Expected string
}

func (w *WatchedState) IsHealthy() bool {
	return *(w.Value) == w.Expected
}

type Report struct {
	At     time.Time
	States []State `json:"states"`
}

type State struct {
	Value   string `json:"value"`
	Healthy bool   `json:"healthy"`
}

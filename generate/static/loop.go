package dom

type Loop struct {
	stop chan bool
}

func NewLoop() *Loop {
	return &Loop{
		stop: make(chan bool),
	}
}

func (l *Loop) Stop() { l.stop <- true }

func (l *Loop) Loop() {
	<-l.stop
}

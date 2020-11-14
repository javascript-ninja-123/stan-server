package listener

import (
	"time"

	subject "github.com/javascript-ninja-123/stan-server/subject"
	stan "github.com/nats-io/stan.go"
)

// NewBaseListener is a build function
func NewBaseListener(subject subject.Subject, queueName string, ackWait time.Duration, onMessage func(msg *stan.Msg), sc stan.Conn) Listener {
	return &BaseListener{subject, queueName, ackWait, onMessage, sc}
}

// Listener is an interface
type Listener interface {
	Listen() (stan.Subscription, error)
}

//BaseListener will be exported
type BaseListener struct {
	subject   subject.Subject
	queueName string
	ackWait   time.Duration
	onMessage func(msg *stan.Msg)
	sc        stan.Conn
}

func (b *BaseListener) getSubscriptions() []stan.SubscriptionOption {
	return []stan.SubscriptionOption{
		stan.AckWait(b.ackWait),
		stan.DurableName(b.queueName),
		stan.SetManualAckMode(),
		stan.DeliverAllAvailable(),
	}
}

// Listen be called first
func (b *BaseListener) Listen() (stan.Subscription, error) {
	return b.sc.QueueSubscribe(string(b.subject), b.queueName, b.onMessage, b.getSubscriptions()...)
}

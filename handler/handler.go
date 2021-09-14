package handler

import (
	"time"

	"github.com/crshao/nsq_to_postgres/client"
	"github.com/nsqio/go-nsq"
	"github.com/segmentio/go-stats"
)

// Handler.
type Handler struct {
	stats *stats.Stats
	db    *client.Client
}

// New Handler with the given db client.
func New(db *client.Client) *Handler {
	stats := stats.New()
	go stats.TickEvery(10 * time.Second)
	return &Handler{
		stats: stats,
		db:    db,
	}
}

// HandleMessage inserts the message body into postgres.
func (h *Handler) HandleMessage(msg *nsq.Message) error {
	h.stats.Incr("inserts")
	return h.db.Insert(msg.Body)
}

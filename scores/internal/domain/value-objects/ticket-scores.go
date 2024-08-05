package vo

// TicketScore aggregates ticket scores by category.
type TicketScore struct {
	TicketID   TicketID
	Categories []CategoryScores
}

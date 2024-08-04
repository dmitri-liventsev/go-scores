package design

import (
	. "goa.design/goa/v3/dsl"
)

// Service: Categories
var _ = Service("get by categories", func() {
	Description("Aggregated category scores over a period of time")

	Method("getAggregatedScores", func() {
		Description("Get aggregated scores for categories within a specified period.")
		Payload(func() {
			Field(1, "from", String, "Start date (YYYY-MM-DD)")
			Field(2, "to", String, "End date (YYYY-MM-DD)")
			Required("from", "to")
		})
		Result(func() {
			Field(1, "meta", CategoryMeta)
			Field(2, "data", ArrayOf(CategoryScore))
			Required("meta", "data")
		})
		HTTP(func() {
			GET("/categories/{from}/{to}")
			Params(func() {
				Param("from", String)
				Param("to", String)
			})
			Response(StatusOK)
		})
		GRPC(func() {
			Message(func() {
				Field(1, "from", String)
				Field(2, "to", String)
			})

			Response(CodeOK, func() {
				Message(func() {
					Field(1, "meta", CategoryMeta)
					Field(2, "data", ArrayOf(CategoryScore))
				})
			})
		})
	})
})

// Service: Tickets

var _ = Service("get by tickets", func() {
	Description("Aggregate scores for categories within defined period by ticket")

	Method("getAggregatedScoresByTicket", func() {
		Description("Get aggregate category scores by ticket for a specified period.")
		Payload(func() {
			Field(1, "from", String, "Start date (YYYY-MM-DD)")
			Field(2, "to", String, "End date (YYYY-MM-DD)")
			Required("from", "to")
		})
		Result(func() {
			Field(1, "meta", TicketMeta)
			Field(2, "data", ArrayOf(TicketScore))
			Required("meta", "data")
		})
		HTTP(func() {
			GET("/tickets/{from}/{to}")
			Params(func() {
				Param("from", String)
				Param("to", String)
			})
			Response(StatusOK)
		})
		GRPC(func() {
			Message(func() {
				Field(1, "from", String)
				Field(2, "to", String)
			})
			Response(CodeOK, func() {
				Message(func() {
					Field(1, "meta", TicketMeta)
					Field(2, "data", ArrayOf(TicketScore))
				})
			})
		})
	})
})

// Service: Overall

var _ = Service("get overall", func() {
	Description("Overall aggregate score for a period")

	Method("getOverallScore", func() {
		Description("Get the overall aggregate score for a specified period.")
		Payload(func() {
			Field(1, "from", String, "Start date (YYYY-MM-DD)")
			Field(2, "to", String, "End date (YYYY-MM-DD)")
			Required("from", "to")
		})
		Result(Float32)
		HTTP(func() {
			GET("/overall/{from}/{to}")
			Params(func() {
				Param("from", String)
				Param("to", String)
			})
			Response(StatusOK)
		})
		GRPC(func() {
			Message(func() {
				Field(1, "from", String)
				Field(2, "to", String)
			})
			Response(CodeOK, func() {
				Message(func() {
					Field(1, "result", Float32)
				})
			})
		})
	})
})

// Service: Periods
var _ = Service("get changes by periods", func() {
	Description("Period over period score change")

	Method("getChangesByPeriods", func() {
		Description("Get the score change from a selected period over the previous period.")
		Payload(func() {
			Field(1, "period", String, "The period type (e.g., week, month, year)")
			Required("period")
		})
		Result(func() {
			Field(1, "meta", PeriodsMeta)
			Field(2, "data", ArrayOf(PeriodScoreChange))
			Required("meta", "data")
		})
		HTTP(func() {
			GET("/periods/{period}")
			Params(func() {
				Param("period", String)
			})
			Response(StatusOK)
		})
		GRPC(func() {
			Message(func() {
				Field(1, "period", PeriodType)
			})
			Response(CodeOK, func() {
				Message(func() {
					Field(1, "meta", PeriodsMeta)
					Field(2, "data", ArrayOf(PeriodScoreChange))
				})
			})
		})
	})
})

var PeriodType = Type("PeriodType", func() {
	Description("Period type")
	OneOf("PeriodType", func() {
		Description("The owner's pet")
		Field(1, "week", String)
		Field(2, "month", String)
		Field(3, "year", String)
	})
})

// Define common types used in multiple services
var Period = Type("Period", func() {
	Field(1, "id", Int)
	Field(3, "start", Int64)
	Field(4, "end", Int64)
})

var Category = Type("Category", func() {
	Field(1, "id", Int)
	Field(2, "name", String)
})

var CategoryScore = Type("CategoryScore", func() {
	Field(1, "category_id", Int)
	Field(2, "num_of_reviews", Int)
	Field(3, "periods", ArrayOf(ScorePeriod))
	Field(4, "total_score", Float32)
})

var ScorePeriod = Type("ScorePeriod", func() {
	Field(1, "id", Int)
	Field(2, "score", Float32)
})

var TicketScore = Type("TicketScore", func() {
	Field(1, "ticket_id", Int)
	Field(2, "categories", ArrayOf(CategoryScoreDetail))
})

var CategoryScoreDetail = Type("CategoryScoreDetail", func() {
	Field(1, "id", Int)
	Field(2, "score", Float32)
})

var PeriodScoreChange = Type("PeriodScoreChange", func() {
	Field(1, "period_id", Int)
	Field(2, "score_diff", Float32)
})

// Define a type for CategoryMeta data
var CategoryMeta = Type("CategoryMeta", func() {
	Field(1, "periods", ArrayOf(Period))
	Field(2, "categories", ArrayOf(Category))
})

// Define a type for TicketMeta data
var TicketMeta = Type("TicketMeta", func() {
	Field(1, "categories", ArrayOf(Category))
})

// Define a type for PeriodsMeta data
var PeriodsMeta = Type("PeriodsMeta", func() {
	Field(1, "periods", ArrayOf(Period))
})

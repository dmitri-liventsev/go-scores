package tests

import (
	"context"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	getbyperiods "go-scores/gen/get_changes_by_periods"
	"go-scores/scores/interfaces/controllers"
	pkgtime "go-scores/scores/pkg/time"
)

var _ = Describe("get scores evolution by periods", func() {
	Context("given a week period type", func() {
		var client *getbyperiods.Client
		periodType := pkgtime.PeriodTypeWeek
		numberOfWeeks := getNumberOfweeks()

		BeforeEach(func() {
			endpoint := getbyperiods.NewGetChangesByPeriodsEndpoint(controllers.NewGetByPeriods(DB))
			client = getbyperiods.NewClient(endpoint)
		})

		When("calculating scores", func() {
			var result *getbyperiods.GetChangesByPeriodsResult
			var err error
			BeforeEach(func(ctx context.Context) {
				payload := &getbyperiods.GetChangesByPeriodsPayload{
					Period: periodType,
				}

				result, err = client.GetChangesByPeriods(ctx, payload)
				Expect(err).To(BeNil())
			})

			It("meta should includes all periods", func() {
				Expect(result.Meta.Periods).To(HaveLen(numberOfWeeks))
			})

			It("data should includes all periods", func() {
				Expect(result.Data).To(HaveLen(numberOfWeeks))
			})

			It("data have correct delta", func() {
				tolerance := 0.0001

				Expect(result.Data[0].ScoreDelta).To(HaveValue(BeNumerically("~", float32(100), tolerance)))
				Expect(result.Data[1].ScoreDelta).To(HaveValue(BeNumerically("~", float32(3.3730597496032715), tolerance)))
				Expect(result.Data[2].ScoreDelta).To(HaveValue(BeNumerically("~", float32(-5.912367820739746), tolerance)))
			})
		})
	})
})

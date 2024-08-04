package tests

import (
	"context"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	getbycategories "go-scores/gen/get_by_categories"
	"go-scores/scores/interfaces/controllers"
)

var _ = Describe("get scores by categories", func() {
	Context("given a daily periods and ratings list", func() {
		var client *getbycategories.Client
		var from, to string
		var err error

		BeforeEach(func() {
			from = "2019-08-01"
			to = "2019-08-02"

			endpoint := getbycategories.NewGetAggregatedScoresEndpoint(controllers.NewGetByCategories(DB))
			client = getbycategories.NewClient(endpoint)
		})

		When("calculating scores", func() {
			var categoryScores *getbycategories.GetAggregatedScoresResult
			BeforeEach(func(ctx context.Context) {
				payload := getbycategories.GetAggregatedScoresPayload{
					From: from,
					To:   to,
				}

				categoryScores, err = client.GetAggregatedScores(ctx, &payload)
				Expect(err).NotTo(HaveOccurred())
			})

			It("meta should includes all categories", func() {
				Expect(categoryScores.Meta.Categories).To(HaveLen(4))
				Expect(categoryScores.Meta.Categories).To(ContainElement(HaveField("ID", HaveValue(Equal(1)))))
				Expect(categoryScores.Meta.Categories).To(ContainElement(HaveField("ID", HaveValue(Equal(2)))))
				Expect(categoryScores.Meta.Categories).To(ContainElement(HaveField("ID", HaveValue(Equal(3)))))
				Expect(categoryScores.Meta.Categories).To(ContainElement(HaveField("ID", HaveValue(Equal(4)))))
			})

			It("meta should includes correct periods", func() {
				Expect(categoryScores.Meta.Periods).To(HaveLen(2))

				Expect(categoryScores.Meta.Periods[0].Start).To(HaveValue(Equal(int64(1564617600))))
				Expect(categoryScores.Meta.Periods[0].End).To(HaveValue(Equal(int64(1564617600))))
				Expect(categoryScores.Meta.Periods[1].Start).To(HaveValue(Equal(int64(1564704000))))
				Expect(categoryScores.Meta.Periods[1].End).To(HaveValue(Equal(int64(1564704000))))
			})

			It("data should includes data for each categories", func() {
				Expect(categoryScores.Data).To(HaveLen(4))
				Expect(categoryScores.Data).To(ContainElement(HaveField("CategoryID", HaveValue(Equal(1)))))
				Expect(categoryScores.Data).To(ContainElement(HaveField("CategoryID", HaveValue(Equal(2)))))
				Expect(categoryScores.Data).To(ContainElement(HaveField("CategoryID", HaveValue(Equal(3)))))
				Expect(categoryScores.Data).To(ContainElement(HaveField("CategoryID", HaveValue(Equal(4)))))
			})

			It("each categories should includes all periods", func() {
				Expect(categoryScores.Data[0].Periods).To(HaveLen(2))
				Expect(categoryScores.Data[0].Periods).To(ContainElement(HaveField("ID", HaveValue(Equal(1)))))
				Expect(categoryScores.Data[0].Periods).To(ContainElement(HaveField("ID", HaveValue(Equal(2)))))
				Expect(categoryScores.Data[1].Periods).To(ContainElement(HaveField("ID", HaveValue(Equal(1)))))
				Expect(categoryScores.Data[1].Periods).To(ContainElement(HaveField("ID", HaveValue(Equal(2)))))
				Expect(categoryScores.Data[2].Periods).To(ContainElement(HaveField("ID", HaveValue(Equal(1)))))
				Expect(categoryScores.Data[2].Periods).To(ContainElement(HaveField("ID", HaveValue(Equal(2)))))
				Expect(categoryScores.Data[3].Periods).To(ContainElement(HaveField("ID", HaveValue(Equal(1)))))
				Expect(categoryScores.Data[3].Periods).To(ContainElement(HaveField("ID", HaveValue(Equal(2)))))
			})

			It("scores are correct", func() {
				Expect(categoryScores.Data[0].Periods[0].Score).To(HaveValue(Equal(float32(50.714290618896484))))
				Expect(categoryScores.Data[0].Periods[1].Score).To(HaveValue(Equal(float32(42.9411735534668))))
				Expect(categoryScores.Data[1].Periods[0].Score).To(HaveValue(Equal(float32(43.86249923706055))))
				Expect(categoryScores.Data[1].Periods[1].Score).To(HaveValue(Equal(float32(40.813236236572266))))
				Expect(categoryScores.Data[2].Periods[0].Score).To(HaveValue(Equal(float32(49.71428680419922))))
				Expect(categoryScores.Data[2].Periods[1].Score).To(HaveValue(Equal(float32(60.04706573486328))))
				Expect(categoryScores.Data[3].Periods[0].Score).To(HaveValue(Equal(float32(23.214284896850586))))
				Expect(categoryScores.Data[3].Periods[1].Score).To(HaveValue(Equal(float32(24.705883026123047))))
			})
		})
	})
	Context("given a weekly periods and ratings list", func() {
		var client *getbycategories.Client
		var from, to string
		var err error

		BeforeEach(func() {
			from = "2019-08-01"
			to = "2019-09-01"

			endpoint := getbycategories.NewGetAggregatedScoresEndpoint(controllers.NewGetByCategories(DB))
			client = getbycategories.NewClient(endpoint)
		})

		When("calculating scores", func() {
			var categoryScores *getbycategories.GetAggregatedScoresResult
			BeforeEach(func(ctx context.Context) {
				payload := getbycategories.GetAggregatedScoresPayload{
					From: from,
					To:   to,
				}

				categoryScores, err = client.GetAggregatedScores(ctx, &payload)
				Expect(err).NotTo(HaveOccurred())
			})

			It("meta should includes all categories", func() {
				Expect(categoryScores.Meta.Categories).To(HaveLen(4))
				Expect(categoryScores.Meta.Categories).To(ContainElement(HaveField("ID", HaveValue(Equal(1)))))
				Expect(categoryScores.Meta.Categories).To(ContainElement(HaveField("ID", HaveValue(Equal(2)))))
				Expect(categoryScores.Meta.Categories).To(ContainElement(HaveField("ID", HaveValue(Equal(3)))))
				Expect(categoryScores.Meta.Categories).To(ContainElement(HaveField("ID", HaveValue(Equal(4)))))
			})

			It("meta should includes correct periods", func() {
				Expect(categoryScores.Meta.Periods).To(HaveLen(5))

				Expect(categoryScores.Meta.Periods[0].Start).To(HaveValue(Equal(int64(1564617600))))
				Expect(categoryScores.Meta.Periods[0].End).To(HaveValue(Equal(int64(1564876800))))
				Expect(categoryScores.Meta.Periods[1].Start).To(HaveValue(Equal(int64(1564963200))))
				Expect(categoryScores.Meta.Periods[1].End).To(HaveValue(Equal(int64(1565481600))))
			})

			It("data should includes data for each categories", func() {
				Expect(categoryScores.Data).To(HaveLen(4))
				Expect(categoryScores.Data).To(ContainElement(HaveField("CategoryID", HaveValue(Equal(1)))))
				Expect(categoryScores.Data).To(ContainElement(HaveField("CategoryID", HaveValue(Equal(2)))))
				Expect(categoryScores.Data).To(ContainElement(HaveField("CategoryID", HaveValue(Equal(3)))))
				Expect(categoryScores.Data).To(ContainElement(HaveField("CategoryID", HaveValue(Equal(4)))))
			})

			It("each categories should includes correct periods length", func() {
				Expect(categoryScores.Data[0].Periods).To(HaveLen(5))
			})

			DescribeTable("each categories should includes all periods", func(i int) {
				Expect(categoryScores.Data[i].Periods).To(ContainElement(HaveField("ID", HaveValue(Equal(1)))))
				Expect(categoryScores.Data[i].Periods).To(ContainElement(HaveField("ID", HaveValue(Equal(2)))))
				Expect(categoryScores.Data[i].Periods).To(ContainElement(HaveField("ID", HaveValue(Equal(3)))))
				Expect(categoryScores.Data[i].Periods).To(ContainElement(HaveField("ID", HaveValue(Equal(4)))))
				Expect(categoryScores.Data[i].Periods).To(ContainElement(HaveField("ID", HaveValue(Equal(5)))))
			},
				Entry("Spelling", 0),
				Entry("Grammar", 1),
				Entry("GDPR", 2),
				Entry("Randomness", 3),
			)

			DescribeTable("scores are correct", func(cIdx, pIdx int, expectedValue float32) {
				Expect(categoryScores.Data[cIdx].Periods[pIdx].Score).To(HaveValue(Equal(expectedValue)))
			},
				Entry("Spelling, period 1", 0, 0, float32(50.4504508972168)),
				Entry("Spelling, period 2", 0, 1, float32(49.58580017089844)),
				Entry("Spelling, period 3", 0, 2, float32(52.70000076293945)),
				Entry("Spelling, period 4", 0, 3, float32(46.40776443481445)),
				Entry("Spelling, period 5", 0, 4, float32(46.89655303955078)),
				Entry("Grammar, period 1", 1, 0, float32(38.797298431396484)),
				Entry("Grammar, period 2", 1, 1, float32(39.35591506958008)),
				Entry("Grammar, period 3", 1, 2, float32(40.35350036621094)),
				Entry("Grammar, period 4", 1, 3, float32(36.93276596069336)),
				Entry("Grammar, period 5", 1, 4, float32(40.5428581237793)),
				Entry("GDPR, period 1", 2, 0, float32(52.25225830078125)),
				Entry("GDPR, period 2", 2, 1, float32(57.65680694580078)),
				Entry("GDPR, period 3", 2, 2, float32(54.75199890136719)),
				Entry("GDPR, period 4", 2, 3, float32(51.355342864990234)),
				Entry("GDPR, period 5", 2, 4, float32(57.25714874267578)),
				Entry("Randomness, period 1", 2, 0, float32(52.25225830078125)),
				Entry("Randomness, period 2", 2, 1, float32(57.65680694580078)),
				Entry("Randomness, period 3", 2, 2, float32(54.75199890136719)),
				Entry("Randomness, period 4", 2, 3, float32(51.355342864990234)),
				Entry("Randomness, period 5", 2, 4, float32(57.25714874267578)),
			)
		})
	})
})

// Code generated by goa v3.18.2, DO NOT EDIT.
//
// get overall gRPC client CLI support package
//
// Command:
// $ goa gen go-scores/design

package cli

import (
	"flag"
	"fmt"
	getbycategoriesc "go-scores/gen/grpc/get_by_categories/client"
	getbyperiodsc "go-scores/gen/grpc/get_by_periods/client"
	getbyticketsc "go-scores/gen/grpc/get_by_tickets/client"
	getoverallc "go-scores/gen/grpc/get_overall/client"
	"os"

	goa "goa.design/goa/v3/pkg"
	grpc "google.golang.org/grpc"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `get-overall get-overall-score
get-by-categories get-aggregated-scores
get-by-tickets get-aggregated-scores-by-ticket
get-by-periods get-by-periods
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` get-overall get-overall-score --message '{
      "from": "Labore neque.",
      "to": "Consequatur facere."
   }'` + "\n" +
		os.Args[0] + ` get-by-categories get-aggregated-scores --message '{
      "from": "Rem consequuntur magnam suscipit.",
      "to": "Qui illum accusantium."
   }'` + "\n" +
		os.Args[0] + ` get-by-tickets get-aggregated-scores-by-ticket --message '{
      "from": "Similique cupiditate ut facilis.",
      "to": "Quod vitae amet quos sint nulla."
   }'` + "\n" +
		os.Args[0] + ` get-by-periods get-by-periods --message '{
      "period": "Ab sed recusandae qui ex molestiae consectetur."
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(cc *grpc.ClientConn, opts ...grpc.CallOption) (goa.Endpoint, any, error) {
	var (
		getOverallFlags = flag.NewFlagSet("get-overall", flag.ContinueOnError)

		getOverallGetOverallScoreFlags       = flag.NewFlagSet("get-overall-score", flag.ExitOnError)
		getOverallGetOverallScoreMessageFlag = getOverallGetOverallScoreFlags.String("message", "", "")

		getByCategoriesFlags = flag.NewFlagSet("get-by-categories", flag.ContinueOnError)

		getByCategoriesGetAggregatedScoresFlags       = flag.NewFlagSet("get-aggregated-scores", flag.ExitOnError)
		getByCategoriesGetAggregatedScoresMessageFlag = getByCategoriesGetAggregatedScoresFlags.String("message", "", "")

		getByTicketsFlags = flag.NewFlagSet("get-by-tickets", flag.ContinueOnError)

		getByTicketsGetAggregatedScoresByTicketFlags       = flag.NewFlagSet("get-aggregated-scores-by-ticket", flag.ExitOnError)
		getByTicketsGetAggregatedScoresByTicketMessageFlag = getByTicketsGetAggregatedScoresByTicketFlags.String("message", "", "")

		getByPeriodsFlags = flag.NewFlagSet("get-by-periods", flag.ContinueOnError)

		getByPeriodsGetByPeriodsFlags       = flag.NewFlagSet("get-by-periods", flag.ExitOnError)
		getByPeriodsGetByPeriodsMessageFlag = getByPeriodsGetByPeriodsFlags.String("message", "", "")
	)
	getOverallFlags.Usage = getOverallUsage
	getOverallGetOverallScoreFlags.Usage = getOverallGetOverallScoreUsage

	getByCategoriesFlags.Usage = getByCategoriesUsage
	getByCategoriesGetAggregatedScoresFlags.Usage = getByCategoriesGetAggregatedScoresUsage

	getByTicketsFlags.Usage = getByTicketsUsage
	getByTicketsGetAggregatedScoresByTicketFlags.Usage = getByTicketsGetAggregatedScoresByTicketUsage

	getByPeriodsFlags.Usage = getByPeriodsUsage
	getByPeriodsGetByPeriodsFlags.Usage = getByPeriodsGetByPeriodsUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "get-overall":
			svcf = getOverallFlags
		case "get-by-categories":
			svcf = getByCategoriesFlags
		case "get-by-tickets":
			svcf = getByTicketsFlags
		case "get-by-periods":
			svcf = getByPeriodsFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "get-overall":
			switch epn {
			case "get-overall-score":
				epf = getOverallGetOverallScoreFlags

			}

		case "get-by-categories":
			switch epn {
			case "get-aggregated-scores":
				epf = getByCategoriesGetAggregatedScoresFlags

			}

		case "get-by-tickets":
			switch epn {
			case "get-aggregated-scores-by-ticket":
				epf = getByTicketsGetAggregatedScoresByTicketFlags

			}

		case "get-by-periods":
			switch epn {
			case "get-by-periods":
				epf = getByPeriodsGetByPeriodsFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "get-overall":
			c := getoverallc.NewClient(cc, opts...)
			switch epn {
			case "get-overall-score":
				endpoint = c.GetOverallScore()
				data, err = getoverallc.BuildGetOverallScorePayload(*getOverallGetOverallScoreMessageFlag)
			}
		case "get-by-categories":
			c := getbycategoriesc.NewClient(cc, opts...)
			switch epn {
			case "get-aggregated-scores":
				endpoint = c.GetAggregatedScores()
				data, err = getbycategoriesc.BuildGetAggregatedScoresPayload(*getByCategoriesGetAggregatedScoresMessageFlag)
			}
		case "get-by-tickets":
			c := getbyticketsc.NewClient(cc, opts...)
			switch epn {
			case "get-aggregated-scores-by-ticket":
				endpoint = c.GetAggregatedScoresByTicket()
				data, err = getbyticketsc.BuildGetAggregatedScoresByTicketPayload(*getByTicketsGetAggregatedScoresByTicketMessageFlag)
			}
		case "get-by-periods":
			c := getbyperiodsc.NewClient(cc, opts...)
			switch epn {
			case "get-by-periods":
				endpoint = c.GetByPeriods()
				data, err = getbyperiodsc.BuildGetByPeriodsPayload(*getByPeriodsGetByPeriodsMessageFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
} // get-overallUsage displays the usage of the get-overall command and its
// subcommands.
func getOverallUsage() {
	fmt.Fprintf(os.Stderr, `Overall aggregate score for a period
Usage:
    %[1]s [globalflags] get-overall COMMAND [flags]

COMMAND:
    get-overall-score: Get the overall aggregate score for a specified period.

Additional help:
    %[1]s get-overall COMMAND --help
`, os.Args[0])
}
func getOverallGetOverallScoreUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] get-overall get-overall-score -message JSON

Get the overall aggregate score for a specified period.
    -message JSON: 

Example:
    %[1]s get-overall get-overall-score --message '{
      "from": "Labore neque.",
      "to": "Consequatur facere."
   }'
`, os.Args[0])
}

// get-by-categoriesUsage displays the usage of the get-by-categories command
// and its subcommands.
func getByCategoriesUsage() {
	fmt.Fprintf(os.Stderr, `Aggregated category scores over a period of time
Usage:
    %[1]s [globalflags] get-by-categories COMMAND [flags]

COMMAND:
    get-aggregated-scores: Get aggregated scores for categories within a specified period.

Additional help:
    %[1]s get-by-categories COMMAND --help
`, os.Args[0])
}
func getByCategoriesGetAggregatedScoresUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] get-by-categories get-aggregated-scores -message JSON

Get aggregated scores for categories within a specified period.
    -message JSON: 

Example:
    %[1]s get-by-categories get-aggregated-scores --message '{
      "from": "Rem consequuntur magnam suscipit.",
      "to": "Qui illum accusantium."
   }'
`, os.Args[0])
}

// get-by-ticketsUsage displays the usage of the get-by-tickets command and its
// subcommands.
func getByTicketsUsage() {
	fmt.Fprintf(os.Stderr, `Aggregate scores for categories within defined period by ticket
Usage:
    %[1]s [globalflags] get-by-tickets COMMAND [flags]

COMMAND:
    get-aggregated-scores-by-ticket: Get aggregate category scores by ticket for a specified period.

Additional help:
    %[1]s get-by-tickets COMMAND --help
`, os.Args[0])
}
func getByTicketsGetAggregatedScoresByTicketUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] get-by-tickets get-aggregated-scores-by-ticket -message JSON

Get aggregate category scores by ticket for a specified period.
    -message JSON: 

Example:
    %[1]s get-by-tickets get-aggregated-scores-by-ticket --message '{
      "from": "Similique cupiditate ut facilis.",
      "to": "Quod vitae amet quos sint nulla."
   }'
`, os.Args[0])
}

// get-by-periodsUsage displays the usage of the get-by-periods command and its
// subcommands.
func getByPeriodsUsage() {
	fmt.Fprintf(os.Stderr, `Period over period score change
Usage:
    %[1]s [globalflags] get-by-periods COMMAND [flags]

COMMAND:
    get-by-periods: Get the score change from a selected period over the previous period.

Additional help:
    %[1]s get-by-periods COMMAND --help
`, os.Args[0])
}
func getByPeriodsGetByPeriodsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] get-by-periods get-by-periods -message JSON

Get the score change from a selected period over the previous period.
    -message JSON: 

Example:
    %[1]s get-by-periods get-by-periods --message '{
      "period": "Ab sed recusandae qui ex molestiae consectetur."
   }'
`, os.Args[0])
}

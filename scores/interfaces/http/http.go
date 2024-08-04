package http

import (
	"context"
	getbycategories "go-scores/gen/get_by_categories"
	getbyperiods "go-scores/gen/get_by_periods"
	getbytickets "go-scores/gen/get_by_tickets"
	getoverall "go-scores/gen/get_overall"
	getbycategoriessvr "go-scores/gen/http/get_by_categories/server"
	getbyperiodssvr "go-scores/gen/http/get_by_periods/server"
	getbyticketssvr "go-scores/gen/http/get_by_tickets/server"
	getoverallsvr "go-scores/gen/http/get_overall/server"
	"net/http"
	"net/url"
	"sync"
	"time"

	"goa.design/clue/debug"
	"goa.design/clue/log"
	goahttp "goa.design/goa/v3/http"
)

// handleHTTPServer starts configures and starts a HTTP server on the given
// URL. It shuts down the server if any error is received in the error channel.
func Handle(ctx context.Context, u *url.URL, getOverallEndpoints *getoverall.Endpoints, getByCategoriesEndpoints *getbycategories.Endpoints, getByTicketsEndpoints *getbytickets.Endpoints, getByPeriodsEndpoints *getbyperiods.Endpoints, wg *sync.WaitGroup, errc chan error, dbg bool) {

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/implement/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and mount debug and profiler
	// endpoints in debug mode.
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
		if dbg {
			// Mount pprof handlers for memory profiling under /debug/pprof.
			debug.MountPprofHandlers(debug.Adapt(mux))
			// Mount /debug endpoint to enable or disable debug logs at runtime.
			debug.MountDebugLogEnabler(debug.Adapt(mux))
		}
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		getOverallServer      *getoverallsvr.Server
		getByCategoriesServer *getbycategoriessvr.Server
		getByTicketsServer    *getbyticketssvr.Server
		getByPeriodsServer    *getbyperiodssvr.Server
	)
	{
		eh := errorHandler(ctx)
		getOverallServer = getoverallsvr.New(getOverallEndpoints, mux, dec, enc, eh, nil)
		getByCategoriesServer = getbycategoriessvr.New(getByCategoriesEndpoints, mux, dec, enc, eh, nil)
		getByTicketsServer = getbyticketssvr.New(getByTicketsEndpoints, mux, dec, enc, eh, nil)
		getByPeriodsServer = getbyperiodssvr.New(getByPeriodsEndpoints, mux, dec, enc, eh, nil)
	}

	// Configure the mux.
	getoverallsvr.Mount(mux, getOverallServer)
	getbycategoriessvr.Mount(mux, getByCategoriesServer)
	getbyticketssvr.Mount(mux, getByTicketsServer)
	getbyperiodssvr.Mount(mux, getByPeriodsServer)

	var handler http.Handler = mux
	if dbg {
		// Log query and response bodies if debug logs are enabled.
		handler = debug.HTTP()(handler)
	}
	handler = log.HTTP(ctx)(handler)

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{Addr: u.Host, Handler: handler, ReadHeaderTimeout: time.Second * 60}
	for _, m := range getOverallServer.Mounts {
		log.Printf(ctx, "HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range getByCategoriesServer.Mounts {
		log.Printf(ctx, "HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range getByTicketsServer.Mounts {
		log.Printf(ctx, "HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range getByPeriodsServer.Mounts {
		log.Printf(ctx, "HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start HTTP server in a separate goroutine.
		go func() {
			log.Printf(ctx, "HTTP server listening on %q", u.Host)
			errc <- srv.ListenAndServe()
		}()

		<-ctx.Done()
		log.Printf(ctx, "shutting down HTTP server at %q", u.Host)

		// Shutdown gracefully with a 30s timeout.
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			log.Printf(ctx, "failed to shutdown: %v", err)
		}
	}()
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique PeriodID so that it's possible
// to correlate.
func errorHandler(logCtx context.Context) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		log.Printf(logCtx, "ERROR: %s", err.Error())
	}
}

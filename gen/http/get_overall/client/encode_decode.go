// Code generated by goa v3.18.2, DO NOT EDIT.
//
// get overall HTTP client encoders and decoders
//
// Command:
// $ goa gen go-scores/design

package client

import (
	"bytes"
	"context"
	getoverall "go-scores/gen/get_overall"
	"io"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildGetOverallScoreRequest instantiates a HTTP request object with method
// and path set to call the "get overall" service "getOverallScore" endpoint
func (c *Client) BuildGetOverallScoreRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		from string
		to   string
	)
	{
		p, ok := v.(*getoverall.GetOverallScorePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("get overall", "getOverallScore", "*getoverall.GetOverallScorePayload", v)
		}
		from = p.From
		to = p.To
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetOverallScoreGetOverallPath(from, to)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("get overall", "getOverallScore", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetOverallScoreResponse returns a decoder for responses returned by
// the get overall getOverallScore endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeGetOverallScoreResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body float32
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("get overall", "getOverallScore", err)
			}
			return body, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("get overall", "getOverallScore", resp.StatusCode, string(body))
		}
	}
}

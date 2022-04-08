package endpoints

import (
	"context"
	"errors"
	"os"

	"watermarksvc/internal"
	"watermarksvc/pkg/shorty"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

type Set struct {
	GetEndpoint     endpoint.Endpoint
	ShortenEndpoint endpoint.Endpoint
}

func NewEndpointSet(svc shorty.Service) Set {
	return Set{
		GetEndpoint:     MakeGetEndpoint(svc),
		ShortenEndpoint: MakeShortenEndpoint(svc),
	}
}

func MakeGetEndpoint(svc shorty.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		url, err := svc.Get(ctx, req.Shortname)
		if err != nil {
			return GetResponse{"", err.Error()}, nil
		}
		return GetResponse{url.OriginalURL, ""}, nil
	}
}

func MakeShortenEndpoint(svc shorty.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ShortenRequest)
		url, err := svc.Shorten(ctx, req.OriginalURL)
		if err != nil {
			return ShortenResponse{Shortname: "", Err: err.Error()}, nil
		}
		return ShortenResponse{Shortname: url.Shortname, Err: ""}, nil
	}
}

func (s *Set) Get(ctx context.Context, filters ...internal.Filter) ([]internal.Document, error) {
	resp, err := s.GetEndpoint(ctx, GetRequest{Filters: filters})
	if err != nil {
		return []internal.Document{}, err
	}
	getResp := resp.(GetResponse)
	if getResp.Err != "" {
		return []internal.Document{}, errors.New(getResp.Err)
	}
	return getResp.Documents, nil
}

func (s *Set) ServiceStatus(ctx context.Context) (int, error) {
	resp, err := s.ServiceStatusEndpoint(ctx, ServiceStatusRequest{})
	svcStatusResp := resp.(ServiceStatusResponse)
	if err != nil {
		return svcStatusResp.Code, err
	}
	if svcStatusResp.Err != "" {
		return svcStatusResp.Code, errors.New(svcStatusResp.Err)
	}
	return svcStatusResp.Code, nil
}

func (s *Set) AddDocument(ctx context.Context, doc *internal.Document) (string, error) {
	resp, err := s.AddDocumentEndpoint(ctx, AddDocumentRequest{Document: doc})
	if err != nil {
		return "", err
	}
	adResp := resp.(AddDocumentResponse)
	if adResp.Err != "" {
		return "", errors.New(adResp.Err)
	}
	return adResp.TicketID, nil
}

func (s *Set) Status(ctx context.Context, ticketID string) (internal.Status, error) {
	resp, err := s.StatusEndpoint(ctx, StatusRequest{TicketID: ticketID})
	if err != nil {
		return internal.Failed, err
	}
	stsResp := resp.(StatusResponse)
	if stsResp.Err != "" {
		return internal.Failed, errors.New(stsResp.Err)
	}
	return stsResp.Status, nil
}

func (s *Set) Watermark(ctx context.Context, ticketID, mark string) (int, error) {
	resp, err := s.WatermarkEndpoint(ctx, WatermarkRequest{TicketID: ticketID, Mark: mark})
	wmResp := resp.(WatermarkResponse)
	if err != nil {
		return wmResp.Code, err
	}
	if wmResp.Err != "" {
		return wmResp.Code, errors.New(wmResp.Err)
	}
	return wmResp.Code, nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}

package transport

import (
	"context"

	"watermarksvc/api/v1/pb/shorty"
	"watermarksvc/internal"
	"watermarksvc/pkg/shorty/endpoints"

	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	get     grpctransport.Handler
	shorten grpctransport.Handler
}

func NewGRPCServer(ep endpoints.Set) shorty.ShortenServer {
	return &grpcServer{
		get: grpctransport.NewServer(
			ep.GetEndpoint,
			decodeGRPCGetRequest,
			decodeGRPCGetResponse,
		),
		shorten: grpctransport.NewServer(
			ep.ShortenEndpoint,
			decodeGRPCGetRequest,
			decodeGRPCGetResponse,
		),
	}
}

func (g *grpcServer) Get(ctx context.Context, r *endpoints.GetRequest) (*endpoints.GetResponse, error) {
	_, rep, err := g.get.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*endpoints.GetResponse), nil
}

func (g *grpcServer) Shorten(ctx context.Context, r *endpoints.ShortenRequest) (*endpoints.ShortenResponse, error) {
	_, rep, err := g.get.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*endpoints.ShortenResponse), nil
}

func decodeGRPCGetRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*shorty.GetRequest)
	var filters []internal.Filter
	for _, f := range req.Filters {
		filters = append(filters, internal.Filter{Key: f.Key, Value: f.Value})
	}
	return endpoints.GetRequest{Filters: filters}, nil
}

func decodeGRPCStatusRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*shorty.StatusRequest)
	return endpoints.StatusRequest{TicketID: req.TicketID}, nil
}

func decodeGRPCWatermarkRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*shorty.WatermarkRequest)
	return endpoints.WatermarkRequest{TicketID: req.TicketID, Mark: req.Mark}, nil
}

func decodeGRPCAddDocumentRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*shorty.AddDocumentRequest)
	doc := &internal.Document{
		Content:   req.Document.Content,
		Title:     req.Document.Title,
		Author:    req.Document.Author,
		Topic:     req.Document.Topic,
		Watermark: req.Document.Watermark,
	}
	return endpoints.AddDocumentRequest{Document: doc}, nil
}

func decodeGRPCServiceStatusRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return endpoints.ServiceStatusRequest{}, nil
}

func decodeGRPCGetResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	reply := grpcResponse.(*shorty.GetResponse)
	var docs []internal.Document
	for _, d := range reply.Documents {
		doc := internal.Document{
			Content:   d.Content,
			Title:     d.Title,
			Author:    d.Author,
			Topic:     d.Topic,
			Watermark: d.Watermark,
		}
		docs = append(docs, doc)
	}
	return endpoints.GetResponse{Documents: docs, Err: reply.Err}, nil
}

func decodeGRPCStatusResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	reply := grpcResponse.(*shorty.StatusResponse)
	return endpoints.StatusResponse{Status: internal.Status(reply.Status), Err: reply.Err}, nil
}

func decodeGRPCWatermarkResponse(ctx context.Context, grpcResponse interface{}) (interface{}, error) {
	reply := grpcResponse.(*shorty.WatermarkResponse)
	return endpoints.WatermarkResponse{Code: int(reply.Code), Err: reply.Err}, nil
}

func decodeGRPCAddDocumentResponse(ctx context.Context, grpcResponse interface{}) (interface{}, error) {
	reply := grpcResponse.(*shorty.AddDocumentResponse)
	return endpoints.AddDocumentResponse{TicketID: reply.TicketID, Err: reply.Err}, nil
}

func decodeGRPCServiceStatusResponse(ctx context.Context, grpcResponse interface{}) (interface{}, error) {
	reply := grpcResponse.(*shorty.ServiceStatusResponse)
	return endpoints.ServiceStatusResponse{Code: int(reply.Code), Err: reply.Err}, nil
}

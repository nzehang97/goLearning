// Code generated by Kitex v0.1.4. DO NOT EDIT.

package noteservice

import (
	"context"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/kitex_gen/notedemo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateNote(ctx context.Context, req *notedemo.CreateNoteRequest, callOptions ...callopt.Option) (r *notedemo.CreateNoteResponse, err error)
	MGetNote(ctx context.Context, req *notedemo.MGetNoteRequest, callOptions ...callopt.Option) (r *notedemo.MGetNoteResponse, err error)
	DeleteNote(ctx context.Context, req *notedemo.DeleteNoteRequest, callOptions ...callopt.Option) (r *notedemo.DeleteNoteResponse, err error)
	QueryNote(ctx context.Context, req *notedemo.QueryNoteRequest, callOptions ...callopt.Option) (r *notedemo.QueryNoteResponse, err error)
	UpdateNote(ctx context.Context, req *notedemo.UpdateNoteRequest, callOptions ...callopt.Option) (r *notedemo.UpdateNoteResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kNoteServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kNoteServiceClient struct {
	*kClient
}

func (p *kNoteServiceClient) CreateNote(ctx context.Context, req *notedemo.CreateNoteRequest, callOptions ...callopt.Option) (r *notedemo.CreateNoteResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateNote(ctx, req)
}

func (p *kNoteServiceClient) MGetNote(ctx context.Context, req *notedemo.MGetNoteRequest, callOptions ...callopt.Option) (r *notedemo.MGetNoteResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MGetNote(ctx, req)
}

func (p *kNoteServiceClient) DeleteNote(ctx context.Context, req *notedemo.DeleteNoteRequest, callOptions ...callopt.Option) (r *notedemo.DeleteNoteResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteNote(ctx, req)
}

func (p *kNoteServiceClient) QueryNote(ctx context.Context, req *notedemo.QueryNoteRequest, callOptions ...callopt.Option) (r *notedemo.QueryNoteResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryNote(ctx, req)
}

func (p *kNoteServiceClient) UpdateNote(ctx context.Context, req *notedemo.UpdateNoteRequest, callOptions ...callopt.Option) (r *notedemo.UpdateNoteResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateNote(ctx, req)
}

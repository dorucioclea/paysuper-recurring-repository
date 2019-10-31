// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: repository/repository.proto

/*
Package repository is a generated protocol buffer package.

It is generated from these files:
	repository/repository.proto

It has these top-level messages:
	Result
	FindByStringValue
	SavedCardRequest
	SavedCardList
	DeleteSavedCardRequest
	DeleteSavedCardResponse
*/
package repository

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import entity "github.com/paysuper/paysuper-recurring-repository/pkg/proto/entity"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = entity.SavedCard{}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Repository service

type RepositoryService interface {
	InsertSavedCard(ctx context.Context, in *SavedCardRequest, opts ...client.CallOption) (*Result, error)
	DeleteSavedCard(ctx context.Context, in *DeleteSavedCardRequest, opts ...client.CallOption) (*DeleteSavedCardResponse, error)
	FindSavedCards(ctx context.Context, in *SavedCardRequest, opts ...client.CallOption) (*SavedCardList, error)
	FindSavedCardById(ctx context.Context, in *FindByStringValue, opts ...client.CallOption) (*entity.SavedCard, error)
}

type repositoryService struct {
	c    client.Client
	name string
}

func NewRepositoryService(name string, c client.Client) RepositoryService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "repository"
	}
	return &repositoryService{
		c:    c,
		name: name,
	}
}

func (c *repositoryService) InsertSavedCard(ctx context.Context, in *SavedCardRequest, opts ...client.CallOption) (*Result, error) {
	req := c.c.NewRequest(c.name, "Repository.InsertSavedCard", in)
	out := new(Result)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryService) DeleteSavedCard(ctx context.Context, in *DeleteSavedCardRequest, opts ...client.CallOption) (*DeleteSavedCardResponse, error) {
	req := c.c.NewRequest(c.name, "Repository.DeleteSavedCard", in)
	out := new(DeleteSavedCardResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryService) FindSavedCards(ctx context.Context, in *SavedCardRequest, opts ...client.CallOption) (*SavedCardList, error) {
	req := c.c.NewRequest(c.name, "Repository.FindSavedCards", in)
	out := new(SavedCardList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryService) FindSavedCardById(ctx context.Context, in *FindByStringValue, opts ...client.CallOption) (*entity.SavedCard, error) {
	req := c.c.NewRequest(c.name, "Repository.FindSavedCardById", in)
	out := new(entity.SavedCard)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Repository service

type RepositoryHandler interface {
	InsertSavedCard(context.Context, *SavedCardRequest, *Result) error
	DeleteSavedCard(context.Context, *DeleteSavedCardRequest, *DeleteSavedCardResponse) error
	FindSavedCards(context.Context, *SavedCardRequest, *SavedCardList) error
	FindSavedCardById(context.Context, *FindByStringValue, *entity.SavedCard) error
}

func RegisterRepositoryHandler(s server.Server, hdlr RepositoryHandler, opts ...server.HandlerOption) error {
	type repository interface {
		InsertSavedCard(ctx context.Context, in *SavedCardRequest, out *Result) error
		DeleteSavedCard(ctx context.Context, in *DeleteSavedCardRequest, out *DeleteSavedCardResponse) error
		FindSavedCards(ctx context.Context, in *SavedCardRequest, out *SavedCardList) error
		FindSavedCardById(ctx context.Context, in *FindByStringValue, out *entity.SavedCard) error
	}
	type Repository struct {
		repository
	}
	h := &repositoryHandler{hdlr}
	return s.Handle(s.NewHandler(&Repository{h}, opts...))
}

type repositoryHandler struct {
	RepositoryHandler
}

func (h *repositoryHandler) InsertSavedCard(ctx context.Context, in *SavedCardRequest, out *Result) error {
	return h.RepositoryHandler.InsertSavedCard(ctx, in, out)
}

func (h *repositoryHandler) DeleteSavedCard(ctx context.Context, in *DeleteSavedCardRequest, out *DeleteSavedCardResponse) error {
	return h.RepositoryHandler.DeleteSavedCard(ctx, in, out)
}

func (h *repositoryHandler) FindSavedCards(ctx context.Context, in *SavedCardRequest, out *SavedCardList) error {
	return h.RepositoryHandler.FindSavedCards(ctx, in, out)
}

func (h *repositoryHandler) FindSavedCardById(ctx context.Context, in *FindByStringValue, out *entity.SavedCard) error {
	return h.RepositoryHandler.FindSavedCardById(ctx, in, out)
}

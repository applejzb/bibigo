// Code generated by protoc-gen-liverpc v0.1, DO NOT EDIT.
// source: v2/AppIndex.proto

/*
Package v2 is a generated liverpc stub package.
This code was generated with go-common/app/tool/liverpc/protoc-gen-liverpc v0.1.

It is generated from these files:
	v2/AppIndex.proto
	v2/Room.proto
*/
package v2

import context "context"

import proto "github.com/golang/protobuf/proto"
import "go-common/library/net/rpc/liverpc"

var _ proto.Message // generate to suppress unused imports
// Imports only used by utility functions:

// ==================
// AppIndex Interface
// ==================

type AppIndexRPCClient interface {
	// * 全部列表数据
	//
	GetAllList(ctx context.Context, req *AppIndexGetAllListReq, opts ...liverpc.CallOption) (resp *AppIndexGetAllListResp, err error)

	// * 全部模块基础信息 5.35网关层调用
	//
	GetBaseMInfoList(ctx context.Context, req *AppIndexGetBaseMInfoListReq, opts ...liverpc.CallOption) (resp *AppIndexGetBaseMInfoListResp, err error)

	// * 根据moduleId查common房间列表(for app-interface 5.35+), 包括运营分区、一级分区、推荐（通用RoomList类的）
	//
	GetRoomListByIds(ctx context.Context, req *AppIndexGetRoomListByIdsReq, opts ...liverpc.CallOption) (resp *AppIndexGetRoomListByIdsResp, err error)

	// * 根据moduleId查common房间列表(for app-interface 5.35+), 包括分区入口（通用picList类的）
	//
	GetPicListByIds(ctx context.Context, req *AppIndexGetPicListByIdsReq, opts ...liverpc.CallOption) (resp *AppIndexGetPicListByIdsResp, err error)

	// * 首页banner
	//
	GetIndexBanner(ctx context.Context, req *AppIndexGetIndexBannerReq, opts ...liverpc.CallOption) (resp *AppIndexGetIndexBannerResp, err error)

	// * 全部列表原始数据 for app-interface 5.33+
	//
	GetAllRawList(ctx context.Context, req *AppIndexGetAllRawListReq, opts ...liverpc.CallOption) (resp *AppIndexGetAllRawListResp, err error)

	// * 5.32+获取多个分区房间列表-对推荐第一刷去重
	//
	GetMultiRoomList(ctx context.Context, req *AppIndexGetMultiRoomListReq, opts ...liverpc.CallOption) (resp *AppIndexGetMultiRoomListResp, err error)

	// * 获取活动信息
	// go网关层调
	GetActivityCard(ctx context.Context, req *AppIndexGetActivityCardReq, opts ...liverpc.CallOption) (resp *AppIndexGetActivityCardResp, err error)
}

// ========================
// AppIndex Live Rpc Client
// ========================

type appIndexRPCClient struct {
	client *liverpc.Client
}

// NewAppIndexRPCClient creates a client that implements the AppIndexRPCClient interface.
func NewAppIndexRPCClient(client *liverpc.Client) AppIndexRPCClient {
	return &appIndexRPCClient{
		client: client,
	}
}

func (c *appIndexRPCClient) GetAllList(ctx context.Context, in *AppIndexGetAllListReq, opts ...liverpc.CallOption) (*AppIndexGetAllListResp, error) {
	out := new(AppIndexGetAllListResp)
	err := doRPCRequest(ctx, c.client, 2, "AppIndex.getAllList", in, out, opts)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appIndexRPCClient) GetBaseMInfoList(ctx context.Context, in *AppIndexGetBaseMInfoListReq, opts ...liverpc.CallOption) (*AppIndexGetBaseMInfoListResp, error) {
	out := new(AppIndexGetBaseMInfoListResp)
	err := doRPCRequest(ctx, c.client, 2, "AppIndex.getBaseMInfoList", in, out, opts)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appIndexRPCClient) GetRoomListByIds(ctx context.Context, in *AppIndexGetRoomListByIdsReq, opts ...liverpc.CallOption) (*AppIndexGetRoomListByIdsResp, error) {
	out := new(AppIndexGetRoomListByIdsResp)
	err := doRPCRequest(ctx, c.client, 2, "AppIndex.getRoomListByIds", in, out, opts)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appIndexRPCClient) GetPicListByIds(ctx context.Context, in *AppIndexGetPicListByIdsReq, opts ...liverpc.CallOption) (*AppIndexGetPicListByIdsResp, error) {
	out := new(AppIndexGetPicListByIdsResp)
	err := doRPCRequest(ctx, c.client, 2, "AppIndex.getPicListByIds", in, out, opts)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appIndexRPCClient) GetIndexBanner(ctx context.Context, in *AppIndexGetIndexBannerReq, opts ...liverpc.CallOption) (*AppIndexGetIndexBannerResp, error) {
	out := new(AppIndexGetIndexBannerResp)
	err := doRPCRequest(ctx, c.client, 2, "AppIndex.getIndexBanner", in, out, opts)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appIndexRPCClient) GetAllRawList(ctx context.Context, in *AppIndexGetAllRawListReq, opts ...liverpc.CallOption) (*AppIndexGetAllRawListResp, error) {
	out := new(AppIndexGetAllRawListResp)
	err := doRPCRequest(ctx, c.client, 2, "AppIndex.getAllRawList", in, out, opts)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appIndexRPCClient) GetMultiRoomList(ctx context.Context, in *AppIndexGetMultiRoomListReq, opts ...liverpc.CallOption) (*AppIndexGetMultiRoomListResp, error) {
	out := new(AppIndexGetMultiRoomListResp)
	err := doRPCRequest(ctx, c.client, 2, "AppIndex.getMultiRoomList", in, out, opts)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appIndexRPCClient) GetActivityCard(ctx context.Context, in *AppIndexGetActivityCardReq, opts ...liverpc.CallOption) (*AppIndexGetActivityCardResp, error) {
	out := new(AppIndexGetActivityCardResp)
	err := doRPCRequest(ctx, c.client, 2, "AppIndex.getActivityCard", in, out, opts)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// =====
// Utils
// =====

func doRPCRequest(ctx context.Context, client *liverpc.Client, version int, method string, in, out proto.Message, opts []liverpc.CallOption) (err error) {
	err = client.Call(ctx, version, method, in, out, opts...)
	return
}

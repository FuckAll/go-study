package protobufhttp

import "golang.org/x/net/context"

type HelloServiceImpl struct {
}

func (h *HelloServiceImpl) Echo(ctx context.Context, req *StringMessage) (*StringMessage, error) {
	return &StringMessage{Value: "hello:" + req.Value}, nil
}

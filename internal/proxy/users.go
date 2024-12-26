package proxy

import (
	"context"

	"github.com/emilekm/go-prbf2/prism"
	"github.com/emilekm/prism-proxy/prismproxy"
)

func (p *Proxy) UserList(ctx context.Context, _ *prismproxy.Empty) (*prismproxy.UserListResp, error) {
	users, err := p.c.Users.List(ctx)
	if err != nil {
		return nil, err
	}

	return &prismproxy.UserListResp{
		Users: prismUsersToProto(users),
	}, nil
}

func (p *Proxy) AddUser(ctx context.Context, req *prismproxy.AddUserReq) (*prismproxy.UserListResp, error) {
	users, err := p.c.Users.Add(ctx, prism.AddUser{
		Name:     req.Name,
		Password: req.Password,
		Power:    int(req.Power),
	})
	if err != nil {
		return nil, err
	}

	return &prismproxy.UserListResp{
		Users: prismUsersToProto(users),
	}, nil
}

func (p *Proxy) ChangeUser(ctx context.Context, req *prismproxy.ChangeUserReq) (*prismproxy.UserListResp, error) {
	users, err := p.c.Users.Change(ctx, prism.ChangeUser{
		Name:        req.Name,
		NewName:     req.NewName,
		NewPassword: req.NewPassword,
		NewPower:    int(req.NewPower),
	})
	if err != nil {
		return nil, err
	}

	return &prismproxy.UserListResp{
		Users: prismUsersToProto(users),
	}, nil
}

func (p *Proxy) DeleteUser(ctx context.Context, req *prismproxy.DeleteUserReq) (*prismproxy.UserListResp, error) {
	users, err := p.c.Users.Delete(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	return &prismproxy.UserListResp{
		Users: prismUsersToProto(users),
	}, nil
}

func prismUsersToProto(users []prism.User) []*prismproxy.User {
	var res []*prismproxy.User
	for _, u := range users {
		res = append(res, &prismproxy.User{
			Name:  u.Name,
			Power: int32(u.Power),
		})
	}
	return res
}

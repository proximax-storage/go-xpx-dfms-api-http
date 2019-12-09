package apihttp

import (
	"context"
	"encoding/json"
	"io"

	api "github.com/proximax-storage/go-xpx-dfms-api"
	drive "github.com/proximax-storage/go-xpx-dfms-drive"
)

type inviteSub struct {
	ctx    context.Context
	cancel context.CancelFunc

	msgs   chan *inviteSubMsg
	stream io.ReadCloser
}

func newInviteSub(ctx context.Context, stream io.ReadCloser) api.InviteSubscription {
	ctx, cancel := context.WithCancel(ctx)
	sub := &inviteSub{
		ctx:    ctx,
		cancel: cancel,
		msgs:   make(chan *inviteSubMsg),
		stream: stream,
	}

	go sub.handle()
	return sub
}

func (sub *inviteSub) Next(ctx context.Context) (*drive.Invite, error) {
	select {
	case msg, ok := <-sub.msgs:
		if !ok {
			return nil, io.EOF
		}

		return msg.resp.Invite, msg.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (sub *inviteSub) Close() error {
	sub.cancel()
	return sub.stream.Close()
}

func (sub *inviteSub) handle() {
	defer close(sub.msgs)

	dec := json.NewDecoder(sub.stream)
	for {
		msg := &inviteSubMsg{
			resp: &inviteResponse{
				Invite: &drive.Invite{},
			},
		}

		err := dec.Decode(msg.resp)
		if err != nil {
			if err == io.EOF {
				return
			}

			msg.err = err
		}

		select {
		case sub.msgs <- msg:
		case <-sub.ctx.Done():
			return
		}
	}
}

type inviteSubMsg struct {
	resp *inviteResponse
	err  error
}

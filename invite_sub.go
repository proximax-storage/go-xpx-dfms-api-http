package apihttp

import (
	"context"
	"encoding/json"
	"io"

	drive "github.com/proximax-storage/go-xpx-dfms-drive"
)

type inviteSub struct {
	ctx    context.Context
	cancel context.CancelFunc

	msgs   chan *inviteSubMsg
	stream io.ReadCloser
}

func newInviteSub(ctx context.Context, stream io.ReadCloser) drive.InviteSubscription {
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

func (sub *inviteSub) Next(ctx context.Context) (drive.Invite, error) {
	select {
	case msg, ok := <-sub.msgs:
		if !ok {
			return drive.NilInvite, io.EOF
		}

		return *msg.invite, msg.err
	case <-ctx.Done():
		return drive.NilInvite, ctx.Err()
	}
}

func (sub *inviteSub) Cancel() {
	panic("implement me")
}

func (sub *inviteSub) handle() {
	defer sub.stream.Close()

	dec := json.NewDecoder(sub.stream)
	for {
		msg := &inviteSubMsg{
			invite: new(drive.Invite),
		}

		err := dec.Decode(msg.invite)
		if err != nil {
			if err == io.EOF {
				close(sub.msgs)
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
	invite *drive.Invite
	err    error
}

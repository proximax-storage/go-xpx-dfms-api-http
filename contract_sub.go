package apihttp

import (
	"context"
	"encoding/json"
	"io"

	drive "github.com/proximax-storage/go-xpx-dfms-drive"
)

type contractSub struct {
	ctx    context.Context
	cancel context.CancelFunc

	msgs   chan *contractSubMsg
	stream io.ReadCloser
}

func newContractSub(ctx context.Context, stream io.ReadCloser) drive.ContractSubscription {
	ctx, cancel := context.WithCancel(ctx)
	sub := &contractSub{
		ctx:    ctx,
		cancel: cancel,
		msgs:   make(chan *contractSubMsg),
		stream: stream,
	}

	go sub.handle()
	return sub
}

func (sub *contractSub) Next(ctx context.Context) (drive.Contract, error) {
	select {
	case msg, ok := <-sub.msgs:
		if !ok {
			return nil, io.EOF
		}

		return msg.ctr, msg.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// TODO Return error
func (sub *contractSub) Cancel() {
	sub.cancel()
}

func (sub *contractSub) handle() {
	defer sub.stream.Close()

	dec := json.NewDecoder(sub.stream)
	for {
		msg := &contractSubMsg{
			ctr: new(drive.BasicContract),
		}

		err := dec.Decode(msg.ctr)
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

type contractSubMsg struct {
	ctr drive.Contract
	err error
}

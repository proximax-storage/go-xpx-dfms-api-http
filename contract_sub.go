package apihttp

import (
	"context"
	"encoding/json"
	"io"

	api "github.com/proximax-storage/go-xpx-dfms-api"
	drive "github.com/proximax-storage/go-xpx-dfms-drive"
)

type contractSub struct {
	ctx    context.Context
	cancel context.CancelFunc

	msgs   chan *contractSubMsg
	stream io.ReadCloser
}

func newContractSub(ctx context.Context, stream io.ReadCloser) api.ContractSubscription {
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

		return msg.resp.Contract, msg.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// TODO Return error
func (sub *contractSub) Close() error {
	sub.cancel()
	return sub.stream.Close()
}

func (sub *contractSub) handle() {
	defer close(sub.msgs)

	dec := json.NewDecoder(sub.stream)
	for {
		msg := &contractSubMsg{
			resp: new(contractResponse),
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

type contractSubMsg struct {
	resp *contractResponse
	err  error
}

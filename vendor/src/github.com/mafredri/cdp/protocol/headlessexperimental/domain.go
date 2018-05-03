// Code generated by cdpgen. DO NOT EDIT.

// Package headlessexperimental implements the HeadlessExperimental domain. This domain provides experimental commands only supported in headless mode.
package headlessexperimental

import (
	"context"

	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// domainClient is a client for the HeadlessExperimental domain. This domain provides experimental commands only supported in headless mode.
type domainClient struct{ conn *rpcc.Conn }

// NewClient returns a client for the HeadlessExperimental domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// BeginFrame invokes the HeadlessExperimental method. Sends a BeginFrame to the target and returns when the frame was completed. Optionally captures a screenshot from the resulting frame. Requires that the target was created with enabled BeginFrameControl.
func (d *domainClient) BeginFrame(ctx context.Context, args *BeginFrameArgs) (reply *BeginFrameReply, err error) {
	reply = new(BeginFrameReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "HeadlessExperimental.beginFrame", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "HeadlessExperimental.beginFrame", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "HeadlessExperimental", Op: "BeginFrame", Err: err}
	}
	return
}

// Disable invokes the HeadlessExperimental method. Disables headless events for the target.
func (d *domainClient) Disable(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "HeadlessExperimental.disable", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "HeadlessExperimental", Op: "Disable", Err: err}
	}
	return
}

// Enable invokes the HeadlessExperimental method. Enables headless events for the target.
func (d *domainClient) Enable(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "HeadlessExperimental.enable", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "HeadlessExperimental", Op: "Enable", Err: err}
	}
	return
}

func (d *domainClient) MainFrameReadyForScreenshots(ctx context.Context) (MainFrameReadyForScreenshotsClient, error) {
	s, err := rpcc.NewStream(ctx, "HeadlessExperimental.mainFrameReadyForScreenshots", d.conn)
	if err != nil {
		return nil, err
	}
	return &mainFrameReadyForScreenshotsClient{Stream: s}, nil
}

type mainFrameReadyForScreenshotsClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *mainFrameReadyForScreenshotsClient) GetStream() rpcc.Stream { return c.Stream }

func (c *mainFrameReadyForScreenshotsClient) Recv() (*MainFrameReadyForScreenshotsReply, error) {
	event := new(MainFrameReadyForScreenshotsReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "HeadlessExperimental", Op: "MainFrameReadyForScreenshots Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) NeedsBeginFramesChanged(ctx context.Context) (NeedsBeginFramesChangedClient, error) {
	s, err := rpcc.NewStream(ctx, "HeadlessExperimental.needsBeginFramesChanged", d.conn)
	if err != nil {
		return nil, err
	}
	return &needsBeginFramesChangedClient{Stream: s}, nil
}

type needsBeginFramesChangedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *needsBeginFramesChangedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *needsBeginFramesChangedClient) Recv() (*NeedsBeginFramesChangedReply, error) {
	event := new(NeedsBeginFramesChangedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "HeadlessExperimental", Op: "NeedsBeginFramesChanged Recv", Err: err}
	}
	return event, nil
}

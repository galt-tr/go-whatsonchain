package whatsonchain

import (
	"log"
	"testing"

	"github.com/centrifugal/centrifuge-go"
	"github.com/stretchr/testify/assert"
)

type testHandler struct {
	HandlerName string
}

func (h *testHandler) OnPublish(subscription *centrifuge.Subscription, event centrifuge.PublishEvent) {
	//TODO implement me
	panic("implement me")
}

func (h *testHandler) OnJoin(subscription *centrifuge.Subscription, event centrifuge.JoinEvent) {
	//TODO implement me
	panic("implement me")
}

func (h *testHandler) OnLeave(subscription *centrifuge.Subscription, event centrifuge.LeaveEvent) {
	//TODO implement me
	panic("implement me")
}

func (h *testHandler) OnSubscribeSuccess(subscription *centrifuge.Subscription, event centrifuge.SubscribeSuccessEvent) {
	//TODO implement me
	panic("implement me")
}

func (h *testHandler) OnSubscribeError(subscription *centrifuge.Subscription, event centrifuge.SubscribeErrorEvent) {
	//TODO implement me
	panic("implement me")
}

func (h *testHandler) OnUnsubscribe(subscription *centrifuge.Subscription, event centrifuge.UnsubscribeEvent) {
	//TODO implement me
	panic("implement me")
}

func (h *testHandler) OnConnect(_ *centrifuge.Client, _ centrifuge.ConnectEvent) {
	log.Printf("connected to socket")
}
func (h *testHandler) OnDisconnect(_ *centrifuge.Client, _ centrifuge.DisconnectEvent) {
	log.Printf("disconnected from socket")
}
func (h *testHandler) OnMessage(_ *centrifuge.Client, _ centrifuge.MessageEvent) {
	log.Printf("received message")
}
func (h *testHandler) OnServerPublish(_ *centrifuge.Client, _ centrifuge.ServerPublishEvent) {
	log.Printf("received server publish event")
}
func (h *testHandler) OnError(_ *centrifuge.Client, _ centrifuge.ErrorEvent) {
	log.Printf("error")
}
func (h *testHandler) OnServerJoin(_ *centrifuge.Client, _ centrifuge.ServerJoinEvent) {
	log.Printf("server join event")
}
func (h *testHandler) OnServerLeave(_ *centrifuge.Client, _ centrifuge.ServerLeaveEvent) {
	log.Printf("server leave event")
}
func (h *testHandler) OnServerSubscribe(_ *centrifuge.Client, _ centrifuge.ServerSubscribeEvent) {
	log.Printf("server subscribe event")
}
func (h *testHandler) OnServerUnsubscribe(_ *centrifuge.Client, _ centrifuge.ServerUnsubscribeEvent) {
	log.Printf("server unsubscribe event")
}

func TestClient_NewMempoolWebsocket(t *testing.T) {
	type args struct {
		handler SocketHandler
	}
	tests := []struct {
		name string
		args args
		want *centrifuge.Client
	}{
		{"nil handler should set nil client",
			args{
				handler: nil,
			},
			nil,
		},
	}
	client := newMockClient(&mockHTTPAddresses{})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, client.NewMempoolWebsocket(tt.args.handler), "NewMempoolWebsocket(%v)", tt.args.handler)
		})
	}
}

func Test_newWebsocketClient(t *testing.T) {
	type args struct {
		url     string
		handler SocketHandler
	}
	tests := []struct {
		name         string
		args         args
		nilClient    bool
		connectError bool
	}{
		{
			name: "empty url should set nil client",
			args: args{
				url:     "",
				handler: nil,
			},
			nilClient:    true,
			connectError: false,
		},
		{
			name: "nil handler should set nil client",
			args: args{
				url:     "wss://socket.whatsonchain.com/mempool",
				handler: nil,
			},
			nilClient:    true,
			connectError: false,
		},
		{
			name: "valid handler should successfully connect and disconnect",
			args: args{
				url: "wss://socket.whatsonchain.com/mempool",
				handler: &testHandler{
					HandlerName: "test handler",
				},
			},
			connectError: false,
		},
		{
			name: "valid handler with malformed url should not successfully connect",
			args: args{
				url: "wss//socket.whatsonchain.com/mempool",
				handler: &testHandler{
					HandlerName: "test handler",
				},
			},
			connectError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newWebsocketClient(tt.args.url, tt.args.handler)
			if c == nil && tt.nilClient {
				return
			}
			if c == nil {
				t.Fatalf("expected a valid client, got %v", c)
			}
			err := c.Connect()
			if err != nil && !tt.connectError {
				t.Fatalf("unexpectedly failed to connect to websocket: %v", err)
			}
			err = c.Disconnect()
			if err != nil {
				t.Fatalf("failed to disconnect from websocket: %v", err)
			}

		})
	}
}

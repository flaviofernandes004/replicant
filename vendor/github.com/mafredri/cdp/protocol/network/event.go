// Code generated by cdpgen. DO NOT EDIT.

package network

import (
	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// DataReceivedClient is a client for DataReceived events. Fired when data
// chunk was received over the network.
type DataReceivedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*DataReceivedReply, error)
	rpcc.Stream
}

// DataReceivedReply is the reply for DataReceived events.
type DataReceivedReply struct {
	RequestID         RequestID     `json:"requestId"`         // Request identifier.
	Timestamp         MonotonicTime `json:"timestamp"`         // Timestamp.
	DataLength        int           `json:"dataLength"`        // Data chunk length.
	EncodedDataLength int           `json:"encodedDataLength"` // Actual bytes received (might be less than dataLength for compressed encodings).
}

// EventSourceMessageReceivedClient is a client for EventSourceMessageReceived events.
// Fired when EventSource message is received.
type EventSourceMessageReceivedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*EventSourceMessageReceivedReply, error)
	rpcc.Stream
}

// EventSourceMessageReceivedReply is the reply for EventSourceMessageReceived events.
type EventSourceMessageReceivedReply struct {
	RequestID RequestID     `json:"requestId"` // Request identifier.
	Timestamp MonotonicTime `json:"timestamp"` // Timestamp.
	EventName string        `json:"eventName"` // Message type.
	EventID   string        `json:"eventId"`   // Message identifier.
	Data      string        `json:"data"`      // Message content.
}

// LoadingFailedClient is a client for LoadingFailed events. Fired when HTTP
// request has failed to load.
type LoadingFailedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*LoadingFailedReply, error)
	rpcc.Stream
}

// LoadingFailedReply is the reply for LoadingFailed events.
type LoadingFailedReply struct {
	RequestID     RequestID     `json:"requestId"`               // Request identifier.
	Timestamp     MonotonicTime `json:"timestamp"`               // Timestamp.
	Type          ResourceType  `json:"type"`                    // Resource type.
	ErrorText     string        `json:"errorText"`               // User friendly error message.
	Canceled      *bool         `json:"canceled,omitempty"`      // True if loading was canceled.
	BlockedReason BlockedReason `json:"blockedReason,omitempty"` // The reason why loading was blocked, if any.
}

// LoadingFinishedClient is a client for LoadingFinished events. Fired when
// HTTP request has finished loading.
type LoadingFinishedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*LoadingFinishedReply, error)
	rpcc.Stream
}

// LoadingFinishedReply is the reply for LoadingFinished events.
type LoadingFinishedReply struct {
	RequestID                RequestID     `json:"requestId"`                          // Request identifier.
	Timestamp                MonotonicTime `json:"timestamp"`                          // Timestamp.
	EncodedDataLength        float64       `json:"encodedDataLength"`                  // Total number of bytes received for this request.
	ShouldReportCorbBlocking *bool         `json:"shouldReportCorbBlocking,omitempty"` // Set when 1) response was blocked by Cross-Origin Read Blocking and also 2) this needs to be reported to the DevTools console.
}

// RequestInterceptedClient is a client for RequestIntercepted events. Details
// of an intercepted HTTP request, which must be either allowed, blocked,
// modified or mocked. Deprecated, use Fetch.requestPaused instead.
type RequestInterceptedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*RequestInterceptedReply, error)
	rpcc.Stream
}

// RequestInterceptedReply is the reply for RequestIntercepted events.
type RequestInterceptedReply struct {
	InterceptionID      InterceptionID       `json:"interceptionId"`                // Each request the page makes will have a unique id, however if any redirects are encountered while processing that fetch, they will be reported with the same id as the original fetch. Likewise if HTTP authentication is needed then the same fetch id will be used.
	Request             Request              `json:"request"`                       // No description.
	FrameID             internal.PageFrameID `json:"frameId"`                       // The id of the frame that initiated the request.
	ResourceType        ResourceType         `json:"resourceType"`                  // How the requested resource will be used.
	IsNavigationRequest bool                 `json:"isNavigationRequest"`           // Whether this is a navigation request, which can abort the navigation completely.
	IsDownload          *bool                `json:"isDownload,omitempty"`          // Set if the request is a navigation that will result in a download. Only present after response is received from the server (i.e. HeadersReceived stage).
	RedirectURL         *string              `json:"redirectUrl,omitempty"`         // Redirect location, only sent if a redirect was intercepted.
	AuthChallenge       *AuthChallenge       `json:"authChallenge,omitempty"`       // Details of the Authorization Challenge encountered. If this is set then continueInterceptedRequest must contain an authChallengeResponse.
	ResponseErrorReason ErrorReason          `json:"responseErrorReason,omitempty"` // Response error if intercepted at response stage or if redirect occurred while intercepting request.
	ResponseStatusCode  *int                 `json:"responseStatusCode,omitempty"`  // Response code if intercepted at response stage or if redirect occurred while intercepting request or auth retry occurred.
	ResponseHeaders     Headers              `json:"responseHeaders,omitempty"`     // Response headers if intercepted at the response stage or if redirect occurred while intercepting request or auth retry occurred.
	RequestID           *RequestID           `json:"requestId,omitempty"`           // If the intercepted request had a corresponding requestWillBeSent event fired for it, then this requestId will be the same as the requestId present in the requestWillBeSent event.
}

// RequestServedFromCacheClient is a client for RequestServedFromCache events.
// Fired if request ended up loading from cache.
type RequestServedFromCacheClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*RequestServedFromCacheReply, error)
	rpcc.Stream
}

// RequestServedFromCacheReply is the reply for RequestServedFromCache events.
type RequestServedFromCacheReply struct {
	RequestID RequestID `json:"requestId"` // Request identifier.
}

// RequestWillBeSentClient is a client for RequestWillBeSent events. Fired
// when page is about to send HTTP request.
type RequestWillBeSentClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*RequestWillBeSentReply, error)
	rpcc.Stream
}

// RequestWillBeSentReply is the reply for RequestWillBeSent events.
type RequestWillBeSentReply struct {
	RequestID        RequestID             `json:"requestId"`                  // Request identifier.
	LoaderID         LoaderID              `json:"loaderId"`                   // Loader identifier. Empty string if the request is fetched from worker.
	DocumentURL      string                `json:"documentURL"`                // URL of the document this request is loaded for.
	Request          Request               `json:"request"`                    // Request data.
	Timestamp        MonotonicTime         `json:"timestamp"`                  // Timestamp.
	WallTime         TimeSinceEpoch        `json:"wallTime"`                   // Timestamp.
	Initiator        Initiator             `json:"initiator"`                  // Request initiator.
	RedirectResponse *Response             `json:"redirectResponse,omitempty"` // Redirect response data.
	Type             ResourceType          `json:"type,omitempty"`             // Type of this resource.
	FrameID          *internal.PageFrameID `json:"frameId,omitempty"`          // Frame identifier.
	HasUserGesture   *bool                 `json:"hasUserGesture,omitempty"`   // Whether the request is initiated by a user gesture. Defaults to false.
}

// ResourceChangedPriorityClient is a client for ResourceChangedPriority events.
// Fired when resource loading priority is changed
type ResourceChangedPriorityClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ResourceChangedPriorityReply, error)
	rpcc.Stream
}

// ResourceChangedPriorityReply is the reply for ResourceChangedPriority events.
type ResourceChangedPriorityReply struct {
	RequestID   RequestID        `json:"requestId"`   // Request identifier.
	NewPriority ResourcePriority `json:"newPriority"` // New priority
	Timestamp   MonotonicTime    `json:"timestamp"`   // Timestamp.
}

// SignedExchangeReceivedClient is a client for SignedExchangeReceived events.
// Fired when a signed exchange was received over the network
type SignedExchangeReceivedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*SignedExchangeReceivedReply, error)
	rpcc.Stream
}

// SignedExchangeReceivedReply is the reply for SignedExchangeReceived events.
type SignedExchangeReceivedReply struct {
	RequestID RequestID          `json:"requestId"` // Request identifier.
	Info      SignedExchangeInfo `json:"info"`      // Information about the signed exchange response.
}

// ResponseReceivedClient is a client for ResponseReceived events. Fired when
// HTTP response is available.
type ResponseReceivedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ResponseReceivedReply, error)
	rpcc.Stream
}

// ResponseReceivedReply is the reply for ResponseReceived events.
type ResponseReceivedReply struct {
	RequestID RequestID             `json:"requestId"`         // Request identifier.
	LoaderID  LoaderID              `json:"loaderId"`          // Loader identifier. Empty string if the request is fetched from worker.
	Timestamp MonotonicTime         `json:"timestamp"`         // Timestamp.
	Type      ResourceType          `json:"type"`              // Resource type.
	Response  Response              `json:"response"`          // Response data.
	FrameID   *internal.PageFrameID `json:"frameId,omitempty"` // Frame identifier.
}

// WebSocketClosedClient is a client for WebSocketClosed events. Fired when
// WebSocket is closed.
type WebSocketClosedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*WebSocketClosedReply, error)
	rpcc.Stream
}

// WebSocketClosedReply is the reply for WebSocketClosed events.
type WebSocketClosedReply struct {
	RequestID RequestID     `json:"requestId"` // Request identifier.
	Timestamp MonotonicTime `json:"timestamp"` // Timestamp.
}

// WebSocketCreatedClient is a client for WebSocketCreated events. Fired upon
// WebSocket creation.
type WebSocketCreatedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*WebSocketCreatedReply, error)
	rpcc.Stream
}

// WebSocketCreatedReply is the reply for WebSocketCreated events.
type WebSocketCreatedReply struct {
	RequestID RequestID  `json:"requestId"`           // Request identifier.
	URL       string     `json:"url"`                 // WebSocket request URL.
	Initiator *Initiator `json:"initiator,omitempty"` // Request initiator.
}

// WebSocketFrameErrorClient is a client for WebSocketFrameError events. Fired
// when WebSocket message error occurs.
type WebSocketFrameErrorClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*WebSocketFrameErrorReply, error)
	rpcc.Stream
}

// WebSocketFrameErrorReply is the reply for WebSocketFrameError events.
type WebSocketFrameErrorReply struct {
	RequestID    RequestID     `json:"requestId"`    // Request identifier.
	Timestamp    MonotonicTime `json:"timestamp"`    // Timestamp.
	ErrorMessage string        `json:"errorMessage"` // WebSocket error message.
}

// WebSocketFrameReceivedClient is a client for WebSocketFrameReceived events.
// Fired when WebSocket message is received.
type WebSocketFrameReceivedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*WebSocketFrameReceivedReply, error)
	rpcc.Stream
}

// WebSocketFrameReceivedReply is the reply for WebSocketFrameReceived events.
type WebSocketFrameReceivedReply struct {
	RequestID RequestID      `json:"requestId"` // Request identifier.
	Timestamp MonotonicTime  `json:"timestamp"` // Timestamp.
	Response  WebSocketFrame `json:"response"`  // WebSocket response data.
}

// WebSocketFrameSentClient is a client for WebSocketFrameSent events. Fired
// when WebSocket message is sent.
type WebSocketFrameSentClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*WebSocketFrameSentReply, error)
	rpcc.Stream
}

// WebSocketFrameSentReply is the reply for WebSocketFrameSent events.
type WebSocketFrameSentReply struct {
	RequestID RequestID      `json:"requestId"` // Request identifier.
	Timestamp MonotonicTime  `json:"timestamp"` // Timestamp.
	Response  WebSocketFrame `json:"response"`  // WebSocket response data.
}

// WebSocketHandshakeResponseReceivedClient is a client for WebSocketHandshakeResponseReceived events.
// Fired when WebSocket handshake response becomes available.
type WebSocketHandshakeResponseReceivedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*WebSocketHandshakeResponseReceivedReply, error)
	rpcc.Stream
}

// WebSocketHandshakeResponseReceivedReply is the reply for WebSocketHandshakeResponseReceived events.
type WebSocketHandshakeResponseReceivedReply struct {
	RequestID RequestID         `json:"requestId"` // Request identifier.
	Timestamp MonotonicTime     `json:"timestamp"` // Timestamp.
	Response  WebSocketResponse `json:"response"`  // WebSocket response data.
}

// WebSocketWillSendHandshakeRequestClient is a client for WebSocketWillSendHandshakeRequest events.
// Fired when WebSocket is about to initiate handshake.
type WebSocketWillSendHandshakeRequestClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*WebSocketWillSendHandshakeRequestReply, error)
	rpcc.Stream
}

// WebSocketWillSendHandshakeRequestReply is the reply for WebSocketWillSendHandshakeRequest events.
type WebSocketWillSendHandshakeRequestReply struct {
	RequestID RequestID        `json:"requestId"` // Request identifier.
	Timestamp MonotonicTime    `json:"timestamp"` // Timestamp.
	WallTime  TimeSinceEpoch   `json:"wallTime"`  // UTC Timestamp.
	Request   WebSocketRequest `json:"request"`   // WebSocket request data.
}

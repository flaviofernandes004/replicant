// Code generated by cdpgen. DO NOT EDIT.

package audits

import (
	"github.com/mafredri/cdp/protocol/network"
)

// GetEncodedResponseArgs represents the arguments for GetEncodedResponse in the Audits domain.
type GetEncodedResponseArgs struct {
	RequestID network.RequestID `json:"requestId"` // Identifier of the network request to get content for.
	// Encoding The encoding to use.
	//
	// Values: "webp", "jpeg", "png".
	Encoding string   `json:"encoding"`
	Quality  *float64 `json:"quality,omitempty"`  // The quality of the encoding (0-1). (defaults to 1)
	SizeOnly *bool    `json:"sizeOnly,omitempty"` // Whether to only return the size information (defaults to false).
}

// NewGetEncodedResponseArgs initializes GetEncodedResponseArgs with the required arguments.
func NewGetEncodedResponseArgs(requestID network.RequestID, encoding string) *GetEncodedResponseArgs {
	args := new(GetEncodedResponseArgs)
	args.RequestID = requestID
	args.Encoding = encoding
	return args
}

// SetQuality sets the Quality optional argument. The quality of the
// encoding (0-1). (defaults to 1)
func (a *GetEncodedResponseArgs) SetQuality(quality float64) *GetEncodedResponseArgs {
	a.Quality = &quality
	return a
}

// SetSizeOnly sets the SizeOnly optional argument. Whether to only
// return the size information (defaults to false).
func (a *GetEncodedResponseArgs) SetSizeOnly(sizeOnly bool) *GetEncodedResponseArgs {
	a.SizeOnly = &sizeOnly
	return a
}

// GetEncodedResponseReply represents the return values for GetEncodedResponse in the Audits domain.
type GetEncodedResponseReply struct {
	Body         *string `json:"body,omitempty"` // The encoded body as a base64 string. Omitted if sizeOnly is true.
	OriginalSize int     `json:"originalSize"`   // Size before re-encoding.
	EncodedSize  int     `json:"encodedSize"`    // Size after re-encoding.
}

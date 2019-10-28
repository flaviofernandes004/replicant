// Code generated by cdpgen. DO NOT EDIT.

package applicationcache

import (
	"github.com/mafredri/cdp/protocol/page"
)

// GetApplicationCacheForFrameArgs represents the arguments for GetApplicationCacheForFrame in the ApplicationCache domain.
type GetApplicationCacheForFrameArgs struct {
	FrameID page.FrameID `json:"frameId"` // Identifier of the frame containing document whose application cache is retrieved.
}

// NewGetApplicationCacheForFrameArgs initializes GetApplicationCacheForFrameArgs with the required arguments.
func NewGetApplicationCacheForFrameArgs(frameID page.FrameID) *GetApplicationCacheForFrameArgs {
	args := new(GetApplicationCacheForFrameArgs)
	args.FrameID = frameID
	return args
}

// GetApplicationCacheForFrameReply represents the return values for GetApplicationCacheForFrame in the ApplicationCache domain.
type GetApplicationCacheForFrameReply struct {
	ApplicationCache ApplicationCache `json:"applicationCache"` // Relevant application cache data for the document in given frame.
}

// GetFramesWithManifestsReply represents the return values for GetFramesWithManifests in the ApplicationCache domain.
type GetFramesWithManifestsReply struct {
	FrameIDs []FrameWithManifest `json:"frameIds"` // Array of frame identifiers with manifest urls for each frame containing a document associated with some application cache.
}

// GetManifestForFrameArgs represents the arguments for GetManifestForFrame in the ApplicationCache domain.
type GetManifestForFrameArgs struct {
	FrameID page.FrameID `json:"frameId"` // Identifier of the frame containing document whose manifest is retrieved.
}

// NewGetManifestForFrameArgs initializes GetManifestForFrameArgs with the required arguments.
func NewGetManifestForFrameArgs(frameID page.FrameID) *GetManifestForFrameArgs {
	args := new(GetManifestForFrameArgs)
	args.FrameID = frameID
	return args
}

// GetManifestForFrameReply represents the return values for GetManifestForFrame in the ApplicationCache domain.
type GetManifestForFrameReply struct {
	ManifestURL string `json:"manifestURL"` // Manifest URL for document in the given frame.
}

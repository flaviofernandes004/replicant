// Code generated by cdpgen. DO NOT EDIT.

package domsnapshot

// GetSnapshotArgs represents the arguments for GetSnapshot in the DOMSnapshot domain.
type GetSnapshotArgs struct {
	ComputedStyleWhitelist     []string `json:"computedStyleWhitelist"`               // Whitelist of computed styles to return.
	IncludeEventListeners      *bool    `json:"includeEventListeners,omitempty"`      // Whether or not to retrieve details of DOM listeners (default false).
	IncludePaintOrder          *bool    `json:"includePaintOrder,omitempty"`          // Whether to determine and include the paint order index of LayoutTreeNodes (default false).
	IncludeUserAgentShadowTree *bool    `json:"includeUserAgentShadowTree,omitempty"` // Whether to include UA shadow tree in the snapshot (default false).
}

// NewGetSnapshotArgs initializes GetSnapshotArgs with the required arguments.
func NewGetSnapshotArgs(computedStyleWhitelist []string) *GetSnapshotArgs {
	args := new(GetSnapshotArgs)
	args.ComputedStyleWhitelist = computedStyleWhitelist
	return args
}

// SetIncludeEventListeners sets the IncludeEventListeners optional argument.
// Whether or not to retrieve details of DOM listeners (default false).
func (a *GetSnapshotArgs) SetIncludeEventListeners(includeEventListeners bool) *GetSnapshotArgs {
	a.IncludeEventListeners = &includeEventListeners
	return a
}

// SetIncludePaintOrder sets the IncludePaintOrder optional argument.
// Whether to determine and include the paint order index of
// LayoutTreeNodes (default false).
func (a *GetSnapshotArgs) SetIncludePaintOrder(includePaintOrder bool) *GetSnapshotArgs {
	a.IncludePaintOrder = &includePaintOrder
	return a
}

// SetIncludeUserAgentShadowTree sets the IncludeUserAgentShadowTree optional argument.
// Whether to include UA shadow tree in the snapshot (default false).
func (a *GetSnapshotArgs) SetIncludeUserAgentShadowTree(includeUserAgentShadowTree bool) *GetSnapshotArgs {
	a.IncludeUserAgentShadowTree = &includeUserAgentShadowTree
	return a
}

// GetSnapshotReply represents the return values for GetSnapshot in the DOMSnapshot domain.
type GetSnapshotReply struct {
	DOMNodes        []DOMNode        `json:"domNodes"`        // The nodes in the DOM tree. The DOMNode at index 0 corresponds to the root document.
	LayoutTreeNodes []LayoutTreeNode `json:"layoutTreeNodes"` // The nodes in the layout tree.
	ComputedStyles  []ComputedStyle  `json:"computedStyles"`  // Whitelisted ComputedStyle properties for each node in the layout tree.
}

// CaptureSnapshotArgs represents the arguments for CaptureSnapshot in the DOMSnapshot domain.
type CaptureSnapshotArgs struct {
	ComputedStyles  []string `json:"computedStyles"`            // Whitelist of computed styles to return.
	IncludeDOMRects *bool    `json:"includeDOMRects,omitempty"` // Whether to include DOM rectangles (offsetRects, clientRects, scrollRects) into the snapshot
}

// NewCaptureSnapshotArgs initializes CaptureSnapshotArgs with the required arguments.
func NewCaptureSnapshotArgs(computedStyles []string) *CaptureSnapshotArgs {
	args := new(CaptureSnapshotArgs)
	args.ComputedStyles = computedStyles
	return args
}

// SetIncludeDOMRects sets the IncludeDOMRects optional argument.
// Whether to include DOM rectangles (offsetRects, clientRects,
// scrollRects) into the snapshot
func (a *CaptureSnapshotArgs) SetIncludeDOMRects(includeDOMRects bool) *CaptureSnapshotArgs {
	a.IncludeDOMRects = &includeDOMRects
	return a
}

// CaptureSnapshotReply represents the return values for CaptureSnapshot in the DOMSnapshot domain.
type CaptureSnapshotReply struct {
	Documents []DocumentSnapshot `json:"documents"` // The nodes in the DOM tree. The DOMNode at index 0 corresponds to the root document.
	Strings   []string           `json:"strings"`   // Shared string table that all string properties refer to with indexes.
}

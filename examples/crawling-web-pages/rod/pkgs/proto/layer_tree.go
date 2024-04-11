// This file is generated by "./lib/proto/generate"

package proto

import "github.com/energye/energy/v2/examples/crawling-web-pages/rod/pkgs/gson"

/*

LayerTree

*/

// LayerTreeLayerID Unique Layer identifier.
type LayerTreeLayerID string

// LayerTreeSnapshotID Unique snapshot identifier.
type LayerTreeSnapshotID string

// LayerTreeScrollRectType enum.
type LayerTreeScrollRectType string

const (
	// LayerTreeScrollRectTypeRepaintsOnScroll enum const.
	LayerTreeScrollRectTypeRepaintsOnScroll LayerTreeScrollRectType = "RepaintsOnScroll"

	// LayerTreeScrollRectTypeTouchEventHandler enum const.
	LayerTreeScrollRectTypeTouchEventHandler LayerTreeScrollRectType = "TouchEventHandler"

	// LayerTreeScrollRectTypeWheelEventHandler enum const.
	LayerTreeScrollRectTypeWheelEventHandler LayerTreeScrollRectType = "WheelEventHandler"
)

// LayerTreeScrollRect Rectangle where scrolling happens on the main thread.
type LayerTreeScrollRect struct {
	// Rectangle itself.
	Rect *DOMRect `json:"rect"`

	// Type Reason for rectangle to force scrolling on the main thread
	Type LayerTreeScrollRectType `json:"type"`
}

// LayerTreeStickyPositionConstraint Sticky position constraints.
type LayerTreeStickyPositionConstraint struct {
	// StickyBoxRect Layout rectangle of the sticky element before being shifted
	StickyBoxRect *DOMRect `json:"stickyBoxRect"`

	// ContainingBlockRect Layout rectangle of the containing block of the sticky element
	ContainingBlockRect *DOMRect `json:"containingBlockRect"`

	// NearestLayerShiftingStickyBox (optional) The nearest sticky layer that shifts the sticky box
	NearestLayerShiftingStickyBox LayerTreeLayerID `json:"nearestLayerShiftingStickyBox,omitempty"`

	// NearestLayerShiftingContainingBlock (optional) The nearest sticky layer that shifts the containing block
	NearestLayerShiftingContainingBlock LayerTreeLayerID `json:"nearestLayerShiftingContainingBlock,omitempty"`
}

// LayerTreePictureTile Serialized fragment of layer picture along with its offset within the layer.
type LayerTreePictureTile struct {
	// X Offset from owning layer left boundary
	X float64 `json:"x"`

	// Y Offset from owning layer top boundary
	Y float64 `json:"y"`

	// Picture Base64-encoded snapshot data.
	Picture []byte `json:"picture"`
}

// LayerTreeLayer Information about a compositing layer.
type LayerTreeLayer struct {
	// LayerID The unique id for this layer.
	LayerID LayerTreeLayerID `json:"layerId"`

	// ParentLayerID (optional) The id of parent (not present for root).
	ParentLayerID LayerTreeLayerID `json:"parentLayerId,omitempty"`

	// BackendNodeID (optional) The backend id for the node associated with this layer.
	BackendNodeID DOMBackendNodeID `json:"backendNodeId,omitempty"`

	// OffsetX Offset from parent layer, X coordinate.
	OffsetX float64 `json:"offsetX"`

	// OffsetY Offset from parent layer, Y coordinate.
	OffsetY float64 `json:"offsetY"`

	// Width Layer width.
	Width float64 `json:"width"`

	// Height Layer height.
	Height float64 `json:"height"`

	// Transform (optional) Transformation matrix for layer, default is identity matrix
	Transform []float64 `json:"transform,omitempty"`

	// AnchorX (optional) Transform anchor point X, absent if no transform specified
	AnchorX *float64 `json:"anchorX,omitempty"`

	// AnchorY (optional) Transform anchor point Y, absent if no transform specified
	AnchorY *float64 `json:"anchorY,omitempty"`

	// AnchorZ (optional) Transform anchor point Z, absent if no transform specified
	AnchorZ *float64 `json:"anchorZ,omitempty"`

	// PaintCount Indicates how many time this layer has painted.
	PaintCount int `json:"paintCount"`

	// DrawsContent Indicates whether this layer hosts any content, rather than being used for
	// transform/scrolling purposes only.
	DrawsContent bool `json:"drawsContent"`

	// Invisible (optional) Set if layer is not visible.
	Invisible bool `json:"invisible,omitempty"`

	// ScrollRects (optional) Rectangles scrolling on main thread only.
	ScrollRects []*LayerTreeScrollRect `json:"scrollRects,omitempty"`

	// StickyPositionConstraint (optional) Sticky position constraint information
	StickyPositionConstraint *LayerTreeStickyPositionConstraint `json:"stickyPositionConstraint,omitempty"`
}

// LayerTreePaintProfile Array of timings, one per paint step.
type LayerTreePaintProfile []float64

// LayerTreeCompositingReasons Provides the reasons why the given layer was composited.
type LayerTreeCompositingReasons struct {
	// LayerID The id of the layer for which we want to get the reasons it was composited.
	LayerID LayerTreeLayerID `json:"layerId"`
}

// ProtoReq name.
func (m LayerTreeCompositingReasons) ProtoReq() string { return "LayerTree.compositingReasons" }

// Call the request.
func (m LayerTreeCompositingReasons) Call(c Client) (*LayerTreeCompositingReasonsResult, error) {
	var res LayerTreeCompositingReasonsResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// LayerTreeCompositingReasonsResult ...
type LayerTreeCompositingReasonsResult struct {
	// CompositingReasons A list of strings specifying reasons for the given layer to become composited.
	CompositingReasons []string `json:"compositingReasons"`

	// CompositingReasonIDs A list of strings specifying reason IDs for the given layer to become composited.
	CompositingReasonIDs []string `json:"compositingReasonIds"`
}

// LayerTreeDisable Disables compositing tree inspection.
type LayerTreeDisable struct{}

// ProtoReq name.
func (m LayerTreeDisable) ProtoReq() string { return "LayerTree.disable" }

// Call sends the request.
func (m LayerTreeDisable) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// LayerTreeEnable Enables compositing tree inspection.
type LayerTreeEnable struct{}

// ProtoReq name.
func (m LayerTreeEnable) ProtoReq() string { return "LayerTree.enable" }

// Call sends the request.
func (m LayerTreeEnable) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// LayerTreeLoadSnapshot Returns the snapshot identifier.
type LayerTreeLoadSnapshot struct {
	// Tiles An array of tiles composing the snapshot.
	Tiles []*LayerTreePictureTile `json:"tiles"`
}

// ProtoReq name.
func (m LayerTreeLoadSnapshot) ProtoReq() string { return "LayerTree.loadSnapshot" }

// Call the request.
func (m LayerTreeLoadSnapshot) Call(c Client) (*LayerTreeLoadSnapshotResult, error) {
	var res LayerTreeLoadSnapshotResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// LayerTreeLoadSnapshotResult ...
type LayerTreeLoadSnapshotResult struct {
	// SnapshotID The id of the snapshot.
	SnapshotID LayerTreeSnapshotID `json:"snapshotId"`
}

// LayerTreeMakeSnapshot Returns the layer snapshot identifier.
type LayerTreeMakeSnapshot struct {
	// LayerID The id of the layer.
	LayerID LayerTreeLayerID `json:"layerId"`
}

// ProtoReq name.
func (m LayerTreeMakeSnapshot) ProtoReq() string { return "LayerTree.makeSnapshot" }

// Call the request.
func (m LayerTreeMakeSnapshot) Call(c Client) (*LayerTreeMakeSnapshotResult, error) {
	var res LayerTreeMakeSnapshotResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// LayerTreeMakeSnapshotResult ...
type LayerTreeMakeSnapshotResult struct {
	// SnapshotID The id of the layer snapshot.
	SnapshotID LayerTreeSnapshotID `json:"snapshotId"`
}

// LayerTreeProfileSnapshot ...
type LayerTreeProfileSnapshot struct {
	// SnapshotID The id of the layer snapshot.
	SnapshotID LayerTreeSnapshotID `json:"snapshotId"`

	// MinRepeatCount (optional) The maximum number of times to replay the snapshot (1, if not specified).
	MinRepeatCount *int `json:"minRepeatCount,omitempty"`

	// MinDuration (optional) The minimum duration (in seconds) to replay the snapshot.
	MinDuration *float64 `json:"minDuration,omitempty"`

	// ClipRect (optional) The clip rectangle to apply when replaying the snapshot.
	ClipRect *DOMRect `json:"clipRect,omitempty"`
}

// ProtoReq name.
func (m LayerTreeProfileSnapshot) ProtoReq() string { return "LayerTree.profileSnapshot" }

// Call the request.
func (m LayerTreeProfileSnapshot) Call(c Client) (*LayerTreeProfileSnapshotResult, error) {
	var res LayerTreeProfileSnapshotResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// LayerTreeProfileSnapshotResult ...
type LayerTreeProfileSnapshotResult struct {
	// Timings The array of paint profiles, one per run.
	Timings []LayerTreePaintProfile `json:"timings"`
}

// LayerTreeReleaseSnapshot Releases layer snapshot captured by the back-end.
type LayerTreeReleaseSnapshot struct {
	// SnapshotID The id of the layer snapshot.
	SnapshotID LayerTreeSnapshotID `json:"snapshotId"`
}

// ProtoReq name.
func (m LayerTreeReleaseSnapshot) ProtoReq() string { return "LayerTree.releaseSnapshot" }

// Call sends the request.
func (m LayerTreeReleaseSnapshot) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// LayerTreeReplaySnapshot Replays the layer snapshot and returns the resulting bitmap.
type LayerTreeReplaySnapshot struct {
	// SnapshotID The id of the layer snapshot.
	SnapshotID LayerTreeSnapshotID `json:"snapshotId"`

	// FromStep (optional) The first step to replay from (replay from the very start if not specified).
	FromStep *int `json:"fromStep,omitempty"`

	// ToStep (optional) The last step to replay to (replay till the end if not specified).
	ToStep *int `json:"toStep,omitempty"`

	// Scale (optional) The scale to apply while replaying (defaults to 1).
	Scale *float64 `json:"scale,omitempty"`
}

// ProtoReq name.
func (m LayerTreeReplaySnapshot) ProtoReq() string { return "LayerTree.replaySnapshot" }

// Call the request.
func (m LayerTreeReplaySnapshot) Call(c Client) (*LayerTreeReplaySnapshotResult, error) {
	var res LayerTreeReplaySnapshotResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// LayerTreeReplaySnapshotResult ...
type LayerTreeReplaySnapshotResult struct {
	// DataURL A data: URL for resulting image.
	DataURL string `json:"dataURL"`
}

// LayerTreeSnapshotCommandLog Replays the layer snapshot and returns canvas log.
type LayerTreeSnapshotCommandLog struct {
	// SnapshotID The id of the layer snapshot.
	SnapshotID LayerTreeSnapshotID `json:"snapshotId"`
}

// ProtoReq name.
func (m LayerTreeSnapshotCommandLog) ProtoReq() string { return "LayerTree.snapshotCommandLog" }

// Call the request.
func (m LayerTreeSnapshotCommandLog) Call(c Client) (*LayerTreeSnapshotCommandLogResult, error) {
	var res LayerTreeSnapshotCommandLogResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// LayerTreeSnapshotCommandLogResult ...
type LayerTreeSnapshotCommandLogResult struct {
	// CommandLog The array of canvas function calls.
	CommandLog []map[string]gson.JSON `json:"commandLog"`
}

// LayerTreeLayerPainted ...
type LayerTreeLayerPainted struct {
	// LayerID The id of the painted layer.
	LayerID LayerTreeLayerID `json:"layerId"`

	// Clip rectangle.
	Clip *DOMRect `json:"clip"`
}

// ProtoEvent name.
func (evt LayerTreeLayerPainted) ProtoEvent() string {
	return "LayerTree.layerPainted"
}

// LayerTreeLayerTreeDidChange ...
type LayerTreeLayerTreeDidChange struct {
	// Layers (optional) Layer tree, absent if not in the compositing mode.
	Layers []*LayerTreeLayer `json:"layers,omitempty"`
}

// ProtoEvent name.
func (evt LayerTreeLayerTreeDidChange) ProtoEvent() string {
	return "LayerTree.layerTreeDidChange"
}

package differ

import (
	"github.com/zclconf/go-cty/cty"

	"github.com/hashicorp/terraform/internal/command/jsonformat/computed"

	"github.com/hashicorp/terraform/internal/command/jsonformat/computed/renderers"
)

func (change Change) computeAttributeDiffAsTuple(elementTypes []cty.Type) computed.Diff {
	var elements []computed.Diff
	current := change.getDefaultActionForIteration()
	sliceValue := change.asSlice()
	for ix, elementType := range elementTypes {
		childValue := sliceValue.getChild(ix, ix, false)
		element := childValue.computeDiffForType(elementType)
		elements = append(elements, element)
		current = compareActions(current, element.Action)
	}
	return computed.NewDiff(renderers.List(elements), current, change.replacePath())
}
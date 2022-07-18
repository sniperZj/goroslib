//autogenerated:yes
//nolint:revive,lll
package vision_msgs

import (
	"github.com/aler9/goroslib/pkg/msg"
	"github.com/aler9/goroslib/pkg/msgs/std_msgs"
)

type BoundingBox2DArray struct {
	msg.Package `ros:"vision_msgs"`
	Header      std_msgs.Header
	Boxes       []BoundingBox2D
}

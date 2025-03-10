//autogenerated:yes
//nolint:revive,lll
package mavros_msgs

import (
	"github.com/aler9/goroslib/pkg/msg"
	"github.com/aler9/goroslib/pkg/msgs/std_msgs"
)

type BatteryStatus struct {
	msg.Package `ros:"mavros_msgs"`
	Header      std_msgs.Header
	Voltage     float32
	Current     float32
	Remaining   float32
}

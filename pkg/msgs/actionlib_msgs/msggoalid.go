//autogenerated:yes
//nolint:golint,lll
package actionlib_msgs

import (
	"time"

	"github.com/aler9/goroslib/pkg/msg"
)

type GoalID struct {
	msg.Package `ros:"actionlib_msgs"`
	Stamp       time.Time
	Id          string
}

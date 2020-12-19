package std_msgs //nolint:golint

import (
	"github.com/aler9/goroslib/pkg/msg"
)

type Int16 struct { //nolint:golint
	msg.Package `ros:"std_msgs"`
	Data        int16 //nolint:golint
}

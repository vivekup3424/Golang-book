package task

import (
	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
)

// State represents the status of a task.
type State int

// Enumeration of possible task states.
const (
	Pending   State = iota // Task is pending.
	Scheduled              // Task is scheduled.
	Running                // Task is currently running.
	Completed              // Task has been completed.
	Failed                 // Task has failed.
)

// Note: For the ID of a task, we are using UUID.

// What is a UUID?
// A UUID (Universally Unique Identifier) is a 128-bit number used to uniquely
// identify information in computer systems. The probability of generating two
// identical UUIDs is extremely low, making it an excellent choice for unique
// identifiers. For more details about UUIDs, refer to RFC 4122:
// https://tools.ietf.org/html/rfc4122

type Task struct {
	ID            uuid.UUID          `json:"id"`
	Name          string             `json:"name"`
	State         State              `json:"state"`
	Image         string             `json:"image"`
	Memory        int                `json:"memory"`
	Disk          int                `json:"disk"`
	ExposedPorts  nat.PortSet        `json:"exposedPorts"`
	PortBindings  map[strings]string `json:"portBindings"`
	RestartPolicy string             `json:"restartPolicy"`
}

//we’re going to limit our orchestrator to dealing with Docker
//containers

package state

// Redis state represent the full state of the connection with Redis
type Redis struct {
	Healthy           bool                  // are we able to connect to Redis?
	Ready             bool                  // are we ready to provide state for the reconciler?
	Replication       RedisReplicationState // current replication data
	ReplicationString string                // raw replication info
	Stopped           bool
}

// isRedisMaster return whether the Redis under management currently
// see itself as a master instance or not
func (r *Redis) IsRedisMaster() bool {
	return r.Replication.Role == "master"
}

func (r *Redis) IsUnhealthy() bool {
	return r.Healthy == false
}

type RedisReplicationState struct {
	Role       string // current redis role (master or slave)
	MasterHost string // if slave, the master hostname its replicating from
	MasterPort int    // if slave, the master port its replicating from
}

// changed will test if the current replication state is different from
// the new one passed in as argument
func (r *RedisReplicationState) Changed(new RedisReplicationState) bool {
	if r.Role != new.Role {
		return true
	}

	if r.MasterHost != new.MasterHost {
		return true
	}

	if r.MasterPort != new.MasterPort {
		return true
	}

	return false
}

package v1alpha

type State string
type ClusterState string

const (
	GracefulUpscaleRequested    State        = "GracefulUpscaleRequested"
	GracefulUpscaleSucceeded    State        = "GracefulUpscaleSucceeded"
	GracefulUpscaleFailed       State        = "GracefulUpscaleFailed"
	GracefulUpscaleInProgress   State        = "GracefulUpscaleInProgress"
	GracefulDownscaleRequested  State        = "GracefulDownscaleRequested"
	GracefulDownscaleSucceeded  State        = "GracefulDownscaleSucceeded"
	GracefulDownscaleFailed     State        = "GracefulDownscaleFailed"
	GracefulDownscaleInProgress State        = "GracefulDownscaleInProgress"
	RedisClusterInitializing    ClusterState = "RedisClusterInitializing"
	RedisClusterInitialized     ClusterState = "RedisClusterInitialized"
)

func (s State) IsUpscale() bool {
	return s == GracefulUpscaleRequested || s == GracefulUpscaleSucceeded || s == GracefulUpscaleFailed
}

func (s State) IsDownscale() bool {
	return s == GracefulDownscaleRequested || s == GracefulDownscaleSucceeded || s == GracefulDownscaleFailed
}

func (s State) IsInProgress() bool {
	return s == GracefulUpscaleInProgress || s == GracefulDownscaleInProgress
}

func (s State) IsRequested() bool {
	return s == GracefulUpscaleRequested || s == GracefulDownscaleRequested
}

func (s State) Complete() State {
	switch s {
	case GracefulUpscaleRequested, GracefulUpscaleInProgress:
		return GracefulUpscaleSucceeded
	case GracefulDownscaleRequested, GracefulDownscaleInProgress:
		return GracefulDownscaleSucceeded
	default:
		return s
	}
}

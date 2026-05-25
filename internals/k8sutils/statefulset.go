package k8sutils

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

type StatefulSet interface {
	IsStatefulSetReady(ctx context.Context, namespace string, name string) bool
	GetStatefulSetReplicas(ctx context.Context, namespace string, name string) int32
}

type StatefulSetImplementation struct {
	clientset kubernetes.Interface
}

func NewStatefulSetImplementation(clientset kubernetes.Interface) *StatefulSetImplementation {
	return &StatefulSetImplementation{
		clientset: clientset,
	}
}

func (s *StatefulSetImplementation) IsStatefulSetReady(ctx context.Context, namespace string, name string) bool {
	var (
		partition int32 = 0
		replicas  int32 = 1
	)

	sts, err := s.clientset.AppsV1().StatefulSets(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		log.FromContext(ctx).Error(err, "Failed to get StatefulSet for RedisCluster")
		return false
	}

	if sts.Spec.UpdateStrategy.Type == "RollingUpdate" && sts.Spec.UpdateStrategy.RollingUpdate != nil && sts.Spec.UpdateStrategy.RollingUpdate.Partition != nil {
		partition = int32(*sts.Spec.UpdateStrategy.RollingUpdate.Partition)
	}

	if sts.Spec.Replicas != nil {
		replicas = *sts.Spec.Replicas
	}

	if expectedUpdateReplicas := replicas - partition; sts.Status.UpdatedReplicas < expectedUpdateReplicas {
		log.FromContext(ctx).Info("StatefulSet is not ready", "Status.UpdatedReplicas", sts.Status.UpdatedReplicas, "ExpectedUpdateReplicas", expectedUpdateReplicas)
		return false
	}

	if partition == 0 && sts.Status.CurrentRevision != sts.Status.UpdateRevision {
		log.FromContext(ctx).Info("StatefulSet is not Ready", "Status.CurrentRevision", sts.Status.CurrentRevision, "Status.UpdateRevision", sts.Status.UpdateRevision)
		return false
	}

	if sts.Status.ObservedGeneration < sts.Generation {
		log.FromContext(ctx).Info("StatefulSet is not Ready", "Status.ObservedGeneration", sts.Status.ObservedGeneration, "Generation", sts.Generation)
		return false
	}

	if sts.Status.ReadyReplicas < replicas {
		log.FromContext(ctx).Info("StatefulSet is not Ready", "Status.ReadyReplicas", sts.Status.ReadyReplicas, "CurrentReplicas", replicas)
		return false
	}
	return true
}

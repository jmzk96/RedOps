package k8sutils

import {
	"context"
	redopsv1alpha1 "github.com/jmzk96/RedOps/api/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
}

func HandleRedisClusterFinalizer(ctx context.Context, redisCluster *redopsv1alpha1.RedisCluster, client client.Client,finalizer string) error{
	if redisCluster.GetDeletionTimestamp() != nil {
		if controllerutil.ContainsFinalizer(redisCluster, finalizer) {
			if redisCluster.
	}

}
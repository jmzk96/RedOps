package controller

import (
	"context"
	"go.uber.org/zap"
	"fmt"
	"strings"
	"time"
	"sigs.k8s.io/controller-runtime/pkg/client"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/tools/record"
)


type RedisClusterReconciler struct {
	client.Client
	DirectClient client.Reader
	Log zap.Logger
	Scheme *runtime.Scheme
	Namespaces []string
	Recorder record.EventRecorder
	RequeueIntervals map[string]int
	RequeueOffset int
}


func (r *RedisClusterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	redisCluster := &v1alpha1.RedisCluster{}
	if err := r.Get(ctx, req.NamespacedName, redisCluster); err != nil {
		if apierrors.IsNotFound(err) {
			r.Log.Info("RedisCluster resource not found. Ignoring since object must be deleted.")
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

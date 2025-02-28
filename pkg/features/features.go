package features

import (
	"github.com/golang/glog"

	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/component-base/featuregate"
)

const (
	// Every feature gate should add method here following this template:
	//
	// // owner: @username
	// // alpha:
	// MyFeature featuregate.Feature = "MyFeature"

	// PersistentVolumes owner: @irfanurrehman
	// beta:
	//
	// Persistent volumes support.
	PersistentVolumes featuregate.Feature = "PersistentVolumes"

	// ThrottlingMetrics owner: @irfanurrehman
	// beta:
	//
	// Throttling Metrics support.
	ThrottlingMetrics featuregate.Feature = "ThrottlingMetrics"

	// GitopsApps owner: @irfanurrehman
	// alpha:
	//
	// Gitops application support.
	// This gate will enable discovery of gitops pipeline applications and
	// the action execution based on the same.
	GitopsApps featuregate.Feature = "GitopsApps"

	// HonorRegionZoneLabels owner: @kevinwang
	// alpha:
	//
	// Honor the region/zone labels of the node.
	// This gate will enable honorinig the labels topology.kubernetes.io/region and topology.kubernetes.io/zone
	// of the node which the pod is currently running on and also enable honoring the PV affninity on a pod move
	HonorAzLabelPvAffinity featuregate.Feature = "HonorAzLabelPvAffinity"

	// GoMemLimit (MemoryOptimisations) owner: @mengding @irfanurrehman
	// alpha:
	// This flag enables below optimisations
	//
	// Go runtime soft memory limit support
	// This gate enables Go runtime soft memory limit as explained in
	// https://pkg.go.dev/runtime/debug#SetMemoryLimit
	//
	// Pagination support for list API calls to API server querying workload controllers
	// Without this feature gate the whole list is requested in a single list API call.
	GoMemLimit featuregate.Feature = "GoMemLimit"
	// AllowIncreaseNsQuota4Resizing owner: @kevinwang
	// alpha:
	//
	// This gate will enable the temporary namespace quota increase when
	// kubeturbo execute a resize action on a workload controller
	AllowIncreaseNsQuota4Resizing featuregate.Feature = "AllowIncreaseNsQuota4Resizing"
	// IgnoreAffinities owner: @irfanurrehman
	// alpha:
	//
	// This gate will simply ignore processing of affinities in the cluster.
	// This is to try out temporarily disable processing of affinities to be able to
	// try out a POV in a customer environment, which is held up because of inefficiency
	// in out code, where affinity processing alone takes a really long time.
	// https://vmturbo.atlassian.net/browse/OM-93635?focusedCommentId=771727
	IgnoreAffinities featuregate.Feature = "IgnoreAffinities"
)

func init() {
	if err := utilfeature.DefaultMutableFeatureGate.Add(DefaultKubeturboFeatureGates); err != nil {
		glog.Fatalf("Unexpected error: %v", err)
	}
}

// DefaultKubeturboFeatureGates consists of all known kubeturbo-specific
// feature keys.  To add a new feature, define a key for it above and
// add it here.
// Ref: https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates/
// Note: We use the config to feed the values, not the command line params.
var DefaultKubeturboFeatureGates = map[featuregate.Feature]featuregate.FeatureSpec{
	PersistentVolumes:             {Default: true, PreRelease: featuregate.Beta},
	ThrottlingMetrics:             {Default: true, PreRelease: featuregate.Beta},
	GitopsApps:                    {Default: false, PreRelease: featuregate.Alpha},
	HonorAzLabelPvAffinity:        {Default: true, PreRelease: featuregate.Alpha},
	GoMemLimit:                    {Default: true, PreRelease: featuregate.Alpha},
	AllowIncreaseNsQuota4Resizing: {Default: true, PreRelease: featuregate.Alpha},
	IgnoreAffinities:              {Default: false, PreRelease: featuregate.Alpha},
}

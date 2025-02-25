// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package bgpv1

import (
	"github.com/cilium/cilium/pkg/bgpv1/agent"
	"github.com/cilium/cilium/pkg/bgpv1/agent/signaler"
	"github.com/cilium/cilium/pkg/bgpv1/manager"
	"github.com/cilium/cilium/pkg/hive"
	"github.com/cilium/cilium/pkg/hive/cell"
	"github.com/cilium/cilium/pkg/k8s"
	v2alpha1api "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2alpha1"
	"github.com/cilium/cilium/pkg/k8s/client"
	"github.com/cilium/cilium/pkg/k8s/resource"
	slim_core_v1 "github.com/cilium/cilium/pkg/k8s/slim/k8s/api/core/v1"
	"github.com/cilium/cilium/pkg/k8s/utils"
	"github.com/cilium/cilium/pkg/option"
)

var Cell = cell.Module(
	"bgp-cp",
	"BGP Control Plane",

	// The Controller which is the entry point of the module
	cell.Provide(agent.NewController, signaler.NewBGPCPSignaler),
	cell.ProvidePrivate(
		// Local Node Store Specer provides the module with information about the current node.
		agent.NewNodeSpecer,
		// BGP Peering Policy resource provides the module with a stream of events for the BGPPeeringPolicy resource.
		newBGPPeeringPolicyResource,
		// goBGP is currently the only supported RouterManager, if more are
		// implemented, provide the manager via a Cell that pics implementation based on configuration.
		manager.NewBGPRouterManager,
		// Create a slim service DiffStore
		manager.NewDiffStore[*slim_core_v1.Service],
		// Create a endpoints DiffStore
		manager.NewDiffStore[*k8s.Endpoints],
	),
	// Provides the reconcilers used by the route manager to update the config
	manager.ConfigReconcilers,
)

func newBGPPeeringPolicyResource(lc hive.Lifecycle, c client.Clientset, dc *option.DaemonConfig) resource.Resource[*v2alpha1api.CiliumBGPPeeringPolicy] {
	// Do not create this resource if the BGP Control Plane is disabled
	if !dc.BGPControlPlaneEnabled() {
		return nil
	}

	if !c.IsEnabled() {
		return nil
	}

	return resource.New[*v2alpha1api.CiliumBGPPeeringPolicy](
		lc, utils.ListerWatcherFromTyped[*v2alpha1api.CiliumBGPPeeringPolicyList](
			c.CiliumV2alpha1().CiliumBGPPeeringPolicies(),
		), resource.WithMetric("CiliumBGPPeeringPolicy"))
}

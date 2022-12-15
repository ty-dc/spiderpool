// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0
package annotation_test

import (
	"context"
	"testing"

	spiderpool "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v1"
	"github.com/spidernet-io/spiderpool/test/e2e/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	e2e "github.com/spidernet-io/e2eframework/framework"
	"k8s.io/apimachinery/pkg/runtime"
)

func TestAnnotation(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Annotation Suite")
}

var frame *e2e.Framework
var globalDefaultV4IpoolList, globalDefaultV6IpoolList []string

var _ = BeforeSuite(func() {
	defer GinkgoRecover()
	var e error
	frame, e = e2e.NewFramework(GinkgoT(), []func(*runtime.Scheme) error{spiderpool.AddToScheme})
	Expect(e).NotTo(HaveOccurred())

	if frame.Info.SpiderSubnetEnabled {
		clusterDefaultV4SubnetList, clusterDefaultV6SubnetList, e := common.GetClusterDefaultSubnet(frame)
		Expect(e).NotTo(HaveOccurred())
		ctx, cancel := context.WithTimeout(context.Background(), common.PodStartTimeout)
		defer cancel()
		if frame.Info.IpV4Enabled && len(clusterDefaultV4SubnetList) == 0 {
			Fail("failed to find cluster ipv4 subnet")
		} else {
			globalV4PoolName, v4Pool := common.GenerateExampleIpv4poolObject(1)
			globalDefaultV4IpoolList = append(globalDefaultV4IpoolList, globalV4PoolName)
			e = common.CreateIppoolInSpiderSubnet(ctx, frame, clusterDefaultV4SubnetList[0], v4Pool, 2)
			Expect(e).NotTo(HaveOccurred())
		}
		if frame.Info.IpV6Enabled && len(clusterDefaultV6SubnetList) == 0 {
			Fail("failed to find cluster ipv6 subnet")
		} else {
			globalV6PoolName, v6Pool := common.GenerateExampleIpv6poolObject(1)
			globalDefaultV6IpoolList = append(globalDefaultV6IpoolList, globalV6PoolName)
			e = common.CreateIppoolInSpiderSubnet(ctx, frame, clusterDefaultV6SubnetList[0], v6Pool, 2)
			Expect(e).NotTo(HaveOccurred())
		}
	} else {
		globalDefaultV4IpoolList, globalDefaultV6IpoolList, e = common.GetClusterDefaultIppool(frame)
		Expect(e).NotTo(HaveOccurred())
		if frame.Info.IpV4Enabled && len(globalDefaultV4IpoolList) == 0 {
			Fail("failed to find cluster ipv4 ippool")
		}
		if frame.Info.IpV6Enabled && len(globalDefaultV6IpoolList) == 0 {
			Fail("failed to find cluster ipv6 ippool")
		}
	}
})

var _ = AfterSuite(func() {
	if frame.Info.SpiderSubnetEnabled {
		if frame.Info.IpV4Enabled {
			for _, v := range globalDefaultV4IpoolList {
				Expect(common.DeleteIPPoolByName(frame, v)).NotTo(HaveOccurred())
			}
		}
		if frame.Info.IpV6Enabled {
			for _, v := range globalDefaultV6IpoolList {
				Expect(common.DeleteIPPoolByName(frame, v)).NotTo(HaveOccurred())
			}
		}
	}
})

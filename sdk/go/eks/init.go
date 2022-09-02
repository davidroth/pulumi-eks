// Code generated by pulumi-gen-eks DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eks

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "eks:index:Cluster":
		r = &Cluster{}
	case "eks:index:ClusterCreationRoleProvider":
		r = &ClusterCreationRoleProvider{}
	case "eks:index:ManagedNodeGroup":
		r = &ManagedNodeGroup{}
	case "eks:index:NodeGroup":
		r = &NodeGroup{}
	case "eks:index:NodeGroupSecurityGroup":
		r = &NodeGroupSecurityGroup{}
	case "eks:index:NodeGroupV2":
		r = &NodeGroupV2{}
	case "eks:index:VpcCni":
		r = &VpcCni{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:eks" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, _ := PkgVersion()
	pulumi.RegisterResourceModule(
		"eks",
		"index",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"eks",
		&pkg{version},
	)
}

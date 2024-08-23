package eksnet

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateSubnets(ctx *pulumi.Context, vpcId pulumi.IDOutput) ([]*ec2.Subnet, []*ec2.Subnet, error) {
	config := ctx.Config()

	// Leer configuración de subnets
	createNatSubnet := config.GetBool("createNatSubnet")
	subnetNatCidr := config.Require("subnetNatCidr")
	subnetNatAz := config.Require("subnetNatAz")

	createIgwSubnet := config.GetBool("createIgwSubnet")
	subnetIgwCidr := config.Require("subnetIgwCidr")
	subnetIgwAz := config.Require("subnetIgwAz")

	tags := pulumi.StringMap{
		"Name":           pulumi.String(ctx.Stack()),
		"tech:Component": pulumi.String("subnet"),
		"tech:Layer":     pulumi.String("networking"),
		"iac:ManagedBy":  pulumi.String("Pulumi"),
		"org:Owner":      pulumi.String("DevOps"),
		"org:Project":    pulumi.String(ctx.Stack()),
	}

	var natSubnets []*ec2.Subnet
	var igwSubnets []*ec2.Subnet

	if createNatSubnet {
		natSubnet, err := ec2.NewSubnet(ctx, "natSubnet", &ec2.SubnetArgs{
			VpcId:            vpcId,
			CidrBlock:        pulumi.String(subnetNatCidr),
			AvailabilityZone: pulumi.String(subnetNatAz),
			Tags:             tags,
		})
		if err != nil {
			return nil, nil, err
		}
		natSubnets = append(natSubnets, natSubnet)
	}

	if createIgwSubnet {
		igwSubnet, err := ec2.NewSubnet(ctx, "igwSubnet", &ec2.SubnetArgs{
			VpcId:            vpcId,
			CidrBlock:        pulumi.String(subnetIgwCidr),
			AvailabilityZone: pulumi.String(subnetIgwAz),
			Tags:             tags,
		})
		if err != nil {
			return nil, nil, err
		}
		igwSubnets = append(igwSubnets, igwSubnet)
	}

	return natSubnets, igwSubnets, nil
}

package eksnet

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateNatGateway(ctx *pulumi.Context, subnetId pulumi.IDOutput) (*ec2.NatGateway, error) {
	// Crear Elastic IP para NAT Gateway
	natEip, err := ec2.NewEip(ctx, "natEip", &ec2.EipArgs{
		Vpc: pulumi.Bool(true),
	})
	if err != nil {
		return nil, err
	}

	// Crear NAT Gateway
	natGateway, err := ec2.NewNatGateway(ctx, "natGateway", &ec2.NatGatewayArgs{
		AllocationId: natEip.ID(),
		SubnetId:     subnetId,
		Tags: pulumi.StringMap{
			"Name":           pulumi.String(ctx.Stack()),
			"tech:Component": pulumi.String("nat_gateway"),
			"tech:Layer":     pulumi.String("networking"),
			"iac:ManagedBy":  pulumi.String("Pulumi"),
			"org:Owner":      pulumi.String("DevOps"),
			"org:Project":    pulumi.String(ctx.Stack()),
		},
	})
	if err != nil {
		return nil, err
	}

	return natGateway, nil
}

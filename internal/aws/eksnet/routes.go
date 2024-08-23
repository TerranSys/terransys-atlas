package eksnet

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateRouteTables(ctx *pulumi.Context, vpcId pulumi.IDOutput, natGatewayId pulumi.IDOutput, natSubnets []*ec2.Subnet, igwSubnets []*ec2.Subnet) error {
	// Crear Route Table para subnets privadas
	privateRouteTable, err := ec2.NewRouteTable(ctx, "privateRouteTable", &ec2.RouteTableArgs{
		VpcId: vpcId,
		Routes: ec2.RouteTableRouteArray{
			&ec2.RouteTableRouteArgs{
				CidrBlock:    pulumi.String("0.0.0.0/0"),
				NatGatewayId: natGatewayId,
			},
		},
		Tags: pulumi.StringMap{
			"Name":           pulumi.String(ctx.Stack() + "-private"),
			"tech:Component": pulumi.String("route_table"),
			"tech:Layer":     pulumi.String("networking"),
			"iac:ManagedBy":  pulumi.String("Pulumi"),
			"org:Owner":      pulumi.String("DevOps"),
			"org:Project":    pulumi.String(ctx.Stack()),
		},
	})
	if err != nil {
		return err
	}

	for _, subnet := range natSubnets {
		_, err := ec2.NewRouteTableAssociation(ctx, "privateRouteTableAssociation", &ec2.RouteTableAssociationArgs{
			SubnetId:     subnet.ID(),
			RouteTableId: privateRouteTable.ID(),
		})
		if err != nil {
			return err
		}
	}

	// Crear Route Table para subnets públicas
	publicRouteTable, err := ec2.NewRouteTable(ctx, "publicRouteTable", &ec2.RouteTableArgs{
		VpcId: vpcId,
		Routes: ec2.RouteTableRouteArray{
			&ec2.RouteTableRouteArgs{
				CidrBlock: pulumi.String("0.0.0.0/0"),
				GatewayId: pulumi.String("<internet-gateway-id>"), // Este debe ser el ID del IGW creado previamente
			},
		},
		Tags: pulumi.StringMap{
			"Name":           pulumi.String(ctx.Stack() + "-public"),
			"tech:Component": pulumi.String("route_table"),
			"tech:Layer":     pulumi.String("networking"),
			"iac:ManagedBy":  pulumi.String("Pulumi"),
			"org:Owner":      pulumi.String("DevOps"),
			"org:Project":    pulumi.String(ctx.Stack()),
		},
	})
	if err != nil {
		return err
	}

	for _, subnet := range igwSubnets {
		_, err := ec2.NewRouteTableAssociation(ctx, "publicRouteTableAssociation", &ec2.RouteTableAssociationArgs{
			SubnetId:     subnet.ID(),
			RouteTableId: publicRouteTable.ID(),
		})
		if err != nil {
			return err
		}
	}

	return nil
}

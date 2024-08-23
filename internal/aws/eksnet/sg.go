package eksnet

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateSecurityGroups(ctx *pulumi.Context, vpcId pulumi.IDOutput) (*ec2.SecurityGroup, error) {
	// Crear el Security Group por defecto para la VPC
	defaultSecurityGroup, err := ec2.NewSecurityGroup(ctx, "defaultSecurityGroup", &ec2.SecurityGroupArgs{
		VpcId: vpcId,
		Egress: ec2.SecurityGroupEgressArray{
			&ec2.SecurityGroupEgressArgs{
				Protocol: pulumi.String("-1"),
				FromPort: pulumi.Int(0),
				ToPort:   pulumi.Int(0),
				CidrBlocks: pulumi.StringArray{
					pulumi.String("0.0.0.0/0"),
				},
			},
		},
		Ingress: ec2.SecurityGroupIngressArray{
			&ec2.SecurityGroupIngressArgs{
				Protocol: pulumi.String("-1"),
				FromPort: pulumi.Int(0),
				ToPort:   pulumi.Int(0),
				Self:     pulumi.Bool(true),
			},
		},
		Tags: pulumi.StringMap{
			"Name":           pulumi.String(ctx.Stack() + "-default-sg"),
			"tech:Component": pulumi.String("security_group"),
			"tech:Layer":     pulumi.String("networking"),
			"iac:ManagedBy":  pulumi.String("Pulumi"),
			"org:Owner":      pulumi.String("DevOps"),
			"org:Project":    pulumi.String(ctx.Stack()),
		},
	})
	if err != nil {
		return nil, err
	}

	return defaultSecurityGroup, nil
}

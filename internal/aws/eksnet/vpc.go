package eksnet

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateVPC(ctx *pulumi.Context) (*ec2.Vpc, error) {
	config := ctx.Config()

	// Leer variables de configuración
	mainVpcCidr := config.Require("mainVpcCidr")
	instanceTenancy := config.Get("instanceTenancy")
	if instanceTenancy == "" {
		instanceTenancy = "default"
	}
	enableDnsSupport := config.GetBool("enableDnsSupport")
	enableDnsHostnames := config.GetBool("enableDnsHostnames")

	// Crear la VPC
	vpc, err := ec2.NewVpc(ctx, "mainVpc", &ec2.VpcArgs{
		CidrBlock:          pulumi.String(mainVpcCidr),
		InstanceTenancy:    pulumi.String(instanceTenancy),
		EnableDnsSupport:   pulumi.Bool(enableDnsSupport),
		EnableDnsHostnames: pulumi.Bool(enableDnsHostnames),
		Tags: pulumi.StringMap{
			"Name":           pulumi.String(ctx.Stack()),
			"tech:Component": pulumi.String("VPC"),
			"tech:Layer":     pulumi.String("networking"),
			"iac:ManagedBy":  pulumi.String("Pulumi"),
			"org:Owner":      pulumi.String("DevOps"),
			"org:Project":    pulumi.String(ctx.Stack()),
		},
	})
	if err != nil {
		return nil, err
	}

	return vpc, nil
}

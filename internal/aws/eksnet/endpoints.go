package eksnet

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateVpcEndpoints(ctx *pulumi.Context, vpcId pulumi.IDOutput, subnetIds pulumi.StringArrayOutput) error {
	// Crear VPC Endpoint para S3
	_, err := ec2.NewVpcEndpoint(ctx, "s3VpcEndpoint", &ec2.VpcEndpointArgs{
		VpcId:           vpcId,
		ServiceName:     pulumi.String("com.amazonaws." + ctx.Stack() + ".s3"), // Asegúrate de ajustar esto al región adecuada
		VpcEndpointType: pulumi.String("Gateway"),
		RouteTableIds:   subnetIds, // Asociar con las tablas de rutas de las subnets
		Tags: pulumi.StringMap{
			"Name":           pulumi.String(ctx.Stack() + "-s3-endpoint"),
			"tech:Component": pulumi.String("vpc_endpoint"),
			"tech:Layer":     pulumi.String("networking"),
			"iac:ManagedBy":  pulumi.String("Pulumi"),
			"org:Owner":      pulumi.String("DevOps"),
			"org:Project":    pulumi.String(ctx.Stack()),
		},
	})
	if err != nil {
		return err
	}

	// Puedes añadir más endpoints aquí para otros servicios como DynamoDB, etc.

	return nil
}

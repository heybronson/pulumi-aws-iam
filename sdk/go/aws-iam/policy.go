// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package awsiam

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource helps you create an IAM policy.
//
// ## Example Usage
// ## Policy
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	iam "github.com/pulumi/pulumi-aws-iam/sdk/go/aws-iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//	    pulumi.Run(func(ctx *pulumi.Context) error {
//	        policyJSON, err := json.Marshal(map[string]interface{}{
//	            "Version": "2012-10-17",
//	            "Statement": []interface{}{
//	                map[string]interface{}{
//	                    "Effect":   "Allow",
//	                    "Action":   []string{"ec2:Describe"},
//	                    "Resource": []string{"*"},
//	                },
//	            },
//	        })
//	        if err != nil {
//	            return err
//	        }
//
//	        policy, err := iam.NewPolicy(ctx, "policy", &iam.PolicyArgs{
//	            Name:           pulumi.String("example"),
//	            Path:           pulumi.String("/"),
//	            Description:    pulumi.String("My example policy"),
//	            PolicyDocument: pulumi.String(string(policyJSON)),
//	        })
//	        if err != nil {
//	            return err
//	        }
//
//	        ctx.Export("policy", policy)
//
//	        return nil
//	    })
//	}
//
// ```
// {{ /example }}
type Policy struct {
	pulumi.ResourceState

	// The ARN assigned by AWS to this policy.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description of the policy.
	Description pulumi.StringOutput `pulumi:"description"`
	// The policy's ID.
	Id pulumi.StringOutput `pulumi:"id"`
	// The name of the policy.
	Name pulumi.StringOutput `pulumi:"name"`
	// The path of the policy in IAM.
	Path pulumi.StringOutput `pulumi:"path"`
	// The policy document.
	PolicyDocument pulumi.StringOutput `pulumi:"policyDocument"`
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.PolicyDocument == nil {
		return nil, errors.New("invalid value for required argument 'PolicyDocument'")
	}
	if args.Description == nil {
		args.Description = pulumi.StringPtr("IAM Policy")
	}
	if args.Path == nil {
		args.Path = pulumi.StringPtr("/")
	}
	var resource Policy
	err := ctx.RegisterRemoteComponentResource("aws-iam:index:Policy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type policyArgs struct {
	// The description of the policy.
	Description *string `pulumi:"description"`
	// The name of the policy.
	Name string `pulumi:"name"`
	// The path of the policy in IAM.
	Path *string `pulumi:"path"`
	// The policy document.
	PolicyDocument string `pulumi:"policyDocument"`
	// A map of tags to add.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// The description of the policy.
	Description pulumi.StringPtrInput
	// The name of the policy.
	Name pulumi.StringInput
	// The path of the policy in IAM.
	Path pulumi.StringPtrInput
	// The policy document.
	PolicyDocument pulumi.StringInput
	// A map of tags to add.
	Tags pulumi.StringMapInput
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}

type PolicyInput interface {
	pulumi.Input

	ToPolicyOutput() PolicyOutput
	ToPolicyOutputWithContext(ctx context.Context) PolicyOutput
}

func (*Policy) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil)).Elem()
}

func (i *Policy) ToPolicyOutput() PolicyOutput {
	return i.ToPolicyOutputWithContext(context.Background())
}

func (i *Policy) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyOutput)
}

// PolicyArrayInput is an input type that accepts PolicyArray and PolicyArrayOutput values.
// You can construct a concrete instance of `PolicyArrayInput` via:
//
//	PolicyArray{ PolicyArgs{...} }
type PolicyArrayInput interface {
	pulumi.Input

	ToPolicyArrayOutput() PolicyArrayOutput
	ToPolicyArrayOutputWithContext(context.Context) PolicyArrayOutput
}

type PolicyArray []PolicyInput

func (PolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Policy)(nil)).Elem()
}

func (i PolicyArray) ToPolicyArrayOutput() PolicyArrayOutput {
	return i.ToPolicyArrayOutputWithContext(context.Background())
}

func (i PolicyArray) ToPolicyArrayOutputWithContext(ctx context.Context) PolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyArrayOutput)
}

// PolicyMapInput is an input type that accepts PolicyMap and PolicyMapOutput values.
// You can construct a concrete instance of `PolicyMapInput` via:
//
//	PolicyMap{ "key": PolicyArgs{...} }
type PolicyMapInput interface {
	pulumi.Input

	ToPolicyMapOutput() PolicyMapOutput
	ToPolicyMapOutputWithContext(context.Context) PolicyMapOutput
}

type PolicyMap map[string]PolicyInput

func (PolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Policy)(nil)).Elem()
}

func (i PolicyMap) ToPolicyMapOutput() PolicyMapOutput {
	return i.ToPolicyMapOutputWithContext(context.Background())
}

func (i PolicyMap) ToPolicyMapOutputWithContext(ctx context.Context) PolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyMapOutput)
}

type PolicyOutput struct{ *pulumi.OutputState }

func (PolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil)).Elem()
}

func (o PolicyOutput) ToPolicyOutput() PolicyOutput {
	return o
}

func (o PolicyOutput) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return o
}

// The ARN assigned by AWS to this policy.
func (o PolicyOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The description of the policy.
func (o PolicyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The policy's ID.
func (o PolicyOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.Id }).(pulumi.StringOutput)
}

// The name of the policy.
func (o PolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The path of the policy in IAM.
func (o PolicyOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// The policy document.
func (o PolicyOutput) PolicyDocument() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.PolicyDocument }).(pulumi.StringOutput)
}

type PolicyArrayOutput struct{ *pulumi.OutputState }

func (PolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Policy)(nil)).Elem()
}

func (o PolicyArrayOutput) ToPolicyArrayOutput() PolicyArrayOutput {
	return o
}

func (o PolicyArrayOutput) ToPolicyArrayOutputWithContext(ctx context.Context) PolicyArrayOutput {
	return o
}

func (o PolicyArrayOutput) Index(i pulumi.IntInput) PolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Policy {
		return vs[0].([]*Policy)[vs[1].(int)]
	}).(PolicyOutput)
}

type PolicyMapOutput struct{ *pulumi.OutputState }

func (PolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Policy)(nil)).Elem()
}

func (o PolicyMapOutput) ToPolicyMapOutput() PolicyMapOutput {
	return o
}

func (o PolicyMapOutput) ToPolicyMapOutputWithContext(ctx context.Context) PolicyMapOutput {
	return o
}

func (o PolicyMapOutput) MapIndex(k pulumi.StringInput) PolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Policy {
		return vs[0].(map[string]*Policy)[vs[1].(string)]
	}).(PolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyInput)(nil)).Elem(), &Policy{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyArrayInput)(nil)).Elem(), PolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyMapInput)(nil)).Elem(), PolicyMap{})
	pulumi.RegisterOutputType(PolicyOutput{})
	pulumi.RegisterOutputType(PolicyArrayOutput{})
	pulumi.RegisterOutputType(PolicyMapOutput{})
}

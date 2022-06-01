// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package awsiam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EKSRole struct {
	pulumi.ResourceState

	// ARN of IAM role.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Name of IAM role.
	Name pulumi.StringOutput `pulumi:"name"`
	// Path of IAM role.
	Path pulumi.StringOutput `pulumi:"path"`
	// Unique ID of IAM role.
	UniqueId pulumi.StringOutput `pulumi:"uniqueId"`
}

// NewEKSRole registers a new resource with the given unique name, arguments, and options.
func NewEKSRole(ctx *pulumi.Context,
	name string, args *EKSRoleArgs, opts ...pulumi.ResourceOption) (*EKSRole, error) {
	if args == nil {
		args = &EKSRoleArgs{}
	}

	if isZero(args.ForceDetachPolicies) {
		args.ForceDetachPolicies = pulumi.BoolPtr(false)
	}
	if isZero(args.MaxSessionDuration) {
		args.MaxSessionDuration = pulumi.IntPtr(3600)
	}
	if args.Role != nil {
		args.Role = args.Role.ToRolePtrOutput().ApplyT(func(v *Role) *Role { return v.Defaults() }).(RolePtrOutput)
	}
	var resource EKSRole
	err := ctx.RegisterRemoteComponentResource("aws-iam:index:EKSRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type eksroleArgs struct {
	// EKS cluster and k8s ServiceAccount pairs. Each EKS cluster can have multiple k8s ServiceAccount. See README for details
	ClusterServiceAccounts map[string][]string `pulumi:"clusterServiceAccounts"`
	// Whether policies should be detached from this role when destroying.
	ForceDetachPolicies *bool `pulumi:"forceDetachPolicies"`
	// Maximum CLI/API session duration in seconds between 3600 and 43200.
	MaxSessionDuration *int `pulumi:"maxSessionDuration"`
	// OIDC provider URL and k8s ServiceAccount pairs. If the assume role policy requires a mix of EKS clusters and other OIDC providers then this can be used
	ProviderUrlSaPairs map[string][]string `pulumi:"providerUrlSaPairs"`
	Role               *Role               `pulumi:"role"`
	// ARNs of any policies to attach to the IAM role.
	RolePolicyArns []string `pulumi:"rolePolicyArns"`
	// A map of tags to add.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a EKSRole resource.
type EKSRoleArgs struct {
	// EKS cluster and k8s ServiceAccount pairs. Each EKS cluster can have multiple k8s ServiceAccount. See README for details
	ClusterServiceAccounts pulumi.StringArrayMapInput
	// Whether policies should be detached from this role when destroying.
	ForceDetachPolicies pulumi.BoolPtrInput
	// Maximum CLI/API session duration in seconds between 3600 and 43200.
	MaxSessionDuration pulumi.IntPtrInput
	// OIDC provider URL and k8s ServiceAccount pairs. If the assume role policy requires a mix of EKS clusters and other OIDC providers then this can be used
	ProviderUrlSaPairs pulumi.StringArrayMapInput
	Role               RolePtrInput
	// ARNs of any policies to attach to the IAM role.
	RolePolicyArns pulumi.StringArrayInput
	// A map of tags to add.
	Tags pulumi.StringMapInput
}

func (EKSRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eksroleArgs)(nil)).Elem()
}

type EKSRoleInput interface {
	pulumi.Input

	ToEKSRoleOutput() EKSRoleOutput
	ToEKSRoleOutputWithContext(ctx context.Context) EKSRoleOutput
}

func (*EKSRole) ElementType() reflect.Type {
	return reflect.TypeOf((**EKSRole)(nil)).Elem()
}

func (i *EKSRole) ToEKSRoleOutput() EKSRoleOutput {
	return i.ToEKSRoleOutputWithContext(context.Background())
}

func (i *EKSRole) ToEKSRoleOutputWithContext(ctx context.Context) EKSRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EKSRoleOutput)
}

// EKSRoleArrayInput is an input type that accepts EKSRoleArray and EKSRoleArrayOutput values.
// You can construct a concrete instance of `EKSRoleArrayInput` via:
//
//          EKSRoleArray{ EKSRoleArgs{...} }
type EKSRoleArrayInput interface {
	pulumi.Input

	ToEKSRoleArrayOutput() EKSRoleArrayOutput
	ToEKSRoleArrayOutputWithContext(context.Context) EKSRoleArrayOutput
}

type EKSRoleArray []EKSRoleInput

func (EKSRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EKSRole)(nil)).Elem()
}

func (i EKSRoleArray) ToEKSRoleArrayOutput() EKSRoleArrayOutput {
	return i.ToEKSRoleArrayOutputWithContext(context.Background())
}

func (i EKSRoleArray) ToEKSRoleArrayOutputWithContext(ctx context.Context) EKSRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EKSRoleArrayOutput)
}

// EKSRoleMapInput is an input type that accepts EKSRoleMap and EKSRoleMapOutput values.
// You can construct a concrete instance of `EKSRoleMapInput` via:
//
//          EKSRoleMap{ "key": EKSRoleArgs{...} }
type EKSRoleMapInput interface {
	pulumi.Input

	ToEKSRoleMapOutput() EKSRoleMapOutput
	ToEKSRoleMapOutputWithContext(context.Context) EKSRoleMapOutput
}

type EKSRoleMap map[string]EKSRoleInput

func (EKSRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EKSRole)(nil)).Elem()
}

func (i EKSRoleMap) ToEKSRoleMapOutput() EKSRoleMapOutput {
	return i.ToEKSRoleMapOutputWithContext(context.Background())
}

func (i EKSRoleMap) ToEKSRoleMapOutputWithContext(ctx context.Context) EKSRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EKSRoleMapOutput)
}

type EKSRoleOutput struct{ *pulumi.OutputState }

func (EKSRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EKSRole)(nil)).Elem()
}

func (o EKSRoleOutput) ToEKSRoleOutput() EKSRoleOutput {
	return o
}

func (o EKSRoleOutput) ToEKSRoleOutputWithContext(ctx context.Context) EKSRoleOutput {
	return o
}

// ARN of IAM role.
func (o EKSRoleOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *EKSRole) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Name of IAM role.
func (o EKSRoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EKSRole) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Path of IAM role.
func (o EKSRoleOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *EKSRole) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// Unique ID of IAM role.
func (o EKSRoleOutput) UniqueId() pulumi.StringOutput {
	return o.ApplyT(func(v *EKSRole) pulumi.StringOutput { return v.UniqueId }).(pulumi.StringOutput)
}

type EKSRoleArrayOutput struct{ *pulumi.OutputState }

func (EKSRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EKSRole)(nil)).Elem()
}

func (o EKSRoleArrayOutput) ToEKSRoleArrayOutput() EKSRoleArrayOutput {
	return o
}

func (o EKSRoleArrayOutput) ToEKSRoleArrayOutputWithContext(ctx context.Context) EKSRoleArrayOutput {
	return o
}

func (o EKSRoleArrayOutput) Index(i pulumi.IntInput) EKSRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EKSRole {
		return vs[0].([]*EKSRole)[vs[1].(int)]
	}).(EKSRoleOutput)
}

type EKSRoleMapOutput struct{ *pulumi.OutputState }

func (EKSRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EKSRole)(nil)).Elem()
}

func (o EKSRoleMapOutput) ToEKSRoleMapOutput() EKSRoleMapOutput {
	return o
}

func (o EKSRoleMapOutput) ToEKSRoleMapOutputWithContext(ctx context.Context) EKSRoleMapOutput {
	return o
}

func (o EKSRoleMapOutput) MapIndex(k pulumi.StringInput) EKSRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EKSRole {
		return vs[0].(map[string]*EKSRole)[vs[1].(string)]
	}).(EKSRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EKSRoleInput)(nil)).Elem(), &EKSRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*EKSRoleArrayInput)(nil)).Elem(), EKSRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EKSRoleMapInput)(nil)).Elem(), EKSRoleMap{})
	pulumi.RegisterOutputType(EKSRoleOutput{})
	pulumi.RegisterOutputType(EKSRoleArrayOutput{})
	pulumi.RegisterOutputType(EKSRoleMapOutput{})
}

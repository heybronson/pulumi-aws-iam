// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package awsiam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AssumableRolesWithSAML struct {
	pulumi.ResourceState

	Admin     pulumi.StringMapOutput `pulumi:"admin"`
	Poweruser pulumi.StringMapOutput `pulumi:"poweruser"`
	Readonly  pulumi.StringMapOutput `pulumi:"readonly"`
}

// NewAssumableRolesWithSAML registers a new resource with the given unique name, arguments, and options.
func NewAssumableRolesWithSAML(ctx *pulumi.Context,
	name string, args *AssumableRolesWithSAMLArgs, opts ...pulumi.ResourceOption) (*AssumableRolesWithSAML, error) {
	if args == nil {
		args = &AssumableRolesWithSAMLArgs{}
	}

	if args.Admin != nil {
		args.Admin = args.Admin.ToAdminRolePtrOutput().ApplyT(func(v *AdminRole) *AdminRole { return v.Defaults() }).(AdminRolePtrOutput)
	}
	if isZero(args.AwsSamlEndpoint) {
		args.AwsSamlEndpoint = pulumi.StringPtr("https://signin.aws.amazon.com/saml")
	}
	if isZero(args.ForceDetachPolicies) {
		args.ForceDetachPolicies = pulumi.BoolPtr(false)
	}
	if isZero(args.MaxSessionDuration) {
		args.MaxSessionDuration = pulumi.IntPtr(3600)
	}
	if args.Poweruser != nil {
		args.Poweruser = args.Poweruser.ToPoweruserRolePtrOutput().ApplyT(func(v *PoweruserRole) *PoweruserRole { return v.Defaults() }).(PoweruserRolePtrOutput)
	}
	if args.Readonly != nil {
		args.Readonly = args.Readonly.ToReadonlyRolePtrOutput().ApplyT(func(v *ReadonlyRole) *ReadonlyRole { return v.Defaults() }).(ReadonlyRolePtrOutput)
	}
	var resource AssumableRolesWithSAML
	err := ctx.RegisterRemoteComponentResource("aws-iam:index:AssumableRolesWithSAML", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type assumableRolesWithSAMLArgs struct {
	Admin *AdminRole `pulumi:"admin"`
	// AWS SAML Endpoint.
	AwsSamlEndpoint *string `pulumi:"awsSamlEndpoint"`
	// Whether policies should be detached from this role when destroying.
	ForceDetachPolicies *bool `pulumi:"forceDetachPolicies"`
	// Maximum CLI/API session duration in seconds between 3600 and 43200.
	MaxSessionDuration *int           `pulumi:"maxSessionDuration"`
	Poweruser          *PoweruserRole `pulumi:"poweruser"`
	// List of SAML Provider IDs.
	ProviderIds []string      `pulumi:"providerIds"`
	Readonly    *ReadonlyRole `pulumi:"readonly"`
}

// The set of arguments for constructing a AssumableRolesWithSAML resource.
type AssumableRolesWithSAMLArgs struct {
	Admin AdminRolePtrInput
	// AWS SAML Endpoint.
	AwsSamlEndpoint pulumi.StringPtrInput
	// Whether policies should be detached from this role when destroying.
	ForceDetachPolicies pulumi.BoolPtrInput
	// Maximum CLI/API session duration in seconds between 3600 and 43200.
	MaxSessionDuration pulumi.IntPtrInput
	Poweruser          PoweruserRolePtrInput
	// List of SAML Provider IDs.
	ProviderIds pulumi.StringArrayInput
	Readonly    ReadonlyRolePtrInput
}

func (AssumableRolesWithSAMLArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assumableRolesWithSAMLArgs)(nil)).Elem()
}

type AssumableRolesWithSAMLInput interface {
	pulumi.Input

	ToAssumableRolesWithSAMLOutput() AssumableRolesWithSAMLOutput
	ToAssumableRolesWithSAMLOutputWithContext(ctx context.Context) AssumableRolesWithSAMLOutput
}

func (*AssumableRolesWithSAML) ElementType() reflect.Type {
	return reflect.TypeOf((**AssumableRolesWithSAML)(nil)).Elem()
}

func (i *AssumableRolesWithSAML) ToAssumableRolesWithSAMLOutput() AssumableRolesWithSAMLOutput {
	return i.ToAssumableRolesWithSAMLOutputWithContext(context.Background())
}

func (i *AssumableRolesWithSAML) ToAssumableRolesWithSAMLOutputWithContext(ctx context.Context) AssumableRolesWithSAMLOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssumableRolesWithSAMLOutput)
}

// AssumableRolesWithSAMLArrayInput is an input type that accepts AssumableRolesWithSAMLArray and AssumableRolesWithSAMLArrayOutput values.
// You can construct a concrete instance of `AssumableRolesWithSAMLArrayInput` via:
//
//          AssumableRolesWithSAMLArray{ AssumableRolesWithSAMLArgs{...} }
type AssumableRolesWithSAMLArrayInput interface {
	pulumi.Input

	ToAssumableRolesWithSAMLArrayOutput() AssumableRolesWithSAMLArrayOutput
	ToAssumableRolesWithSAMLArrayOutputWithContext(context.Context) AssumableRolesWithSAMLArrayOutput
}

type AssumableRolesWithSAMLArray []AssumableRolesWithSAMLInput

func (AssumableRolesWithSAMLArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AssumableRolesWithSAML)(nil)).Elem()
}

func (i AssumableRolesWithSAMLArray) ToAssumableRolesWithSAMLArrayOutput() AssumableRolesWithSAMLArrayOutput {
	return i.ToAssumableRolesWithSAMLArrayOutputWithContext(context.Background())
}

func (i AssumableRolesWithSAMLArray) ToAssumableRolesWithSAMLArrayOutputWithContext(ctx context.Context) AssumableRolesWithSAMLArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssumableRolesWithSAMLArrayOutput)
}

// AssumableRolesWithSAMLMapInput is an input type that accepts AssumableRolesWithSAMLMap and AssumableRolesWithSAMLMapOutput values.
// You can construct a concrete instance of `AssumableRolesWithSAMLMapInput` via:
//
//          AssumableRolesWithSAMLMap{ "key": AssumableRolesWithSAMLArgs{...} }
type AssumableRolesWithSAMLMapInput interface {
	pulumi.Input

	ToAssumableRolesWithSAMLMapOutput() AssumableRolesWithSAMLMapOutput
	ToAssumableRolesWithSAMLMapOutputWithContext(context.Context) AssumableRolesWithSAMLMapOutput
}

type AssumableRolesWithSAMLMap map[string]AssumableRolesWithSAMLInput

func (AssumableRolesWithSAMLMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AssumableRolesWithSAML)(nil)).Elem()
}

func (i AssumableRolesWithSAMLMap) ToAssumableRolesWithSAMLMapOutput() AssumableRolesWithSAMLMapOutput {
	return i.ToAssumableRolesWithSAMLMapOutputWithContext(context.Background())
}

func (i AssumableRolesWithSAMLMap) ToAssumableRolesWithSAMLMapOutputWithContext(ctx context.Context) AssumableRolesWithSAMLMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssumableRolesWithSAMLMapOutput)
}

type AssumableRolesWithSAMLOutput struct{ *pulumi.OutputState }

func (AssumableRolesWithSAMLOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssumableRolesWithSAML)(nil)).Elem()
}

func (o AssumableRolesWithSAMLOutput) ToAssumableRolesWithSAMLOutput() AssumableRolesWithSAMLOutput {
	return o
}

func (o AssumableRolesWithSAMLOutput) ToAssumableRolesWithSAMLOutputWithContext(ctx context.Context) AssumableRolesWithSAMLOutput {
	return o
}

func (o AssumableRolesWithSAMLOutput) Admin() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AssumableRolesWithSAML) pulumi.StringMapOutput { return v.Admin }).(pulumi.StringMapOutput)
}

func (o AssumableRolesWithSAMLOutput) Poweruser() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AssumableRolesWithSAML) pulumi.StringMapOutput { return v.Poweruser }).(pulumi.StringMapOutput)
}

func (o AssumableRolesWithSAMLOutput) Readonly() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AssumableRolesWithSAML) pulumi.StringMapOutput { return v.Readonly }).(pulumi.StringMapOutput)
}

type AssumableRolesWithSAMLArrayOutput struct{ *pulumi.OutputState }

func (AssumableRolesWithSAMLArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AssumableRolesWithSAML)(nil)).Elem()
}

func (o AssumableRolesWithSAMLArrayOutput) ToAssumableRolesWithSAMLArrayOutput() AssumableRolesWithSAMLArrayOutput {
	return o
}

func (o AssumableRolesWithSAMLArrayOutput) ToAssumableRolesWithSAMLArrayOutputWithContext(ctx context.Context) AssumableRolesWithSAMLArrayOutput {
	return o
}

func (o AssumableRolesWithSAMLArrayOutput) Index(i pulumi.IntInput) AssumableRolesWithSAMLOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AssumableRolesWithSAML {
		return vs[0].([]*AssumableRolesWithSAML)[vs[1].(int)]
	}).(AssumableRolesWithSAMLOutput)
}

type AssumableRolesWithSAMLMapOutput struct{ *pulumi.OutputState }

func (AssumableRolesWithSAMLMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AssumableRolesWithSAML)(nil)).Elem()
}

func (o AssumableRolesWithSAMLMapOutput) ToAssumableRolesWithSAMLMapOutput() AssumableRolesWithSAMLMapOutput {
	return o
}

func (o AssumableRolesWithSAMLMapOutput) ToAssumableRolesWithSAMLMapOutputWithContext(ctx context.Context) AssumableRolesWithSAMLMapOutput {
	return o
}

func (o AssumableRolesWithSAMLMapOutput) MapIndex(k pulumi.StringInput) AssumableRolesWithSAMLOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AssumableRolesWithSAML {
		return vs[0].(map[string]*AssumableRolesWithSAML)[vs[1].(string)]
	}).(AssumableRolesWithSAMLOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AssumableRolesWithSAMLInput)(nil)).Elem(), &AssumableRolesWithSAML{})
	pulumi.RegisterInputType(reflect.TypeOf((*AssumableRolesWithSAMLArrayInput)(nil)).Elem(), AssumableRolesWithSAMLArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AssumableRolesWithSAMLMapInput)(nil)).Elem(), AssumableRolesWithSAMLMap{})
	pulumi.RegisterOutputType(AssumableRolesWithSAMLOutput{})
	pulumi.RegisterOutputType(AssumableRolesWithSAMLArrayOutput{})
	pulumi.RegisterOutputType(AssumableRolesWithSAMLMapOutput{})
}

// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsIam
{
    [AwsIamResourceType("aws-iam:index:AssumableRole")]
    public partial class AssumableRole : Pulumi.ComponentResource
    {
        [Output("instanceProfile")]
        public Output<ImmutableDictionary<string, string>> InstanceProfile { get; private set; } = null!;

        [Output("role")]
        public Output<ImmutableDictionary<string, string>> Role { get; private set; } = null!;


        /// <summary>
        /// Create a AssumableRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AssumableRole(string name, AssumableRoleArgs? args = null, ComponentResourceOptions? options = null)
            : base("aws-iam:index:AssumableRole", name, args ?? new AssumableRoleArgs(), MakeResourceOptions(options, ""), remote: true)
        {
        }

        private static ComponentResourceOptions MakeResourceOptions(ComponentResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new ComponentResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = ComponentResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class AssumableRoleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to attach an admin policy to a role.
        /// </summary>
        [Input("attachAdminPolicy")]
        public Input<bool>? AttachAdminPolicy { get; set; }

        /// <summary>
        /// Whether to attach a poweruser policy to a role.
        /// </summary>
        [Input("attachPoweruserPolicy")]
        public Input<bool>? AttachPoweruserPolicy { get; set; }

        /// <summary>
        /// Whether to attach a readonly policy to a role.
        /// </summary>
        [Input("attachReadonlyPolicy")]
        public Input<bool>? AttachReadonlyPolicy { get; set; }

        /// <summary>
        /// A custom role trust policy.
        /// </summary>
        [Input("customRoleTrustPolicy")]
        public Input<string>? CustomRoleTrustPolicy { get; set; }

        /// <summary>
        /// Whether policies should be detached from this role when destroying.
        /// </summary>
        [Input("forceDetachPolicies")]
        public Input<bool>? ForceDetachPolicies { get; set; }

        /// <summary>
        /// Maximum CLI/API session duration in seconds between 3600 and 43200.
        /// </summary>
        [Input("maxSessionDuration")]
        public Input<int>? MaxSessionDuration { get; set; }

        /// <summary>
        /// Max age of valid MFA (in seconds) for roles which require MFA.
        /// </summary>
        [Input("mfaAge")]
        public Input<int>? MfaAge { get; set; }

        /// <summary>
        /// An IAM role that requires MFA.
        /// </summary>
        [Input("role")]
        public Input<Inputs.RoleWithMFAArgs>? Role { get; set; }

        [Input("roleStsExternalIds")]
        private InputList<string>? _roleStsExternalIds;

        /// <summary>
        /// STS ExternalId condition values to use with a role (when MFA is not required).
        /// </summary>
        public InputList<string> RoleStsExternalIds
        {
            get => _roleStsExternalIds ?? (_roleStsExternalIds = new InputList<string>());
            set => _roleStsExternalIds = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to add.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("trustedRoleActions")]
        private InputList<string>? _trustedRoleActions;

        /// <summary>
        /// Actions of STS.
        /// </summary>
        public InputList<string> TrustedRoleActions
        {
            get => _trustedRoleActions ?? (_trustedRoleActions = new InputList<string>());
            set => _trustedRoleActions = value;
        }

        [Input("trustedRoleArns")]
        private InputList<string>? _trustedRoleArns;

        /// <summary>
        /// ARNs of AWS entities who can assume these roles.
        /// </summary>
        public InputList<string> TrustedRoleArns
        {
            get => _trustedRoleArns ?? (_trustedRoleArns = new InputList<string>());
            set => _trustedRoleArns = value;
        }

        [Input("trustedRoleServices")]
        private InputList<string>? _trustedRoleServices;

        /// <summary>
        /// AWS Services that can assume these roles.
        /// </summary>
        public InputList<string> TrustedRoleServices
        {
            get => _trustedRoleServices ?? (_trustedRoleServices = new InputList<string>());
            set => _trustedRoleServices = value;
        }

        public AssumableRoleArgs()
        {
            AttachAdminPolicy = false;
            AttachPoweruserPolicy = false;
            AttachReadonlyPolicy = false;
            CustomRoleTrustPolicy = "";
            ForceDetachPolicies = false;
            MaxSessionDuration = 3600;
            MfaAge = 86400;
        }
    }
}

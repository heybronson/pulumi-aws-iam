// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsIam.Inputs
{

    /// <summary>
    /// The Appmesh policies.
    /// </summary>
    public sealed class EKSAppmeshPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines whether to attach the Appmesh Controller policy to the role.
        /// </summary>
        [Input("controller")]
        public Input<bool>? Controller { get; set; }

        /// <summary>
        /// Determines whether to attach the Appmesh envoy proxy policy to the role.
        /// </summary>
        [Input("envoyProxy")]
        public Input<bool>? EnvoyProxy { get; set; }

        public EKSAppmeshPolicyArgs()
        {
        }
        public static new EKSAppmeshPolicyArgs Empty => new EKSAppmeshPolicyArgs();
    }
}

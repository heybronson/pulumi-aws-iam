// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * This resources helps you create an IAM role which can be assumed by AWS EKS ServiceAccounts with optional policies for
 * commonly used controllers/custom resources within EKS. The optional policies you can specify are:
 *
 * - Cert-Manager
 * - Cluster Autoscaler
 * - EBS CSI Driver
 * - EFS CSI Driver
 * - External DNS
 * - External Secrets
 * - FSx for Lustre CSI Driver
 * - Karpenter
 * - Load Balancer Controller
 * - Load Balancer Controller Target Group Binding Only
 * - App Mesh Controller
 * - App Mesh Envoy Proxy
 * - Managed Service for Prometheus
 * - Node Termination Handler
 * - Velero
 * - VPC CNI
 *
 * ## Example Usage
 * ## VPC CNI
 *
 * ```typescript
 * import * as iam from "@pulumi/aws-iam";
 *
 * export const roleForServiceAccountsEks = new iam.RoleForServiceAccountsEks("aws-iam-example-role-for-service-accounts-eks", {
 *     role: {
 *         name: "vpc-cni"
 *     },
 *     tags: {
 *         Name: "vpc-cni-irsa",
 *     },
 *     oidcProviders: {
 *         main: {
 *             providerArn: "arn:aws:iam::012345678901:oidc-provider/oidc.eks.us-east-1.amazonaws.com/id/5C54DDF35ER19312844C7333374CC09D",
 *             namespaceServiceAccounts: ["default:my-app", "canary:my-app"],
 *         }
 *     },
 *     policies: {
 *         vpnCni: {
 *             attach: true,
 *             enableIpv4: true,
 *         },
 *     },
 * });
 * ```
 * {{ /example }}
 */
export class RoleForServiceAccountsEks extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'aws-iam:index:RoleForServiceAccountsEks';

    /**
     * Returns true if the given object is an instance of RoleForServiceAccountsEks.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RoleForServiceAccountsEks {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RoleForServiceAccountsEks.__pulumiType;
    }

    public readonly role!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a RoleForServiceAccountsEks resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RoleForServiceAccountsEksArgs, opts?: pulumi.ComponentResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["assumeRoleConditionTest"] = (args ? args.assumeRoleConditionTest : undefined) ?? "StringEquals";
            resourceInputs["forceDetachPolicies"] = (args ? args.forceDetachPolicies : undefined) ?? false;
            resourceInputs["maxSessionDuration"] = (args ? args.maxSessionDuration : undefined) ?? 3600;
            resourceInputs["oidcProviders"] = args ? args.oidcProviders : undefined;
            resourceInputs["policies"] = args ? (args.policies ? pulumi.output(args.policies).apply(inputs.eksrolePoliciesArgsProvideDefaults) : undefined) : undefined;
            resourceInputs["policyNamePrefix"] = (args ? args.policyNamePrefix : undefined) ?? "AmazonEKS_";
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
        } else {
            resourceInputs["role"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RoleForServiceAccountsEks.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }
}

/**
 * The set of arguments for constructing a RoleForServiceAccountsEks resource.
 */
export interface RoleForServiceAccountsEksArgs {
    /**
     * Name of the IAM condition operator to evaluate when assuming the role.
     */
    assumeRoleConditionTest?: pulumi.Input<string>;
    /**
     * Whether policies should be detached from this role when destroying.
     */
    forceDetachPolicies?: pulumi.Input<boolean>;
    /**
     * Maximum CLI/API session duration in seconds between 3600 and 43200.
     */
    maxSessionDuration?: pulumi.Input<number>;
    /**
     * Map of OIDC providers.
     */
    oidcProviders?: pulumi.Input<{[key: string]: pulumi.Input<inputs.OIDCProviderArgs>}>;
    policies?: pulumi.Input<inputs.EKSRolePoliciesArgs>;
    /**
     * IAM policy name prefix.
     */
    policyNamePrefix?: pulumi.Input<string>;
    role?: pulumi.Input<inputs.EKSServiceAccountRoleArgs>;
    /**
     * A map of tags to add.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

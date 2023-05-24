# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ReadOnlyPolicyArgs', 'ReadOnlyPolicy']

@pulumi.input_type
class ReadOnlyPolicyArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 additional_policy_json: Optional[pulumi.Input[str]] = None,
                 allow_cloudwatch_logs_query: Optional[pulumi.Input[bool]] = None,
                 allow_predefined_sts_actions: Optional[pulumi.Input[bool]] = None,
                 allow_web_console_services: Optional[pulumi.Input[bool]] = None,
                 allowed_services: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 web_console_services: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a ReadOnlyPolicy resource.
        :param pulumi.Input[str] name: The name of the policy.
        :param pulumi.Input[str] additional_policy_json: JSON policy document if you want to add custom actions.
        :param pulumi.Input[bool] allow_cloudwatch_logs_query: Allows StartQuery/StopQuery/FilterLogEvents CloudWatch actions.
        :param pulumi.Input[bool] allow_predefined_sts_actions: Allows GetCallerIdentity/GetSessionToken/GetAccessKeyInfo sts actions.
        :param pulumi.Input[bool] allow_web_console_services: Allows List/Get/Describe/View actions for services used when browsing AWS console (e.g. resource-groups, tag, health services).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_services: List of services to allow Get/List/Describe/View options. Service name should be the same as corresponding service IAM prefix. See what it is for each service here https://docs.aws.amazon.com/service-authorization/latest/reference/reference_policies_actions-resources-contextkeys.html.
        :param pulumi.Input[str] description: The description of the policy.
        :param pulumi.Input[str] path: The path of the policy in IAM.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to add.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] web_console_services: List of web console services to allow.
        """
        pulumi.set(__self__, "name", name)
        if additional_policy_json is None:
            additional_policy_json = '{}'
        if additional_policy_json is not None:
            pulumi.set(__self__, "additional_policy_json", additional_policy_json)
        if allow_cloudwatch_logs_query is None:
            allow_cloudwatch_logs_query = True
        if allow_cloudwatch_logs_query is not None:
            pulumi.set(__self__, "allow_cloudwatch_logs_query", allow_cloudwatch_logs_query)
        if allow_predefined_sts_actions is None:
            allow_predefined_sts_actions = True
        if allow_predefined_sts_actions is not None:
            pulumi.set(__self__, "allow_predefined_sts_actions", allow_predefined_sts_actions)
        if allow_web_console_services is None:
            allow_web_console_services = True
        if allow_web_console_services is not None:
            pulumi.set(__self__, "allow_web_console_services", allow_web_console_services)
        if allowed_services is not None:
            pulumi.set(__self__, "allowed_services", allowed_services)
        if description is None:
            description = 'IAM Policy'
        if description is not None:
            pulumi.set(__self__, "description", description)
        if path is None:
            path = '/'
        if path is not None:
            pulumi.set(__self__, "path", path)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if web_console_services is not None:
            pulumi.set(__self__, "web_console_services", web_console_services)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the policy.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="additionalPolicyJson")
    def additional_policy_json(self) -> Optional[pulumi.Input[str]]:
        """
        JSON policy document if you want to add custom actions.
        """
        return pulumi.get(self, "additional_policy_json")

    @additional_policy_json.setter
    def additional_policy_json(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "additional_policy_json", value)

    @property
    @pulumi.getter(name="allowCloudwatchLogsQuery")
    def allow_cloudwatch_logs_query(self) -> Optional[pulumi.Input[bool]]:
        """
        Allows StartQuery/StopQuery/FilterLogEvents CloudWatch actions.
        """
        return pulumi.get(self, "allow_cloudwatch_logs_query")

    @allow_cloudwatch_logs_query.setter
    def allow_cloudwatch_logs_query(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_cloudwatch_logs_query", value)

    @property
    @pulumi.getter(name="allowPredefinedStsActions")
    def allow_predefined_sts_actions(self) -> Optional[pulumi.Input[bool]]:
        """
        Allows GetCallerIdentity/GetSessionToken/GetAccessKeyInfo sts actions.
        """
        return pulumi.get(self, "allow_predefined_sts_actions")

    @allow_predefined_sts_actions.setter
    def allow_predefined_sts_actions(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_predefined_sts_actions", value)

    @property
    @pulumi.getter(name="allowWebConsoleServices")
    def allow_web_console_services(self) -> Optional[pulumi.Input[bool]]:
        """
        Allows List/Get/Describe/View actions for services used when browsing AWS console (e.g. resource-groups, tag, health services).
        """
        return pulumi.get(self, "allow_web_console_services")

    @allow_web_console_services.setter
    def allow_web_console_services(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_web_console_services", value)

    @property
    @pulumi.getter(name="allowedServices")
    def allowed_services(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of services to allow Get/List/Describe/View options. Service name should be the same as corresponding service IAM prefix. See what it is for each service here https://docs.aws.amazon.com/service-authorization/latest/reference/reference_policies_actions-resources-contextkeys.html.
        """
        return pulumi.get(self, "allowed_services")

    @allowed_services.setter
    def allowed_services(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "allowed_services", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[str]]:
        """
        The path of the policy in IAM.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to add.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="webConsoleServices")
    def web_console_services(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of web console services to allow.
        """
        return pulumi.get(self, "web_console_services")

    @web_console_services.setter
    def web_console_services(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "web_console_services", value)


class ReadOnlyPolicy(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_policy_json: Optional[pulumi.Input[str]] = None,
                 allow_cloudwatch_logs_query: Optional[pulumi.Input[bool]] = None,
                 allow_predefined_sts_actions: Optional[pulumi.Input[bool]] = None,
                 allow_web_console_services: Optional[pulumi.Input[bool]] = None,
                 allowed_services: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 web_console_services: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        This resource helps you create an IAM read-only policy for the services you specify. The default AWS
        read-only policies may not include services you need or may contain services you do not need access to.
        This resource helps ensure your read-only policy has permissions to exactly what you specify.

        ## Example Usage
        ## RDS and Dynamo Read Only Policy

        ```python
        import pulumi
        import pulumi_aws_iam as iam

        read_only_policy = iam.ReadOnlyPolicy(
            'read_only_policy',
            name='example',
            path='/',
            description='My example read only policy',
            allowed_services=['rds','dynamodb'],
        )

        pulumi.export('read_only_policy', read_only_policy)
        ```
        {{ /example }}

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] additional_policy_json: JSON policy document if you want to add custom actions.
        :param pulumi.Input[bool] allow_cloudwatch_logs_query: Allows StartQuery/StopQuery/FilterLogEvents CloudWatch actions.
        :param pulumi.Input[bool] allow_predefined_sts_actions: Allows GetCallerIdentity/GetSessionToken/GetAccessKeyInfo sts actions.
        :param pulumi.Input[bool] allow_web_console_services: Allows List/Get/Describe/View actions for services used when browsing AWS console (e.g. resource-groups, tag, health services).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_services: List of services to allow Get/List/Describe/View options. Service name should be the same as corresponding service IAM prefix. See what it is for each service here https://docs.aws.amazon.com/service-authorization/latest/reference/reference_policies_actions-resources-contextkeys.html.
        :param pulumi.Input[str] description: The description of the policy.
        :param pulumi.Input[str] name: The name of the policy.
        :param pulumi.Input[str] path: The path of the policy in IAM.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to add.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] web_console_services: List of web console services to allow.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ReadOnlyPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource helps you create an IAM read-only policy for the services you specify. The default AWS
        read-only policies may not include services you need or may contain services you do not need access to.
        This resource helps ensure your read-only policy has permissions to exactly what you specify.

        ## Example Usage
        ## RDS and Dynamo Read Only Policy

        ```python
        import pulumi
        import pulumi_aws_iam as iam

        read_only_policy = iam.ReadOnlyPolicy(
            'read_only_policy',
            name='example',
            path='/',
            description='My example read only policy',
            allowed_services=['rds','dynamodb'],
        )

        pulumi.export('read_only_policy', read_only_policy)
        ```
        {{ /example }}

        :param str resource_name: The name of the resource.
        :param ReadOnlyPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReadOnlyPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_policy_json: Optional[pulumi.Input[str]] = None,
                 allow_cloudwatch_logs_query: Optional[pulumi.Input[bool]] = None,
                 allow_predefined_sts_actions: Optional[pulumi.Input[bool]] = None,
                 allow_web_console_services: Optional[pulumi.Input[bool]] = None,
                 allowed_services: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 web_console_services: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is not None:
            raise ValueError('ComponentResource classes do not support opts.id')
        else:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ReadOnlyPolicyArgs.__new__(ReadOnlyPolicyArgs)

            if additional_policy_json is None:
                additional_policy_json = '{}'
            __props__.__dict__["additional_policy_json"] = additional_policy_json
            if allow_cloudwatch_logs_query is None:
                allow_cloudwatch_logs_query = True
            __props__.__dict__["allow_cloudwatch_logs_query"] = allow_cloudwatch_logs_query
            if allow_predefined_sts_actions is None:
                allow_predefined_sts_actions = True
            __props__.__dict__["allow_predefined_sts_actions"] = allow_predefined_sts_actions
            if allow_web_console_services is None:
                allow_web_console_services = True
            __props__.__dict__["allow_web_console_services"] = allow_web_console_services
            __props__.__dict__["allowed_services"] = allowed_services
            if description is None:
                description = 'IAM Policy'
            __props__.__dict__["description"] = description
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            if path is None:
                path = '/'
            __props__.__dict__["path"] = path
            __props__.__dict__["tags"] = tags
            __props__.__dict__["web_console_services"] = web_console_services
            __props__.__dict__["arn"] = None
            __props__.__dict__["id"] = None
            __props__.__dict__["policy"] = None
            __props__.__dict__["policy_json"] = None
        super(ReadOnlyPolicy, __self__).__init__(
            'aws-iam:index:ReadOnlyPolicy',
            resource_name,
            __props__,
            opts,
            remote=True)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN assigned by AWS to this policy.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The description of the policy.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> pulumi.Output[str]:
        """
        The policy's ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the policy.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[str]:
        """
        The path of the policy in IAM.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Output[str]:
        """
        The policy document.
        """
        return pulumi.get(self, "policy")

    @property
    @pulumi.getter(name="policyJson")
    def policy_json(self) -> pulumi.Output[str]:
        """
        Policy document as json. Useful if you need document but do not want to create IAM policy itself. For example for SSO Permission Set inline policies.
        """
        return pulumi.get(self, "policy_json")


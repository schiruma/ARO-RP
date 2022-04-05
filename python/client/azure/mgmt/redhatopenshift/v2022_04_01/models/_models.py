# coding=utf-8
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#   http://www.apache.org/licenses/LICENSE-2.0
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
# 
# Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code is regenerated.
# --------------------------------------------------------------------------

import msrest.serialization


class APIServerProfile(msrest.serialization.Model):
    """APIServerProfile represents an API server profile.

    :ivar visibility: API server visibility. Possible values include: "Private", "Public".
    :vartype visibility: str or ~azure.mgmt.redhatopenshift.v2022_04_01.models.Visibility
    :ivar url: The URL to access the cluster API server.
    :vartype url: str
    :ivar ip: The IP of the cluster API server.
    :vartype ip: str
    """

    _attribute_map = {
        'visibility': {'key': 'visibility', 'type': 'str'},
        'url': {'key': 'url', 'type': 'str'},
        'ip': {'key': 'ip', 'type': 'str'},
    }

    def __init__(
        self,
        **kwargs
    ):
        """
        :keyword visibility: API server visibility. Possible values include: "Private", "Public".
        :paramtype visibility: str or ~azure.mgmt.redhatopenshift.v2022_04_01.models.Visibility
        :keyword url: The URL to access the cluster API server.
        :paramtype url: str
        :keyword ip: The IP of the cluster API server.
        :paramtype ip: str
        """
        super(APIServerProfile, self).__init__(**kwargs)
        self.visibility = kwargs.get('visibility', None)
        self.url = kwargs.get('url', None)
        self.ip = kwargs.get('ip', None)


class CloudErrorBody(msrest.serialization.Model):
    """CloudErrorBody represents the body of a cloud error.

    :ivar code: An identifier for the error. Codes are invariant and are intended to be consumed
     programmatically.
    :vartype code: str
    :ivar message: A message describing the error, intended to be suitable for display in a user
     interface.
    :vartype message: str
    :ivar target: The target of the particular error. For example, the name of the property in
     error.
    :vartype target: str
    :ivar details: A list of additional details about the error.
    :vartype details: list[~azure.mgmt.redhatopenshift.v2022_04_01.models.CloudErrorBody]
    """

    _attribute_map = {
        'code': {'key': 'code', 'type': 'str'},
        'message': {'key': 'message', 'type': 'str'},
        'target': {'key': 'target', 'type': 'str'},
        'details': {'key': 'details', 'type': '[CloudErrorBody]'},
    }

    def __init__(
        self,
        **kwargs
    ):
        """
        :keyword code: An identifier for the error. Codes are invariant and are intended to be consumed
         programmatically.
        :paramtype code: str
        :keyword message: A message describing the error, intended to be suitable for display in a user
         interface.
        :paramtype message: str
        :keyword target: The target of the particular error. For example, the name of the property in
         error.
        :paramtype target: str
        :keyword details: A list of additional details about the error.
        :paramtype details: list[~azure.mgmt.redhatopenshift.v2022_04_01.models.CloudErrorBody]
        """
        super(CloudErrorBody, self).__init__(**kwargs)
        self.code = kwargs.get('code', None)
        self.message = kwargs.get('message', None)
        self.target = kwargs.get('target', None)
        self.details = kwargs.get('details', None)


class ClusterProfile(msrest.serialization.Model):
    """ClusterProfile represents a cluster profile.

    :ivar pull_secret: The pull secret for the cluster.
    :vartype pull_secret: str
    :ivar domain: The domain for the cluster.
    :vartype domain: str
    :ivar version: The version of the cluster.
    :vartype version: str
    :ivar resource_group_id: The ID of the cluster resource group.
    :vartype resource_group_id: str
    :ivar fips_validated_modules: If FIPS validated crypto modules are used. Possible values
     include: "Disabled", "Enabled".
    :vartype fips_validated_modules: str or
     ~azure.mgmt.redhatopenshift.v2022_04_01.models.FipsValidatedModules
    """

    _attribute_map = {
        'pull_secret': {'key': 'pullSecret', 'type': 'str'},
        'domain': {'key': 'domain', 'type': 'str'},
        'version': {'key': 'version', 'type': 'str'},
        'resource_group_id': {'key': 'resourceGroupId', 'type': 'str'},
        'fips_validated_modules': {'key': 'fipsValidatedModules', 'type': 'str'},
    }

    def __init__(
        self,
        **kwargs
    ):
        """
        :keyword pull_secret: The pull secret for the cluster.
        :paramtype pull_secret: str
        :keyword domain: The domain for the cluster.
        :paramtype domain: str
        :keyword version: The version of the cluster.
        :paramtype version: str
        :keyword resource_group_id: The ID of the cluster resource group.
        :paramtype resource_group_id: str
        :keyword fips_validated_modules: If FIPS validated crypto modules are used. Possible values
         include: "Disabled", "Enabled".
        :paramtype fips_validated_modules: str or
         ~azure.mgmt.redhatopenshift.v2022_04_01.models.FipsValidatedModules
        """
        super(ClusterProfile, self).__init__(**kwargs)
        self.pull_secret = kwargs.get('pull_secret', None)
        self.domain = kwargs.get('domain', None)
        self.version = kwargs.get('version', None)
        self.resource_group_id = kwargs.get('resource_group_id', None)
        self.fips_validated_modules = kwargs.get('fips_validated_modules', None)


class ConsoleProfile(msrest.serialization.Model):
    """ConsoleProfile represents a console profile.

    :ivar url: The URL to access the cluster console.
    :vartype url: str
    """

    _attribute_map = {
        'url': {'key': 'url', 'type': 'str'},
    }

    def __init__(
        self,
        **kwargs
    ):
        """
        :keyword url: The URL to access the cluster console.
        :paramtype url: str
        """
        super(ConsoleProfile, self).__init__(**kwargs)
        self.url = kwargs.get('url', None)


class Display(msrest.serialization.Model):
    """Display represents the display details of an operation.

    :ivar provider: Friendly name of the resource provider.
    :vartype provider: str
    :ivar resource: Resource type on which the operation is performed.
    :vartype resource: str
    :ivar operation: Operation type: read, write, delete, listKeys/action, etc.
    :vartype operation: str
    :ivar description: Friendly name of the operation.
    :vartype description: str
    """

    _attribute_map = {
        'provider': {'key': 'provider', 'type': 'str'},
        'resource': {'key': 'resource', 'type': 'str'},
        'operation': {'key': 'operation', 'type': 'str'},
        'description': {'key': 'description', 'type': 'str'},
    }

    def __init__(
        self,
        **kwargs
    ):
        """
        :keyword provider: Friendly name of the resource provider.
        :paramtype provider: str
        :keyword resource: Resource type on which the operation is performed.
        :paramtype resource: str
        :keyword operation: Operation type: read, write, delete, listKeys/action, etc.
        :paramtype operation: str
        :keyword description: Friendly name of the operation.
        :paramtype description: str
        """
        super(Display, self).__init__(**kwargs)
        self.provider = kwargs.get('provider', None)
        self.resource = kwargs.get('resource', None)
        self.operation = kwargs.get('operation', None)
        self.description = kwargs.get('description', None)


class IngressProfile(msrest.serialization.Model):
    """IngressProfile represents an ingress profile.

    :ivar name: The ingress profile name.
    :vartype name: str
    :ivar visibility: Ingress visibility. Possible values include: "Private", "Public".
    :vartype visibility: str or ~azure.mgmt.redhatopenshift.v2022_04_01.models.Visibility
    :ivar ip: The IP of the ingress.
    :vartype ip: str
    """

    _attribute_map = {
        'name': {'key': 'name', 'type': 'str'},
        'visibility': {'key': 'visibility', 'type': 'str'},
        'ip': {'key': 'ip', 'type': 'str'},
    }

    def __init__(
        self,
        **kwargs
    ):
        """
        :keyword name: The ingress profile name.
        :paramtype name: str
        :keyword visibility: Ingress visibility. Possible values include: "Private", "Public".
        :paramtype visibility: str or ~azure.mgmt.redhatopenshift.v2022_04_01.models.Visibility
        :keyword ip: The IP of the ingress.
        :paramtype ip: str
        """
        super(IngressProfile, self).__init__(**kwargs)
        self.name = kwargs.get('name', None)
        self.visibility = kwargs.get('visibility', None)
        self.ip = kwargs.get('ip', None)


class MasterProfile(msrest.serialization.Model):
    """MasterProfile represents a master profile.

    :ivar vm_size: The size of the master VMs.
    :vartype vm_size: str
    :ivar subnet_id: The Azure resource ID of the master subnet.
    :vartype subnet_id: str
    :ivar encryption_at_host: Whether master virtual machines are encrypted at host. Possible
     values include: "Disabled", "Enabled".
    :vartype encryption_at_host: str or
     ~azure.mgmt.redhatopenshift.v2022_04_01.models.EncryptionAtHost
    :ivar disk_encryption_set_id: The resource ID of an associated DiskEncryptionSet, if
     applicable.
    :vartype disk_encryption_set_id: str
    """

    _attribute_map = {
        'vm_size': {'key': 'vmSize', 'type': 'str'},
        'subnet_id': {'key': 'subnetId', 'type': 'str'},
        'encryption_at_host': {'key': 'encryptionAtHost', 'type': 'str'},
        'disk_encryption_set_id': {'key': 'diskEncryptionSetId', 'type': 'str'},
    }

    def __init__(
        self,
        **kwargs
    ):
        """
        :keyword vm_size: The size of the master VMs.
        :paramtype vm_size: str
        :keyword subnet_id: The Azure resource ID of the master subnet.
        :paramtype subnet_id: str
        :keyword encryption_at_host: Whether master virtual machines are encrypted at host. Possible
         values include: "Disabled", "Enabled".
        :paramtype encryption_at_host: str or
         ~azure.mgmt.redhatopenshift.v2022_04_01.models.EncryptionAtHost
        :keyword disk_encryption_set_id: The resource ID of an associated DiskEncryptionSet, if
         applicable.
        :paramtype disk_encryption_set_id: str
        """
        super(MasterProfile, self).__init__(**kwargs)
        self.vm_size = kwargs.get('vm_size', None)
        self.subnet_id = kwargs.get('subnet_id', None)
        self.encryption_at_host = kwargs.get('encryption_at_host', None)
        self.disk_encryption_set_id = kwargs.get('disk_encryption_set_id', None)


class NetworkProfile(msrest.serialization.Model):
    """NetworkProfile represents a network profile.

    :ivar pod_cidr: The CIDR used for OpenShift/Kubernetes Pods.
    :vartype pod_cidr: str
    :ivar service_cidr: The CIDR used for OpenShift/Kubernetes Services.
    :vartype service_cidr: str
    """

    _attribute_map = {
        'pod_cidr': {'key': 'podCidr', 'type': 'str'},
        'service_cidr': {'key': 'serviceCidr', 'type': 'str'},
    }

    def __init__(
        self,
        **kwargs
    ):
        """
        :keyword pod_cidr: The CIDR used for OpenShift/Kubernetes Pods.
        :paramtype pod_cidr: str
        :keyword service_cidr: The CIDR used for OpenShift/Kubernetes Services.
        :paramtype service_cidr: str
        """
        super(NetworkProfile, self).__init__(**kwargs)
        self.pod_cidr = kwargs.get('pod_cidr', None)
        self.service_cidr = kwargs.get('service_cidr', None)


class Resource(msrest.serialization.Model):
    """Common fields that are returned in the response for all Azure Resource Manager resources.

    Variables are only populated by the server, and will be ignored when sending a request.

    :ivar id: Fully qualified resource ID for the resource. Ex -
     /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}.
    :vartype id: str
    :ivar name: The name of the resource.
    :vartype name: str
    :ivar type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or
     "Microsoft.Storage/storageAccounts".
    :vartype type: str
    """

    _validation = {
        'id': {'readonly': True},
        'name': {'readonly': True},
        'type': {'readonly': True},
    }

    _attribute_map = {
        'id': {'key': 'id', 'type': 'str'},
        'name': {'key': 'name', 'type': 'str'},
        'type': {'key': 'type', 'type': 'str'},
    }

    def __init__(
        self,
        **kwargs
    ):
        """
        """
        super(Resource, self).__init__(**kwargs)
        self.id = None
        self.name = None
        self.type = None


class TrackedResource(Resource):
    """The resource model definition for an Azure Resource Manager tracked top level resource.

    Variables are only populated by the server, and will be ignored when sending a request.

    All required parameters must be populated in order to send to Azure.

    :ivar id: Fully qualified resource ID for the resource. Ex -
     /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}.
    :vartype id: str
    :ivar name: The name of the resource.
    :vartype name: str
    :ivar type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or
     "Microsoft.Storage/storageAccounts".
    :vartype type: str
    :ivar tags: A set of tags. Resource tags.
    :vartype tags: dict[str, str]
    :ivar location: Required. The geo-location where the resource lives.
    :vartype location: str
    """

    _validation = {
        'id': {'readonly': True},
        'name': {'readonly': True},
        'type': {'readonly': True},
        'location': {'required': True},
    }

    _attribute_map = {
        'id': {'key': 'id', 'type': 'str'},
        'name': {'key': 'name', 'type': 'str'},
        'type': {'key': 'type', 'type': 'str'},
        'tags': {'key': 'tags', 'type': '{str}'},
        'location': {'key': 'location', 'type': 'str'},
    }

    def __init__(
        self,
        **kwargs
    ):
        """
        :keyword tags: A set of tags. Resource tags.
        :paramtype tags: dict[str, str]
        :keyword location: Required. The geo-location where the resource lives.
        :paramtype location: str
        """
        super(TrackedResource, self).__init__(**kwargs)
        self.tags = kwargs.get('tags', None)
        self.location = kwargs['location']


class OpenShiftCluster(TrackedResource):
    """OpenShiftCluster represents an Azure Red Hat OpenShift cluster.

    Variables are only populated by the server, and will be ignored when sending a request.

    All required parameters must be populated in order to send to Azure.

    :ivar id: Fully qualified resource ID for the resource. Ex -
     /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}.
    :vartype id: str
    :ivar name: The name of the resource.
    :vartype name: str
    :ivar type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or
     "Microsoft.Storage/storageAccounts".
    :vartype type: str
    :ivar tags: A set of tags. Resource tags.
    :vartype tags: dict[str, str]
    :ivar location: Required. The geo-location where the resource lives.
    :vartype location: str
    :ivar system_data: The system meta data relating to this resource.
    :vartype system_data: ~azure.mgmt.redhatopenshift.v2022_04_01.models.SystemData
    :ivar provisioning_state: The cluster provisioning state. Possible values include:
     "AdminUpdating", "Creating", "Deleting", "Failed", "Succeeded", "Updating".
    :vartype provisioning_state: str or
     ~azure.mgmt.redhatopenshift.v2022_04_01.models.ProvisioningState
    :ivar cluster_profile: The cluster profile.
    :vartype cluster_profile: ~azure.mgmt.redhatopenshift.v2022_04_01.models.ClusterProfile
    :ivar console_profile: The console profile.
    :vartype console_profile: ~azure.mgmt.redhatopenshift.v2022_04_01.models.ConsoleProfile
    :ivar service_principal_profile: The cluster service principal profile.
    :vartype service_principal_profile:
     ~azure.mgmt.redhatopenshift.v2022_04_01.models.ServicePrincipalProfile
    :ivar network_profile: The cluster network profile.
    :vartype network_profile: ~azure.mgmt.redhatopenshift.v2022_04_01.models.NetworkProfile
    :ivar master_profile: The cluster master profile.
    :vartype master_profile: ~azure.mgmt.redhatopenshift.v2022_04_01.models.MasterProfile
    :ivar worker_profiles: The cluster worker profiles.
    :vartype worker_profiles: list[~azure.mgmt.redhatopenshift.v2022_04_01.models.WorkerProfile]
    :ivar apiserver_profile: The cluster API server profile.
    :vartype apiserver_profile: ~azure.mgmt.redhatopenshift.v2022_04_01.models.APIServerProfile
    :ivar ingress_profiles: The cluster ingress profiles.
    :vartype ingress_profiles: list[~azure.mgmt.redhatopenshift.v2022_04_01.models.IngressProfile]
    """

    _validation = {
        'id': {'readonly': True},
        'name': {'readonly': True},
        'type': {'readonly': True},
        'location': {'required': True},
        'system_data': {'readonly': True},
    }

    _attribute_map = {
        'id': {'key': 'id', 'type': 'str'},
        'name': {'key': 'name', 'type': 'str'},
        'type': {'key': 'type', 'type': 'str'},
        'tags': {'key': 'tags', 'type': '{str}'},
        'location': {'key': 'location', 'type': 'str'},
        'system_data': {'key': 'systemData', 'type': 'SystemData'},
        'provisioning_state': {'key': 'properties.provisioningState', 'type': 'str'},
        'cluster_profile': {'key': 'properties.clusterProfile', 'type': 'ClusterProfile'},
        'console_profile': {'key': 'properties.consoleProfile', 'type': 'ConsoleProfile'},
        'service_principal_profile': {'key': 'properties.servicePrincipalProfile', 'type': 'ServicePrincipalProfile'},
        'network_profile': {'key': 'properties.networkProfile', 'type': 'NetworkProfile'},
        'master_profile': {'key': 'properties.masterProfile', 'type': 'MasterProfile'},
        'worker_profiles': {'key': 'properties.workerProfiles', 'type': '[WorkerProfile]'},
        'apiserver_profile': {'key': 'properties.apiserverProfile', 'type': 'APIServerProfile'},
        'ingress_profiles': {'key': 'properties.ingressProfiles', 'type': '[IngressProfile]'},
    }

    def __init__(
        self,
        **kwargs
    ):
        """
        :keyword tags: A set of tags. Resource tags.
        :paramtype tags: dict[str, str]
        :keyword location: Required. The geo-location where the resource lives.
        :paramtype location: str
        :keyword provisioning_state: The cluster provisioning state. Possible values include:
         "AdminUpdating", "Creating", "Deleting", "Failed", "Succeeded", "Updating".
        :paramtype provisioning_state: str or
         ~azure.mgmt.redhatopenshift.v2022_04_01.models.ProvisioningState
        :keyword cluster_profile: The cluster profile.
        :paramtype cluster_profile: ~azure.mgmt.redhatopenshift.v2022_04_01.models.ClusterProfile
        :keyword console_profile: The console profile.
        :paramtype console_profile: ~azure.mgmt.redhatopenshift.v2022_04_01.models.ConsoleProfile
        :keyword service_principal_profile: The cluster service principal profile.
        :paramtype service_principal_profile:
         ~azure.mgmt.redhatopenshift.v2022_04_01.models.ServicePrincipalProfile
        :keyword network_profile: The cluster network profile.
        :paramtype network_profile: ~azure.mgmt.redhatopenshift.v2022_04_01.models.NetworkProfile
        :keyword master_profile: The cluster master profile.
        :paramtype master_profile: ~azure.mgmt.redhatopenshift.v2022_04_01.models.MasterProfile
        :keyword worker_profiles: The cluster worker profiles.
        :paramtype worker_profiles: list[~azure.mgmt.redhatopenshift.v2022_04_01.models.WorkerProfile]
        :keyword apiserver_profile: The cluster API server profile.
        :paramtype apiserver_profile: ~azure.mgmt.redhatopenshift.v2022_04_01.models.APIServerProfile
        :keyword ingress_profiles: The cluster ingress profiles.
        :paramtype ingress_profiles:
         list[~azure.mgmt.redhatopenshift.v2022_04_01.models.IngressProfile]
        """
        super(OpenShiftCluster, self).__init__(**kwargs)
        self.system_data = None
        self.provisioning_state = kwargs.get('provisioning_state', None)
        self.cluster_profile = kwargs.get('cluster_profile', None)
        self.console_profile = kwargs.get('console_profile', None)
        self.service_principal_profile = kwargs.get('service_principal_profile', None)
        self.network_profile = kwargs.get('network_profile', None)
        self.master_profile = kwargs.get('master_profile', None)
        self.worker_profiles = kwargs.get('worker_profiles', None)
        self.apiserver_profile = kwargs.get('apiserver_profile', None)
        self.ingress_profiles = kwargs.get('ingress_profiles', None)


class OpenShiftClusterAdminKubeconfig(msrest.serialization.Model):
    """OpenShiftClusterAdminKubeconfig represents an OpenShift cluster's admin kubeconfig.

    :ivar kubeconfig: The base64-encoded kubeconfig file.
    :vartype kubeconfig: str
    """

    _attribute_map = {
        'kubeconfig': {'key': 'kubeconfig', 'type': 'str'},
    }

    def __init__(
        self,
        **kwargs
    ):
        """
        :keyword kubeconfig: The base64-encoded kubeconfig file.
        :paramtype kubeconfig: str
        """
        super(OpenShiftClusterAdminKubeconfig, self).__init__(**kwargs)
        self.kubeconfig = kwargs.get('kubeconfig', None)


class OpenShiftClusterCredentials(msrest.serialization.Model):
    """OpenShiftClusterCredentials represents an OpenShift cluster's credentials.

    :ivar kubeadmin_username: The username for the kubeadmin user.
    :vartype kubeadmin_username: str
    :ivar kubeadmin_password: The password for the kubeadmin user.
    :vartype kubeadmin_password: str
    """

    _attribute_map = {
        'kubeadmin_username': {'key': 'kubeadminUsername', 'type': 'str'},
        'kubeadmin_password': {'key': 'kubeadminPassword', 'type': 'str'},
    }

    def __init__(
        self,
        **kwargs
    ):
        """
        :keyword kubeadmin_username: The username for the kubeadmin user.
        :paramtype kubeadmin_username: str
        :keyword kubeadmin_password: The password for the kubeadmin user.
        :paramtype kubeadmin_password: str
        """
        super(OpenShiftClusterCredentials, self).__init__(**kwargs)
        self.kubeadmin_username = kwargs.get('kubeadmin_username', None)
        self.kubeadmin_password = kwargs.get('kubeadmin_password', None)


class OpenShiftClusterList(msrest.serialization.Model):
    """OpenShiftClusterList represents a list of OpenShift clusters.

    :ivar value: The list of OpenShift clusters.
    :vartype value: list[~azure.mgmt.redhatopenshift.v2022_04_01.models.OpenShiftCluster]
    :ivar next_link: The link used to get the next page of operations.
    :vartype next_link: str
    """

    _attribute_map = {
        'value': {'key': 'value', 'type': '[OpenShiftCluster]'},
        'next_link': {'key': 'nextLink', 'type': 'str'},
    }

    def __init__(
        self,
        **kwargs
    ):
        """
        :keyword value: The list of OpenShift clusters.
        :paramtype value: list[~azure.mgmt.redhatopenshift.v2022_04_01.models.OpenShiftCluster]
        :keyword next_link: The link used to get the next page of operations.
        :paramtype next_link: str
        """
        super(OpenShiftClusterList, self).__init__(**kwargs)
        self.value = kwargs.get('value', None)
        self.next_link = kwargs.get('next_link', None)


class OpenShiftClusterUpdate(msrest.serialization.Model):
    """OpenShiftCluster represents an Azure Red Hat OpenShift cluster.

    Variables are only populated by the server, and will be ignored when sending a request.

    :ivar tags: A set of tags. The resource tags.
    :vartype tags: dict[str, str]
    :ivar system_data: The system meta data relating to this resource.
    :vartype system_data: ~azure.mgmt.redhatopenshift.v2022_04_01.models.SystemData
    :ivar provisioning_state: The cluster provisioning state. Possible values include:
     "AdminUpdating", "Creating", "Deleting", "Failed", "Succeeded", "Updating".
    :vartype provisioning_state: str or
     ~azure.mgmt.redhatopenshift.v2022_04_01.models.ProvisioningState
    :ivar cluster_profile: The cluster profile.
    :vartype cluster_profile: ~azure.mgmt.redhatopenshift.v2022_04_01.models.ClusterProfile
    :ivar console_profile: The console profile.
    :vartype console_profile: ~azure.mgmt.redhatopenshift.v2022_04_01.models.ConsoleProfile
    :ivar service_principal_profile: The cluster service principal profile.
    :vartype service_principal_profile:
     ~azure.mgmt.redhatopenshift.v2022_04_01.models.ServicePrincipalProfile
    :ivar network_profile: The cluster network profile.
    :vartype network_profile: ~azure.mgmt.redhatopenshift.v2022_04_01.models.NetworkProfile
    :ivar master_profile: The cluster master profile.
    :vartype master_profile: ~azure.mgmt.redhatopenshift.v2022_04_01.models.MasterProfile
    :ivar worker_profiles: The cluster worker profiles.
    :vartype worker_profiles: list[~azure.mgmt.redhatopenshift.v2022_04_01.models.WorkerProfile]
    :ivar apiserver_profile: The cluster API server profile.
    :vartype apiserver_profile: ~azure.mgmt.redhatopenshift.v2022_04_01.models.APIServerProfile
    :ivar ingress_profiles: The cluster ingress profiles.
    :vartype ingress_profiles: list[~azure.mgmt.redhatopenshift.v2022_04_01.models.IngressProfile]
    """

    _validation = {
        'system_data': {'readonly': True},
    }

    _attribute_map = {
        'tags': {'key': 'tags', 'type': '{str}'},
        'system_data': {'key': 'systemData', 'type': 'SystemData'},
        'provisioning_state': {'key': 'properties.provisioningState', 'type': 'str'},
        'cluster_profile': {'key': 'properties.clusterProfile', 'type': 'ClusterProfile'},
        'console_profile': {'key': 'properties.consoleProfile', 'type': 'ConsoleProfile'},
        'service_principal_profile': {'key': 'properties.servicePrincipalProfile', 'type': 'ServicePrincipalProfile'},
        'network_profile': {'key': 'properties.networkProfile', 'type': 'NetworkProfile'},
        'master_profile': {'key': 'properties.masterProfile', 'type': 'MasterProfile'},
        'worker_profiles': {'key': 'properties.workerProfiles', 'type': '[WorkerProfile]'},
        'apiserver_profile': {'key': 'properties.apiserverProfile', 'type': 'APIServerProfile'},
        'ingress_profiles': {'key': 'properties.ingressProfiles', 'type': '[IngressProfile]'},
    }

    def __init__(
        self,
        **kwargs
    ):
        """
        :keyword tags: A set of tags. The resource tags.
        :paramtype tags: dict[str, str]
        :keyword provisioning_state: The cluster provisioning state. Possible values include:
         "AdminUpdating", "Creating", "Deleting", "Failed", "Succeeded", "Updating".
        :paramtype provisioning_state: str or
         ~azure.mgmt.redhatopenshift.v2022_04_01.models.ProvisioningState
        :keyword cluster_profile: The cluster profile.
        :paramtype cluster_profile: ~azure.mgmt.redhatopenshift.v2022_04_01.models.ClusterProfile
        :keyword console_profile: The console profile.
        :paramtype console_profile: ~azure.mgmt.redhatopenshift.v2022_04_01.models.ConsoleProfile
        :keyword service_principal_profile: The cluster service principal profile.
        :paramtype service_principal_profile:
         ~azure.mgmt.redhatopenshift.v2022_04_01.models.ServicePrincipalProfile
        :keyword network_profile: The cluster network profile.
        :paramtype network_profile: ~azure.mgmt.redhatopenshift.v2022_04_01.models.NetworkProfile
        :keyword master_profile: The cluster master profile.
        :paramtype master_profile: ~azure.mgmt.redhatopenshift.v2022_04_01.models.MasterProfile
        :keyword worker_profiles: The cluster worker profiles.
        :paramtype worker_profiles: list[~azure.mgmt.redhatopenshift.v2022_04_01.models.WorkerProfile]
        :keyword apiserver_profile: The cluster API server profile.
        :paramtype apiserver_profile: ~azure.mgmt.redhatopenshift.v2022_04_01.models.APIServerProfile
        :keyword ingress_profiles: The cluster ingress profiles.
        :paramtype ingress_profiles:
         list[~azure.mgmt.redhatopenshift.v2022_04_01.models.IngressProfile]
        """
        super(OpenShiftClusterUpdate, self).__init__(**kwargs)
        self.tags = kwargs.get('tags', None)
        self.system_data = None
        self.provisioning_state = kwargs.get('provisioning_state', None)
        self.cluster_profile = kwargs.get('cluster_profile', None)
        self.console_profile = kwargs.get('console_profile', None)
        self.service_principal_profile = kwargs.get('service_principal_profile', None)
        self.network_profile = kwargs.get('network_profile', None)
        self.master_profile = kwargs.get('master_profile', None)
        self.worker_profiles = kwargs.get('worker_profiles', None)
        self.apiserver_profile = kwargs.get('apiserver_profile', None)
        self.ingress_profiles = kwargs.get('ingress_profiles', None)


class Operation(msrest.serialization.Model):
    """Operation represents an RP operation.

    :ivar name: Operation name: {provider}/{resource}/{operation}.
    :vartype name: str
    :ivar display: The object that describes the operation.
    :vartype display: ~azure.mgmt.redhatopenshift.v2022_04_01.models.Display
    :ivar origin: Sources of requests to this operation.  Comma separated list with valid values
     user or system, e.g. "user,system".
    :vartype origin: str
    """

    _attribute_map = {
        'name': {'key': 'name', 'type': 'str'},
        'display': {'key': 'display', 'type': 'Display'},
        'origin': {'key': 'origin', 'type': 'str'},
    }

    def __init__(
        self,
        **kwargs
    ):
        """
        :keyword name: Operation name: {provider}/{resource}/{operation}.
        :paramtype name: str
        :keyword display: The object that describes the operation.
        :paramtype display: ~azure.mgmt.redhatopenshift.v2022_04_01.models.Display
        :keyword origin: Sources of requests to this operation.  Comma separated list with valid values
         user or system, e.g. "user,system".
        :paramtype origin: str
        """
        super(Operation, self).__init__(**kwargs)
        self.name = kwargs.get('name', None)
        self.display = kwargs.get('display', None)
        self.origin = kwargs.get('origin', None)


class OperationList(msrest.serialization.Model):
    """OperationList represents an RP operation list.

    :ivar value: List of operations supported by the resource provider.
    :vartype value: list[~azure.mgmt.redhatopenshift.v2022_04_01.models.Operation]
    :ivar next_link: The link used to get the next page of operations.
    :vartype next_link: str
    """

    _attribute_map = {
        'value': {'key': 'value', 'type': '[Operation]'},
        'next_link': {'key': 'nextLink', 'type': 'str'},
    }

    def __init__(
        self,
        **kwargs
    ):
        """
        :keyword value: List of operations supported by the resource provider.
        :paramtype value: list[~azure.mgmt.redhatopenshift.v2022_04_01.models.Operation]
        :keyword next_link: The link used to get the next page of operations.
        :paramtype next_link: str
        """
        super(OperationList, self).__init__(**kwargs)
        self.value = kwargs.get('value', None)
        self.next_link = kwargs.get('next_link', None)


class ServicePrincipalProfile(msrest.serialization.Model):
    """ServicePrincipalProfile represents a service principal profile.

    :ivar client_id: The client ID used for the cluster.
    :vartype client_id: str
    :ivar client_secret: The client secret used for the cluster.
    :vartype client_secret: str
    """

    _attribute_map = {
        'client_id': {'key': 'clientId', 'type': 'str'},
        'client_secret': {'key': 'clientSecret', 'type': 'str'},
    }

    def __init__(
        self,
        **kwargs
    ):
        """
        :keyword client_id: The client ID used for the cluster.
        :paramtype client_id: str
        :keyword client_secret: The client secret used for the cluster.
        :paramtype client_secret: str
        """
        super(ServicePrincipalProfile, self).__init__(**kwargs)
        self.client_id = kwargs.get('client_id', None)
        self.client_secret = kwargs.get('client_secret', None)


class SystemData(msrest.serialization.Model):
    """Metadata pertaining to creation and last modification of the resource.

    :ivar created_by: The identity that created the resource.
    :vartype created_by: str
    :ivar created_by_type: The type of identity that created the resource. Possible values include:
     "User", "Application", "ManagedIdentity", "Key".
    :vartype created_by_type: str or ~azure.mgmt.redhatopenshift.v2022_04_01.models.CreatedByType
    :ivar created_at: The timestamp of resource creation (UTC).
    :vartype created_at: ~datetime.datetime
    :ivar last_modified_by: The identity that last modified the resource.
    :vartype last_modified_by: str
    :ivar last_modified_by_type: The type of identity that last modified the resource. Possible
     values include: "User", "Application", "ManagedIdentity", "Key".
    :vartype last_modified_by_type: str or
     ~azure.mgmt.redhatopenshift.v2022_04_01.models.CreatedByType
    :ivar last_modified_at: The type of identity that last modified the resource.
    :vartype last_modified_at: ~datetime.datetime
    """

    _attribute_map = {
        'created_by': {'key': 'createdBy', 'type': 'str'},
        'created_by_type': {'key': 'createdByType', 'type': 'str'},
        'created_at': {'key': 'createdAt', 'type': 'iso-8601'},
        'last_modified_by': {'key': 'lastModifiedBy', 'type': 'str'},
        'last_modified_by_type': {'key': 'lastModifiedByType', 'type': 'str'},
        'last_modified_at': {'key': 'lastModifiedAt', 'type': 'iso-8601'},
    }

    def __init__(
        self,
        **kwargs
    ):
        """
        :keyword created_by: The identity that created the resource.
        :paramtype created_by: str
        :keyword created_by_type: The type of identity that created the resource. Possible values
         include: "User", "Application", "ManagedIdentity", "Key".
        :paramtype created_by_type: str or ~azure.mgmt.redhatopenshift.v2022_04_01.models.CreatedByType
        :keyword created_at: The timestamp of resource creation (UTC).
        :paramtype created_at: ~datetime.datetime
        :keyword last_modified_by: The identity that last modified the resource.
        :paramtype last_modified_by: str
        :keyword last_modified_by_type: The type of identity that last modified the resource. Possible
         values include: "User", "Application", "ManagedIdentity", "Key".
        :paramtype last_modified_by_type: str or
         ~azure.mgmt.redhatopenshift.v2022_04_01.models.CreatedByType
        :keyword last_modified_at: The type of identity that last modified the resource.
        :paramtype last_modified_at: ~datetime.datetime
        """
        super(SystemData, self).__init__(**kwargs)
        self.created_by = kwargs.get('created_by', None)
        self.created_by_type = kwargs.get('created_by_type', None)
        self.created_at = kwargs.get('created_at', None)
        self.last_modified_by = kwargs.get('last_modified_by', None)
        self.last_modified_by_type = kwargs.get('last_modified_by_type', None)
        self.last_modified_at = kwargs.get('last_modified_at', None)


class WorkerProfile(msrest.serialization.Model):
    """WorkerProfile represents a worker profile.

    :ivar name: The worker profile name.
    :vartype name: str
    :ivar vm_size: The size of the worker VMs.
    :vartype vm_size: str
    :ivar disk_size_gb: The disk size of the worker VMs.
    :vartype disk_size_gb: int
    :ivar subnet_id: The Azure resource ID of the worker subnet.
    :vartype subnet_id: str
    :ivar count: The number of worker VMs.
    :vartype count: int
    :ivar encryption_at_host: Whether master virtual machines are encrypted at host. Possible
     values include: "Disabled", "Enabled".
    :vartype encryption_at_host: str or
     ~azure.mgmt.redhatopenshift.v2022_04_01.models.EncryptionAtHost
    :ivar disk_encryption_set_id: The resource ID of an associated DiskEncryptionSet, if
     applicable.
    :vartype disk_encryption_set_id: str
    """

    _attribute_map = {
        'name': {'key': 'name', 'type': 'str'},
        'vm_size': {'key': 'vmSize', 'type': 'str'},
        'disk_size_gb': {'key': 'diskSizeGB', 'type': 'int'},
        'subnet_id': {'key': 'subnetId', 'type': 'str'},
        'count': {'key': 'count', 'type': 'int'},
        'encryption_at_host': {'key': 'encryptionAtHost', 'type': 'str'},
        'disk_encryption_set_id': {'key': 'diskEncryptionSetId', 'type': 'str'},
    }

    def __init__(
        self,
        **kwargs
    ):
        """
        :keyword name: The worker profile name.
        :paramtype name: str
        :keyword vm_size: The size of the worker VMs.
        :paramtype vm_size: str
        :keyword disk_size_gb: The disk size of the worker VMs.
        :paramtype disk_size_gb: int
        :keyword subnet_id: The Azure resource ID of the worker subnet.
        :paramtype subnet_id: str
        :keyword count: The number of worker VMs.
        :paramtype count: int
        :keyword encryption_at_host: Whether master virtual machines are encrypted at host. Possible
         values include: "Disabled", "Enabled".
        :paramtype encryption_at_host: str or
         ~azure.mgmt.redhatopenshift.v2022_04_01.models.EncryptionAtHost
        :keyword disk_encryption_set_id: The resource ID of an associated DiskEncryptionSet, if
         applicable.
        :paramtype disk_encryption_set_id: str
        """
        super(WorkerProfile, self).__init__(**kwargs)
        self.name = kwargs.get('name', None)
        self.vm_size = kwargs.get('vm_size', None)
        self.disk_size_gb = kwargs.get('disk_size_gb', None)
        self.subnet_id = kwargs.get('subnet_id', None)
        self.count = kwargs.get('count', None)
        self.encryption_at_host = kwargs.get('encryption_at_host', None)
        self.disk_encryption_set_id = kwargs.get('disk_encryption_set_id', None)

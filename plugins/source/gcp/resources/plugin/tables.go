package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/aiplatform"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/apigateway"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/apikeys"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/appengine"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/artifactregistry"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/baremetalsolution"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/batch"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/beyondcorp"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/bigquery"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/bigtableadmin"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/billing"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/binaryauthorization"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/certificatemanager"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/clouddeploy"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/clouderrorreporting"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/cloudiot"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/cloudresourcemanager"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/cloudscheduler"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/cloudsupport"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/compute"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/container"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/containeranalysis"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/dns"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/domains"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/functions"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/iam"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/kms"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/livestream"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/logging"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/monitoring"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/networkconnectivity"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/redis"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/resourcemanager"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/run"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/secretmanager"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/securitycenter"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/serviceusage"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/sql"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/storage"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/translate"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/videotranscoder"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/vision"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/vmmigration"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/vpcaccess"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/websecurityscanner"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/resources/services/workflows"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func PluginAutoGeneratedTables() schema.Tables {
	tables := []*schema.Table{
		aiplatform.JobLocations(),
		aiplatform.DatasetLocations(),
		aiplatform.EndpointLocations(),
		aiplatform.FeaturestoreLocations(),
		aiplatform.IndexendpointLocations(),
		aiplatform.IndexLocations(),
		aiplatform.MetadataLocations(),
		aiplatform.ModelLocations(),
		aiplatform.PipelineLocations(),
		aiplatform.SpecialistpoolLocations(),
		aiplatform.VizierLocations(),
		aiplatform.TensorboardLocations(),
		aiplatform.Operations(),
		apigateway.Apis(),
		apigateway.Gateways(),
		apikeys.Keys(),
		appengine.Apps(),
		appengine.Services(),
		appengine.AuthorizedCertificates(),
		appengine.AuthorizedDomains(),
		appengine.DomainMappings(),
		appengine.FirewallIngressRules(),
		artifactregistry.Locations(),
		baremetalsolution.Instances(),
		baremetalsolution.Networks(),
		baremetalsolution.NfsShares(),
		baremetalsolution.Volumes(),
		batch.Jobs(),
		batch.TaskGroups(),
		beyondcorp.AppConnections(),
		beyondcorp.AppConnectors(),
		beyondcorp.AppGateways(),
		beyondcorp.ClientConnectorServices(),
		beyondcorp.ClientGateways(),
		bigquery.Datasets(),
		bigtableadmin.Instances(),
		billing.BillingAccounts(),
		billing.Services(),
		binaryauthorization.Assertors(),
		certificatemanager.CertificateIssuanceConfigs(),
		certificatemanager.CertificateMaps(),
		certificatemanager.Certificates(),
		certificatemanager.DnsAuthorizations(),
		clouddeploy.Targets(),
		clouddeploy.DeliveryPipelines(),
		clouderrorreporting.ErrorGroupStats(),
		cloudiot.DeviceRegistries(),
		cloudresourcemanager.Organizations(),
		cloudscheduler.Locations(),
		cloudsupport.Cases(),
		compute.Addresses(),
		compute.Autoscalers(),
		compute.BackendServices(),
		compute.DiskTypes(),
		compute.Disks(),
		compute.ForwardingRules(),
		compute.Instances(),
		compute.SslCertificates(),
		compute.Subnetworks(),
		compute.TargetHttpProxies(),
		compute.UrlMaps(),
		compute.VpnGateways(),
		compute.VpnTunnels(),
		compute.TargetVpnGateways(),
		compute.InstanceGroups(),
		compute.Images(),
		compute.Firewalls(),
		compute.Networks(),
		compute.SslPolicies(),
		compute.Interconnects(),
		compute.TargetSslProxies(),
		compute.Projects(),
		compute.Routers(),
		compute.Routes(),
		compute.Zones(),
		container.Clusters(),
		containeranalysis.Occurrences(),
		dns.Policies(),
		dns.ManagedZones(),
		domains.Registrations(),
		functions.Functions(),
		iam.Roles(),
		iam.ServiceAccounts(),
		iam.DenyPolicies(),
		kms.Locations(),
		livestream.Channels(),
		livestream.Inputs(),
		logging.Metrics(),
		logging.Sinks(),
		monitoring.AlertPolicies(),
		networkconnectivity.Locations(),
		redis.Instances(),
		resourcemanager.Folders(),
		resourcemanager.Projects(),
		resourcemanager.ProjectPolicies(),
		resourcemanager.OrganizationTagKeys(),
		resourcemanager.ProjectTagKeys(),
		resourcemanager.ProjectTagBindings(),
		run.Locations(),
		secretmanager.Secrets(),
		securitycenter.ProjectFindings(),
		securitycenter.OrganizationFindings(),
		securitycenter.FolderFindings(),
		serviceusage.Services(),
		services.Projects(),
		sql.Instances(),
		storage.Buckets(),
		translate.Glossaries(),
		videotranscoder.Jobs(),
		videotranscoder.JobTemplates(),
		vision.Products(),
		vmmigration.Groups(),
		vmmigration.Sources(),
		vmmigration.TargetProjects(),
		vpcaccess.Locations(),
		websecurityscanner.ScanConfigs(),
		workflows.Workflows(),
	}
	if err := transformers.TransformTables(tables); err != nil {
		panic(err)
	}
	if err := transformers.Apply(tables, titleTransformer); err != nil {
		panic(err)
	}
	for _, table := range tables {
		schema.AddCqIDs(table)
	}
	return tables
}

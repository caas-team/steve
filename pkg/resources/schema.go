package resources

import (
	"context"

	corecontrollers "github.com/rancher/wrangler/pkg/generated/controllers/core/v1"
	"gitlab.devops.telekom.de/caas/rancher/apiserver/pkg/store/apiroot"
	"gitlab.devops.telekom.de/caas/rancher/apiserver/pkg/subscribe"
	"gitlab.devops.telekom.de/caas/rancher/apiserver/pkg/types"
	"gitlab.devops.telekom.de/caas/rancher/steve/pkg/accesscontrol"
	"gitlab.devops.telekom.de/caas/rancher/steve/pkg/client"
	"gitlab.devops.telekom.de/caas/rancher/steve/pkg/clustercache"
	"gitlab.devops.telekom.de/caas/rancher/steve/pkg/resources/apigroups"
	"gitlab.devops.telekom.de/caas/rancher/steve/pkg/resources/cluster"
	"gitlab.devops.telekom.de/caas/rancher/steve/pkg/resources/common"
	"gitlab.devops.telekom.de/caas/rancher/steve/pkg/resources/counts"
	"gitlab.devops.telekom.de/caas/rancher/steve/pkg/resources/formatters"
	"gitlab.devops.telekom.de/caas/rancher/steve/pkg/resources/userpreferences"
	"gitlab.devops.telekom.de/caas/rancher/steve/pkg/schema"
	steveschema "gitlab.devops.telekom.de/caas/rancher/steve/pkg/schema"
	"gitlab.devops.telekom.de/caas/rancher/steve/pkg/stores/proxy"
	"gitlab.devops.telekom.de/caas/rancher/steve/pkg/summarycache"
	"k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/client-go/discovery"
)

func DefaultSchemas(ctx context.Context, baseSchema *types.APISchemas, ccache clustercache.ClusterCache,
	cg proxy.ClientGetter, schemaFactory steveschema.Factory, serverVersion string) error {
	counts.Register(baseSchema, ccache)
	subscribe.Register(baseSchema, func(apiOp *types.APIRequest) *types.APISchemas {
		user, ok := request.UserFrom(apiOp.Context())
		if ok {
			schemas, err := schemaFactory.Schemas(user)
			if err == nil {
				return schemas
			}
		}
		return apiOp.Schemas
	}, serverVersion)
	apiroot.Register(baseSchema, []string{"v1"}, "proxy:/apis")
	cluster.Register(ctx, baseSchema, cg, schemaFactory)
	userpreferences.Register(baseSchema)
	return nil
}

func DefaultSchemaTemplates(cf *client.Factory,
	baseSchemas *types.APISchemas,
	summaryCache *summarycache.SummaryCache,
	lookup accesscontrol.AccessSetLookup,
	discovery discovery.DiscoveryInterface,
	namespaceCache corecontrollers.NamespaceCache) []schema.Template {
	return []schema.Template{
		common.DefaultTemplate(cf, summaryCache, lookup, namespaceCache),
		apigroups.Template(discovery),
		{
			ID:        "configmap",
			Formatter: formatters.DropHelmData,
		},
		{
			ID:        "secret",
			Formatter: formatters.DropHelmData,
		},
		{
			ID:        "pod",
			Formatter: formatters.Pod,
		},
		{
			ID: "management.cattle.io.cluster",
			Customize: func(apiSchema *types.APISchema) {
				cluster.AddApply(baseSchemas, apiSchema)
			},
		},
	}
}

package resources

import (
	"context"

	"github.com/caas-team/apiserver/pkg/store/apiroot"
	"github.com/caas-team/apiserver/pkg/subscribe"
	"github.com/caas-team/apiserver/pkg/types"
	"github.com/caas-team/steve/pkg/accesscontrol"
	"github.com/caas-team/steve/pkg/client"
	"github.com/caas-team/steve/pkg/clustercache"
	"github.com/caas-team/steve/pkg/resources/apigroups"
	"github.com/caas-team/steve/pkg/resources/cluster"
	"github.com/caas-team/steve/pkg/resources/common"
	"github.com/caas-team/steve/pkg/resources/counts"
	"github.com/caas-team/steve/pkg/resources/formatters"
	"github.com/caas-team/steve/pkg/resources/userpreferences"
	"github.com/caas-team/steve/pkg/schema"
	"github.com/caas-team/steve/pkg/stores/proxy"
	"github.com/caas-team/steve/pkg/summarycache"
	corecontrollers "github.com/rancher/wrangler/v2/pkg/generated/controllers/core/v1"
	"k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/client-go/discovery"
)

func DefaultSchemas(ctx context.Context, baseSchema *types.APISchemas, ccache clustercache.ClusterCache,
	cg proxy.ClientGetter, schemaFactory schema.Factory, serverVersion string) error {
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

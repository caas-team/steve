package handler

import (
	"github.com/caas-team/steve/pkg/attributes"
	"github.com/caas-team/steve/pkg/schema"
	"github.com/gorilla/mux"
	"github.com/rancher/apiserver/pkg/types"
)

func k8sAPI(sf schema.Factory, apiOp *types.APIRequest) {
	vars := mux.Vars(apiOp.Request)
	apiOp.Name = vars["name"]
	apiOp.Type = vars["type"]

	nOrN := vars["nameorns"]
	if nOrN != "" {
		schema := apiOp.Schemas.LookupSchema(apiOp.Type)
		if attributes.Namespaced(schema) {
			vars["namespace"] = nOrN
		} else {
			vars["name"] = nOrN
		}
	}

	if namespace := vars["namespace"]; namespace != "" {
		apiOp.Namespace = namespace
	}
}

func apiRoot(sf schema.Factory, apiOp *types.APIRequest) {
	apiOp.Type = "apiRoot"
}

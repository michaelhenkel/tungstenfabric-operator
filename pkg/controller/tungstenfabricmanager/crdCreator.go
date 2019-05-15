package tungstenfabricmanager

import (
	"strings"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PropertiesStruct struct {
	Properties map[string]apiextensionsv1beta1.JSONSchemaProps
}

func getProperties(crdName string) PropertiesStruct {

	var PropertiesStructMap map[string]PropertiesStruct
	PropertiesStructMap = make(map[string]PropertiesStruct)

	vrouterPropertiesStruct := PropertiesStruct{
		Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
			"aaamode": {
				Type: "string",
			},
		},
	}

	PropertiesStructMap["Vrouter"] = vrouterPropertiesStruct
	
	return PropertiesStructMap[crdName]
}


func (r *ReconcileTungstenfabricManager) CreateCrd(crdName string,
	crdVersion string,
	crdGroup string,
	crdNamespace string) *apiextensionsv1beta1.CustomResourceDefinition {

	propertiesStruct := getProperties(crdName)


	kind := crdName
	listKind := crdName + "List"
	singular := strings.ToLower(kind)
	plural := singular + "s"

	scope := apiextensionsv1beta1.NamespaceScoped
	crd := &apiextensionsv1beta1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: plural + "." + crdGroup,
			Namespace: crdNamespace,
		},
		Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
			Group: crdGroup,
			Version: crdVersion,
			Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
				Plural: plural,
				Singular: singular,
				Kind: kind,
				ListKind: listKind,
			},
			Scope: scope,
			Validation: &apiextensionsv1beta1.CustomResourceValidation{
				OpenAPIV3Schema: &apiextensionsv1beta1.JSONSchemaProps{
					Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
						"apiVersion": {
							Description: crdVersion,
							Type: "string",
						},
						"kind": {
							Description: crdName,
							Type: "string",
						},
						"metadata": {
							Type: "object",
						},
						"spec": {
							Type: "object",
							Properties: propertiesStruct.Properties,
						},
						"status": {
							Type: "object",
						},
					},
				},
			},
		},
	}
	return crd
}
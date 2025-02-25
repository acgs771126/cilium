// +build !ignore_autogenerated

/*
Copyright2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by helpgen. DO NOT EDIT.

package markers

import (
	"sigs.k8s.io/controller-tools/pkg/markers"
)

func (Default) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "CRD validation",
		DetailedHelp: markers.DetailedHelp{
			Summary: "sets the default value for this field. ",
			Details: "A default value will be accepted as any value valid for the field. Formatting for common types include: boolean: `true`, string: `Cluster`, numerical: `1.24`, array: `{1,2}`, object: `{policy: \"delete\"}`). Defaults should be defined in pruned form, and only best-effort validation will be performed. Full validation of a default requires submission of the containing CRD to an apiserver.",
		},
		FieldHelp: map[string]markers.DetailedHelp{
			"Value": {
				Summary: "",
				Details: "",
			},
		},
	}
}

func (DeprecatedVersion) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "CRD",
		DetailedHelp: markers.DetailedHelp{
			Summary: "marks this version as deprecated.",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{
			"Warning": {
				Summary: "message to be shown on the deprecated version",
				Details: "",
			},
		},
	}
}

func (Enum) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "CRD validation",
		DetailedHelp: markers.DetailedHelp{
			Summary: "specifies that this (scalar) field is restricted to the *exact* values specified here.",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (ExclusiveMaximum) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "CRD validation",
		DetailedHelp: markers.DetailedHelp{
			Summary: "indicates that the maximum is \"up to\" but not including that value.",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (ExclusiveMinimum) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "CRD validation",
		DetailedHelp: markers.DetailedHelp{
			Summary: "indicates that the minimum is \"up to\" but not including that value.",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (Format) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "CRD validation",
		DetailedHelp: markers.DetailedHelp{
			Summary: "specifies additional \"complex\" formatting for this field. ",
			Details: "For example, a date-time field would be marked as \"type: string\" and \"format: date-time\".",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (ListMapKey) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "CRD processing",
		DetailedHelp: markers.DetailedHelp{
			Summary: "specifies the keys to map listTypes. ",
			Details: "It indicates the index of a map list. They can be repeated if multiple keys must be used. It can only be used when ListType is set to map, and the keys should be scalar types.",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (ListType) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "CRD processing",
		DetailedHelp: markers.DetailedHelp{
			Summary: "specifies the type of data-structure that the list represents (map, set, atomic). ",
			Details: "Possible data-structure types of a list are: \n - \"map\": it needs to have a key field, which will be used to build an   associative list. A typical example is a the pod container list,   which is indexed by the container name. \n - \"set\": Fields need to be \"scalar\", and there can be only one   occurrence of each. \n - \"atomic\": All the fields in the list are treated as a single value,   are typically manipulated together by the same actor.",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (MapType) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "CRD processing",
		DetailedHelp: markers.DetailedHelp{
			Summary: "specifies the level of atomicity of the map; i.e. whether each item in the map is independent of the others, or all fields are treated as a single unit. ",
			Details: "Possible values: \n - \"granular\": items in the map are independent of each other,   and can be manipulated by different actors.   This is the default behavior. \n - \"atomic\": all fields are treated as one unit.   Any changes have to replace the entire map.",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (MaxItems) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "CRD validation",
		DetailedHelp: markers.DetailedHelp{
			Summary: "specifies the maximum length for this list.",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (MaxLength) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "CRD validation",
		DetailedHelp: markers.DetailedHelp{
			Summary: "specifies the maximum length for this string.",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (MaxProperties) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "CRD validation",
		DetailedHelp: markers.DetailedHelp{
			Summary: "restricts the number of keys in an object",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (Maximum) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "CRD validation",
		DetailedHelp: markers.DetailedHelp{
			Summary: "specifies the maximum numeric value that this field can have.",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (MinItems) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "CRD validation",
		DetailedHelp: markers.DetailedHelp{
			Summary: "specifies the minimun length for this list.",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (MinLength) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "CRD validation",
		DetailedHelp: markers.DetailedHelp{
			Summary: "specifies the minimum length for this string.",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (MinProperties) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "CRD validation",
		DetailedHelp: markers.DetailedHelp{
			Summary: "restricts the number of keys in an object",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (Minimum) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "CRD validation",
		DetailedHelp: markers.DetailedHelp{
			Summary: "specifies the minimum numeric value that this field can have. Negative integers are supported.",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (MultipleOf) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "CRD validation",
		DetailedHelp: markers.DetailedHelp{
			Summary: "specifies that this field must have a numeric value that's a multiple of this one.",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (Nullable) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "CRD validation",
		DetailedHelp: markers.DetailedHelp{
			Summary: "marks this field as allowing the \"null\" value. ",
			Details: "This is often not necessary, but may be helpful with custom serialization.",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (Pattern) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "CRD validation",
		DetailedHelp: markers.DetailedHelp{
			Summary: "specifies that this string must match the given regular expression.",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (PrintColumn) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "CRD",
		DetailedHelp: markers.DetailedHelp{
			Summary: "adds a column to \"kubectl get\" output for this CRD.",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{
			"Name": {
				Summary: "specifies the name of the column.",
				Details: "",
			},
			"Type": {
				Summary: "indicates the type of the column. ",
				Details: "It may be any OpenAPI data type listed at https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#data-types.",
			},
			"JSONPath": {
				Summary: "specifies the jsonpath expression used to extract the value of the column.",
				Details: "",
			},
			"Description": {
				Summary: "specifies the help/description for this column.",
				Details: "",
			},
			"Format": {
				Summary: "specifies the format of the column. ",
				Details: "It may be any OpenAPI data format corresponding to the type, listed at https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#data-types.",
			},
			"Priority": {
				Summary: "indicates how important it is that this column be displayed. ",
				Details: "Lower priority (*higher* numbered) columns will be hidden if the terminal width is too small.",
			},
		},
	}
}

func (Resource) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "CRD",
		DetailedHelp: markers.DetailedHelp{
			Summary: "configures naming and scope for a CRD.",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{
			"Path": {
				Summary: "specifies the plural \"resource\" for this CRD. ",
				Details: "It generally corresponds to a plural, lower-cased version of the Kind. See https://book.kubebuilder.io/cronjob-tutorial/gvks.html.",
			},
			"ShortName": {
				Summary: "specifies aliases for this CRD. ",
				Details: "Short names are often used when people have work with your resource over and over again.  For instance, \"rs\" for \"replicaset\" or \"crd\" for customresourcedefinition.",
			},
			"Categories": {
				Summary: "specifies which group aliases this resource is part of. ",
				Details: "Group aliases are used to work with groups of resources at once. The most common one is \"all\" which covers about a third of the base resources in Kubernetes, and is generally used for \"user-facing\" resources.",
			},
			"Singular": {
				Summary: "overrides the singular form of your resource. ",
				Details: "The singular form is otherwise defaulted off the plural (path).",
			},
			"Scope": {
				Summary: "overrides the scope of the CRD (Cluster vs Namespaced). ",
				Details: "Scope defaults to \"Namespaced\".  Cluster-scoped (\"Cluster\") resources don't exist in namespaces.",
			},
		},
	}
}

func (Schemaless) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "CRD validation",
		DetailedHelp: markers.DetailedHelp{
			Summary: "marks a field as being a schemaless object. ",
			Details: "Schemaless objects are not introspected, so you must provide any type and validation information yourself. One use for this tag is for embedding fields that hold JSONSchema typed objects. Because this field disables all type checking, it is recommended to be used only as a last resort.",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (SkipVersion) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "CRD",
		DetailedHelp: markers.DetailedHelp{
			Summary: "removes the particular version of the CRD from the CRDs spec. ",
			Details: "This is useful if you need to skip generating and listing version entries for 'internal' resource versions, which typically exist if using the Kubernetes upstream conversion-gen tool.",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (StorageVersion) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "CRD",
		DetailedHelp: markers.DetailedHelp{
			Summary: "marks this version as the \"storage version\" for the CRD for conversion. ",
			Details: "When conversion is enabled for a CRD (i.e. it's not a trivial-versions/single-version CRD), one version is set as the \"storage version\" to be stored in etcd.  Attempting to store any other version will result in conversion to the storage version via a conversion webhook.",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (StructType) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "CRD processing",
		DetailedHelp: markers.DetailedHelp{
			Summary: "specifies the level of atomicity of the struct; i.e. whether each field in the struct is independent of the others, or all fields are treated as a single unit. ",
			Details: "Possible values: \n - \"granular\": fields in the struct are independent of each other,   and can be manipulated by different actors.   This is the default behavior. \n - \"atomic\": all fields are treated as one unit.   Any changes have to replace the entire struct.",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (SubresourceScale) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "CRD",
		DetailedHelp: markers.DetailedHelp{
			Summary: "enables the \"/scale\" subresource on a CRD.",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{
			"SpecPath": {
				Summary: "specifies the jsonpath to the replicas field for the scale's spec.",
				Details: "",
			},
			"StatusPath": {
				Summary: "specifies the jsonpath to the replicas field for the scale's status.",
				Details: "",
			},
			"SelectorPath": {
				Summary: "specifies the jsonpath to the pod label selector field for the scale's status. ",
				Details: "The selector field must be the *string* form (serialized form) of a selector. Setting a pod label selector is necessary for your type to work with the HorizontalPodAutoscaler.",
			},
		},
	}
}

func (SubresourceStatus) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "CRD",
		DetailedHelp: markers.DetailedHelp{
			Summary: "enables the \"/status\" subresource on a CRD.",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (Type) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "CRD validation",
		DetailedHelp: markers.DetailedHelp{
			Summary: "overrides the type for this field (which defaults to the equivalent of the Go type). ",
			Details: "This generally must be paired with custom serialization.  For example, the metav1.Time field would be marked as \"type: string\" and \"format: date-time\".",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (UniqueItems) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "CRD validation",
		DetailedHelp: markers.DetailedHelp{
			Summary: "specifies that all items in this list must be unique.",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (UnservedVersion) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "CRD",
		DetailedHelp: markers.DetailedHelp{
			Summary: "does not serve this version. ",
			Details: "This is useful if you need to drop support for a version in favor of a newer version.",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (XEmbeddedResource) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "CRD validation",
		DetailedHelp: markers.DetailedHelp{
			Summary: "EmbeddedResource marks a fields as an embedded resource with apiVersion, kind and metadata fields. ",
			Details: "An embedded resource is a value that has apiVersion, kind and metadata fields. They are validated implicitly according to the semantics of the currently running apiserver. It is not necessary to add any additional schema for these field, yet it is possible. This can be combined with PreserveUnknownFields.",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (XPreserveUnknownFields) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "CRD processing",
		DetailedHelp: markers.DetailedHelp{
			Summary: "PreserveUnknownFields stops the apiserver from pruning fields which are not specified. ",
			Details: "By default the apiserver drops unknown fields from the request payload during the decoding step. This marker stops the API server from doing so. It affects fields recursively, but switches back to normal pruning behaviour if nested  properties or additionalProperties are specified in the schema. This can either be true or undefined. False is forbidden. \n NB: The kubebuilder:validation:XPreserveUnknownFields variant is deprecated in favor of the kubebuilder:pruning:PreserveUnknownFields variant.  They function identically.",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

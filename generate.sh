#!/bin/sh

go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config oapi.yaml https://demo.openhab.org/rest/spec

# Use a new type for EnrichedSemanticTagDTO instead of an alias. Kubebuilder will try to generate methods on
# json.RawMessage otherwise.
sed -i '' 's/type EnrichedSemanticTagDTO = json.RawMessage/type EnrichedSemanticTagDTO json.RawMessage/' api.go

# Mark raw JSON fields as schemaless for Kubebuilder.
sed -i '' '/json.RawMessage/i\
	\/\/ +kubebuilder:pruning:PreserveUnknownFields\
	\/\/ +kubebuilder:validation:Schemaless
' api.go

go run sigs.k8s.io/controller-tools/cmd/controller-gen +object

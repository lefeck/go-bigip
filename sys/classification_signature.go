package sys

// ClassificationSignatureConfigList holds a list of ClassificationSignature configuration.
type ClassificationSignatureConfigList struct {
	Items    []ClassificationSignatureConfig `json:"items"`
	Kind     string                          `json:"kind"`
	SelfLink string                          `json:"selflink"`
}

// ClassificationSignatureConfig holds the configuration of a single ClassificationSignature.
type ClassificationSignatureConfig struct {
}

// ClassificationSignatureEndpoint represents the REST resource for managing ClassificationSignature.
const ClassificationSignatureEndpoint = "/classification-signature"

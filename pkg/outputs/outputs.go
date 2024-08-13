package outputs

import (
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/gcp/gcswebsite"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
)

func PulumiOutputsToStackOutputsConverter(stackOutput auto.OutputMap,
	input *gcswebsite.GcsWebsiteStackInput) *gcswebsite.GcsWebsiteStackOutputs {
	return &gcswebsite.GcsWebsiteStackOutputs{}
}

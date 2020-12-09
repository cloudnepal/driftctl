// GENERATED, DO NOT EDIT THIS FILE
package aws

const AwsIamUserPolicyAttachmentResourceType = "aws_iam_user_policy_attachment"

type AwsIamUserPolicyAttachment struct {
	Id        string  `cty:"id" computed:"true"`
	PolicyArn *string `cty:"policy_arn"`
	User      *string `cty:"user"`
}

func (r *AwsIamUserPolicyAttachment) TerraformId() string {
	return r.Id
}

func (r *AwsIamUserPolicyAttachment) TerraformType() string {
	return AwsIamUserPolicyAttachmentResourceType
}
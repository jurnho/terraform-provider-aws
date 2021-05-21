package aws

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAwsDefaultTags() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAwsDefaultTagsRead,

		Schema: map[string]*schema.Schema{
			"tags": tagsSchemaComputed(),
		},
	}
}

func dataSourceAwsDefaultTagsRead(d *schema.ResourceData, meta interface{}) error {
	defaultTagsConfig := meta.(*AWSClient).DefaultTagsConfig
	ignoreTagsConfig := meta.(*AWSClient).IgnoreTagsConfig

	d.SetId(meta.(*AWSClient).partition)

	if defaultTagsConfig != nil && defaultTagsConfig.Tags != nil {
		if err := d.Set("tags", defaultTagsConfig.Tags.IgnoreAws().IgnoreConfig(ignoreTagsConfig).Map()); err != nil {
			return err
		}
	} else {
		d.Set("tags", nil)
	}

	return nil
}

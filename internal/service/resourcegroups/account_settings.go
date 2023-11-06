package resourcegroups

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/resourcegroups"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

const (
	AccountSettingsOperationTimeout = 5 * time.Minute
)

// @SDKResource("aws_resourcegroups_account_settings", name="AccountSettings")
func accountSettings() *schema.Resource {
	return &schema.Resource{
		CreateWithoutTimeout: accountSettingsCreate,
		ReadWithoutTimeout:   accountSettingsRead,
		DeleteWithoutTimeout: accountSettingsDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"GroupLifecycleEventsDesiredStatus": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ACTIVE", "INACTIVE"}, false),
			},
			"GroupLifecycleEventsStatus": {
				ValidateFunc: validation.StringInSlice([]string{"ACTIVE", "INACTIVE", "IN_PROGRESS", "ERROR"}, false),
			},
			"GroupLifecycleEventsStatusMessage": {
				ValidateFunc: validation.StringInSlice([]string{"ACTIVE", "INACTIVE"}, false),
			},
		},
	}
}

func accountSettingsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).ResourceGroupsClient(ctx)

	input := &resourcegroups.UpdateAccountSettingsInput{
		GroupLifecycleEventsDesiredStatus: d.Get("GroupLifecycleEventsDesiredStatus").(types.GroupLifecycleEventsDesiredStatus),
	}

	settings, err := conn.UpdateAccountSettings(ctx, input)

	if err != nil {
		return diag.Errorf("creating ResourceGroups Account Settings (%s)", err)
	}

	waitError := accountSettingsWaitForState(ctx, conn, types.GroupLifecycleEventsStatusActive)

	if waitError != nil {
		return diag.Errorf("creating ResourceGroups Account Settings: wait error: %s", err)
	}

	d.SetId("ResourceGroupsAcountSettings") //ResourceGroups Account Settings are at account-level (regional), no specific ARN/resource ID.

	d.Set("GroupLifecycleEventsDesiredStatus", settings.AccountSettings.GroupLifecycleEventsDesiredStatus)
	d.Set("GroupLifecycleEventsStatus", settings.AccountSettings.GroupLifecycleEventsStatus)
	d.Set("GroupLifecycleEventsStatusMessage", settings.AccountSettings.GroupLifecycleEventsStatusMessage)

	return nil
}

func accountSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).ResourceGroupsClient(ctx)

	input := &resourcegroups.GetAccountSettingsInput{}

	settings, err := conn.GetAccountSettings(ctx, input)

	if err != nil {
		return diag.Errorf("reading ResourceGroups Account Settings: %s", err)
	}

	if !d.IsNewResource() {
		d.SetId("")
		return nil
	}

	d.Set("GroupLifecycleEventsDesiredStatus", settings.AccountSettings.GroupLifecycleEventsDesiredStatus)
	d.Set("GroupLifecycleEventsStatus", settings.AccountSettings.GroupLifecycleEventsStatus)
	d.Set("GroupLifecycleEventsStatusMessage", settings.AccountSettings.GroupLifecycleEventsStatusMessage)

	return nil
}

func accountSettingsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).ResourceGroupsClient(ctx)

	input := &resourcegroups.UpdateAccountSettingsInput{
		GroupLifecycleEventsDesiredStatus: types.GroupLifecycleEventsDesiredStatusInactive,
	}

	_, err := conn.UpdateAccountSettings(ctx, input)

	if err != nil {
		return diag.Errorf("reading ResourceGroups Account Settings (%s)", err)
	}

	waitError := accountSettingsWaitForState(ctx, conn, types.GroupLifecycleEventsStatusInactive)

	if waitError != nil {
		return diag.Errorf("creating ResourceGroups Account Settings: wait error: %s", err)
	}

	d.SetId("")

	return nil
}

func accountSettingsWaitForState(ctx context.Context, conn *resourcegroups.Client, desiredState types.GroupLifecycleEventsStatus) error {

	stateConf := &retry.StateChangeConf{
		Pending: []string{string(types.GroupLifecycleEventsStatusInProgress)},
		Target:  []string{string(desiredState)},
		Refresh: refreshSettingsState(ctx, conn),
		Timeout: AccountSettingsOperationTimeout,
	}

	_, err := stateConf.WaitForStateContext(ctx)

	if err != nil {
		return err
	}

	return nil
}

func refreshSettingsState(ctx context.Context, conn *resourcegroups.Client) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		input := &resourcegroups.GetAccountSettingsInput{}

		settings, err := conn.GetAccountSettings(ctx, input)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return settings, string(settings.AccountSettings.GroupLifecycleEventsStatus), err
	}
}

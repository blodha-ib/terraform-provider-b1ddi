// Code generated by go-swagger; DO NOT EDIT.

package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// IpamsvcDHCPInfo DHCPInfo
//
// The __DHCPInfo__ object represents the DHCP information associated with an address object.
//
// swagger:model ipamsvcDHCPInfo
func schemaIpamsvcDHCPInfo() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// The DHCP host name associated with this client.
			// Read Only: true
			"client_hostname": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The DHCP host name associated with this client.",
			},

			// The hardware address associated with this client.
			// Read Only: true
			"client_hwaddr": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The hardware address associated with this client.",
			},

			// The ID associated with this client.
			// Read Only: true
			"client_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID associated with this client.",
			},

			// The timestamp at which the _state_, when set to _leased_, will be changed to _free_.
			// Read Only: true
			// Format: date-time
			"end": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp at which the _state_, when set to _leased_, will be changed to _free_.",
			},

			// The DHCP fingerprint for the associated lease.
			// Read Only: true
			"fingerprint": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The DHCP fingerprint for the associated lease.",
			},

			// The remaining time, in seconds, until which the _state_, when set to _leased_, will remain in that state.
			// Read Only: true
			"remain": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The remaining time, in seconds, until which the _state_, when set to _leased_, will remain in that state.",
			},

			// The timestamp at which _state_ was first set to _leased_.
			// Read Only: true
			// Format: date-time
			"start": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp at which _state_ was first set to _leased_.",
			},

			// Indicates the status of this IP address from a DHCP protocol standpoint as:
			//   * _none_: Address is not under DHCP control.
			//   * _free_: Address is under DHCP control but has no lease currently assigned.
			//   * _leased_: Address is under DHCP control and has a lease currently assigned. The lease details are contained in the matching _dhcp/lease_ resource.
			// Read Only: true
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Indicates the status of this IP address from a DHCP protocol standpoint as:\n  * _none_: Address is not under DHCP control.\n  * _free_: Address is under DHCP control but has no lease currently assigned.\n  * _leased_: Address is under DHCP control and has a lease currently assigned. The lease details are contained in the matching _dhcp/lease_ resource.",
			},

			// The timestamp at which the _state_ was last reported.
			// Read Only: true
			// Format: date-time
			"state_ts": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp at which the _state_ was last reported.",
			},
		},
	}
}

func flattenIpamsvcDHCPInfo(r *models.IpamsvcDHCPInfo) []interface{} {
	if r == nil {
		return nil
	}
	return []interface{}{
		map[string]interface{}{
			"client_hostname": r.ClientHostname,
			"client_hwaddr":   r.ClientHwaddr,
			"client_id":       r.ClientID,
			"end":             r.End,
			"fingerprint":     r.Fingerprint,
			"remain":          r.Remain,
			"start":           r.Start,
			"state":           r.State,
			"state_ts":        r.StateTs,
		},
	}
}
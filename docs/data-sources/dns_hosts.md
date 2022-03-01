# b1ddi_dns_hosts (Data Source)

## Example Usage

```terraform
terraform {
  required_providers {
    b1ddi = {
      version = "0.1"
      source  = "infobloxopen/b1ddi"
    }
  }
}

# Get all DNS hosts
data "b1ddi_dns_hosts" "all_hosts" {}

# Get DNS Host by name
data "b1ddi_dns_hosts" "dns_host_by_name" {
  filters = {
    name = "dns_host_name"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **filters** (Map of String) Configure a map of filters to be applied on the search result.
- **id** (String) The ID of this resource.

### Read-Only

- **results** (List of Object) List of DNS Hosts matching filters. (see [below for nested schema](#nestedatt--results))

<a id="nestedatt--results"></a>
### Nested Schema for `results`

Read-Only:

- **absolute_name** (String) Host FQDN.
- **address** (String) Host's primary IP Address.
- **anycast_addresses** (List of String) Anycast address configured to the host. Order is not significant.
- **associated_server** (List of Object) Host associated server configuration. (see [below for nested schema](#nestedobjatt--results--associated_server))
- **comment** (String) Host description.
- **current_version** (String) Host current version.
- **dfp** (Boolean) Below _dfp_ field is deprecated and not supported anymore.
  The indication whether or not BloxOne DDI DNS and BloxOne TD DFP are both active on the host will be migrated into the 
  new _dfp_service_ field.
- **dfp_service** (String) DFP service indicates whether or not BloxOne DDI DNS and BloxOne TD DFP are both active on the host. 
  If so, BloxOne DDI DNS will augment recursive queries and forward them to BloxOne TD DFP.
  
  Allowed values:
  * _unavailable_: BloxOne TD DFP application is not available,
  * _enabled_: BloxOne TD DFP application is available and enabled,
  * _disabled_: BloxOne TD DFP application is available but disabled.

- **id** (String) The resource identifier.
- **inheritance_sources** (List of Object) Optional. Inheritance configuration. (see [below for nested schema](#nestedobjatt--results--inheritance_sources))
- **kerberos_keys** (List of Object) Optional. _kerberos_keys_ contains a list of keys for GSS-TSIG signed dynamic updates. Defaults to empty. (see [below for nested schema](#nestedobjatt--results--kerberos_keys))

- **name** (String) Host display name.
- **ophid** (String) On-Prem Host ID.
- **protocol_absolute_name** (String) Host FQDN in punycode.
- **server** (String) The resource identifier.
- **site_id** (String) Host site ID.
- **tags** (Map of String) Host tagging specifics.

<a id="nestedobjatt--results--associated_server"></a>
### Nested Schema for `results.associated_server`

Read-Only:

- **name** (String) DNS server name.


<a id="nestedobjatt--results--inheritance_sources"></a>
### Nested Schema for `results.inheritance_sources`

Read-Only:

- **kerberos_keys** (List of Object) Optional. Field config for _kerberos_keys_ field from _Host_ object. (see [below for nested schema](#nestedobjatt--results--inheritance_sources--kerberos_keys))

<a id="nestedobjatt--results--inheritance_sources--kerberos_keys"></a>
### Nested Schema for `results.inheritance_sources.kerberos_keys`

Read-Only:

- **action** (String) Optional. Inheritance setting for a field.
  Defaults to _inherit_.
- **display_name** (String) Human-readable display name for the object referred to by _source_.
- **source** (String) The resource identifier.
- **value** (List of Object) Inherited value. (see [below for nested schema](#nestedobjatt--results--inheritance_sources--kerberos_keys--value))

<a id="nestedobjatt--results--inheritance_sources--kerberos_keys--value"></a>
### Nested Schema for `results.inheritance_sources.kerberos_keys.value`

Read-Only:

- **algorithm** (String) Encryption algorithm of the key in accordance with RFC 3961.
- **domain** (String) Kerberos realm of the principal.
- **key** (String) The resource identifier.
- **principal** (String) Kerberos principal associated with key.
- **uploaded_at** (String) Upload time for the key.
- **version** (Number) The version number (KVNO) of the key.




<a id="nestedobjatt--results--kerberos_keys"></a>
### Nested Schema for `results.kerberos_keys`

Read-Only:

- **algorithm** (String) Encryption algorithm of the key in accordance with RFC 3961.
- **domain** (String) Kerberos realm of the principal.
- **key** (String) The resource identifier.
- **principal** (String) Kerberos principal associated with key.
- **uploaded_at** (String) Upload time for the key.
- **version** (Number) The version number (KVNO) of the key.
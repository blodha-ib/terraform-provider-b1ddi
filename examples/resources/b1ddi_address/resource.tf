terraform {
  required_providers {
    b1ddi = {
      version = "0.1"
      source  = "github.com/infobloxopen/b1ddi"
    }
  }
}

resource "b1ddi_ip_space" "example_tf_space" {
  name    = "example_tf_space"
  comment = "Example IP space created by the terraform provider"
}

resource "b1ddi_subnet" "example_tf_subnet" {
  name    = "example_tf_subnet"
  space   = b1ddi_ip_space.example_tf_space.id
  address = "192.168.1.0"
  cidr    = 24
  comment = "Example Subnet created by the terraform provider"
}

# Address can be specified manually via the address field
resource "b1ddi_address" "example_tf_address" {
  address    = "192.168.1.45"
  comment    = "Example Address created by the terraform provider"
  space      = b1ddi_ip_space.example_tf_space.id
  depends_on = [b1ddi_subnet.example_tf_subnet]
}

# Address can be allocated from a subnet via the parent field
resource "b1ddi_address" "next_available_ip" {
  parent  = b1ddi_subnet.example_tf_subnet.id
  comment = "Example Address automatically allocated from the subnet"
  space   = b1ddi_ip_space.example_tf_space.id
}
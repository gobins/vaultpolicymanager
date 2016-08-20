path "/sys/mounts" {
  capabilities = ["read"]
}

path "/sys/seal" {
  capabilities = ["sudo","update"]
}

path "/sys/mounts/*" {
  capabilities = ["read","update","create","delete"]
}

path "/sys/remount" {
  capabilities = ["sudo","update"]
}

path "/sys/auth" {
  capabilities = ["read"]
}

path "/sys/auth/*" {
  capabilities = ["sudo","create","update","delete"]
}

path "/sys/policy" {
  capabilities = ["read"]
}

path "/sys/policy/*" {
  capabilities = ["sudo","create","delete","update"]
}

path "/sys/capabilities" {
  capabilities = ["sudo", "update"]
}

path "/sys/capabilities-self" {
  capabilities = ["sudo", "update"]
}

path "/sys/capabilities-accesor" {
  capabilities = ["sudo", "update"]
}

path "/sys/audit" {
  capabilities = ["sudo","read"]
}

path "/sys/audit/*" {
  capabilities = ["sudo","update","delete"]
}

path "/sys/revoke" {
  capabilities = ["create"]
}

path "/sys/key-status" {
  capabilities = ["read"]
}


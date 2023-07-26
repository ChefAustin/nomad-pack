# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

app {
  url    = ""
  author = "Nomad Team"
}

pack {
  name        = "deps_test"
  description = "This pack tests dependencies"
  url         = "github.com/hashicorp/nomad-pack/fixtures/test_registry/packs/deps_test"
  version     = "0.0.1"
}

dependency "child1" {}
dependency "child2" {}

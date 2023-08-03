# Copyright 2023 Jimmi Dyson.
# SPDX-License-Identifier: Apache-2.0

.PHONY: dev.run-on-kind
dev.run-on-kind: export KUBECONFIG := $(KIND_KUBECONFIG)
dev.run-on-kind: kind.create
	if ! kubectl get namespace cert-manager &>/dev/null; then cmctl experimental install; fi
ifndef SKIP_BUILD
	$(MAKE) release-snapshot
endif
	kind load docker-image --name $(KIND_CLUSTER_NAME) \
		$$(gojq -r '.[] | select(.type=="Docker Image") | select(.goarch=="$(GOARCH)") | .name' dist/artifacts.json)
	helm upgrade --install dummy-controller ./charts/dummy-controller \
		--set-string image.tag=$$(gojq -r .version dist/metadata.json) \
		--wait --wait-for-jobs
	kubectl rollout restart deployment dummy-controller
	kubectl rollout status deployment dummy-controller

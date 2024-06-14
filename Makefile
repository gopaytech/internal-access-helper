.PHONY: helm.create.releases
helm.create.releases:
	helm package charts/internal-access-helper --destination charts/releases
	helm repo index charts/releases

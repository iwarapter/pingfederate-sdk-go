# Makefile

pf-init:
	@docker run --rm -d --hostname pingfederate --name pingfederate -e PING_IDENTITY_DEVOPS_KEY=$(PING_IDENTITY_DEVOPS_KEY) -e PING_IDENTITY_DEVOPS_USER=$(PING_IDENTITY_DEVOPS_USER) -e PING_IDENTITY_ACCEPT_EULA=YES --publish 9999:9999 pingidentity/pingfederate:10.0.2-edge


test:
	@rm -f report.json coverage.out
	@go test -mod=vendor ./... -v -trimpath -coverprofile=coverage.out && go tool cover -func=coverage.out

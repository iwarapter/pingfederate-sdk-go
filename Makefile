# Makefile

pf-init:
	@docker run --rm -d --hostname pingfederate --name pingfederate -e PING_IDENTITY_DEVOPS_KEY=$(PING_IDENTITY_DEVOPS_KEY) -e PING_IDENTITY_DEVOPS_USER=$(PING_IDENTITY_DEVOPS_USER) -e PING_IDENTITY_ACCEPT_EULA=YES --publish 9999:9999 pingidentity/pingfederate:10.0.2-edge


test:
	@rm -f report.json coverage.out
	@go test ./... -v -coverprofile=coverage.out -json > report.json

sonar:
	@sonar-scanner \
		-Dsonar.projectKey=github.com.iwarapter.pingfederate-sdk-go \
		-Dsonar.sources=. \
		-Dsonar.host.url=http://localhost:9001 \
		-Dsonar.login=28d86a90f2e4ae9563b4501cbc99de7522219c88 \
		-Dsonar.go.coverage.reportPaths=coverage.out \
		-Dsonar.go.tests.reportPaths=report.json \
		-Dsonar.exclusions=vendor \
		-Dsonar.tests="." \
		-Dsonar.test.inclusions="**/*_test.go"
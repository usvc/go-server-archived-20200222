deps:
	@GO111MODULE=on go mod vendor
release.github:
	@if [ "${GITHUB_REPOSITORY_URL}" = "" ]; then exit 1; fi;
	@git remote set-url origin ${GITHUB_REPOSITORY_URL}
	@git checkout --f master
	@git fetch
	@$(MAKE) version.bump VERSION=${BUMP}
	@git tag "v$$($(MAKE) version.get | grep '[0-9]*\.[0-9]*\.[0-9]*')"
	@git commit -m "released $$($(MAKE) version.get | grep '[0-9]*\.[0-9]*\.[0-9]*') [skip ci]"
	@git push
	@git push --tags
secrets:
	@mkdir -p secrets
	@openssl genrsa \
		-out secrets/tls.key \
		4192
	@openssl req -new -x509 -sha256 \
		-subj '/C=SG/ST=Singapore/L=Singapore/O=go-server/OU=testing/CN=localhost' \
		-days 365 \
		-key secrets/tls.key \
		-out secrets/tls.crt
	@openssl genrsa \
		-out secrets/tls2.key \
		4192
	@openssl req -new -x509 -sha256 \
		-subj '/C=SG/ST=Singapore/L=Singapore/O=go-server/OU=testing/CN=localhost' \
		-days 365 \
		-key secrets/tls2.key \
		-out secrets/tls2.crt
secrets.clean:
	@rm -rf secrets
ssh.keys:
	@mkdir -p ./bin
	@ssh-keygen -t rsa -b 8192 -f ./bin/${PREFIX}_id_rsa -q -N ''
	@cat ./bin/${PREFIX}_id_rsa | base64 -w 0 > ./bin/${PREFIX}_id_rsa_b64
test:
	@go test ./... -coverprofile c.out
test.watch:
	@godev test
tls: secrets
	@mkdir -p tls
	@ln -s $$(pwd)/secrets/tls.crt $$(pwd)/tls/server.crt
	@ln -s $$(pwd)/secrets/tls.key $$(pwd)/tls/server.key
version.get:
	@docker run -v "$(CURDIR):/app" zephinzer/vtscripts:latest get-latest -q
version.bump: 
	@docker run -v "$(CURDIR):/app" zephinzer/vtscripts:latest iterate ${VERSION} -i

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
tls: secrets
	@mkdir -p tls
	@ln -s $$(pwd)/secrets/tls.crt $$(pwd)/tls/server.crt
	@ln -s $$(pwd)/secrets/tls.key $$(pwd)/tls/server.key
secrets.clean:
	@rm -rf secrets
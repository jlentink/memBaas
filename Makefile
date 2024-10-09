.PHONY: gen

SSL_COUNTY?="NL"
SSL_STATE?=MyState
SSL_CITY?=MyCity
SSL_ORG?=MyORg
SSL_OU?=Computer
SSL_CN?=*.fobar.com
SSL_EMAIL?=foo@foobar.com

gen:
	protoc --go_out=. --go-grpc_out=. membaas.proto

tls:
	-rm *.pem
	openssl genrsa -out private.key.pem 2048
	openssl req -new -key private.key.pem -out request.csr.pem -subj "/C=${SSL_COUNTY}/ST=${SSL_STATE}/L=${SSL_CITY}/O=${SSL_ORG}/OU=${SSL_OU}/CN=${SSL_CN}/emailAddress=${SSL_EMAIL}"
	openssl x509 -req -days 365 -in request.csr.pem -signkey private.key.pem -out selfsigned.crt.pem

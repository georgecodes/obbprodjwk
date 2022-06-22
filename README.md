OBB Prod JWK generator

This command line utility generates an Open Banking Brasil Prod TPP-compliant JWK/JWKS from a certificate.

We use a slightly customised version of JWK for production TPP tests, to appear superficially like private key JWK 
but are in fact backed by a KMS-held private key.

USAGE:

```shell
./obbprodjwk --in <certfile> --alias <KMS alias> [--jwks=false]
```

This takes an Open Banking Brasil brcac or brseal, a KMS alias and generates the JWKS for it. The optional --jwks parameter
allows you to output as either a JWK or a JWKS with a single JWK in. 
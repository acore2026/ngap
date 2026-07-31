# ASN.1 Schema

`ngap-15.8.0.asn1` is the repository-owned NGAP compatibility schema used by
`rust/build.rs`. It originated from a local UERANSIM NGAP schema based on 3GPP
TS 38.413 release 15.8.0 and was normalized to UTF-8. It includes the
Selected-PLMN-Identity and extension identifiers required by the current Go
codec.

The Rust build generates types into Cargo's `OUT_DIR`; generated source is
never checked in or edited directly. When changing the schema, add or update
canonical vectors under `testdata/interop/` and verify both:

```text
cargo test --manifest-path rust/Cargo.toml
go test ./...
```

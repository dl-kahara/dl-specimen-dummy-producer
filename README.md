# A dummy producer

This program generates an endless sequence of POST requests towards an HTTP endpoint
configured with `CONSUMER`.
The requests' JSON payloads contain an RFC3339-formatted timestamp and a random number
with a length of `PRIMEBITS` number of bits, that is
[likely to be a prime](https://pkg.go.dev/crypto/rand#Prime):

```json
{
  "timestamp": "2024-11-12T08:55:54.586005Z",
  "prime":"fecf8a4ca29a6839c4f8297260347a78aec61e2b8..."
}
```

The HTTP endpoint is expected to respond with a SHA256 of the received JSON payload,
in "raw" binary form (Base64-encoded here for convenience):

```
PHTZ7JEXHlJKCRz33UXzz+nslt77jK5o3VDbtxIaq6s=
```

All exceptional situations lead to program termination. This includes not being able to
connect to, or not receiving the correct SHA256 back from the endpoint.

Following Prometheus-formatted counters (in addition to Golang's builtin metrics) are
exposed on `METRICS_ADDRESS`:

```
datalounges_producer_requests_total
datalounges_producer_bytes_total
```

## Runtime configuration

```console
# Configuration takes place over environment variables, and the defaults are as follows:
PRIMEBITS=4096 \
    CONSUMER=http://localhost:8080 \
    METRICS_ADDRESS=:9108 \
    go run .
```

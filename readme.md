# traefik-header-class-marker

Traefik middleware plugin which creates HTTP header based on the value found in another header. This can be used to
implement different throttling class when combined with a `rateLimiter` middleware.

## Example configuration

### Kubernetes
``` tab="File (Kubernetes)"
apiVersion: traefik.containo.us/v1alpha1
kind: Middleware
metadata:
  name:throttling-classes
spec:
  plugin:
    traefik-header-class-marker:
      sourceHeader: x-jwt-preferred_username
      destinationHeaderPrefix: x-throttling-class-
      classes:
        degraded:
          - alice
          - john
        minimal:
          - bob
        rejected:
          - sebastien
```

When an HTTP request contains the header `x-jwt-preferred_username: john`, an additional header
`x-throttling-class-degraded: john` is created and injected in the request, because `john` is listed in the class
`degraded`.



## Running and testing

Requires `yaegi` (used by Traefik itself to run plugins).

```
go install github.com/traefik/yaegi/cmd/yaegi@latest
# or
curl -sfL https://raw.githubusercontent.com/traefik/yaegi/master/install.sh | bash -s -- -b $GOPATH/bin v0.9.0
```

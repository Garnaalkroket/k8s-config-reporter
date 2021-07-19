# k8s-config-reporter
After providing:
 * An Authentication method (kubeconfig / cluster)
 * The name of a ConfigMap
 * A comma seperated list of keys to check in this ConfigMap

The application will provide a metrics **(/metrics)** endpoint.
Each requested key will have an associated Gauge metric, which will carry the value '1' if the exists.
A label 'status' will show the actual value of the key.
_TODO: Add Missing key support (e.g. value = 0, label as missing)_

The application will provide an additional endpoint **(/resources)** which will represent the key/value pair(s) for each namespace as plain JSON.

By default the application will run on port 8080.

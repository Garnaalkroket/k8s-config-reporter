# k8s-config-reporter
After providing:
 * An Authentication method (kubeconfig / cluster)
 * The name of a ConfigMap
 * A comma seperated list of keys to check in this ConfigMap

The application will provide a metrics endpoint, each key will have an associated Gauge metric (which will always be 1). A label 'status' will show the actual value of the key.
The application will provide an additional endpoints (/resources) which will represent the key/value pair(s) for each namespace as plain JSON.

By default the application will run on port 8080.

oc project cert-manager;
oc patch certmanagers.operator.openshift.io cluster   --type='json'   -p="[
    {
      "op": "add",
      "path": "/spec/unsupportedConfigOverrides",
      "value": {
        "controller": {
          "args": $(oc get deployment cert-manager -o json | jq '.spec.template.spec.containers[0].args += ["--feature-gates=LiteralCertificateSubject=true"]' | jq -c '.spec.template.spec.containers[0].args')
        },
        "webhook": {
          "args": $(oc get deployment cert-manager-webhook -o json | jq '.spec.template.spec.containers[0].args += ["--feature-gates=LiteralCertificateSubject=true"]' | jq -c '.spec.template.spec.containers[0].args')
        }
      }
    }
  ]"

#!/bin/bash
#TBD: create a function to call them with generic strings
echo -ne "Welcome to the world of demo-patterns\033[0K\r"
sleep 2
echo -ne "For demo source codes, refer to https://github.com/TrilokGeer/demo-patterns\033[0K\r"
sleep 2
echo -ne "Click on \"Watch\" to get notifications for new demos, \"fork\" the repository and raise a PR for contribution\033[0K\r"
sleep 2
echo -ne "The showcase here is about usage of literalRDN subject with OpenShift supported cert-manager\033[0K\r"
sleep 2
echo -ne "Pre-requisites: cert-manager operator is installed from OpenShift operator hub\033[0K\r"
sleep 2
echo -ne "To enable the literalRDN feature gate, patch cluster CR managed by the operator \033[0K\r"
sleep 2
sh ./patch.sh
echo -ne "Check for updated arguments in cluster resource under certmanagers.operator.openshift.io API group\033[0K\r"
oc get certmanagers.operator.openshift.io -o json | jq '.items[0].spec.unsupportedConfigOverrides.controller.args' | grep --color '\-\-feature-gates=LiteralCertificateSubject'
sleep 2
echo -ne "Check the deployments for cert-manager controller and webhook contain updated args \033[0K\r"
oc get deployment cert-manager -o json | jq '.spec.template.spec.containers[0].args' | grep --color '\--feature-gates=LiteralCertificateSubject'
oc get deployment cert-manager-webhook -o json | jq '.spec.template.spec.containers[0].args' | grep --color '\--feature-gates=LiteralCertificateSubject'
sleep 2 
echo -ne "Check the deployments for cert-manager controller and webhook are restarted \033[0K\r"
sleep 2


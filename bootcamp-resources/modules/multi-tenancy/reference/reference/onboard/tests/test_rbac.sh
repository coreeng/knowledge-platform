#!/bin/bash
set +e

testConnectivity() {
  nonAccessibleNamespaces=()
  accessibleNamespaces=()
  NAMESPACES=$(kubectl get ns | awk '{print $1}')

  for namespace in $NAMESPACES; do
    kubectl get pods -n ${namespace} --as=system:serviceaccount:$1 >/dev/null 2>&1
    STATUS_EXIT_CODE=$?
    if [ ${STATUS_EXIT_CODE} -eq 1 ]; then
      nonAccessibleNamespaces+=(${namespace})
    else
      accessibleNamespaces+=(${namespace})
    fi
  done

  echo "The service \"${1}\" has access to:"
  for ns in "${accessibleNamespaces[@]}"; do
    echo $ns
  done
  echo "----------"
  echo "The service \"${1}\" does not have access to:"
  for ns in "${nonAccessibleNamespaces[@]}"; do
    echo $ns
  done
  echo "*****************************"
}

testConnectivity $1



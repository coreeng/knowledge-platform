#!/bin/bash -x

for i in {1..10}
do
  if [[ $(kubectl get -n reference-service-ci pipelinerun test-pipeline-run  -o 'jsonpath={..status.conditions[?(@.type=="Succeeded")].status}') != "True" ]]
  then
    echo "Pipeline not complete yet"
    kubectl describe pipelinerun -n reference-service-ci  test-pipeline-run
    kubectl get pods -n reference-service-ci -o wide
    kubectl describe pod -n reference-service-ci -l  tekton.dev/pipelineRun=test-pipeline-run
    kubectl logs -n reference-service-ci -l  tekton.dev/pipelineRun=test-pipeline-run || true
    sleep 120
  else
    echo "Pipeline complete"
    kubectl describe pipelinerun -n reference-service-ci  test-pipeline-run
    kubectl get pods -n reference-service-ci -o wide
    kubectl logs -n reference-service-ci -l  tekton.dev/pipelineRun=test-pipeline-run || true
    echo "PIPELINE SUCCESSFUL!"
    exit 0
  fi
done

echo "Pipeline run not successful"
exit -1

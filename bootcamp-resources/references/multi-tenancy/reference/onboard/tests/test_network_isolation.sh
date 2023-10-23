#!/bin/bash
set +e

testConnection(){

    echo -ne " Connection from [ ${1} ]  Namespace -> [ ${2} ${3} ]     |     Result: "
    result=$(kubectl run -it --rm --restart=Never netcat-pod --image=alpine  -n ${1} --command -- timeout 3s nc -vz ${2} ${3} 2> /dev/null)
    if [[ $result == *"open"* ]]; then
        echo "-------Connected Successfully-------"
    else
        echo  "------Couldn't connect------"
    fi

}

echo "**********************************"
echo "Test inter namespace connectivity"
echo "**********************************"

#testConnection app-1 team-a-monitoring-prom-kube-state-metrics.team-a-monitoring 8080
#testConnection app-2 team-a-monitoring-prom-kube-state-metrics.team-a-monitoring 8080
#testConnection app-3 team-a-monitoring-prom-kube-state-metrics.team-a-monitoring 8080

echo "**********************************"
echo "Test intra namespace connectivity"
echo "**********************************"

testConnection team-a-monitoring alertmanager-operated.team-a-monitoring 9093

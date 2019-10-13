IFS=$'\n'; CLUSTERS=($(kubectl config get-clusters)); unset IFS;

unset 'CLUSTERS[0]'

for i in "${CLUSTERS[@]}"
do
   echo "$i"
  kubectl config use-context "$i"
  kubectl get pods --all-namespaces | grep web
 done

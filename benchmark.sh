
# First argument is used as the pattern
PATTERN=$1

IFS=$'\n'; CLUSTERS=($(kubectl config get-clusters)); unset IFS;

# Remove first line (NAME)
unset 'CLUSTERS[0]'

for i in "${CLUSTERS[@]}"
do
  _=$(kubectl config use-context "$i")
  kubectl get pods --all-namespaces | grep "$PATTERN"
done

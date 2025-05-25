set -euo pipefail
set -x  # print each command as it runs

commands=(
        "oc new-project spire-mgmgt"
	"oc create -f rbac"
	"oc create -f scc"
	"oc create -f service"
	"oc create -f configmaps"
	"oc create -f webhooks"
	"oc create -f statefulset"
	"oc wait --for=condition=ready pod -l app.kubernetes.io/name=server -n spire-mgmt --timeout=300s"
	"oc create -f csidriver"
	"oc create -f daemonset"
	"oc wait --for=condition=ready pod -l app.kubernetes.io/name=agent -n spire-mgmt --timeout=300s"
	"oc wait --for=condition=ready pod -l app.kubernetes.io/name=spiffe-csi-driver -n spire-mgmt --timeout=300s"
	"oc create -f crd/bases"
	"oc create -f apps"
	"oc create -f deployment"
	"oc wait --for=condition=ready pod -l app.kubernetes.io/name=spiffe-oidc-discovery-provider -n spire-mgmt --timeout=300s"
	"oc create -f test/test-namespace.yaml"
	"oc create -f test/sa-svid-test.yaml"
	"oc create -f test/scc-test-pod.yaml"
	"oc create -f test/test-spiffe-id.yaml"
	"oc create -f test/test-pod.yaml"
	"oc wait --for=condition=ready pod -l app.kubernetes.io/name=svid-test -n svid-test --timeout=300s"
)

testCommands=(
	"oc cp ./bin/spire-1.11.2/bin/spire-agent svid-test/spire-svid-test:/spire-agent -c test"
	"oc exec -it spire-svid-test -n svid-test  -- /spire-agent api fetch x509 -socketPath /spiffe-workload-api/api.sock"
	"oc exec -it spire-server-0 -c spire-server -- /opt/spire/bin/spire-server entry show"
)

# Show help if no arguments
if [[ $# -eq 0 ]]; then
  set -- --help
fi

# Handle help or list arguments
for arg in "$@"; do
  case "$arg" in
    --help|-h)
      echo "Usage: $0 [--list] | [[--skip=NUM[,NUM|NUM-NUM]...] --run]"
      echo
      echo "This script runs a sequence of OpenShift (oc) commands. You can skip specific commands by number."
      echo
      echo "Options:"
      echo "  --run                      Execute the command sequence. Also, used in conjunction with skip"
      echo "  --skip=NUM[,NUM|NUM-NUM]   Skip command(s) by number. Supports individual numbers or ranges."
      echo "  --list                     List the numbered command sequence and exit"
      echo "  -h, --help                 Show this help message and exit"
      exit 0
      ;;
    --list)
      echo "Command sequence:"
      for i in "${!commands[@]}"; do
        printf "  %d. %s\n" $((i+1)) "${commands[$i]}"
      done
      exit 0
      ;;
  esac
done


# Parse --skip argument
skip_arg=""
test_requested=false
run_requested=false

for arg in "$@"; do
  if [[ $arg == --skip=* ]]; then
    skip_arg="${arg#--skip=}"
  fi
  if [[ $arg == --run ]]; then
    run_requested=true
  elif [[ $arg == --test ]]; then
    test_requested=true
  fi 
done

# Expand skip list (e.g., 2,4-5 → 2 4 5)
expand_skip_list() {
  local skip_expanded=()
  IFS=',' read -ra parts <<< "$1"
  for part in "${parts[@]}"; do
    if [[ "$part" == *"-"* ]]; then
      IFS='-' read -r start end <<< "$part"
      for ((i=start; i<=end; i++)); do
        skip_expanded+=("$i")
      done
    else
      skip_expanded+=("$part")
    fi
  done
  echo "${skip_expanded[@]}"
}

# Execute commands
run_commands() {
  local skip=(0)  # Always declare it to avoid unbound variable issues

  if [[ -n "${skip_arg:-}" ]]; then
    skip=($(expand_skip_list "$skip_arg"))
  fi

  for i in "${!commands[@]}"; do
    cmd_num=$((i+1))
    if [[ " ${skip[*]} " =~ " $cmd_num " ]]; then
      echo "[SKIP] Command $cmd_num: ${commands[$i]}"
      continue
    fi
    echo "[RUN ] Command $cmd_num: ${commands[$i]}"
    eval "${commands[$i]}"
  done
}

# Run commands if --run was specified
if $run_requested; then
  run_commands
fi

# Execute test commands
test_commands() {
  for i in "${!testCommands[@]}"; do
    cmd_num=$((i+1))
    echo "[RUN ] Command $cmd_num: ${testCommands[$i]}"
    eval "${testCommands[$i]}"
  done
}

# Run test commands if --test was specified
if $test_requested; then
  test_commands
fi

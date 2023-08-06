start=$(date +%s%N)
SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd)
go run $SCRIPT_DIR/$(dirname "$0")/main.go
end=$(date +%s%N)
echo Execution time for day 1 Go_attempt was $(expr $end - $start) nanoseconds.

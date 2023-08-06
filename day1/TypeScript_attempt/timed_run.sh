start=$(date +%s%N)
SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd)
npm start
end=$(date +%s%N)
echo Execution time for day 1 TypeScript_attempt with ts-node was $(expr $end - $start) nanoseconds.
start=$(date +%s%N)
SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd)
npm run start:pre-transpiled
end=$(date +%s%N)
echo Execution time for day 1 TypeScript_attempt with pre-transpilation was $(expr $end - $start) nanoseconds.

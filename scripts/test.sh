#!/usr/bin/env bash

function lint() {
  go run . --config tests/simple.config.yaml lint tests/input."$1"
}

function fix() {
  go run . --config tests/simple.config.yaml fix tests/input."$1"
}

diff <(lint match) <(echo 2)
diff <(fix match) <(echo 3)
diff <(lint shebang) <(echo 1)
diff <(fix shebang) <(echo 4)
diff <(lint none 2>/dev/null) <(echo -n)
diff <(fix none 2>/dev/null) <(echo -n)
diff <(lint none 2>&1) <(echo Ignoring: tests/input.none)
diff <(fix none 2>&1) <(echo Ignoring: tests/input.none)

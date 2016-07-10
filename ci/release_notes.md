# Changes

- delmo doesn't run on `docker-machine` by default anymore.
- `wait:` step now only accepts 1 task name. (Previously it was an array).

# Features

- `--skip-pull` option to prevent delmo from pulling the images prior to building them. Usefull when network connectivity is bad.

# Bugs

- Correct output when using `--parallel`

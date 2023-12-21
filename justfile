test:
    just clojure-aoc/test
    just typescript/test
    just golang/test

lint:
    just --fmt --check --unstable
    just clojure-aoc/lint
    just golang/lint
    just typescript/lint

fmt:
    just --fmt --unstable
    just clojure-aoc/fmt
    just golang/fmt
    just typescript/fmt

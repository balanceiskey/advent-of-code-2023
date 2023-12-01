test:
    just clojure/test
    just typescript/test
    just golang/test

lint:
    just --fmt --check --unstable
    just clojure/lint
    just golang/lint
    just typescript/lint

fmt:
    just --fmt --unstable
    just clojure/fmt
    just golang/fmt
    just typescript/fmt

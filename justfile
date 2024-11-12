set shell := ["bash", "-uec"]

default:
    @just --list

fmt:
    prettier --ignore-path=.prettierignore --config=.prettierrc.json --write .
    just --unstable --fmt

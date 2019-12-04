![AoC Badge](https://img.shields.io/endpoint?url=https%3A%2F%2Fus-central1-hobbs-tinkering.cloudfunctions.net%2Faoc-badge-2019%3Fuser_id%3D4896)
[![Build Status](https://travis-ci.org/jmhobbs/advent-of-code-2019.svg?branch=master)](https://travis-ci.org/jmhobbs/advent-of-code-2019)

# Advent of Code 2019

For this year I've decided to learn [Rust](https://www.rust-lang.org/) as I work on the problems.

Likely a mistake, but hopefully I can get through the Rust book and some problems before Christmas.

## Layout

### ./[0-9]+

These are the problems by day.  Inside are individual Rust crates (?) that can be ran with `cargo run` and tested with `cargo test`

### ./support

This contains a tool to download the input files to avoid copy-paste issues.

Additionally, the `badge` directory has a Google cloud function that is used to generate the badge in this repo showing how many stars I have earned.

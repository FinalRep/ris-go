![CI Status](https://github.com/finalrep/ris-go/actions/workflows/ci.yml/badge.svg)

# ris-go – Relative Index for Streetlifting

**ris-go** is a Go library for calculating the *Relative Index for Streetlifting (RIS)*. The library allows the calculation of the RIS value based on individual strength levels and body weight, and it provides a method for determining the parameters of the underlying mathematical model based on real sports data.

This project was created as part of a bachelor's thesis with the goal of making performance evaluation methods in strength sports comparable and adaptable.

---

## Background

The RIS is a relative scoring system specifically for the sport of **Streetlifting**. It is designed to make athletes of different weight classes comparable by normalizing total performance.

### Formula

&nbsp;
**RIS = (Total × 100) / [A + (K − A) / (1 + Q · e^(−B · (BW − v)))]**

&nbsp;

**Parameters:**

- `Total`: Total performance (e.g., sum of Weighted Pull-Up and Weighted Dip)
- `BW`: Body weight
- `A, K, Q, B, v`: Parameters that are optimized through fitting to real data

---

## Features

- Calculation of the RIS value with given parameters
- Fitting of RIS parameters to performance data via nonlinear optimization
- Importing and processing CSV data
- Modular architecture in Go

---

## Quick Start

### Prerequisites

- Go ≥ 1.24
- [`gonum`](https://github.com/gonum/gonum) for mathematical optimization

### Installation

```bash
go get github.com/finalrep/ris-go/lib
```

### Example

- see our [example](https://github.com/FinalRep/ris-go/tree/main/examples) implementation

## Development and Contribution

- _Test your code properly!_
- use [conventional commit messages](https://www.conventionalcommits.org/en/v1.0.0/)
- use [semantic versioning](https://semver.org/lang/de/) to create tags and versions
- use [pre-commit](https://pre-commit.com/)

### pre-commit hook

This project uses [pre-commit](https://pre-commit.com/).
Install pre-commit and run•
```
pre-commit install
```


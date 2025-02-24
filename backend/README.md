# Income Tax Calculator (ITC)

## Overview
The Income Tax Calculator (ITC) is a backend application built in Go with a React Native frontend. It calculates income tax based on both the **new tax regime** (FY 25-26 onward) and **old tax regime**, considering various deductions and exemptions.

## Features
- Supports both **new and old tax regimes**.
- Calculates **Gross Total Income**, **Employer/Employee PF**, and **NPS Contributions**.
- Includes **capital gains taxation** (short-term and long-term).
- Supports **HRA exemptions** based on rent paid and metro city status.
- Implements **Section 80C & 80D deductions** (PPF, ELSS, LIC, health insurance, medical expenses).
- Determines the appropriate **ITR form** to file (ITR-1, ITR-2).
- Provides structured **JSON-based input and output**.
- Backend built using **Gorilla Mux**, structured in a modular format.
- Fully tested with **unit tests** before building.

## Installation & Setup

### **Backend (Go)**
#### **Prerequisites**
- Install [Go](https://golang.org/dl/)
- Install [Make](https://www.gnu.org/software/make/)

#### **Build and Run**
```sh
make        # Runs tests and builds the binary
make run    # Starts the application

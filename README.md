# Controlled System Node

A simulated long-running Go service demonstrating **system lifecycle management**, **strict state ownership**, and **controlled failure handling**.

This project models how production-grade systems behave under normal operation, degradation, and shutdown — focusing on **clarity**, **safety**, and **observability** over features.

---

##  Objective

This project was developed as part of a task to illustrate key system design concepts:

- Model explicit lifecycle states: `init → running → draining → stopped`  
- Enforce valid state transitions only  
- Separate responsibilities into:
  - Core system control  
  - Observability (read-only access)  
  - HTTP boundary layer  
- Support controlled failure injection without crashing  
- Handle **graceful shutdown** using OS signals  

---

##  Features

-  Controlled system lifecycle states  
-  State transition validation  
-  Graceful shutdown support  
-  Observability via HTTP endpoints  
-  Fault injection support  

---

##  Architecture

The service is structured to emphasize separation of concerns:

Task3/
├── core/ # Core logic and system state management

├── http/ # HTTP API and observability endpoints

├── signals/ # OS signal handling

└── main.go # Program entry point


---

##  Requirements

Make sure you have the following installed:

- Go ≥ 1.18  
- Git  

---

##  How to Run

1. Clone the repository

```sh
git clone https://github.com/Prakashkumarydv3/Task3.git
cd Task3
Build the project

go build
Run the service

./Task3
Monitor logs or access HTTP observability endpoints (if enabled) to track system state and transitions.

How It Works
This system simulates the lifecycle of a long-running service.

Valid lifecycle flow:

init → running → draining → stopped
The system starts in the init state

Moves to running when the service becomes active

Transitions to draining during graceful shutdown

Ends in stopped once all operations are safely terminated

Invalid transitions are blocked to maintain system safety and clarity.

 What I Learned
How to design a system with explicit state ownership

Importance of graceful shutdown handling

Separating core logic from API and system boundaries

Writing cleaner, more maintainable Go services

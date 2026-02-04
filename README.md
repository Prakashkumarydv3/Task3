# Controlled System Node 
**Lifecycle Modeling, State Ownership & Fault Injection in Go**

---

## 📌 Overview

This project is a simulated long-running Go service designed to demonstrate **system lifecycle management, strict state ownership, and controlled failure handling**.  
It models how real production systems behave under normal operation, degradation, and shutdown — with a focus on **clarity, safety, and observability over features**.

This is a **system design exercise**, not a product or platform.

---

## 🎯 Objectives

- Model explicit lifecycle states:  
  `init → running → draining → stopped`
- Enforce **valid state transitions only**
- Separate:
  - Core system control
  - Observability (read-only)
  - HTTP boundary layer
- Inject controlled failure modes without crashing
- Implement graceful shutdown using OS signals

---

## 🏗 Architecture


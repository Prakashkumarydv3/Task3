# Conceptual Boundaries

## Difference Between Data, Signal, Decision, and Execution

### Data
- Data is a raw fact collected by the system.
- It is simply information that exists.
- By itself, data does not mean anything and does not require action.

### Signal
- A signal is data that has been interpreted.
- It indicates that something has changed or crossed a condition.
- A signal only tells what is happening, not what should be done.

### Decision
- A decision is a conscious choice made after evaluating signals.
- It depends on rules, intent, and authority.
- Decisions answer the question: what action should be taken?

### Execution
- Execution is the act of carrying out a decision.
- It changes the state of the system or the environment.
- Execution should not question or reinterpret decisions.

## Why Monitoring Systems Should Never Decide Actions
- Monitoring systems are meant to observe and report.
- They do not have full context or authority.
- If monitoring systems start deciding actions, a wrong or noisy signal can directly cause damage.
- Keeping monitoring separate from decision-making prevents panic reactions and unsafe behavior.

## Why Historical Records Should Never Be Modified
- Historical records represent events that already happened.
- If history is modified, trust in the system is lost.
- It becomes impossible to audit or understand cause and effect.
- If something needs correction, it should be added as a new record, not by changing the past.

## What It Means for a System to Degrade Gracefully
- Graceful degradation means the system does not fail completely when something breaks.
- Instead, it continues operating in a limited but safe way.
- Non-critical features can stop, but core behavior remains stable.
- The goal is controlled behavior, not perfect behavior during failure.
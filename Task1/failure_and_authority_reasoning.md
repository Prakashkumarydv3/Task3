# Failure and Authority Reasoning

## If a Scoring System Goes Offline, Should Execution Stop?
**No.**
- A scoring system helps with evaluation, not execution authority.
- Execution should continue using existing rules or the last known safe state.
- Stopping execution would make a non-critical system block the entire system.

## If a State Tracker Is Delayed, What Should Continue to Function?
- Execution should continue based on the most recently confirmed state.
- Observation may become outdated, but authority should not pause.
- It is better to operate with slightly old information than to stop entirely.

## Danger of Letting Analytics Directly Influence Execution
- Analytics are designed to analyze past behavior and trends.
- They are not meant to control live execution.
- If analytics directly influence execution, hidden feedback loops can form.
- This makes system behavior unpredictable and hard to debug.

## Why Authority Must Always Exist in One Place
- Authority decides who is allowed to make decisions.
- If multiple parts of a system have authority, conflicts are unavoidable.
- This leads to inconsistent behavior and race conditions.
- Keeping authority in one place ensures clarity, accountability, and predictable execution.

## Final Reflection
Through this task, I understood that good systems are not built by making everything intelligent. They are built by respecting boundaries and limiting responsibility. Clear ownership and separation make systems easier to trust, operate, and recover.
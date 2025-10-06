# Full-Stack Challenge

## Overview

This challenge is designed to assess your ability to build a simple remote job system with a Go backend and React frontend. The specific requirements for each part are in the `REQUIREMENTS.md` in the `backend` and `frontend` folders.

## Time Guidelines

***The challenge requirements and criteria are intentionally open ended to allow 4+ hours of work. The goal is to see what you prioritize and can accomplish in a given time constraint.***

- **Core Requirements:** ~1-2 hours
- **With Stretch Goals:** ~3-4 hours

_These estimates are calibrated for mid-senior engineers comfortable with Go and React. If you're less familiar with one stack, focus on quality in the stack you know best._

**Please don't spend more than 4 hours total.** We value quality over quantity — a polished implementation of core features is better than a rushed attempt at everything.

## Evaluation Criteria

_You **may use AI coding agents to help**, but if you do you must **clearly identify where you did not use AI**. You must be able to explain/justify any AI generated code during our review call. **We judge all code to the same standard regardless of human or AI authorship.**_

_If there is no explicit "I did not use any AI" statement in your README, we will assume AI was used for the entire project. We use AI tools at Quadric too — we're simply asking for transparency so we can have an informed discussion about your work._

_Note: Misrepresenting your use of AI tools undermines the evaluation process and reflects poorly on professional integrity._

### Must Have (Core Requirements)

- ✅ All API endpoints functioning correctly
- ✅ Create, list and acquire jobs
- ✅ Register runners
- ✅ Append and save logs
- ✅ Code is clean, organized, and readable

### Nice to Have (Quality Indicators)

- Loading states and UX polish
- Clean separation of concerns
- Responsive design
- Code comments where appropriate
- Logs where appropriate
- Unit or integration tests

### What We're Evaluating

- **Functionality:** Does the application work as specified?
- **Code Quality:** Is the code clean, organized, and maintainable?
- **User Experience:** Is the UI intuitive? Are errors handled gracefully?
- **Communication:** Is your README clear? Are architectural decisions explained?
- **Judgment:** Did you make smart trade-offs given the time constraint?
- **AI vs Self**: Did you use AI? How much? (remember; if you can do better than AI, don't let the AI take credit)

## Submission

Please submit your solution via:
- A link to a GitHub repository (preferred), or
- A zip file containing your code

### Your README should include:

1. **Setup Instructions**
   - How to run the backend
   - How to run the frontend
   - Any dependencies or prerequisites
2. **Time Spent** (honest estimate)
3. **AI Usage** (clearly mark which sections were written without AI assistance)
4. **Implementation Notes**
   - Any assumptions you made
   - Which optional features you implemented
   - Trade-offs or decisions you'd like to explain
5. **Next Steps** (optional)
   - What you would add with more time
   - How you would improve the implementation

## Tips

- We've included a `compose.yml` to simulate external services (`docker compose up` or `podman compose up`)
- The `redis` compatible service (via [DragonflyDB](https://www.dragonflydb.io/dragonfly-vs-redis)) can be used to cache short lived values and/or help reduce the load on the DB
- A test runner is provided in the `/runner` folder — you don't need to modify it unless attempting the job cancellation bonus

## Questions?

If anything is unclear or you have questions about the requirements, please reach out. We're happy to clarify!

Good luck! We're excited to see what you build.
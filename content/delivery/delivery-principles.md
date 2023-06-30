+++
title = "Delivery Principles"
weight = 1
chapter = false
+++

## Culture

#### Operate as an idea meritocracy
* An idea meritocracy is an environment in which the best ideas win, regardless of where or whom they came from.

#### Avoid direct messaging to individuals
* If you message someone directly with information, then only that individual receives it.  They might pass the information on, but often don’t.  Even if the message is intended for an individual, post it in a public space so that everyone received it.  If you are asking for something, the same applies.  You should rely on teams of individuals.  It also promotes openness and allows others to contribute value if they want to.

#### Disagreeing should be done efficiently
* In an idea meritocracy, disagreement can be time consuming, so know when it’s time to stop debating and make a decision

#### Once a decision is made, everyone should get behind it
* There are many ways to get a job done, so if individuals disagree with a decision, they should do their best to make it work even if they feel it’s not the best way to do it.
* ADRs

#### Data Driven Accountability
The success of Targets/Goals/Missions/Visions should be measurable, giving us a clear view of success/failure, and ideally the level of success. This should be applied to both engineering and management.

## Delivery

#### Delivery is king
* Focus on delivery. Encourage technical excellence but not at the expense of delivering for stakeholders
* Prevent perfect from being the enemy of good
* We never want to be accused of being ivory tower architects. We meet the stakeholders where they are and incrementally move them towards where they need to be

#### First class showcases
* Demonstrate  work via showcases aimed at high level stakeholders
* Demonstrate technical work to engineers

#### Do difficult things often
* Committing to increasing the burden caused by manual problems helps prioritise longer term strategic solutions that rely on the right level of automation.  This allows for building systems that scale well without requiring a linear scaling of human resources.   It also reduces risk, by exercising manual tasks (e.g. runbooks) regularly that could otherwise become stale and fail in a time of great need. E.g. backup and restore

#### Definition of done
* Agree a clear definition of done with your peers, including:
  * Peer reviewed
  * Merged into the main branch of version control
  * Built into a release
  * Deployed as far as Continuous Delivery
  * Done does not mean the feature has been “coded”

#### Decouple infrastructure and delivery teams
* New services should be able to be introduced without involvement from an infrastructure team
* At most a PR / automated rollout needs to take place

#### Customer and Delivery Focus
* Delivery teams are the customer of a platform and their satisfaction should be a primary concern. We are not the guardians of an ideal, but the stewards of their platform, serving their needs as a priority.
* Engineers produce their best work if they can determine their own ways of working.  Tools and processes should be driven by the developers experience rather than mandated by managers.  Standardisation is important but again those standards should be determined by the developer experience.

#### Planning the journey matters as much as planning the outcome
Figuring out how we get there from where we are is often harder than designing the final solution, and is often looked over.

#### Limit work in progress
* For a given amount of resource, increasing the work done in parallel will also increase the time it takes to deliver the work.  This applies at the individual and team level.  Optimise for increased velocity rather than parallelism so we can get fast feedback on what we are doing and reduce the management overhead.

#### T-shaped engineers fit best
* Engineering is by nature complicated and therefore we need engineers that are willing to go deeply into a subject.  At the same time, we expect engineers to be able to work across many aspects of the system they work with.  The solution is to have T-shared engineers that have a broad and shallow knowledge across the whole system they work with, and a deep knowledge of the part of the system they are responsible for.

#### Solutionising should be led by engineering teams
* Engineering teams should be given problems to solve instead of solutions to deliver.  Too often the latter is the case and nearly as often the solution is sub-optional.  The team has the deepest knowledge of the technology and the landscape around it, both current and future.  If they are not free to solve problems, then innovation will stop, and engineers will become demotivated.
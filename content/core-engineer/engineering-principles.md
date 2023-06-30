+++
title = "Engineering Principles"
weight = 1
chapter = false
+++

## Software & Infrastructure lifecycle

#### Minimum Viable Product (MVP)
*  Work in an agile way, favour feedback over up front preparation. 
*  Don’t do large releases without dropping continuous value
*  Sprint 0:
   a. New systems should define all their environments, including production, from day 1

#### Production changes are delivered to prod hands free 
*  Manual actions are a major cause of outages strive for no manual actions to deploy to production
*  Deployments can still be triggered manually if true continuous delivery isn't feasible
*  Follow the [testing strategy](/core-p2p/testing-strategy/) for infrastructure and applications

#### Encapsulated service repositories with a single main branch always built into an immutable, releasable artifact
* Use trunk based development over long lives feature / release branches 
* Include everything that is required to run and test a service in a single repository, versioned together. Including:
  * Code
  * Configuration if not environment specific or sensitive
  * Monitoring

## Site Reliability Engineering

Follow many of the principles and tecniques as described in [Site Reliability Engineering](https://sre.google/)

#### Teams should have production ownership of what they work on
* If you own the software you build in production, you will build it safely, and if you don’t, you’ll fix it quickly. This makes the ownership of production software very clear.
* We avoid ivory tower architects that maintain reference architecture, blue-prints, IaC templates, that they don't use themselves in production

#### Systems are unreliable
* Don’t design infrastructure and services as if they were stable. For example:
  * Assume any dependency will fail e.g. another service, a cloud provider service, a database

#### All alerts are actionable or a bug
*  False alerts need to be eliminated
*  Alerts should either have a clear runbook OR 3. They are a bug which is top priority
*  Have a postmortem strategy
*  Every failure resolution includes provision for regression testing


##  Platform Engineering

#### Platforming enables the scaling of DevOps
* Platform engineering teams build products that abstract low level infrastructure concerns for delivery teams, so they can build and run
    services
  * Delivery teams can’t be expected to “build it and run it” if they have to deal with low level ifnrastructure

#### Autonomous infrastructure is preferred to automated infrastructure
* All infrastructure, where possible, should be configured with code. IaC tools such as terraform are what we call automated infrastructure. Everything is created via code but there is nothing ensuring that it stays in the desired state.
* Autonomous infrastructure is the next evolution where there is software constantly reconciling the desired state with the actual state. The Kubernetes resource model is one example of this.
  * Typically, base level infrastructure that rarely changes, such as networks, routers, VPNs, and Cloud interconnects can be configured with IaC tools but all higher level infrastructure components, logging, CI/CD, and Ingress should be configured with constantly running software, a.k.a operators, so that they are autonomous.

#### Destroy and re-create regularly
* Infrastructure should be destroyed regularly, even if just in lower environments.
* It is our experience that environments and infrastructure that are not regularly destroyed can not be re-created from scratch, even when defined with IaC tools and operators.
  * An incrementally developed IaC code base, that isn’t torn down, has never run from scratch in its current state meaning it is unlikely it will
      work.

#### Infrastructure components should be deployed with pipelines
* Production pipelines, such as is commonplace for applications, should be used to deploy infrastructure software, such as operators, and IaC code bases.

## Software Engineering

#### Complexity is the root of all software problems.
* Complexity is incremental and if you don’t actively manage it, it can become overwhelming to fix and lead to systems which are very difficult to change and maintain. Aka “technical debt”.
* Symptoms of complexity: 
  * change amplification (even a simple change requires changing many different places) 
  * cognitive load (needing to  know a large amount of information to make any change)
* Unknown unknowns (not even knowing where to look or about unknown effects of a change).

#### Tactical vs strategic programming
* Tactical approach is to “get it done” as quickly as possible.
  * It’s almost impossible to come up with a good design via purely tactical programming.
* Strategic programming may take more time initially, but it pays off quickly with a codebase that is easier to maintain and understand. 4. Fundamentally strategic programming means placing good design as a priority over getting it done as quickly as possible.
* Spend 10-20% of your time investing in good design.

#### Modules should be deep
* Interface of a module or component should be simpler than its implementation.
* Large numbers of dependent shallow modules introduce significant complexity. Instead, prefer deeper modules with similar logic and code,  and simpler interfaces. This applies both at the macro (projects) and micro (classes/packages/functions).
* Interfaces should be somewhat general purpose. Not tied specifically to a single client, and general enough to provide simpler interfaces
    with a deeper implementation. Too specific interfaces lead to larger cognitive load due to a more complex and larger interface API.

#### When to keep logic together in a module
* When information is shared
* If it will simplify the interface
* If it eliminates duplication. If the same pattern appears over and over, it means we have the wrong abstraction.
* General purpose vs special purpose provides a good boundary between modules. E.g. generic deployment library, vs specific project deployment logic.
* “7 line method” is a myth. Instead, methods should do one thing completely, regardless of length. Many small,
    shallow methods leads to increased complexity (adds indirection, difficult to follow the code flow, increased cognitive load, etc.). We shouldn’t base method size on arbitrary LOC. We want deep methods too that are self contained and provide a simpler interface than their implementation (giving a useful abstraction).

#### Design it Twice
* The classic advice is still valid today.
* One of the best approaches is to design two competing approaches, and then compare them objectively.
* This forces you to think outside of a narrow mindset and consider options you may not have before. It also requires objectively analysing
    the designs to compare them, which may also reveal issues.

#### Consistency

* Is a powerful tool to reduce complexity.
* Once you’ve learned how a system works in one place, you can apply that knowledge to the rest of the system, if it behaves consistently. 3. This applies at many levels: naming, coding style, tooling used, system invariants.
* Automate enforcement of things you can like coding styles.
* Don’t be overzealous either - sometimes it can make sense to break consistency when it doesn’t fit the existing pattern.

#### Code should be Obvious
* Or, the principle of least surprise.
* Your code should behave as expected from its interface and names.
* It should be easy to understand.
* Non obvious code is best exposed in code review, and one of the best times for code review to be useful. It can be difficult to analyse your
    own code since, well, it’s obvious to you.

#### Designing for Performance
* Keep performance in mind when designing. The goal isn’t to hyper optimize everything, but to make sure we don’t take 10x or 100x slower approaches that won’t scale or cause problems later on.

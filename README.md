# Polaris API

## Index

1. [Application Structure](#application-structure)

## Application Structure

```sh
appplication/ # application code
|--- adapters/ # implementation of the ports defined in the domain
|--- commons/  # utils
|--- entrypoints/ # primary adapters, entry points
  |--- api/ # api entry point
  |--- model/ # api model
|--- domain/ # domain to implement business logic using hexagonal architecture
  |--- commands/ # commands on the domain
  |--- events/ # events emitted by the domain
  |--- exceptions/ # exceptions defined on the domain
  |--- model/ # domain model
  |--- ports/ # abstractions used for external communication
infraestructure/ # infrastructure code
```

- **command** contains the command functions that define the information required to perform an operation on the domain.
- **event** contains the events that are emitted through the domain and then routed to other microservices
- **exception** contains the known errors defined within the domain.
- **model** contains the domain entities, value objects, and domain services
- **ports** contains the abstractions through which the domain communicates with databases, APIs, or other external components.
- **tests** contains the test methods (such as business logic tests) that are run on the domain.
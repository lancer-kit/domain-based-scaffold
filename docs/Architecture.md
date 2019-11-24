# Domain-driven-design concept  

## Overview

1. `actions`
2. `config`
3. `docs`
4. `domains`
    1. service
        - delivery
        - entities
        - repo
        - usecases
    2. another-sub-domain
        - delivery
        - entities
        - repo
        - usecases
5. `workers`

##### Example
```text
domain-based-scaffold
├── actions
│   ├── initialization
│   └── migrations
│       └── postgres
├── config
├── docs
├── domains
│   └── service
│       ├── delivery
│       │   └── middlewares
│       ├── entities
│       ├── repo
│       └── usecases
├── info
├── tests
│   ├── requests
│   └── test-data
└── workers
    ├── api
    │   └── handler
    └── foobar
```


## Service structure  

### `actions`  
 Actions stores all packages which are used only once during service "initialization"
 It can include packages like *initialization, migrations etc.*.  

### `config`  
 Stores config structure and config.Init method.
 
### `docs`  

 Project related documentation.

### `domains`  

 This package stores main service codebase divided into separate independent parts that are combined to achieve desired functionality. 

##### `domains/service`

 Main domain of application. Includes subpackages:

1. `delivery`

 This is package that presents a contract we use to communicate with "outer world".
 This package stores all handlers which are used by workers in their routers. 
 If more than one type of transport is used, they must be divided by types(REST, gRPC, graphQL etc.).  

2. `entities` 

 Stores all business models with related methods. 
 These models are used in use cases and repository as part of domain "contract" for other domains(if they are intended to use some logic from this contract).  

3. `repository`

 In general represents an entity we retrieve and store data(DB is a good example)  

4. `use cases`

 All business logic is placed here. This package joins everything together. It gets formatted data from delivery, sets or updates data in the repository and performs some actions with entities. Also it can use usecases of other domains to retrieve their functionality  

##### `domains/some-another-domain`

This domain has the same structure **but** they can omit some packages they don't use. For example, delivery if domain doesn't have external API

### `workers`

Workers are any subprocesses like api servers


# Component Entity System

- World
  - Systems -> stateless functions which performs some task
    - Query
    + Update(query Query, delta float32)

  - Managers
  - Entities -> container for holding data
    - Components -> data


### Component
component is data

### System
system is a stateless unit which periodically gets updated. System can be 
add and remove to/from world.
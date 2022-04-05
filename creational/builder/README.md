# Builder Pattern

## Example - vehicle manufacturing

### Requirements and acceptance criteria

- Must have a manufacturing type that constructs everything that a vehicle needs

- When using a car builder, the `VehicleProduct` with four wheels, five seats, and a structure defined as `Car` must be
  returned

- When using a motorbike builder, the `VehicleProduct` with two wheels, two seats, and a structure defined as `Bike`
  must be returned

- A `VehicleProduct` built by any `BuildProcess` builder must be open to modifications

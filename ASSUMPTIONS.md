# Documentation of assumptions/decisisons as they come up

## Initial Assumptions

   - as no external tools are used, I assume that input size is smaller than RAM
     until requirements change
   - if the input is smaller than RAM, it's safe to also assume untaring the
     data is outside the exercise's scope => test data is unpacked into
     `test-data/`
   - CSV is not the greatest of formats. It would make sense to upstream the
     suggestion to change the data format to something more suitable
     (ideally even zero-allocation parsable in a streaming fashion), but again
     outside the exercise's scope
   - given the above, it makes sense to optimize for readability/maintainability
     at first until CPU or RAM becomes a bottleneck

## Tooling decision
   - Golang 1.17, no external processing tools
   - Optimizing for readibility suggests using explicit structs for the data and
     wrap the native encoding/csv (at the expense of CPU/RAM, see above for
     reasoning)

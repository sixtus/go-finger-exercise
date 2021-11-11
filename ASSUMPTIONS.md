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
     reasoning) (Note: didn't seem necessary actually coding, leaving comment for completeness)

## Analyzing the requirements

### Top 10 active users sorted by amount of PRs created and commits pushed

   1. read all commits to build a map `event_id -> commit count`
   2. read all events to build a map `actor_id -> PullRequestEvent` and and
   `actor_id -> PushEvent`(thought: either build array of push events and compute
   the `actor_id -> commit` count in a 2nd step or pass the map from 1. as
   parameter and build the map `actor_id -> commit count` directly)
   3. find the top 10 of both actor_id maps(thought: that's a worst case access to a map,
   check if there is a better data structure)
   4. lookup the name of the top 10 actors in actors
   5. output

### Top 10 repositories sorted by amount of commits pushed
   1. read all events to build a map `repo_id -> commit count`
      (that's basically the same as the first top 10, just on repo_id instead of actor_id)

### Top 10 repositories sorted by amount of watch events
   1. read all events to build a map `repo_id -> watch count`

### Conclusions from first analysis
   - There is a high overlap between the individual tasks. Also the requirements
     read as if all parts are executed all the time
   - actors and repos are classical lookups, it's probably easiest to just read
     them into a map (assumping they stay small even if we scale out)
   - commits are only relevant to find the commit count per event, that's a map
     too (this will grow faster on scale, but not unreasonably)
   - reading/building the information from events look like a map, however we
     then want a top10 of the values.
   - initial thought: it probably makes sense to keep the values in a slice
     (so the map value is just the slice offset). This will allow to use a
     quickselect w/o extra conversion step from map values to value slice.
     Let's see if we find a beter approach in-flight, but this seems like a
     fairly good trade off between code complexity and runtime efficiency.
   - This problem is more complex than it first appears. Defend with tests and
     revisit our assumptions frequently to make sure we don't have any logic
     error.

## Notes implementing an actors helper
   - The file is not unique, reading into a map will make the last name win.
     Assuming that's ok.
   - The repos helper will be exactly the same, so refactoring the code to be
     generic

## Notes on first code complete
   - As expected, this made a lot more sense in flight
   - Putting everything together was so easy, I actually wrote main before the
     test for events_scanner (pending ofc)
   - The code is fairly straight forward, now that I have a first working
     solution, it makes sense to revisit variable names. Also top_n.go deserves
     documentation
   - I was thinking about threading the intake, but my little laptop runs it in
     0m0,038s
   - Early optimization is the root of all evil, so keeping it single thread

## Notes on completion
   - Removing the mutex from topN almost halfed the runtime, it's not needed atm
   - The design philosophy of this app is "fail early and keep it simple"
   - The internal constructors all take a file name.
     This is "keeping it simple", but likely the first thing to change, when
     requirements change.

# This is a simple finger exercise repo

This repo contains Github event data for 1 hour.

Goal is a CLI application that outputs:

- Top 10 active users sorted by amount of PRs created and commits pushed
- Top 10 repositories sorted by amount of commits pushed
- Top 10 repositories sorted by amount of watch events

# Testing the solution

``` go test ./...` ```

# Running the solution
```
Usage of main:
  -actors string
        the actors.csv file (default "test-data/actors.csv")
  -commits string
        the commits.csv file (default "test-data/commits.csv")
  -events string
        the events.csv file (default "test-data/events.csv")
  -repos string
        the repos.csv file (default "test-data/repots.csv")
```

i.e. to run with the test data, just do a `go run main.go`

# Form3 Take Home Exercise



### Technical assumptions:

- I focused on simplicity and readability 
- I used testcontainers for setting up docker containers during executions of tests, so there is no need to run docker containers out of the box.
 Before you run tests, please make sure that in your local registry there is no containers like: ``vault, accountapi, postgresql`` (please delete already running containers, but even orphaned, not used ones containers)

### API client specification:

    createAccount(accountData AccountData)
    deleteAccount(id string)
    fetchAccount(id string)

### Summary comments:
- This was my first project in Go, I had to learn everything from the scratch
- Packages would look better, but I know that packaging process is quite different than in language that I'm currently using (Java), so I didn't want to spend a lot of time on it.
- I've added ``time.Sleep`` function before run tests to allow all container run up

Contact: Damian Popielarski - damian.popielarski.business@gmail.com

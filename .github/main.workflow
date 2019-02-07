workflow "Linty Check" {
  on = "pull_request"
  resolves = ["ArangoGutierrez/GoLinty-Action"]
}

action "ArangoGutierrez/GoLinty-Action" {
  uses = "ArangoGutierrez/GoLinty-Action@master"
  secrets = ["GITHUB_TOKEN"]
}

freight {
  root = "~/go/src/github.com/ChrisMcKenzie/freight/test/"
}

project "dropship" {
  path = "tools"
  remote = "github.com/ChrisMckenzie/dropship"

  after "script" {
    command = <<CODE
      go install -v
    CODE
  }

  after "script" {
    command = <<CODE
      echo "what's up!"
    CODE
  }
}

project "accord" {
  path = "tools"
  remote = "github.com/ChrisMckenzie/accord"
}

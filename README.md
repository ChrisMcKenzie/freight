# Freight

Freight is a tool for downloading and scripting the bootstrapping of large multi-repo projects. it can download git repos, as well 
as any other assets as well as execute basic tasks relating to those resources.

## Example
```
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
```

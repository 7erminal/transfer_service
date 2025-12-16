data "external_schema" "beego" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./loader",
  ]
}

env "beego" {
  src = data.external_schema.beego.url
  dev = "docker://mysql/8/dev"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
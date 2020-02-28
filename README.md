# data-external-wrapper
The [Terraform External Program Protocol](https://www.terraform.io/docs/providers/external/data_source.html) is a way for Terraform to communicate with external programs. However, the interface that they've chosen to use means that you can't conveniently use any existing external program without some sort of modification or scripting. Terraform expects to provide input via json posted to STDIN and expects to read JSON on STDOUT. This wrapper provides that interface layer for any existing script.

This wrapper will ignore any json sent to it via the "query" argument. It will capture stdout and return it as valid json in the "Result" json node.

```
data "external" "hello_world" {
        program = ["/usr/local/bin/data-external-wrapper", "/usr/local/bin/hello-world","--arg1","--arg2"]        
}
...
output "hello_world_result" {
  value = data.external.hello_world.result.Result
}
```


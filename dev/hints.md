

## Parsing YAML 
```
func TestParseInvalidYaml(t *testing.T) {
	targetDescription := []byte(`---
namespace: sora
containers:
  - type: helm
    releaseName: sora
    path: charts/swiftnav
    valuesPath: files/values.yaml
    values:
      server:
        enabled: true
        service:
          port: 10000
      service:
        port: 5000
        type: ClusterIP
    includes: [ *s, server]
  `)
	description := &types.TargetDescription{}
	err := yaml.Unmarshal(targetDescription, description)
	assert.Error(t, err, "yaml: line 17: did not find expected alphabetic or numeric character")
}
```

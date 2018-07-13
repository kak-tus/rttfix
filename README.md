# rttfix

consul-template plugin to fix issue with frequent rtt changes between nodes.

Usage

```
{{ range $s := service "some-service~_agent" | toJSON | plugin "rttfix" | parseJSON }}
{{ $s.Address }}:{{ $s.Port }}
{{ end }}
```

package global

{{- if .HasGlobal }}

import "github.com/VINDA-98/Tenacity-Solutions/server/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}
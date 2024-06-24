{{/* $Trigger Key */}}
{{/* 01 : Banned response*/}}
{{/* 02 : Wrong Server response */}}
{{/* 03 : Setup response */}}
{{/* 04 : Open Ticket response */}}
{{/* 05 : Premium response */}}
{{/* 06 : Not Responding response */}}
{{/* 07 : Custom Instance response */}}
{{/* 08 : Sef Host response */}}
{{/* 9 : Clyde response */}}
{{/* 10 : Global Ticket response */}}
{{/* 11 : Logging response */}}
{{/* 12 : General Help response */}}

{{ if not (hasPrefix .Message.Content "=") }}
    {{ return }}
{{ end }}

{{/* Regexes to match, pulled from the existing commands */}}
{{ $logging := reFindAllSubmatches `\A(?:\-|<@!?204255221017214977>|!|.)\s*(?:logging|logs|transcript|file|viewer)(?: +|\z)` .Message.Content }}
{{ $help := reFindAllSubmatches `(?i)(?:need (?:support|help|assistance|aid|advice)|(?:help|support) me)` .Message.Content }}


{{ $trigger := 0 }}


{{ if $logging }}
	{{ $trigger = 11 }}
{{ else if $help }}
	{{ $trigger = 12 }}
{{ end }}

{{/* ExecCC to call the main response trigger */}}
{{execCC 75 .Channel.ID 0 (sdict 
	"trigger" $trigger 
	"triggerMsgID" .Message.ID
	"note" "Just some extra info to passthrough if needed."
)}}

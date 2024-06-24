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


{{/* Declaring variables */}}
{{ $trigger := 0 }}
{{ $command := .strippedMsg.Content | lower}}
{{ $prefix := reFindAllSubmatches `^[.,!?]` .Message.Content }}

{{/* Checks if the reaction is the bin emoji */}}
{{ if and (not $prefix) (not .ServerPrefix) }}
	{{ return }}
{{ end }}


{{ if eq $command "banned" }}
	{{ $trigger = 1}}
{{ else if eq $command "wrong server" }}
	{{ $trigger = 2 }}
{{ else if eq $command "setup" }}
	{{ $trigger = 3 }}
{{ else if eq $command "open ticket" }}
	{{ $trigger = 4 }}
{{ else if eq $command "premium" }}
	{{ $trigger = 5 }}
{{ else if eq $command "not responding" }}
	{{ $trigger = 6 }}
{{ else if eq $command "custom instance" }}
	{{ $trigger = 7 }}
{{ else if eq $command "self host" }}
	{{ $trigger = 8 }}
{{ else if eq $command "clyde" }}
	{{ $trigger = 9 }}
{{ else if eq $command "global ticket" }}
	{{ $trigger = 10 }}
{{ else if eq $command "logging" }}
	{{ $trigger = 11 }}
{{ else if eq $command "a2a" }}
	{{ $trigger = 12 }}
{{ end }}

{{/* ExecCC to call the main response trigger */}}
{{execCC 75 .Channel.ID 0 (sdict 
	"trigger" $trigger 
	"message" .ReactionMessage.ID
	"note" "Just some extra info to passthrough if needed."
)}}
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

{{ $command_response_map := sdict
	1 "banned"
	2 "wrong server"
	3 "setup"
	4 "open ticket"
	5 "premium"
	6 "not responding"
	7 "custom instance"
	8 "self host"
	9 "clyde"
	10 "global ticket"
	11 "logging"
	12 "a2a"
}}

{{ range $key, $value := $command_response_map }}
	{{ if eq $value $command }}
		{{ $trigger = $key }}
	{{ else }}
	{{ $embed := sdict }}
	{{ $embed.Set "title" "How do I setup ModMail?" }}
	{{ $embed.Set "description" "I'm sorry, I don't understand that command. Please use one of the following commands:" }}
	{{ $embed.Set "color" 2003199 }}
	{{ $embed.Set "fields" (cslice) }}
	{{ range $key, $value := $command_response_map }}
		{{ $embed.Set "description" (print $embed.description "\n" $value) }}
	{{ end }}
	{{ sendMessageNoEscapeRetID nil (complexMessage "reply" .Message.ID "embed" $embed) }}
	{{ end }}
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
{{ else }}
	$trigger = 0
	{{ $template := sdict "color" 2003199 }}
	{{ $embed := sdict }}
	{{ $embed.Set "title" "How do I setup ModMail?" }}
	{{ $embed.Set "description" "I'm sorry, I don't understand that command. Please use one of the following commands:" }}
	{{ $embed.Set "fields" (cslice
	{{ sendMessageNoEscapeRetID nil (complexMessage "reply" $replytarget "embed" $embed) }}
{{ end }}

{{/* ExecCC to call the main response trigger */}}
{{execCC 75 .Channel.ID 0 (sdict 
	"trigger" $trigger 
	"message" .ReactionMessage.ID
	"note" "Just some extra info to passthrough if needed."
)}}
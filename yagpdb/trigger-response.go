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
{{ $regex : = `^[.,!?]` }}
{{ $prefix := reFindAllSubmatches $regex .Message.Content }}
{{ $command := index (split (reReplace $regex .Message.Content "") " ") 0 }}
{{ $embed := sdict }}
{{ $embed.Set "color" 2003199 }}
{{- $cmdfields := ( cslice ) }}
{{ $cmd_map := sdict
	"1" (sdict "command" "banned" "aliases" (cslice "ban" "racefactory" "bloxburg" "appeal"))
	"2" (sdict "command" "wrongserver" "aliases" (cslice "ws"))
	"3" (sdict "command" "setup" "aliases" (cslice "gs" "getstarted" "config" "configure" "firststep" "fs" ))
	"4" (sdict "command" "ticket" "aliases" (cslice "thread" "message" "contact" "open" "create" "new" "start" "send" "mail" "support" ))
	"5" (sdict "command" "premium" "aliases" (cslice "patreon" "patron" "donate"))
	"6" (sdict "command" "notresponding" "aliases" (cslice "nr" "notworking" "noresponse" "nores" ))
	"7" (sdict "command" "custom" "aliases" (cslice "change" "customize" "instance" "name" "profile" "banner" "icon" "avatar" "pfp" "status" "private" "noverify" "bypass" ))
	"8" (sdict "command" "selfhost" "aliases" (cslice "source" "vps" "sh" "github" ))
	"9" (sdict "command" "clyde" "aliases" (cslice "blocked" "dm" "directmessage" "blockedme" "botblocked" ))
	"10" (sdict "command" "globalticket" "aliases" (cslice "global" "gt" "everyone" "all" "sees" "see" ))
	"11" (sdict "command" "logging" "aliases" (cslice "logging+" "logs" "transcript" "file" "viewer" "loggingplus" "lp" "l+" "log" ))
	"12" (sdict "command" "ask2ask" "aliases" (cslice "a2a" "ask" "support" ))
}}

{{/* Checks if the reaction is the bin emoji */}}
{{ if not (or $prefix .ServerPrefix) }}
	{{ return }}
{{ end }}

{{ range $key, $value := $cmd_map }}
    {{ $cmdfields = $cmdfields.Append (sdict 
		"name" (joinStr " - " (joinStr ": " "ID" $key) (joinStr ": " "Command: " $value.command))  
		"value" (joinStr "\n" "Aliases: " (joinStr ", " $value.aliases)
		"inline" false)) 
	}}
    {{- if eq "taglist" $command }}
        {{ $embed.Set "title" "Tag List" }}
        {{ $embed.Set "description" "Here is a list of all the tags available:" }}
        {{ $embed.Set "fields" $cmdfields }}
    {{- else if eq $value.command $command }}
        {{ $trigger = toInt $key}}
    {{- else if in $value.aliases $command}}
		{{- $trigger = toInt $key}}
    {{- else }}
        {{- $embed.Set "title" "Invalid Command!" }}
        {{- $embed.Set "description" (joinStr "" "I'm sorry, I don't understand that command. Please use one of the commands in `" $prefix "taglist`")}}
    {{- end -}}
{{ end }}

{{ sendMessage nil (complexMessage "reply" .Message.ID "embed" $embed) }}
{{ sendMessage nil (complexMessage "content" (joinStr "" "You triggered response " $trigger)) }}

{{/* ExecCC to call the main response trigger */}}
{{execCC 75 .Channel.ID 0 (sdict 
	"trigger" $trigger 
	"message" .ReactionMessage.ID
	"note" "Just some extra info to passthrough if needed."
)}}


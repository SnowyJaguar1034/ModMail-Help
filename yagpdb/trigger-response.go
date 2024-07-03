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
{{ $regex := `^[.,!?]` }}
{{ $prefix := reFindAllSubmatches $regex .Message.Content }}
{{ $command := index (split (reReplace $regex .Message.Content "") " ") 0 }}
{{ $embed := sdict }}
{{ $embed.Set "color" 2003199 }}
{{- $cmdfields := ( cslice ) }}
{{ $replytarget := (or .Message.ReferencedMessage .Message).ID }}

{{/* Command map */}}
{{ $cmd_map := cslice
	(sdict "trigger" 1 "command" "banned" "aliases" (cslice "ban" "racefactory" "bloxburg" "appeal"))
	(sdict "trigger" 2 "command" "wrongserver" "aliases" (cslice "ws"))
	(sdict "trigger" 3 "command" "setup" "aliases" (cslice "gs" "getstarted" "config" "configure" "firststep" "fs" ))
	(sdict "trigger" 4 "command" "ticket" "aliases" (cslice "thread" "message" "contact" "open" "create" "new" "start" "send" "mail" "support" ))
	(sdict "trigger" 5 "command" "premium" "aliases" (cslice "patreon" "patron" "donate"))
	(sdict "trigger" 6 "command" "notresponding" "aliases" (cslice "nr" "notworking" "noresponse" "nores" ))
	(sdict "trigger" 7 "command" "custom" "aliases" (cslice "change" "customize" "instance" "name" "profile" "banner" "icon" "avatar" "pfp" "status" "private" "noverify" "bypass" ))
	(sdict "trigger" 8 "command" "selfhost" "aliases" (cslice "source" "vps" "sh" "github" ))
	(sdict "trigger" 9 "command" "clyde" "aliases" (cslice "blocked" "dm" "directmessage" "blockedme" "botblocked" ))
	(sdict "trigger" 10 "command" "globalticket" "aliases" (cslice "global" "gt" "everyone" "all" "sees" "see" ))
	(sdict "trigger" 11 "command" "logging" "aliases" (cslice "logging+" "logs" "transcript" "file" "viewer" "loggingplus" "lp" "l+" "log" ))
	(sdict "trigger" 12 "command" "ask2ask" "aliases" (cslice "a2a" "ask" "support" ))
}}

{{/* Escapes the response if the trigger is one of the prefixes, the server prefix or has no command*/}}
{{ if not (and $prefix .ServerPrefix $command ) }}
	{{ return }}
{{ end }}

{{ range $cmd_map }}
	{{ $cmdfields = $cmdfields.Append (sdict 
		"name" (printf "Command: `%s`" .command)  
		"value" (printf "Aliases: `%s`" (joinStr "`, `" .aliases))
		"inline" true
	) }}
    {{- if eq "taglist" $command }}
        {{ $embed.Set "title" "Tag List" }}
        {{ $embed.Set "description" "Here is a list of all the tags available:" }}
        {{ $embed.Set "fields" $cmdfields }}
		{{- else if or (eq .command $command) (in .aliases $command)}}
			{{- $trigger = .trigger }}
    {{- else }}
        {{- $embed.Set "title" "Invalid Command!" }}
        {{- $embed.Set "description" (printf "The command `%s` is not valid. Please use one of the commands in `%s`" $command "taglist") }}
    {{- end -}}
{{ end }}

{{ if eq 0 $trigger}}
	{{ sendMessageNoEscape nil (complexMessage "reply" $replytarget "embed" $embed) }}
{{ end }}
{{ sendMessage nil (complexMessage "content" (joinStr "" "You triggered response " $trigger "\nThe message id is: " $replytarget)) }}

{{/* ExecCC to call the main response trigger */}}
{{execCC 75 .Channel.ID 0 (sdict 
	"trigger" $trigger 
	"message" $replytarget
	"note" "Just some extra info to passthrough if needed."
)}}
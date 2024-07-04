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
{{- $cmdfieldst := ( cslice ) }}
{{ $replytarget := (or .Message.ReferencedMessage .Message).ID }}

{{/* Command map */}}
{{ $cmd_map := cslice
	(sdict "trigger" 1 "command" "banned" "reaction" "banned:1251258151425282289" "aliases" (cslice "ban" "racefactory" "bloxburg" "appeal"))
	(sdict "trigger" 2 "command" "wrongserver" "reaction" "wrong_server:1251257683487494217" "aliases" (cslice "ws"))
	(sdict "trigger" 3 "command" "setup" "reaction" "setup:1251258670981976126" "aliases" (cslice "gs" "getstarted" "config" "configure" "firststep" "fs" ))
	(sdict "trigger" 4 "command" "ticket" "reaction" "ticket~1:1251258339518582787" "aliases" (cslice "thread" "message" "contact" "open" "create" "new" "start" "send" "mail" "support" ))
	(sdict "trigger" 5 "command" "premium" "reaction" "premium:1251273319110414429" "aliases" (cslice "patreon" "patron" "donate"))
	(sdict "trigger" 6 "command" "notresponding" "reaction" "no_response:1251273446860783718" "aliases" (cslice "nr" "notworking" "noresponse" "nores" ))
	(sdict "trigger" 7 "command" "custom" "reaction" "custom_instance:1251256312017457284" "aliases" (cslice "change" "customize" "instance" "name" "profile" "banner" "icon" "avatar" "pfp" "status" "private" "noverify" "bypass" ))
	(sdict "trigger" 8 "command" "selfhost" "reaction" "selfhost:1251257779730124884" "aliases" (cslice "source" "vps" "sh" "github" ))
	(sdict "trigger" 9 "command" "clyde" "reaction" "clyde:1251278199292297289" "aliases" (cslice "blocked" "dm" "directmessage" "blockedme" "botblocked" ))
	(sdict "trigger" 10 "command" "globalticket" "reaction" "global_ticket:1251274307347153027" "aliases" (cslice "global" "gt" "everyone" "all" "sees" "see" ))
	(sdict "trigger" 11 "command" "logging" "reaction" "text_file:1254891424965722122" "aliases" (cslice "logging+" "logs" "transcript" "file" "viewer" "loggingplus" "lp" "l+" "log" ))
	(sdict "trigger" 12 "command" "ask2ask" "reaction" "help:1251274523861581999" "aliases" (cslice "a2a" "ask" "support" ))
}}

{{/* Escapes the response if the trigger is one of the prefixes, the server prefix or has no command*/}}
{{ if not (and $prefix .ServerPrefix $command ) }}
	{{ return }}
{{ end }}

{{ range $cmd_map }}
	{{ $cmdfields = $cmdfields.Append (sdict 
		"name" (print (print "" "<:" .reaction ">") "" (printf "Command: `%s`" .command))  
		"value" (printf "Aliases: `%s`" (joinStr "`, `" .aliases))
		"inline" true
	) }}
	{{ $cmdfieldst = $cmdfieldst.Append (sdict 
		"name" (print (print "" "<:" .reaction ">") " " (printf "`%s`" .command))  
		"value" (printf "`%s`" (joinStr "`, `" .aliases))
		"inline" true
	) }}
    {{- if eq "taglist" $command }}
        {{ $embed.Set "title" "Tag List" }}
        {{ $embed.Set "description" "Here is a list of all the tags available:" }}
        {{ $embed.Set "fields" $cmdfields }}
	{{- else if eq "taglist2" $command }}
        {{ $embed.Set "title" "Tag List" }}
        {{ $embed.Set "description" "Here is a list of all the tags available with their corrosponding reaction and aliases:" }}
        {{ $embed.Set "fields" $cmdfieldst }}
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
{{execCC 104 .Channel.ID 0 (sdict 
	"trigger" $trigger 
	"message" $replytarget
	"note" "Just some extra info to passthrough if needed."
)}}
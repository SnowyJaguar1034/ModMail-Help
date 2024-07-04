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
{{ $replytarget := (or .ReactionMessage.ReferencedMessage .ReactionMessage).ID }}
{{ $msgreaction := .Reaction.Emoji.APIName }}

{{ $reaction_map := cslice
	(sdict "trigger" 1 "reaction" "banned:1251258151425282289" )
	(sdict "trigger" 2 "reaction" "wrong_server:1251257683487494217" )
	(sdict "trigger" 3 "reaction" "setup:1251258670981976126" )
	(sdict "trigger" 4 "reaction" "ticket~1:1251258339518582787" )
	(sdict "trigger" 5 "reaction" "premium:1251273319110414429" )
	(sdict "trigger" 6 "reaction" "no_response:1251273446860783718" )
	(sdict "trigger" 7 "reaction" "custom_instance:1251256312017457284" )
	(sdict "trigger" 8 "reaction" "selfhost:1251257779730124884" )
	(sdict "trigger" 9 "reaction" "clyde:1251278199292297289" )
	(sdict "trigger" 10 "reaction" "global_ticket:1251274307347153027" )
	(sdict "trigger" 11 "reaction" "text_file:1254891424965722122" )
	(sdict "trigger" 12 "reaction" "help:1251274523861581999" )
}}

{{/* Checks if the reaction is the bin emoji */}}
{{ if and (eq .ReactionMessage.Author.ID .BotUser.ID) (not (in $reaction_map .Reaction.Emoji.APIName)) }}
	{{ return }}
{{ end }}


{{ range $reaction_map }}
	
	{{ if eq $msgreaction .reaction }}
		{{ $trigger = .trigger }}
    {{- end -}}
{{ end }}

{{/* ExecCC to call the main response trigger */}}
{{execCC 104 .Channel.ID 0 (sdict 
	"trigger" $trigger 
	"message" $replytarget
	"note" "Just some extra info to passthrough if needed."
)}}




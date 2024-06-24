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


{{/* Checks if the reaction is the bin emoji */}}
{{ if not (eq .ReactionMessage.Author.ID .BotUser.ID) }}
	{{ return }}
{{ end }}

{{/* Declaring variables */}}
{{ $trigger := 0 }}

{{ if eq .Reaction.Emoji.APIName "banned:1251258151425282289" }}
	{{ $trigger = 1}}
{{ else if eq .Reaction.Emoji.APIName "wrong_server:1251257683487494217" }}
	{{ $trigger = 2 }}
{{ else if eq .Reaction.Emoji.APIName "setup:1251258670981976126" }}
	{{ $trigger = 3 }}
{{ else if eq .Reaction.Emoji.APIName "ticket~1:1251258339518582787" }}
	{{ $trigger = 4 }}
{{ else if eq .Reaction.Emoji.APIName "premium:1251273319110414429" }}
	{{ $trigger = 5 }}
{{ else if eq .Reaction.Emoji.APIName "no_response:1251273446860783718" }}
	{{ $trigger = 6 }}
{{ else if eq .Reaction.Emoji.APIName "custom_instance:1251256312017457284" }}
	{{ $trigger = 7 }}
{{ else if eq .Reaction.Emoji.APIName "selfhost:1251257779730124884" }}
	{{ $trigger = 8 }}
{{ else if eq .Reaction.Emoji.APIName "clyde:1251278199292297289" }}
	{{ $trigger = 9 }}
{{ else if eq .Reaction.Emoji.APIName "global_ticket:1251274307347153027" }}
	{{ $trigger = 10 }}
{{ else if eq .Reaction.Emoji.APIName "text_file:1254891424965722122" }}
	{{ $trigger = 11 }}
{{ else if eq .Reaction.Emoji.APIName "help:1251274523861581999" }}
	{{ $trigger = 12 }}
{{ end }}

{{/* ExecCC to call the main response trigger */}}
{{execCC 75 .Channel.ID 0 (sdict 
	"trigger" $trigger 
	"message" .ReactionMessage.ID
	"note" "Just some extra info to passthrough if needed."
)}}




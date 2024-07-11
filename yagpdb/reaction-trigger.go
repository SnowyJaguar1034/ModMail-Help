{{/* $Trigger Key */}}
{{/* 1 : General Help response */}}
{{/* 2 : Banned response*/}}
{{/* 3 : Wrong Server response */}}
{{/* 4 : Setup response */}}
{{/* 5 : Open Ticket response */}}
{{/* 6 : Premium response */}}
{{/* 7 : Not Responding response */}}
{{/* 8 : Custom Instance response */}}
{{/* 9 : Sef Host response */}}
{{/* 10 : Clyde response */}}
{{/* 11 : Global Ticket response */}}
{{/* 12 : Logging response */}}
{{/* 13 : Verfication response */}}

{{/* Declaring variables */}}
{{ $trigger := 0 }}
{{ $replytarget := (or .ReactionMessage.ReferencedMessage .ReactionMessage).ID }}
{{ $msgreaction := .Reaction.Emoji.APIName }}

{{ $response_mapping := cslice
	(sdict 
		"trigger" 1 
		"command" "ask2ask" 
		"reaction" "help:1251274523861581999" 
		"aliases" (cslice "a2a" "ask" "support" "help" ) 
		"regex" ``
	)
	(sdict 
		"trigger" 2 
		"command" "banned" 
		"reaction" "banned:1251258151425282289" 
		"aliases" (cslice "ban" "racefactory" "bloxburg" "appeal") 
		"regex" `(?i)ban|racefactory|bloxburg|appeal`
	)
	(sdict 
		"trigger" 3 
		"command" "wrongserver" 
		"reaction" "wrong_server:1251257683487494217" 
		"aliases" (cslice "ws") 
		"regex" `(?i)wrong server|not the right server|not the server`
	)
	(sdict 
		"trigger" 4 
		"command" "setup" 
		"reaction" "setup:1251258670981976126" 
		"aliases" (cslice "gs" "getstarted" "config" "configure" "firststep" "fs" ) 
		"regex" `(?i:modmail (?i:invite|joined|setup|added)|invite modmail|setup modmail|added modmail|setup bot|bot setup|bot added|setup`
	)
	(sdict 
		"trigger" 5 
		"command" "ticket" 
		"reaction" "ticket~1:1251258339518582787" 
		"aliases" (cslice "thread" "message" "contact" "open" "create" "new" "start" "send" "mail" "support" )
		"regex" `(?:m(?:essage (?:a )?server|sg (?:a )?server)|c(?:reate (?:a )?ticket|ustom commands)|open (?:a )?ticket)`
	)
	(sdict 
		"trigger" 6 
		"command" "premium" 
		"reaction" "premium:1251273319110414429" 
		"aliases" (cslice "patreon" "patron" "donate")
		"regex" `(?:message logs|(?:transcrip|snippe)ts|p(?:atreon|remium)|donate)`
	)
	(sdict 
		"trigger" 7 
		"command" "notresponding" 
		"reaction" "no_response:1251273446860783718" 
		"aliases" (cslice "nr" "notworking" "noresponse" "nores" )
		"regex" `(?:doesn't (?:seem to )?work|doesn't respond|isn(?:'t (?:respond|working)|t (?:respond|working))|no respon(?:se|d))`
	)
	(sdict 
		"trigger" 8 
		"command" "custom" 
		"reaction" "custom_instance:1251256312017457284" 
		"aliases" (cslice "change" "customize" "instance" "name" "profile" "banner" "icon" "avatar" "pfp" "status" "private" "noverify" "bypass" )
		"regex" `(?i)(?:bot(?:'?s)?|(?:change|customi[sz]e)(?: the)?) (?:name|profile|banner|icon|avatar|pfp|status)|bot(?:'?s)? user|customi[sz]e(?: the)? (?:instance|bot)|private (?:instance|bot)|(?:no|bypass) verif(?:ication|y)`
	)
	(sdict 
		"trigger" 9 
		"command" "selfhost" 
		"reaction" "selfhost:1251257779730124884" 
		"aliases" (cslice "source" "vps" "sh" "github" )
		"regex" `(?i)(?:source|modmails?|bots?|bot's?|self(?:-)?host|host (?:modmail|bot)|(?:best|recommended|which) (?:virtual(?: private)? )?server)(?:'s)?(?: code| repo| github)`
	)
	(sdict 
		"trigger" 10 
		"command" "clyde" 
		"reaction" "clyde:1251278199292297289" 
		"aliases" (cslice "blocked" "dm" "directmessage" "blockedme" "botblocked" )
		"regex" `(?i:only accepting (?:direct message|dm)s from friends|message (?:(?:could not be|not) delivered|blocked)|(?:don't share a|no (?:shared|mutual)) server|clyde(?:[- ]bot)?|i(?:'| a)?m blocked|bot blocked me)`
	)
	(sdict 
		"trigger" 11 
		"command" "globalticket" 
		"reaction" "global_ticket:1251274307347153027" 
		"aliases" (cslice "global" "gt" "everyone" "all" "sees" "see" )
		"regex" `(ticket|tickets|everyone) (can|see|sees|see's) (the )?(mail|ticket|tickets|my mail|mod mail|modmail|mod-mail) message`
	)
	(sdict 
		"trigger" 12 
		"command" "logging" 
		"reaction" "text_file:1254891424965722122" 
		"aliases" (cslice "logging+" "logs" "transcript" "file" "viewer" "loggingplus" "lp" "l+" "log" )
		"regex" `(?i)(?:\-|<@!?204255221017214977>|!|.)\s*(?:logging|logs|transcript|file|viewer)(?: +|\z)`
	)
}}


{{/* Checks if the reaction is the bin emoji */}}
{{ if and (eq .ReactionMessage.Author.ID .BotUser.ID) (not (in $response_mapping .Reaction.Emoji.APIName)) }}
	{{ return }}
{{ end }}


{{ range $response_mapping }}
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




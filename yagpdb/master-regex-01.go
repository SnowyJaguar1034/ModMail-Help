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
{{ $banned := reFindAllSubmatches `(?i)ban|racefactory|bloxburg|appeal` .Message.Content }}
{{ $wrongserver := reFindAllSubmatches `(?i)wrong server|not the right server|not the server` .Message.Content }}
{{ $setup := reFindAllSubmatches `(?i:modmail (?i:invite|joined|setup|added)|invite modmail|setup modmail|added modmail|setup bot|bot setup|bot added|setup)` .Message.Content }}
{{ $ticket := reFindAllSubmatches `(?:m(?:essage (?:a )?server|sg (?:a )?server)|c(?:reate (?:a )?ticket|ustom commands)|open (?:a )?ticket)` .Message.Content }}
{{ $premium := reFindAllSubmatches `(?:message logs|(?:transcrip|snippe)ts|p(?:atreon|remium)|donate)` .Message.Content }}
{{ $noresponse := reFindAllSubmatches `(?:doesn't (?:seem to )?work|doesn't respond|isn(?:'t (?:respond|working)|t (?:respond|working))|no respon(?:se|d))` .Message.Content }}
{{ $custom := reFindAllSubmatches `(?i)(?:bot(?:'?s)?|(?:change|customi[sz]e)(?: the)?) (?:name|profile|banner|icon|avatar|pfp|status)|bot(?:'?s)? user|customi[sz]e(?: the)? (?:instance|bot)|private (?:instance|bot)|(?:no|bypass) verif(?:ication|y)` .Message.Content }}
{{ $selfhost := reFindAllSubmatches `(?i)(?:source|modmails?|bots?|bot's?|self(?:-)?host|host (?:modmail|bot)|(?:best|recommended|which) (?:virtual(?: private)? )?server)(?:'s)?(?: code| repo| github)` .Message.Content }}
{{ $clyde := reFindAllSubmatches `(?i:only accepting (?:direct message|dm)s from friends|message (?:(?:could not be|not) delivered|blocked)|(?:don't share a|no (?:shared|mutual)) server|clyde(?:[- ]bot)?|i(?:'| a)?m blocked|bot blocked me)` .Message.Content }}
{{ $globalticket := reFindAllSubmatches `(ticket|tickets|everyone) (can|see|sees|see's) (the )?(mail|ticket|tickets|my mail|mod mail|modmail|mod-mail) message` .Message.Content }}

{{ $trigger := 0 }}

{{ if $banned }}
	{{ $trigger = 1}}
{{ else if $wrongserver }}
	{{ $trigger = 2 }}
{{ else if $setup }}
	{{ $trigger = 3 }}
{{ else if $ticket }}
	{{ $trigger = 4 }}
{{ else if $premium }}
	{{ $trigger = 5 }}
{{ else if $noresponse }}
	{{ $trigger = 6 }}
{{ else if $custom }}
	{{ $trigger = 7 }}
{{ else if $selfhost }}
	{{ $trigger = 8 }}
{{ else if $clyde }}
	{{ $trigger = 9 }}
{{ else if $globalticket }}
	{{ $trigger = 10 }}
{{ end }}

{{/* ExecCC to call the main response trigger */}}
{{execCC 75 .Channel.ID 0 (sdict 
	"trigger" $trigger 
	"message" .Message.ID
	"note" "Just some extra info to passthrough if needed."
)}}

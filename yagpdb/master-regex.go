{{/* Regexes to match, pulled from the existing commands */}}
{{ $banned := reFindAllSubmatches `(?i)ban|racefactory|bloxburg|appeal` .Message.Content }}
{{ $wrongserver := reFindAllSubmatches `(?i)wrong server|not the right server|not the server` .Message.Content }}
{{ $setup := reFindAllSubmatches `(?i:modmail (?i:invite|joined|setup|added)|invite modmail|setup modmail|added modmail|setup bot|bot setup|bot added|setup)` .Message.Content }}
{{ $ticket := reFindAllSubmatches `(?:m(?:essage (?:a )?server|sg (?:a )?server)|c(?:reate (?:a )?ticket|ustom commands)|open (?:a )?ticket)` .Message.Content }}
{{ $premium := reFindAllSubmatches `(?:message logs|(?:transcrip|snippe)ts|p(?:atreon|remium)|donate)` .Message.Content }}
{{ $logging := reFindAllSubmatches `\A(?:\-|<@!?204255221017214977>|!|.)\s*(?:logging|logs|transcript|file|viewer)(?: +|\z)` .Message.Content }}
{{ $noresponse := reFindAllSubmatches `(?:doesn't (?:seem to )?work|doesn't respond|isn(?:'t (?:respond|working)|t (?:respond|working))|no respon(?:se|d))` .Message.Content }}
{{ $custom := reFindAllSubmatches `(?i)(?:bot(?:'?s)?|(?:change|customi[sz]e)(?: the)?) (?:name|profile|banner|icon|avatar|pfp|status)|bot(?:'?s)? user|customi[sz]e(?: the)? (?:instance|bot)|private (?:instance|bot)|(?:no|bypass) verif(?:ication|y)` .Message.Content }}
{{ $selfhost := reFindAllSubmatches `(?i)(?:source|modmails?|bots?|bot's?|self(?:-)?host|host (?:modmail|bot)|(?:best|recommended|which) (?:virtual(?: private)? )?server)(?:'s)?(?: code| repo| github)` .Message.Content }}
{{ $clyde := reFindAllSubmatches `(?i:only accepting (?:direct message|dm)s from friends|message (?:(?:could not be|not) delivered|blocked)|(?:don't share a|no (?:shared|mutual)) server|clyde(?:[- ]bot)?|i(?:'| a)?m blocked|bot blocked me)` .Message.Content }}
{{ $globalticket := reFindAllSubmatches `(ticket|tickets|everyone) (can|see|sees|see's) (the )?(mail|ticket|tickets|my mail|mod mail|modmail|mod-mail) message` .Message.Content }}
{{ $help := reFindAllSubmatches `(?i)(?:need (?:support|help|assistance|aid|advice)|(?:help|support) me)` .Message.Content }}

{{ $trigger := 0 }}

{{ if $banned }}
	{{ $trigger = 1}}
{{ else if $wrongserver }}
	{{ $trigger = 2}}
{{ else if $setup }}
	{{ $trigger = 3}}
{{ else if $ticket }}
	{{ $trigger = 4}}
{{ else if $premium }}
	{{ $trigger = 5}}
{{ else if $logging }}
	{{ $trigger = 6}}
{{ else if $noresponse }}
	{{ $trigger = 7}}
{{ else if $custom }}
	{{ $trigger = 8}}
{{ else if $selfhost }}
	{{ $trigger = 9}}
{{ else if $clyde }}
	{{ $trigger = 10}}
{{ else if $globalticket }}
	{{ $trigger = 11}}
{{ else if $help }}
	{{ $trigger = 12}}
{{ end }}

{{/* ExecCC to call the main response trigger */}}
{{execCC 75 .Channel.ID 0 (sdict 
	"trigger" $trigger 
	"note" "Just some extra info to passthrough if needed."
)}}

{{ $embed := sdict "color" 2003199 }}
{{ $replytarget := .Message.ID }}

{{sendMessage nil "Doing stuff upon trigger!" }}

{{/* Check if already replied to prevent "Wrong Server" reply */}}
{{ $alreadyreplied:=false }}

{{/* Regexes to match, pulled from the existing commands */}}
{{ $banned:= reFindAllSubmatches `(?i)ban|racefactory|bloxburg|appeal` .Message.Content }}
{{ $setup:= reFindAllSubmatches `(?i:modmail (?i:invite|joined|setup|added)|invite modmail|setup modmail|added modmail|setup bot|bot setup|bot added|setup)$` .Message.Content }}
{{ $ticket:= reFindAllSubmatches `(?:m(?:essage (?:a )?server|sg (?:a )?server)|c(?:reate (?:a )?ticket|ustom commands)|open (?:a )?ticket)$` .Message.Content }}
{{ $premium:= reFindAllSubmatches `(?:message logs|(?:transcrip|snippe)ts|p(?:atreon|remium)|donate)$` .Message.Content }}
{{ $noresponse:= reFindAllSubmatches `(?:doesn't (?:seem to )?work|doesn't respond|isn(?:'t (?:respond|working)|t (?:respond|working))|no respon(?:se|d))$` .Message.Content }}
{{ $custom:= reFindAllSubmatches `(?:bypass verif(?:ication|y)|private (?:instance|bot)|no verif(?:ication|y)|custom (?:instance|bot)|bot(?:'s (?:profil|nam)|s (?:profil|nam)| (?:profil|nam))e|bot(?:'s (?:avatar|banner|status|user|pfp)|s (?:avatar|banner|status|user|pfp)| (?:avatar|banner|status|user|pfp))|bot(?:'s|s?) icon)$` .Message.Content }}

{{ if $banned }}
{{sendMessage nil "Doing stuff post check!" }}
{{ $embed.Set "title" "How do I get unbanned?" }}
{{ $embed.Set "description" "You are in the wrong server for what you’re seeking help for.\n\nWe are the Support server for the ModMail **bot**.\n\nWe have no affiliation with the server you are banned from.\n\nYou cannot use ModMail to contact a server you are banned from.\n\nWe cannot help you any further, sorry." }}
{{ sendMessageNoEscape nil (complexMessage "reply" $replytarget "embed" $embed) }}
{{ $alreadyreplied:=false }}
{{ end }}

{{ if $setup }}
{{sendMessage nil "Doing stuff post check!" }}
{{ $embed.Set "title" "How do I setup ModMail" }}
{{ $embed.Set "footer" (sdict "text" "Click the reaction below to see advanced setup information") }}
{{ $embed.Set "fields" (cslice (sdict 
		"name" "Initial Setup"
		"value" "**1.** [Invite the bot](https://modmail.xyz/invite).\n**2.** Run `=setup` in your server.\n**3.** Done! :tada:\n\nYou can use `=help` for a [list of commands.](https://modmail.xyz/commands)"
		"inline" true
		) (sdict 
		"name" "Premium"
		"value" "Please consider purchasing premium for more features!\nThis includes full conversation logging, custom greeting and closing messages, as well as snippets."
		"inline" true
		)
	)}}
{{ $msgID := sendMessageNoEscapeRetID nil (complexMessage "reply" $replytarget "embed" $embed) }}
{{addMessageReactions nil $msgID (cslice ":modmail:702099194701152266")}}
{{ $embed.Set "fields" ($embed.fields.Append (sdict 
		"name" "Advanced Setup" 
		"value" "Some additional commands you could use are:" 
		"inline" false
	) (sdict
		"name" "- `=pingrole <roles>`"
		"value" "For the bot to ping certain roles when tickets are created."
		"inline" true
	) (sdict
		"name" "- `=accessrole <roles>`"
		"value" "For configuring which roles can reply to ModMail tickets."
		"inline" true
	) (sdict
		"name" "- `=anonymous`"
		"value" "For toggling anonymous staff replies to hide the responder's name.\nThis does not work for making your end-user anonymous."
		"inline" true
	) (sdict
		"name" "- `=logging`"
		"value" "For toggling log messages of tickets being opened or closed.\nThis does not log a transcript of the messages."
		"inline" true
	) (sdict
		"name" "- `=commandonly`"
		"value" "For toggling if commands are required to reply to tickets.\nIf **disabled** staff have only to type in the channel for their message to be sent.\nIf **enabled** staff have to reply with `=reply` or `=areply`."
		"inline" true
	) (sdict
		"name" "You can mention the roles, use role IDs or role names."
		"value" "For role names with a space, it needs to be in quotes (e.g. \"Head Admin\")"
		"inline" false
	)
	)}}
{{editMessage nil $msgID (complexMessageEdit "embed" $embed)}}

{{ $alreadyreplied:=false }}
{{ end }}


{{ if $ticket }}
{{sendMessage nil "Doing stuff post check!" }}
{{ $embed.Set "title" "How do I open a ticket?" }}
{{ $embed.Set "footer" (sdict "text" "Click the reaction below to see advanced setup information") }}
{{ $embed.Set "fields" (cslice (sdict 
		"name" "Method One: Message the Bot" 
		"value" "The quickest and simplest way to open a ticket is to DM the bot a message and follow the given prompts."
		"inline" false
	) (sdict 
		"name" "Method Two: Using a command in DMs" 
		"value" "You can use `=send <server ID> <message>` to create a ticket on a specific server. You still actually need to be a member of that server and ModMail still needs to be on the server as well, but this skips the server selection menu, which we know can confuse some people."
	)
	)}}
{{ $msgID := sendMessageNoEscape nil (complexMessage "reply" $replytarget "embed" $embed) }}
{{addMessageReactions nil $msgID (cslice ":modmail:702099194701152266")}}
{{editMessage nil $msgID (complexMessageEdit "embed" ($embed.fields.Append (sdict 
	"name" "Bonus Information: `=confirmation` command" 
	"value" "If you have enabled the `=confirmation` mode, you will not be given the server selection menu immediately and will instead be prompted to resume messaging the last server you contacted.\nThis prompt will contain an option to take you to the server selection menu if it's incorrect but you can also use `=new <message>` to force the server selection menu to appear."
	"inline" false
	) (sdict
		"name" "Note"
		"value" "If you are having trouble with the `=send` command, please ensure you are using the correct server ID. You can find this by right-clicking on the server name and selecting `Copy ID`."
		"inline" false
	) )) }}
{{ $alreadyreplied:=false }}
{{ end }}

{{ if and $premium ( not (hasPrefix .Message.Content "=")) }}
{{sendMessage nil "Doing stuff post check!" }}
{{ $embed.Set "title" "Donation Link" }}
{{ $embed.Set "description" "[Purchase ModMail Premium Here](https://modmail.xyz/premium)" }}
{{ sendMessageNoEscape nil (complexMessage "reply" $replytarget "embed" $embed) }}
{{ $alreadyreplied:=false }}
{{ end }}

{{ if and $noresponse }}
{{sendMessage nil "Doing stuff post check!" }}
{{ $embed.Set "title" "ModMail is not responding" }}
{{ $embed.Set "description" "If ModMail is not responding in your server, please check the following:\n- The bot has Read Messages, Send Messages, and Embed Links permissions.\n- You are using the correct prefix. Use `@ModMail prefix` to check the prefix.\n- The command you are using is valid. Check using `=help <command>`.\n- The bot is online. Discord might be having issues, or the bot might be restarting.\n\nIf the bot still does respond, please let us know your [server ID](https://support.discord.com/hc/en-us/articles/206346498-Where-can-I-find-my-User-Server-Message-ID-)." }}
{{ sendMessageNoEscape nil (complexMessage "reply" $replytarget "embed" $embed) }}
{{ $alreadyreplied:=false }}
{{ end }}

{{ if $custom }}
{{sendMessage nil "Doing stuff post check!" }}
{{ $embed.Set "title" "ModMail Custom Instance" }}
{{ $embed.Set "description" "You can contact <@381998065327931392> (James [a_leon]#6196) or <@365262543872327681> (snowyjaguar#0) for a custom instance. The pricing is $60/year." }}
{{ $embed.Set "fields" (cslice (sdict 
		"name" "Custom Instance Benefits" 
		"value" "- Custom username, avatar, status message and status activity type.\n- All the premium features listed [here](https://modmail.xyz/premium).\n- No confirmation messages.\n- Commands to create tickets with users.\n- Requiring a command to send messages.\n- Showing users roles in 'New Ticket' messages"
		"inline" false
	))}}
{{ sendMessageNoEscape nil (complexMessage "reply" $replytarget "embed" $embed) }}
{{ $alreadyreplied:=false }}
{{ end }}

{{/* {{ if and $help ( not $alreadyreplied) }} */}}

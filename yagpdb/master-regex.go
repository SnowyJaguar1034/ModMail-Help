{{/* Defining varibles for use throughout the script */}}
{{ $template := sdict "color" 2003199 }}
{{ $replytarget := .Message.ID }}
{{ $alreadyreplied := false }}
{{ $bin := ":bin:1251255316121653343" }}
{{ $bookmark := ":bookmark:1251243802207846566" }}
{{ $mail := ":mail:1251255870701047909" }}

{{/* {{sendMessage nil "Doing stuff upon trigger!" }} */}}
{{/* https://docs.yagpdb.xyz/reference/templates#templates.sdict */}}
{{/* https://discord.com/channels/166207328570441728/578976698931085333/1248038506312237127 */}}

{{/* Regexes to match, pulled from the existing commands */}}
{{ $banned := reFindAllSubmatches `(?i)ban|racefactory|bloxburg|appeal` .Message.Content }}
{{ $setup := reFindAllSubmatches `(?i:modmail (?i:invite|joined|setup|added)|invite modmail|setup modmail|added modmail|setup bot|bot setup|bot added|setup)` .Message.Content }}
{{ $ticket := reFindAllSubmatches `(?:m(?:essage (?:a )?server|sg (?:a )?server)|c(?:reate (?:a )?ticket|ustom commands)|open (?:a )?ticket)` .Message.Content }}
{{ $premium := reFindAllSubmatches `(?:message logs|(?:transcrip|snippe)ts|p(?:atreon|remium)|donate)` .Message.Content }}
{{ $noresponse := reFindAllSubmatches `(?:doesn't (?:seem to )?work|doesn't respond|isn(?:'t (?:respond|working)|t (?:respond|working))|no respon(?:se|d))` .Message.Content }}
{{ $custom := reFindAllSubmatches `(?i)(?:bot(?:'?s)?|(?:change|customi[sz]e)(?: the)?) (?:name|profile|banner|icon|avatar|pfp|status)|bot(?:'?s)? user|customi[sz]e(?: the)? (?:instance|bot)|private (?:instance|bot)|(?:no|bypass) verif(?:ication|y)` .Message.Content }}
{{ $selfhost := reFindAllSubmatches `(?i)(?:source|modmails?|bots?|bot's?|self(?:-)?host|host (?:modmail|bot)|(?:best|recommended|which) (?:virtual(?: private)? )?server)(?:'s)?(?: code| repo| github)` .Message.Content }}
{{ $clyde := reFindAllSubmatches `(?i:only accepting (?:direct message|dm)s from friends|message (?:(?:could not be|not) delivered|blocked)|(?:don't share a|no (?:shared|mutual)) server|clyde(?:[- ]bot)?|i(?:'| a)?m blocked|bot blocked me)` .Message.Content }}
{{ $globalticket := reFindAllSubmatches `(ticket|tickets|everyone) (can|see|sees|see's) (the )?(mail|ticket|tickets|my mail|mod mail|modmail|mod-mail) message` .Message.Content }}
{{ $help := reFindAllSubmatches `(?i)(?:need (?:support|help|assistance|aid|advice)|(?:help|support) me)` .Message.Content }}


{{ if $banned }}
	{{ $embed := sdict }}
	{{ range $k, $v := $template }}
		{{ $embed.Set $k $v}}
	{{ end }}
	{{ $embed.Set "title" "How do I get unbanned?" }}
	{{ $embed.Set "description" "You are in the wrong server for what you’re seeking help for." }}
	{{ $embed.Set "fields" (cslice (sdict 
		"name" "We are the Support server for the ModMail __bot__."
		"value" "We have no affiliation with the server you are banned from."
		"inline" true
		) (sdict 
		"name" "You cannot use ModMail to contact a server you are banned from."
		"value" "We cannot help you any further, sorry."
		"inline" true
		)
	)}}
	{{ $msgID := sendMessageNoEscapeRetID nil (complexMessage "reply" $replytarget "embed" $embed) }}
	{{ addMessageReactions nil $msgID (cslice $bin $bookmark) }}
	{{ $alreadyreplied := true }}
{{ end }}

{{ if $setup }}
	{{ $embed := sdict }}
	{{ range $k, $v := $template }}
		{{ $embed.Set $k $v}}
	{{ end }}
	{{ $embed.Set "title" "How do I setup ModMail?" }}
	{{ $embed.Set "footer" (sdict "text" "Click the reaction below to see bonus information") }}
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
	{{ addMessageReactions nil $msgID (cslice $bin $bookmark $mail) }}
	{{ $alreadyreplied := true }}
{{ end }}


{{ if $ticket }}
	{{ $embed := sdict }}
	{{ range $k, $v := $template }}
		{{ $embed.Set $k $v}}
	{{ end }}
	{{ $embed.Set "title" "How do I open a ticket?" }}
	{{ $embed.Set "footer" (sdict "text" "Click the reaction below to see bonus information") }}
	{{ $embed.Set "fields" (cslice (sdict 
			"name" "Method One: Message the Bot" 
			"value" "The quickest and simplest way to open a ticket is to DM the bot a message and follow the given prompts."
			"inline" false
		) (sdict 
			"name" "Method Two: Using a command in DMs" 
			"value" "You can use `=send <server ID> <message>` to create a ticket on a specific server. You still actually need to be a member of that server and ModMail still needs to be on the server as well, but this skips the server selection menu, which we know can confuse some people."
		) (sdict
			"name" "Note"
			"value" "If you are having trouble with the `=send` command, please ensure you are using the correct server ID. You can find this by right-clicking on the server name and selecting `Copy ID`."
			"inline" false
		))}}
	{{ $msgID := sendMessageNoEscapeRetID nil (complexMessage "reply" $replytarget "embed" $embed) }}
	{{ addMessageReactions nil $msgID (cslice $bin $bookmark $mail) }}
	{{ $alreadyreplied := false }}
{{ end }}


{{ if and $premium ( not (hasPrefix .Message.Content "=")) }}
	{{ $embed := sdict }}
	{{ range $k, $v := $template }}
		{{ $embed.Set $k $v}}
	{{ end }}
	{{ $embed.Set "title" "Donation Link" }}
	{{ $embed.Set "description" "[Purchase ModMail Premium Here](https://modmail.xyz/premium)" }}
	{{ $msgID := sendMessageNoEscapeRetID nil (complexMessage "reply" $replytarget "embed" $embed) }}
	{{ addMessageReactions nil $msgID (cslice $bin $bookmark) }}
	{{ $alreadyreplied := true }}
{{ end }}

{{ if and $logging ( not (hasPrefix .Message.Content "=")) }}
	{{ $embed := sdict }}
	{{ range $k, $v := $template }}
		{{ $embed.Set $k $v}}
	{{ end }}
	{{ $embed.Set "title" "Advanced Logging Example" }}
	{{ $embed.Set "description" "This is an example of what you'll get with advanced logging. https://modmail.xyz/logs/d7586c153425000-10d1416086c01033-10d141608b802047" }}
	{{ $file := "
[2024-02-26 23:44:18] scyye#0 (User): Hello! I'm a user in need of assistance, can someone help me?
[2024-02-26 23:44:43] scyye#0 (Comment): I am a staff member writing a comment on the ticket, for other staff to see
[2024-02-26 23:45:16] scyye#0 (Staff): I am now replying to the user, asking them what they need help with.
[2024-02-26 23:45:46] jrwallor#0 (Staff): Another staff member with an anonymous reply.
[2024-02-26 23:45:58] scyye#0 (User): This is the user replying, thanking me for support (I didn't think this through, cut me some slack)
[2024-02-26 23:46:36] scyye#0 (Comment): =c This ticket is resolved, so I'm closing it now.
	" }}
	{{ $msgID := sendMessageNoEscapeRetID nil (complexMessage "reply" $replytarget "embed" $embed "file" $file) }}
	{{ addMessageReactions nil $msgID (cslice $bin $bookmark) }}
	{{ $alreadyreplied := true }}
{{ end }}

{{ if and $noresponse }}
	{{ $embed := sdict }}
	{{ range $k, $v := $template }}
		{{ $embed.Set $k $v}}
	{{ end }}
	{{ $embed.Set "title" "ModMail is not responding" }}
	{{ $embed.Set "description" "If ModMail is not responding in your server, please check the following:\n- The bot has Read Messages, Send Messages, and Embed Links permissions.\n- You are using the correct prefix. Use `@ModMail prefix` to check the prefix.\n- The command you are using is valid. Check using `=help <command>`.\n- The bot is online. Discord might be having issues, or the bot might be restarting.\n\nIf the bot still does respond, please let us know your [server ID](https://support.discord.com/hc/en-us/articles/206346498-Where-can-I-find-my-User-Server-Message-ID-)." }}
	{{ $msgID := sendMessageNoEscapeRetID nil (complexMessage "reply" $replytarget "embed" $embed) }}
	{{ addMessageReactions nil $msgID (cslice $bin $bookmark) }}
	{{ $alreadyreplied := true }}
{{ end }}

{{ if $custom }}
	{{ $embed := sdict }}
	{{ range $k, $v := $template }}
		{{ $embed.Set $k $v}}
	{{ end }}
	{{ $embed.Set "title" "ModMail Custom Instance" }}
	{{ $embed.Set "fields" (cslice (sdict 
		"name" "Custom Instance Benefits" 
		"value" "- Custom username, avatar, status message and status activity type.\n- All the [premium features](https://modmail.xyz/premium).\n- No confirmation messages.\n- Commands to create tickets with users.\n- Bypass verfication."
		"inline" false
	) (sdict 
		"name" "Fee" 
		"value" "The fee is $60/year.\nThis is a single payment and you will not be charged again until the next year."
		"inline" true
	) (sdict 
		"name" "Contact" 
		"value" "|<@381998065327931392> (`James [a_leon]`)\n|or\n|<@365262543872327681> (`snowyjaguar`)"
		"inline" true
	))}}
	{{ $msgID := sendMessageNoEscapeRetID nil (complexMessage "reply" $replytarget "embed" $embed) }}
	{{ addMessageReactions nil $msgID (cslice $bin $bookmark) }}
	{{ $alreadyreplied := true }}
{{ end }}

{{ if $selfhost }}
	{{ $embed := sdict }}
	{{ range $k, $v := $template }}
		{{ $embed.Set $k $v}}
	{{ end }}
	{{ $embed.Set "title" "How do I self-host ModMail?" }}
	{{ $embed.Set "description" "We do not officially support self-hosting.\n\nHowever, you can find the source code for ModMail on [GitHub](https://github.com/chamburr/modmail) and we provide some guides on how to self host." }}
	{{ $embed.Set "fields" (cslice (sdict 
		"name" "ModMail V3 Official Guide"
		"value" "[Click here](https://github.com/chamburr/modmail#self-hosting)"
		"inline" true
	) (sdict 
		"name" "ModMail V2 Official Guide"
		"value" "[Click here](https://github.com/chamburr/modmail/blob/v2.1.2/README.md#self-hosting)"
		"inline" true
	) (sdict 
		"name" "ModMail V2 Community Guide"
		"value" "[Click here](https://gist.github.com/waterflamev8/cab61e680e2fb5ea6027cbf144732925)"
		"inline" true
	))}}
	{{ $msgID := sendMessageNoEscapeRetID nil (complexMessage "reply" $replytarget "embed" $embed) }}
	{{ addMessageReactions nil $msgID (cslice $bin $bookmark $mail) }}
	{{ $alreadyreplied := true }}
{{ end }}

{{ if $clyde }}
	{{ $embed := sdict }}
	{{ range $k, $v := $template }}
		{{ $embed.Set $k $v}}
	{{ end }}
	{{ $embed.Set "title" "My message wasn't delivered!" }}
	{{ $embed.Set "description" "If you receive \"your message could not be delivered\", check your privacy settings for the server you want to contact server. You need to enable the \"allow direct messages from server members\" option." }}
	{{ $embed.Set "image" (sdict "url" "https://media.discordapp.net/attachments/576764854673735680/837129125327011860/unknown.png") }}
	{{ $msgID := sendMessageNoEscapeRetID nil (complexMessage "reply" $replytarget "embed" $embed) }}
	{{ addMessageReactions nil $msgID (cslice $bin $bookmark) }}
	{{ $alreadyreplied := true }}
	{{ end }}

{{ if $globalticket }}
	{{ $embed := sdict }}
	{{ range $k, $v := $template }}
		{{ $embed.Set $k $v}}
	{{ end }}
	{{ $embed.Set "title" "Everyone can see my tickets!" }}
	{{ $embed.Set "description" "If everyone in your server is able to view tickets, there is a chance that another bot in your server is interfering. This is not a problem with ModMail. We recommend checking the audit logs in your server settings to see which bot is changing the channel permissions." }}
	{{ $msgID := sendMessageNoEscapeRetID nil (complexMessage "reply" $replytarget "embed" $embed) }}
	{{ addMessageReactions nil $msgID (cslice $bin $bookmark) }}
	{{ $alreadyreplied := true }}
{{ end }}

{{ if and $help ( not $alreadyreplied) }}
	{{ $embed := sdict }}
	{{ range $k, $v := $template }}
		{{ $embed.Set $k $v}}
	{{ end }}
	{{ $embed.Set "title" "Don't just say `i need help`, tell us what you need help with!" }}
	{{ $embed.Set "description" "[This saves all of us time and we can jump in to provide you with a solution!](https://dontasktoask.com/)" }}
	{{ $msgID := sendMessageNoEscapeRetID nil (complexMessage "reply" $replytarget "embed" $embed) }}
	{{ addMessageReactions nil $msgID (cslice $bin $bookmark $mail) }}
{{ end }}
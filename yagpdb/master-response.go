{{/* Criteria to match */}}
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

{{/* Defining varibles for use throughout the script */}}
{{ $template := sdict "color" 2003199 }}
{{ $replytarget := .ExecData.triggerMsgID }}
{{ $alreadyreplied := false }}
{{ $bin := ":bin:1251255316121653343" }}
{{ $bookmark := ":bookmark:1251243802207846566" }}
{{ $mail := ":mail:1251255870701047909" }}

{{/* https://docs.yagpdb.xyz/reference/templates#templates.sdict */}}
{{/* https://discord.com/channels/166207328570441728/578976698931085333/1248038506312237127 */}}
{{/* https://discord.com/channels/166207328570441728/578976698931085333/1254524326632362036 */}}

{{ if eq .ExecData.trigger 0 }}
    {{ return }}
{{ end }}

{{/* Check if trigger was "banned" or "wrong server" response */}}
{{ if eq .ExecData.trigger 1 2 }}
	{{ $embed := sdict }}
	{{ range $k, $v := $template }}
		{{ $embed.Set $k $v}}
	{{ end }}
	{/* Add "banned" specific title and description */}}
	{{ if eq .ExecData.trigger 1 }}
		{{ $embed.Set "title" "How do I get unbanned?" }}
		{{ $embed.Set "description" "You cannot use ModMail to contact a server you are banned from." }}
	{/* Add "wrong server" specific title and description */}}
	{{ else if eq .ExecData.trigger 2 }}
		{{ $embed.Set "title" "How do I contact *XYZ* Server" }}
		{{ $embed.Set "description" "Please DM the __**bot**__ <@575252669443211264> with your message instead. Make sure __**not**__ to select __ModMail Support__ as that will send it to this server" }}
	{{ end }}
	{{ $embed.Set "fields" (cslice (sdict 
		"name" "You are in the wrong server for what you’re seeking help for."
		"value" "We are the Support server for the ModMail __**bot**__."
		"inline" true
		) (sdict 
		"name" "We have no affiliation with the server/community/game you are seeking help for."
		"value" "We cannot help you any further, sorry."
		"inline" true
		)
	)}}
	{{ $msgID := sendMessageNoEscapeRetID nil (complexMessage "reply" $replytarget "embed" $embed) }}
	{{ addMessageReactions nil $msgID (cslice $bin $bookmark) }}
	{{ $alreadyreplied := true }}
{{ end }}

{{/* Check if trigger was "setup" response */}}
{{ if eq .ExecData.trigger 3 }}
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

{{/* Check if trigger was "open ticket" response */}}
{{ if eq .ExecData.trigger 4 }}
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
	{{ $alreadyreplied := true }}
{{ end }}

{{/* Check if trigger was "premium" response */}}
{{ if eq .ExecData.trigger 5 }}
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

{{/* Check if trigger was "logging" response */}}
{{ if eq .ExecData.trigger 11 }}
	{{ $embed := sdict }}
	{{ range $k, $v := $template }}
		{{ $embed.Set $k $v}}
	{{ end }}
	{{ $embed.Set "title" "Advanced Logging Example" }}
	{{ $embed.Set "description" "This is an example of what you'll get with advanced logging. https://modmail.xyz/logs/d7586c153425000-10d1416086c01033-10d141608b802047" }}
	{{ $file := "[2024-02-26 23:44:18] scyye#0 (User): Hello! I'm a user in need of assistance, can someone help me?\n[2024-02-26 23:44:43] scyye#0 (Comment): I am a staff member writing a comment on the ticket, for other staff to see\n[2024-02-26 23:45:16] scyye#0 (Staff): I am now replying to the user, asking them what they need help with.\n[2024-02-26 23:45:46] jrwallor#0 (Staff): Another staff member with an anonymous reply.\n[2024-02-26 23:45:58] scyye#0 (User): This is the user replying, thanking me for support (I didn't think this through, cut me some slack)\n[2024-02-26 23:46:36] scyye#0 (Comment): =c This ticket is resolved, so I'm closing it now." }}
	{{ $msgID := sendMessageNoEscapeRetID nil (complexMessage "reply" $replytarget "embed" $embed "file" $file) }}
	{{ addMessageReactions nil $msgID (cslice $bin $bookmark) }}
	{{ $alreadyreplied := true }}
{{ end }}

{{/* Check if trigger was "not responding" response */}}
{{ if eq .ExecData.trigger 6 }}
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

{{/* Check if trigger was "custom instance" response */}}
{{ if eq .ExecData.trigger 7 }}
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

{{/* Check if trigger was "self host" response */}}
{{ if eq .ExecData.trigger 8 }}
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

{{/* Check if trigger was "clyde" response */}}
{{ if eq .ExecData.trigger 9 }}
	{{ $embed := sdict }}
	{{ range $k, $v := $template }}
		{{ $embed.Set $k $v}}
	{{ end }}
	{{ $embed.Set "title" "My message wasn't delivered!" }}
	{{ $embed.Set "description" "If you receive \"your message could not be delivered\", check your privacy settings for the server you want to contact. You need to enable the \"allow direct messages from server members\" option." }}
	{{ $embed.Set "image" (sdict "url" "https://i.imgur.com/x5Hcio5.png") }}
	{{ $msgID := sendMessageNoEscapeRetID nil (complexMessage "reply" $replytarget "embed" $embed) }}
	{{ addMessageReactions nil $msgID (cslice $bin $bookmark) }}
	{{ $alreadyreplied := true }}
	{{ end }}

{{/* Check if trigger was "globalticket" response */}}
{{ if eq .ExecData.trigger 10 }}
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

{{/* Check if trigger was "help" response */}}
{{ if and ( eq .ExecData.trigger 12 ) ( not $alreadyreplied) }}
	{{ $embed := sdict }}
	{{ range $k, $v := $template }}
		{{ $embed.Set $k $v}}
	{{ end }}
	{{ $embed.Set "title" "Don't just say `i need help`, tell us what you need help with!" }}
	{{ $embed.Set "description" "[This saves all of us time and we can jump in to provide you with a solution!](https://dontasktoask.com/)" }}
	{{ $msgID := sendMessageNoEscapeRetID nil (complexMessage "reply" $replytarget "embed" $embed) }}
	{{ addMessageReactions nil $msgID (cslice $bin $bookmark $mail) }}
{{ end }}


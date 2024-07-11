{{/* https://docs.yagpdb.xyz/reference/templates#templates.sdict */}}
{{/* https://discord.com/channels/166207328570441728/578976698931085333/1248038506312237127 */}}
{{/* https://discord.com/channels/166207328570441728/578976698931085333/1254524326632362036 */}}

{{/* Criteria to match */}}
{{/* 01 : General Help response */}}
{{/* 02 : Banned response*/}}
{{/* 03 : Wrong Server response */}}
{{/* 04 : Setup response */}}
{{/* 05 : Open Ticket response */}}
{{/* 06 : Premium response */}}
{{/* 07 : Not Responding response */}}
{{/* 08 : Custom Instance response */}}
{{/* 09 : Sef Host response */}}
{{/* 10 : Clyde response */}}
{{/* 11 : Global Ticket response */}}
{{/* 12 : Logging response */}}
{{/* 13 : Verfication response */}}
{{/* 14 : Permission response */}}


{{/* Defining varibles for use throughout the script */}}
{{ $template := sdict "color" 2003199 }}
{{ $replytarget := .ExecData.message }}
{{ $alreadyreplied := false }}
{{ $bin := ":bin:1251255316121653343" }}
{{ $bookmark := ":bookmark:1251243802207846566" }}
{{ $mail := ":mail:1251255870701047909" }}
{{ $discordlogo := "579210587557462021" }}
{{ $deletebutton := cbutton "label" "Delete Response" "custom_id" "support-response-delete" "style" 4 "disabled" true "emoji" (sdict "id" "1251255316121653343") }}
{{ $bookmarkbutton := cbutton "label" "Bookmark Response" "custom_id" "support-response-bookmark" "style" 2 "disabled" true "emoji" (sdict "id" "1251243802207846566") }}
{{ $corebuttons := cslice $deletebutton $bookmarkbutton }}
{{ $extrabuttons := cslice }}

{{ if eq .ExecData.trigger 0 }}
    {{ return }}
{{ end }}

{{/* Check if trigger was "banned" or "wrong server" response */}}
{{ if eq .ExecData.trigger 2 3 }}
	{{ $embed := sdict }}
	{{ range $k, $v := $template }}
		{{ $embed.Set $k $v}}
	{{ end }}
	{{/* Add "banned" specific title and description */}}
	{{ if eq .ExecData.trigger 2 }}
		{{ $embed.Set "title" "How do I get unbanned?" }}
		{{ $embed.Set "description" "You cannot use ModMail to contact a server you are banned from." }}
	{{/* Add "wrong server" specific title and description */}}
	{{ else if eq .ExecData.trigger 3 }}
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
	{{ $msgID := sendMessageNoEscapeRetID nil (complexMessage "reply" $replytarget "embed" $embed "buttons" $corebuttons) }}
	{{ addMessageReactions nil $msgID (cslice $bin $bookmark) }}
	{{ $alreadyreplied := true }}
{{ end }}

{{/* Check if trigger was "setup" response */}}
{{ if eq .ExecData.trigger 4 }}
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
	{{ $corebuttons := $corebuttons.Append (cbutton "label" "Toggle Extra Information" "custom_id" "support-response-toggle" "style" 1 "disabled" true "emoji" (sdict "id" "1258858981372330165")) }}
	{{ $extrabuttons = $extrabuttons.AppendSlice (cslice (cbutton "label" "Invite ModMail" "custom_id" "support-response-invite" "url" "https://modmail.xyz/invite" "style" "link" "emoji" (sdict "id" "1251255870701047909")) (cbutton "label" "ModMail Commands (=help)" "custom_id" "support-response-commands" "url" "https://modmail.xyz/commands" "style" "link" "emoji" (sdict "id" "1258858466081116293"))) }}
	{{ $msgID := sendMessageNoEscapeRetID nil (complexMessage "reply" $replytarget "embed" $embed "buttons" $corebuttons "buttons" $extrabuttons) }}
	{{ addMessageReactions nil $msgID (cslice $bin $bookmark $mail) }}
	{{ $alreadyreplied := true }}
{{ end }}

{{/* Check if trigger was "open ticket" response */}}
{{ if eq .ExecData.trigger 5 }}
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
		{{ $corebuttons = $corebuttons.Append (cbutton "label" "Toggle Extra Information" "custom_id" "support-response-toggle" "style" 1 "disabled" true "emoji" (sdict "id" "1258858981372330165")) }}
	{{ $msgID := sendMessageNoEscapeRetID nil (complexMessage "reply" $replytarget "embed" $embed "buttons" $corebuttons) }}
	{{ addMessageReactions nil $msgID (cslice $bin $bookmark $mail) }}
	{{ $alreadyreplied := true }}
{{ end }}

{{/* Check if trigger was "premium" response */}}
{{ if eq .ExecData.trigger 6 }}
	{{ $embed := sdict }}
	{{ range $k, $v := $template }}
		{{ $embed.Set $k $v}}
	{{ end }}
	{{ $embed.Set "title" "Donation Link" }}
	{{ $embed.Set "description" "[Purchase ModMail Premium Here](https://modmail.xyz/premium)" }}
	{{ $corebuttons = $corebuttons.Append (cbutton "label" "Buy Premium" "custom_id" "support-response-premium" "url" "https://modmail.xyz/premium" "style" "link" "emoji" (sdict "id" "1251273319110414429")) }}
	{{ $msgID := sendMessageNoEscapeRetID nil (complexMessage "reply" $replytarget "embed" $embed "buttons" $corebuttons) }}
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
	{{ $corebuttons = $corebuttons.Append (cbutton "label" "Example Logs" "custom_id" "support-response-logs" "url" "https://modmail.xyz/logs/d7586c153425000-10d1416086c01033-10d141608b802047" "style" "link" "emoji" (sdict "id" "1254891424965722122")) }}
	{{ $msgID := sendMessageNoEscapeRetID nil (complexMessage "reply" $replytarget "embed" $embed "file" $file) }}
	{{ addMessageReactions nil $msgID (cslice $bin $bookmark) }}
	{{ $alreadyreplied := true }}
{{ end }}

{{/* Check if trigger was "not responding" response */}}
{{ if eq .ExecData.trigger 7 }}
	{{ $embed := sdict }}
	{{ range $k, $v := $template }}
		{{ $embed.Set $k $v}}
	{{ end }}
	{{ $embed.Set "title" "ModMail is not responding" }}
	{{ $embed.Set "description" "If ModMail is not responding in your server, please check the following:\n- The bot has Read Messages, Send Messages, and Embed Links permissions.\n- You are using the correct prefix. Use `@ModMail prefix` to check the prefix.\n- The command you are using is valid. Check using `=help <command>`.\n- The bot is online. Discord might be having issues, or the bot might be restarting.\n\nIf the bot still does respond, please let us know your [server ID](https://support.discord.com/hc/en-us/articles/206346498-Where-can-I-find-my-User-Server-Message-ID-)." }}
	{{ $corebuttons = $corebuttons.Append (cbutton "label" "Discord ID Guide" "custom_id" "support-response-idguide" "url" "https://support.discord.com/hc/en-us/articles/206346498-Where-can-I-find-my-User-Server-Message-ID-" "style" "link" "emoji" (sdict "id" $discordlogo)) }}
	{{ $msgID := sendMessageNoEscapeRetID nil (complexMessage "reply" $replytarget "embed" $embed "buttons" $corebuttons) }}
	{{ addMessageReactions nil $msgID (cslice $bin $bookmark) }}
	{{ $alreadyreplied := true }}
{{ end }}

{{/* Check if trigger was "custom instance" response */}}
{{ if eq .ExecData.trigger 8 }}
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
		"value" "| <@381998065327931392> (`James [a_leon]`)\n|or\n| <@365262543872327681> (`snowyjaguar`)"
		"inline" true
	))}}
	{{ $msgID := sendMessageNoEscapeRetID nil (complexMessage "reply" $replytarget "embed" $embed "buttons" $corebuttons) }}
	{{ addMessageReactions nil $msgID (cslice $bin $bookmark) }}
	{{ $alreadyreplied := true }}
{{ end }}

{{/* Check if trigger was "self host" response */}}
{{ if eq .ExecData.trigger 9 }}
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
	{{ $extrabuttons = $extrabuttons.AppendSlice (cslice (cbutton "label" "ModMail GitHub" "custom_id" "support-response-github" "url" "https://github.com/chamburr/modmail" "style" "link" "emoji" (sdict "id" "579211233840857109")) (cbutton "label" "Official V3 Guide" "custom_id" "support-response-officialv3" "url" "https://github.com/chamburr/modmail#self-hosting" "style" "link" "emoji" (sdict "id" "...")) (cbutton "label" "Official V2 Guide" "custom_id" "support-response-officialv3" "url" "https://github.com/chamburr/modmail/blob/v2.1.2/README.md#self-hosting" "style" "link" "emoji" (sdict "id" "..."))) }}
	{{ $msgID := sendMessageNoEscapeRetID nil (complexMessage "reply" $replytarget "embed" $embed "buttons" $corebuttons "buttons" $extrabuttons) }}
	{{ addMessageReactions nil $msgID (cslice $bin $bookmark $mail) }}
	{{ $alreadyreplied := true }}
{{ end }}

{{/* Check if trigger was "clyde" response */}}
{{ if eq .ExecData.trigger 10 }}
	{{ $embed := sdict }}
	{{ range $k, $v := $template }}
		{{ $embed.Set $k $v}}
	{{ end }}
	{{ $embed.Set "title" "My message wasn't delivered!" }}
	{{ $embed.Set "description" "If you receive \"your message could not be delivered\", check your privacy settings for the server you want to contact. You need to enable the \"allow direct messages from server members\" option." }}
	{{ $embed.Set "image" (sdict "url" "https://i.imgur.com/x5Hcio5.png") }}
	{{ $msgID := sendMessageNoEscapeRetID nil (complexMessage "reply" $replytarget "embed" $embed "buttons" $corebuttons) }}
	{{ addMessageReactions nil $msgID (cslice $bin $bookmark) }}
	{{ $alreadyreplied := true }}
	{{ end }}

{{/* Check if trigger was "globalticket" response */}}
{{ if eq .ExecData.trigger 11 }}
	{{ $embed := sdict }}
	{{ range $k, $v := $template }}
		{{ $embed.Set $k $v}}
	{{ end }}
	{{ $embed.Set "title" "Everyone can see my tickets!" }}
	{{ $embed.Set "description" "If everyone in your server is able to view tickets, there is a chance that another bot in your server is interfering. This is not a problem with ModMail. We recommend checking the audit logs in your server settings to see which bot is changing the channel permissions." }}
	{{ $msgID := sendMessageNoEscapeRetID nil (complexMessage "reply" $replytarget "embed" $embed "buttons" $corebuttons) }}
	{{ addMessageReactions nil $msgID (cslice $bin $bookmark) }}
	{{ $alreadyreplied := true }}
{{ end }}

{{/* Check if trigger was "verification" response */}}
{{ if eq .ExecData.trigger 13 }}
	{{ $embed := sdict }}
	{{ range $k, $v := $template }}
		{{ $embed.Set $k $v}}
	{{ end }}
	{{ $embed.Set "title" "ModMail Verfication" }}
	{{ $embed.Set "description" "If you are having trouble verifying, please check the following:" }}
	{{ $embed.Set "image" (sdict "url" "https://media.discordapp.net/attachments/576764484065165313/985238549886562374/unknown.png") }}
	{{ $embed.Set "fields" (cslice (sdict
		"name" "Confirm Correct Discord Account in Browser"
		"value" "Go to [Discord Login](https://discord.com/login) and check check the account shown.\nIf it's incorrect you might have an alt or unclaimed account on your browser.\nTry the following:"
		"inline" false
	) (sdict
		"name" "Open Verfication Link in Incognito"
		"value" "Yes, this can be done on mobile!\nThis ignores saved logins and prompts a fresh login screen."
		"inline" false
	)) }}
	{{ $extrabuttons = $extrabuttons.AppendSlice (cslice (cbutton "label" "Discord Login" "custom_id" "support-response-discordlogin" "url" "https://discord.com/login" "style" "link" "emoji" (sdict "id" $discordlogo)) (cbutton "label" "Incognito Browser Guide" "custom_id" "support-response-incognito" "url" "https://incognitobrowser.io/step-by-step-guide-to-using-incognito-mode-on-chrome-firefox-and-safari/" "style" "link")) }}
	{{ $msgID := sendMessageNoEscapeRetID nil (complexMessage "reply" $replytarget "embed" $embed "buttons" $corebuttons "buttons" $extrabuttons) }}
	{{ addMessageReactions nil $msgID (cslice $bin $bookmark) }}
	{{ $alreadyreplied := true }}
{{ end }}

{{/* Check if trigger was "permission" response */}}
{{ if eq .ExecData.trigger 14 }}
	{{ $embed := sdict }}
	{{ range $k, $v := $template }}
		{{ $embed.Set $k $v}}
	{{ end }}
	{{ $embed.Set "title" "ModMail Permissions" }}
	{{ $embed.Set "description" "If you are having trouble with ModMail, please check the following:" }}
	{{ $embed.Set "image" (sdict "url" "https://media.discordapp.net/attachments/576764854673735680/863863546915979274/unknown-4.png") }}
	{{ $embed.Set "fields" (cslice (sdict
		"name" "Note"
		"value" "The adminatorpermission is not required for ModMail to function. It can be useful for troubleshooting or preventing conflicts with other bots."
		"inline" false
	)) }}
	{{ $extrabuttons = $extrabuttons.Append (cbutton "label" "Discord Permissions FAQ" "custom_id" "support-response-permissions" "url" "https://support.discord.com/hc/en-us/articles/206029707-Setting-Up-Permissions-FAQ" "style" "link" "emoji" (sdict "id" $discordlogo)) }}



{{/* Check if trigger was "help" response */}}
{{ if and ( eq .ExecData.trigger 1 ) ( not $alreadyreplied) }}
	{{ $embed := sdict }}
	{{ range $k, $v := $template }}
		{{ $embed.Set $k $v}}
	{{ end }}
	{{ $embed.Set "title" "Don't just say `i need help`, tell us what you need help with!" }}
	{{ $embed.Set "description" "[This saves all of us time and we can jump in to provide you with a solution!](https://dontasktoask.com/)" }}
	{{ $msgID := sendMessageNoEscapeRetID nil (complexMessage "reply" $replytarget "embed" $embed "buttons" $corebuttons) }}
	{{ addMessageReactions nil $msgID (cslice $bin $bookmark $mail) }}
{{ end }}
{{/* Defining varibles for use throughout the script */}}
{{ $template := sdict "color" 2003199 }}
{{ $replytarget := .Message.ID }}
{{ $alreadyreplied := false }}
{{ $modmaillogo := ":modmail:702099194701152266" }}
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
	{{ $embed.Set "description" "You are in the wrong server for what you’re seeking help for.\n\nWe are the Support server for the ModMail **bot**.\n\nWe have no affiliation with the server you are banned from.\n\nYou cannot use ModMail to contact a server you are banned from.\n\nWe cannot help you any further, sorry." }}
	{{ sendMessageNoEscape nil (complexMessage "reply" $replytarget "embed" $embed) }}
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
	{{ addMessageReactions nil $msgID (cslice $modmaillogo) }}
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
		)
		)}}
	{{ $msgID := sendMessageNoEscapeRetID nil (complexMessage "reply" $replytarget "embed" $embed) }}
	{{ addMessageReactions nil $msgID (cslice $modmaillogo) }}
	{{ $embed.Set "fields" ($embed.fields.Append (sdict
			"name" "Note"
			"value" "If you are having trouble with the `=send` command, please ensure you are using the correct server ID. You can find this by right-clicking on the server name and selecting `Copy ID`."
			"inline" false
		) ) }}
	{{editMessage nil $msgID (complexMessageEdit "embed" $embed)}}
	{{ $alreadyreplied:=false }}
	{{ end }}


{{ if and $premium ( not (hasPrefix .Message.Content "=")) }}
	{{ $embed := sdict }}
	{{ range $k, $v := $template }}
		{{ $embed.Set $k $v}}
		{{ end }}
	{{ $embed.Set "title" "Donation Link" }}
	{{ $embed.Set "description" "[Purchase ModMail Premium Here](https://modmail.xyz/premium)" }}
	{{ sendMessageNoEscape nil (complexMessage "reply" $replytarget "embed" $embed) }}
	{{ $alreadyreplied := true }}
	{{ end }}

{{ if and $noresponse }}
	{{ $embed := sdict }}
		{{ range $k, $v := $template }}
			{{ $embed.Set $k $v}}
			{{ end }}
	{{ $embed.Set "title" "ModMail is not responding" }}
	{{ $embed.Set "description" "If ModMail is not responding in your server, please check the following:\n- The bot has Read Messages, Send Messages, and Embed Links permissions.\n- You are using the correct prefix. Use `@ModMail prefix` to check the prefix.\n- The command you are using is valid. Check using `=help <command>`.\n- The bot is online. Discord might be having issues, or the bot might be restarting.\n\nIf the bot still does respond, please let us know your [server ID](https://support.discord.com/hc/en-us/articles/206346498-Where-can-I-find-my-User-Server-Message-ID-)." }}
	{{ sendMessageNoEscape nil (complexMessage "reply" $replytarget "embed" $embed) }}
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
	{{ sendMessageNoEscape nil (complexMessage "reply" $replytarget "embed" $embed) }}
	{{ $alreadyreplied := true }}
	{{ end }}

	{{ if $selfhost }}
	{{ $embed := sdict }}
		{{ range $k, $v := $template }}
			{{ $embed.Set $k $v}}
			{{ end }}
	{{ $embed.Set "title" "ModMail Self Hosting" }}
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
	) (sdict 
		"name" "Cheap VPS Providers" 
		"value" "- [Scaleway | EU](https://www.scaleway.com)\n- [Linode | US, EU, Asia](https://www.linode.com)\n- [Digital Ocean | US](https://www.digitalocean.com)\n- [Vultr | US, EU, NA, SA, Asia](https://www.vultr.com)\n- [OVH | EU, Canada](https://www.ovh.co.uk)\n- [Hetzner | Germany, US](https://www.hetzner.com)\n- [Time4VPS | Lithuania](https://www.time4vps.eu)"
		"inline" false
	) (sdict 
		"name" "Free Hosting"
		"value" "No. Not even heroku/repl.it"
		"inline" true
	) (sdict 
		"name" "Kinda Free hosting"
		"value" "[GCP](https://cloud.google.com/gcp), [AWS](https://aws.amazon.com/ec2/) and [Azure](https://azure.microsoft.com/en-gb/free/) have one year free micros. Some are blanket free, some give you a pot of credits to use over the year. Many YouTubers run adverts for [Linode](https://www.linode.com) which give you $100 USD to spend over 90 days."
		"inline" true
	) (sdict 
		"name" "Self Host Locally"
		"value" "Any modern hardware should be sufficient for running a bot. An old computer with a few GB of RAM could be suitable, or a Raspberry Pi (any model, except perhaps one of the particularly less powerful ones).\n**Note:** We don't recommend using Windows because it requires a drastic change to the scripts to support Windows file path structure, Windows can also be finicky about releasing the AVX2 instruction set which is required for ModMail to function."
		"inline" false
	) (sdict 
		"name" "🚩 Repli.it Hosting Red Flags 🚩"
		"value" "While this may seem like a nice and free service, it has a lot more caveats than you might think, such as:\n- The machines are super underpowered.\n - This means your bot will lag a lot as it gets bigger.\n- You need to run a webserver alongside your bot to prevent it from being shut off. (I don't think this is an issue for ModMail).\n- Repl.it uses an ephemeral file system.\n - This means any file you saved via your bot will be overwritten when you next launch.\n- They use a shared IP for everything running on the service.\n - This one is important, if someone is running a user bot on their service and gets banned, everyone on that IP will be banned. **Including you.**"
		"inline" true
	) (sdict 
		"name" "🚩 Heroku 🚩"
		"value" "- Bots are not what the platform is designed for.\n - Heroku is designed to provide web servers (like Django, Flask, etc). This is why they give you a domain name and open a port on their local emulator.\n- Heroku's environment is heavily containerized, making it significantly underpowered for a standard use case.\n- Heroku's environment is volatile.\n - In order to handle the insane amount of users trying to use it for their own applications, Heroku will dispose of your environment every time your application dies unless you pay.\n- Heroku has minimal system dependency control.\n - This is the reason why voice doesn't work natively on Heroku.\n- Heroku only offers a limited amount of time on their free program for your applications. If you exceed this limit, which you probably will, they'll shut down your application until your free credit resets."
		"inline" true
	))}}
	{{ sendMessageNoEscape nil (complexMessage "reply" $replytarget "embed" $embed) }}
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
	{{ sendMessageNoEscape nil (complexMessage "reply" $replytarget "embed" $embed) }}
	{{ $alreadyreplied := true }}
	{{ end }}

{{ if $globalticket }}
	{{ $embed := sdict }}
		{{ range $k, $v := $template }}
			{{ $embed.Set $k $v}}
			{{ end }}
	{{ $embed.Set "title" "Everyone can see my tickets!" }}
	{{ $embed.Set "description" "If everyone in your server is able to view tickets, there is a chance that another bot in your server is interfering. This is not a problem with ModMail. We recommend checking the audit logs in your server settings to see which bot is changing the channel permissions." }}
	{{ sendMessageNoEscape nil (complexMessage "reply" $replytarget "embed" $embed) }}
	{{ $alreadyreplied := true }}
	{{ end }}

{{ if and $help ( not $alreadyreplied) }}
	{{ $embed := sdict }}
		{{ range $k, $v := $template }}
			{{ $embed.Set $k $v}}
			{{ end }}
	{{ $embed.Set "title" "Don't just say `i need help`, tell us what you need help with!" }}
	{{ $embed.Set "description" "[This saves all of us time and we can jump in to provide you with a solution!](https://dontasktoask.com/)" }}
	{{ sendMessageNoEscape nil (complexMessage "reply" $replytarget "embed" $embed) }}
	{{ end }}
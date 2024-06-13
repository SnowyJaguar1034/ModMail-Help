{{/* Declares the variables */}}
{{ $modmaillogo := "modmail:702099194701152266" }}
{{ $bin := "bin:1250547674562957313" }}
{{ $redflag := "red_flag:1250907194778583101" }}
{{ $msgID := .ReactionMessage.ID }}

{{/* Declaring the new fields */}}
{{ $setupfields := (cslice (sdict 
	"name" "Advanced Setup" 
	"value" "Some additional commands you could use are:" 
	"inline" false
) (sdict
	"name" "`=pingrole <roles>`"
	"value" "For configuring which roles get pinged upon a ModMail ticket being created."
	"inline" true
) (sdict
	"name" "`=accessrole <roles>`"
	"value" "For configuring which roles can reply to ModMail tickets."
	"inline" true
) (sdict
	"name" "`=commandonly`"
	"value" "For toggling if commands are required to reply to tickets.\nIf **disabled** staff have only to type in the channel for their message to be sent.\nIf **enabled** staff have to reply with `=reply` or `=areply`."
	"inline" false
) (sdict
	"name" "`=anonymous`"
	"value" "For toggling anonymous staff replies to hide the responder's name.\nThis does not work for making your end-user anonymous."
	"inline" true
) (sdict
	"name" "`=logging`"
	"value" "For toggling log messages of tickets being opened or closed.\nThis does not log a transcript of the messages."
	"inline" true
) (sdict
	"name" "You can mention the roles, use role IDs or role names."
	"value" "For role names with a space, it needs to be in quotes (e.g. \"Head Admin\")"
	"inline" false
)) }}

{{ $ticketfields := (cslice (sdict
	"name" "Bonus Information: `=confirmation` command" 
	"value" "If you have enabled the `=confirmation` mode, you will not be given the server selection menu immediately and will instead be prompted to resume messaging the last server you contacted.\nThis prompt will contain an option to take you to the server selection menu if it's incorrect but you can also use `=new <message>` to force the server selection menu to appear."
	"inline" false
)) }}

{{ $selfhosting := (cslice (sdict
	"name" "Cheap EU VPS's"
	"value" "- [Linode](https://www.linode.com)\n- [Vultr](https://www.vultr.com)\n- [Scaleway](https://www.scaleway.com)\n- [OVH](https://www.ovh.co.uk)\n- [Time4VPS](https://www.time4vps.eu)"
	"inline" true
) (sdict
	"name" "Cheap US VPS's"
	"value" "- [Linode](https://www.linode.com)\n- [Vultr](https://www.vultr.com)\n- [Digital Ocean](https://www.digitalocean.com)\n- [Hetzner](https://www.hetzner.com)"
	"inline" true
) (sdict
	"name" "Cheap Asia VPS's"
	"value" "- [Linode](https://www.linode.com)\n- [Vultr](https://www.vultr.com)"
	"inline" true
) (sdict
	"name" "Cheap Africa VPS's"
	"value" "- [Vultr](https://www.vultr.com)"
	"inline" true
) (sdict
	"name" "Cheap Canada VPS's"
	"value" "- [OVH](https://www.ovh.co.uk)"
	"inline" true
) (sdict
	"name" "Cheap Germany VPS's"
	"value" "- [Hetzner](https://www.hetzner.com)"
	"inline" true
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
) )}}

{{ $redflags := (cslice (sdict 
	"name" "🚩 Repli.it Hosting Red Flags 🚩"
	"value" "While this may seem like a nice and free service, it has a lot more caveats than you might think, such as:\n- The machines are super underpowered.\n - This means your bot will lag a lot as it gets bigger.\n- You need to run a webserver alongside your bot to prevent it from being shut off. (I don't think this is an issue for ModMail).\n- Repl.it uses an ephemeral file system.\n - This means any file you saved via your bot will be overwritten when you next launch.\n- They use a shared IP for everything running on the service.\n - This one is important, if someone is running a user bot on their service and gets banned, everyone on that IP will be banned. **Including you.**"
	"inline" true
) (sdict 
	"name" "🚩 Heroku 🚩"
	"value" "- Bots are not what the platform is designed for.\n - Heroku is designed to provide web servers (like Django, Flask, etc). This is why they give you a domain name and open a port on their local emulator.\n- Heroku's environment is heavily containerized, making it significantly underpowered for a standard use case.\n- Heroku's environment is volatile.\n - In order to handle the insane amount of users trying to use it for their own applications, Heroku will dispose of your environment every time your application dies unless you pay.\n- Heroku has minimal system dependency control.\n - This is the reason why voice doesn't work natively on Heroku.\n- Heroku only offers a limited amount of time on their free program for your applications. If you exceed this limit, which you probably will, they'll shut down your application until your free credit resets."
	"inline" true
) )}}

{{/* {{ $titles : cslice (sdict "title" "How do I setup ModMail?" "fields" $setupfields) (sdict "title" "How do I open a ticket?" "fields" $ticketfields) (sdict "title" "How do I self-host ModMail?" "fields" $selfhosting) }} */}}

{{ $titles := sdict "How do I setup ModMail?" $setupfields "How do I open a ticket?" $ticketfields "How do I self-host ModMail?" $selfhosting }}

{{/* Checks if the reaction is the bin emoji */}}
{{ if eq .Reaction.Emoji.APIName $bin }}
	{{ deleteMessage nil .ReactionMessage.ID 0 }}
{{ end }}

{{/* Checks if the reaction is the ModMail logo */}}
{{ if eq .Reaction.Emoji.APIName $modmaillogo }}
	{{ range .ReactionMessage.Embeds }}
		{{ $currentfieldnames := cslice }}
		{{ range .Fields }}
			{{ $currentfieldnames = $currentfieldnames.Append .Name }}
		{{ end }}
		{{ $embed := structToSdict . }}
		{{ $embed.Set "Fields" (cslice.AppendSlice $embed.Fields) }}
		{{ range $titles.Get $embed.Title }}
			{{ if not (in $currentfieldnames .name) }}
				{{ $embed.Set "Fields" ($embed.Fields.Append .) }}
			{{ end }}
		{{ end }}
		{{ addMessageReactions nil $msgID (cslice $redflag) }}
		{{editMessage nil $msgID (complexMessageEdit "embed" $embed)}}
	{{ end}}
{{ end }}

{{ if eq .Reaction.Emoji.APIName $redflag }}
	{{ range .ReactionMessage.Embeds }}
		{{ $currentfieldnames := cslice }}
		{{ range .Fields }}
			{{ $currentfieldnames = $currentfieldnames.Append .Name }}
		{{ end }}
		{{ $embed := structToSdict . }}
		{{ $embed.Set "Fields" (cslice.AppendSlice $embed.Fields) }}
		{{ if eq $embed.Title "How do I self-host ModMail?" }}
			{{ range $redflags }}
				{{ if not (in $currentfieldnames .name) }}
					{{ $embed.Set "Fields" ($embed.Fields.Append .)}}
				{{ end }}
			{{ end }}
		{{ end }}
		{{editMessage nil $msgID (complexMessageEdit "embed" $embed)}}
	{{ end}}
{{ end }}

		{{ if eq $embed.Title "How do I setup ModMail?" }}
			{{ range $setupfields }}
				{{ if not (in $currentfieldnames .name) }}
					{{ $embed.Set "Fields" ($embed.Fields.Append .)}}
				{{ end }}
			{{ end }}
		{{ end }}
		{{ if eq $embed.Title "How do I open a ticket?" }}
			{{ range $ticketfields }}
				{{ if not (in $currentfieldnames .name) }}
					{{ $embed.Set "Fields" ($embed.Fields.Append .)}}
				{{ end }}
			{{ end }}
		{{ end }}
		{{ if eq $embed.Title "How do I self-host ModMail?" }}
			{{ range $selfhosting }}
				{{ if not (in $currentfieldnames .name) }}
					{{ $embed.Set "Fields" ($embed.Fields.Append .)}}
				{{ end }}
			{{ end }}
		{{ end }}
		{{editMessage nil $msgID (complexMessageEdit "embed" $embed)}}
		{{/*{{ sendMessage nil (complexMessage "file" (json $embed) ) }} */}}
	{{ end}}
{{ end }}


(sdict 
	"name" "Cheap VPS Providers" 
	"value" "- [Scaleway | EU](https://www.scaleway.com)\n- [Linode | US, EU, Asia](https://www.linode.com)\n- [Digital Ocean | US](https://www.digitalocean.com)\n- [Vultr | US, EU, NA, SA, Asia](https://www.vultr.com)\n- [OVH | EU, Canada](https://www.ovh.co.uk)\n- [Hetzner | Germany, US](https://www.hetzner.com)\n- [Time4VPS | Lithuania](https://www.time4vps.eu)"
	"inline" false
) 
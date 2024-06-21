{{/* Declares the variables */}}
{{ $msgID := .ReactionMessage.ID }}
{{ $mail := "mail:1251255870701047909" }}
{{ $redflag := "flag:1251303058542039202" }}
{{ $reactionadded := .Reaction.Emoji.APIName }}

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


{{ $titles := sdict 
	"How do I setup ModMail?" $setupfields.name 
	"How do I open a ticket?" $ticketfields.name 
	"How do I self-host ModMail?" $selfhosting.name 
}}

{{ $mtitles := sdict
	"How do I setup ModMail?" (cslice "Advanced Setup" "`=pingrole <roles>`" "`=accessrole <roles>`" "`=commandonly`" "`=anonymous`" "`=logging`" "You can mention the roles, use role IDs or role names.")
	"How do I open a ticket?" (cslice "Bonus Information: `=confirmation` command")
	"How do I self-host ModMail?" (cslice "Cheap EU VPS's" "Cheap US VPS's" "Cheap Asia VPS's" "Cheap Africa VPS's" "Cheap Canada VPS's" "Cheap Germany VPS's" "Free Hosting" "Kinda Free hosting" "Self Host Locally")
}}

{{ fields := sdict
	"Advanced Setup" (cslice "Some additional commands you could use are:" false)
	"`=pingrole <roles>`" (cslice "For configuring which roles get pinged upon a ModMail ticket being created." true)
	"`=accessrole <roles>`" (cslice "For configuring which roles can reply to ModMail tickets." true)
	"`=commandonly`" (cslice "For toggling if commands are required to reply to tickets.\nIf **disabled** staff have only to type in the channel for their message to be sent.\nIf **enabled** staff have to reply with `=reply` or `=areply`." false)
	"`=anonymous`" (cslice "For toggling anonymous staff replies to hide the responder's name.\nThis does not work for making your end-user anonymous." true)
	"`=logging`" (cslice "For toggling log messages of tickets being opened or closed.\nThis does not log a transcript of the messages." true)
	"You can mention the roles, use role IDs or role names." (cslice "For role names with a space, it needs to be in quotes (e.g. \"Head Admin\")" false)
	"Bonus Information: `=confirmation` command" (cslice "If you have enabled the `=confirmation` mode, you will not be given the server selection menu immediately and will instead be prompted to resume messaging the last server you contacted.\nThis prompt will contain an option to take you to the server selection menu if it's incorrect but you can also use `=new <message>` to force the server selection menu to appear." false)
	"Cheap EU VPS's" (cslice "- [Linode](https://www.linode.com)\n- [Vultr](https://www.vultr.com)\n- [Scaleway](https://www.scaleway.com)\n- [OVH](https://www.ovh.co.uk)\n- [Time4VPS](https://www.time4vps.eu)" true)
	"Cheap US VPS's" (cslice "- [Linode](https://www.linode.com)\n- [Vultr](https://www.vultr.com)\n- [Digital Ocean](https://www.digitalocean.com)\n- [Hetzner](https://www.hetzner.com)" true)
	"Cheap Asia VPS's" (cslice "- [Linode](https://www.linode.com)\n- [Vultr](https://www.vultr.com)" true)
	"Cheap Africa VPS's" (cslice "- [Vultr](https://www.vultr.com)" true)
	"Cheap Canada VPS's" (cslice "- [OVH](https://www.ovh.co.uk)" true)
	"Cheap Germany VPS's" (cslice "- [Hetzner](https://www.hetzner.com)" true)
	"Free Hosting" (cslice "No. Not even heroku/repl.it" true)
	"Kinda Free hosting" (cslice "[GCP](https://cloud.google.com/gcp), [AWS](https://aws.amazon.com/ec2/) and [Azure](https://azure.microsoft.com/en-gb/free/) have one year free micros. Some are blanket free, some give you a pot of credits to use over the year. Many YouTubers run adverts for [Linode](https://www.linode.com) which give you $100 USD to spend over 90 days." true)
	"Self Host Locally" (cslice "Any modern hardware should be sufficient for running a bot. An old computer with a few GB of RAM could be suitable, or a Raspberry Pi (any model, except perhaps one of the particularly less powerful ones).\n**Note:** We don't recommend using Windows because it requires a drastic change to the scripts to support Windows file path structure, Windows can also be finicky about releasing the AVX2 instruction set which is required for ModMail to function." false)
}}

{{/* Checks if the reaction is the ModMail logo */}}
{{ if eq $reactionadded $mail }}
    {{ range .ReactionMessage.Embeds }}
        {{ $newfields := cslice }}
        {{ $embed := structToSdict . }}
        {{ range .Fields }}
			{{ json $titles }}
			{{ range $title := $titles }}
				{{ if (in ($title.Get $embed.Title) .Name) }}
					{{ continue }}
				{{ end }}
            {{ end }}
            {{ $newfields = $newfields.Append . }}
        {{ end }}
        {{ $embed.Set "Fields" $newfields }}
		{{editMessage nil $msgID (complexMessageEdit "embed" $embed)}}
		{{ json $embed true }}
	{{ end}}
{{ end }}

{{ range $title := $itles }}
	{{ if (in ($title.Get $embed.Title) .Name) }}

{{/* Declares the variables */}}
{{ $msgID := .ReactionMessage.ID }}
{{ $mail := "mail:1251255870701047909" }}
{{ $redflag := "flag:1251303058542039202" }}
{{ $reactionadded := .Reaction.Emoji.APIName }}

{{/* Declaring the new fields */}}
{{ $titles := sdict
	"How do I setup ModMail?" (cslice (sdict 
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
	))
	"How do I open a ticket?" (cslice (sdict
		"name" "Bonus Information: `=confirmation` command" 
		"value" "If you have enabled the `=confirmation` mode, you will not be given the server selection menu immediately and will instead be prompted to resume messaging the last server you contacted.\nThis prompt will contain an option to take you to the server selection menu if it's incorrect but you can also use `=new <message>` to force the server selection menu to appear."
		"inline" false
	))
	"How do I self-host ModMail?" (cslice (sdict
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
	))
}}

{{/* Checks if the reaction is the ModMail logo */}}
{{ if eq $reactionadded $mail }}
	{{/* Iterates over the titles and removes the fields from the embed */}}
	{{ range $oldEmbed := .ReactionMessage.Embeds }}
		{{ $newEmbed := structToSdict . }}
		{{/* Creates a new slice to store the new fields */}}
		{{ $newFields := cslice }}
		{{/* Gets the slice of dicts that match the title of the embed */}}
		{{ $removeFields := $titles.Get $newEmbed.Title }}
		{{/* Creates a new slice to store the field names to be removed */}}
		{{ $removeFieldsNames := cslice }}
		{{/* Iterates over the $removeFields slice and appends them to the $removeFieldsNames slice to create a list of field names to be removed */}}
		{{ range $fielddict := $removeFields }}
			{{ $removeFieldsNames = $removeFieldsNames.Append $fielddict.name }}
		{{ end }}
		{{/* Iterates over the fields of the old embed and checks if the field name matches the field names to be removed */}}
		{{ range $field := $oldEmbed.Fields }}
			{{ $match := false }}
			{{/* Iterates over the field names to be removed and checks if the field name matches */}}
			{{ range $removeField := $removeFieldsNames }}
				{{/* If the field name matches, breaks the loop */}}
				{{ if $match = eq $removeField $field.Name }}
					{{ break }}
				{{ end }}
			{{ end }}
			{{/* If the field name does not match, appends the field to the new fields slice */}}
			{{ if not $match }}
				{{ $newFields = $newFields.Append $field }}
			{{ end }}
		{{ end }}
		{{/* Sets the new fields to the new embed and edits the message with the new embed */}}
		{{ $newEmbed.Set "Fields" $newFields }}
		{{editMessage nil $msgID (complexMessageEdit "embed" $newEmbed)}}
	{{ end}}
{{ end }}
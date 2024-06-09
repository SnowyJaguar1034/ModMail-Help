{{ $modmaillogo := "modmail:702099194701152266" }}
{{ $msgID := .ReactionMessage.ID }}

{{ if eq .Reaction.Emoji.APIName $modmaillogo }}
	{{ $embeds := .ReactionMessage.Embeds }}
	{{ range $embed := $embeds }}
		{{ if eq $embed.Title "How do I setup ModMail?" }}
			{{ $embed.Set "fields" ($embed.fields.Append (sdict 
				"name" "Advanced Setup" 
				"value" "Some additional commands you could use are:" 
				"inline" false
			))}}
			{{ $embed.Set "fields" ($embed.fields.Append (sdict
				"name" "`=pingrole <roles>`"
				"value" "For configuring which roles get pinged upon a ModMail ticket being created."
				"inline" true
			))}}
			{{ $embed.Set "fields" ($embed.fields.Append (sdict
				"name" "`=accessrole <roles>`"
				"value" "For configuring which roles can reply to ModMail tickets."
				"inline" true
			))}}
			{{ $embed.Set "fields" ($embed.fields.Append (sdict
				"name" "`=commandonly`"
				"value" "For toggling if commands are required to reply to tickets.\nIf **disabled** staff have only to type in the channel for their message to be sent.\nIf **enabled** staff have to reply with `=reply` or `=areply`."
				"inline" false
			))}}
			{{ $embed.Set "fields" ($embed.fields.Append (sdict
				"name" "`=anonymous`"
				"value" "For toggling anonymous staff replies to hide the responder's name.\nThis does not work for making your end-user anonymous."
				"inline" true
			))}}
			{{ $embed.Set "fields" ($embed.fields.Append (sdict
				"name" "`=logging`"
				"value" "For toggling log messages of tickets being opened or closed.\nThis does not log a transcript of the messages."
				"inline" true
			))}}
			{{ $embed.Set "fields" ($embed.fields.Append (sdict
				"name" "You can mention the roles, use role IDs or role names."
				"value" "For role names with a space, it needs to be in quotes (e.g. \"Head Admin\")"
				"inline" false
			))}}
			{{editMessage nil $msgID (complexMessageEdit "embed" $embed)}}
		{{ if eq $embed.Title "How do I open a ticket?" }}
			{{ $embed.Set "fields" ($embed.fields.Append (sdict
				"name" "Bonus Information: `=confirmation` command" 
				"value" "If you have enabled the `=confirmation` mode, you will not be given the server selection menu immediately and will instead be prompted to resume messaging the last server you contacted.\nThis prompt will contain an option to take you to the server selection menu if it's incorrect but you can also use `=new <message>` to force the server selection menu to appear."
				"inline" false
				))}}
			{{editMessage nil $msgID (complexMessageEdit "embed" $embed)}}


{{ end }}
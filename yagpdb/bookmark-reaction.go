{{/* DM the user a copy of the message they reacted to */}}
{{/* Declaring variables */}}
{{ $msgID := .ReactionMessage.ID }}
{{ $bookmark := "bookmark:1251243802207846566" }}

{{/* Checks if the reaction is the bin emoji */}}
{{ if and (eq .Reaction.Emoji.APIName $bookmark) (eq .ReactionMessage.Author.ID .BotUser.ID) }}
	{{ print "Bookmark reaction detected" }}
	{{ print "Attachments: " .ReactionMessage.Attachments }}
	{{ print "Embeds: " .ReactionMessage.Embeds }}
	{{ print "Content: " .ReactionMessage.Content }}
	{{ print "Author: " .ReactionMessage.Author.Mention }}
	{{ print "Link: " (joinStr "" "[Jump to message](https://discord.com/channels/" .Guild.ID "/" .Channel.ID "/" $msgID ")" ) }}
	{{/* DM the user a copy of the message they reacted to */}}
	{{ $msgembed := sdict }}
	{{ $attachments := cslice }}
	{{ $embeds := cslice }}
	{{ if .ReactionMessage.Content }}
		{{ $msgembed.Set "title" "Message Bookmarked" }}
		{{ $msgembed.Set "description" "Here is a copy of the message you bookmarked:" }}
		{{ $msgembed.Set "fields" (cslice 
			(sdict "name" "Message" "value" .ReactionMessage.Content "inline" false)
			(sdict "name" "Author" "value" .ReactionMessage.Author.Mention "inline" true)
			(sdict "name" "Link" "value" (joinStr "" "[Jump to message](https://discord.com/channels/" .Guild.ID "/" .Channel.ID "/" $msgID ")" ) "inline" true)
		) }}
		{{ $embeds = $embeds.Append $msgembed }}
	{{ end }}
	{{ if .ReactionMessage.Embeds }}
		{{ $embeds = $embeds.AppendSlice .ReactionMessage.Embeds }}
	{{ end }}
	{{ if .ReactionMessage.Attachments }}
		{{ range .ReactionMessage.Attachments }}
			{{ if . 0 }}
				{{ $msgembed.Set "image" (sdict "url" .URL) }}
			{{ else }}
				{{ $embeds = $embeds.Append (sdict "image" (sdict "url" .URL)) }}
			{{ end }}
			{{ $attachments = $attachments.Append (sdict "url" .URL "filename" .Filename) }}
		{{ end }}
	{{ end }}
	{{ sendDM (complexMessage "embed" $embeds ) }}
{{ end }}
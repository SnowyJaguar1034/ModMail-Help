{{/* Deletes a message when a reaction is added to it */}}
{{/* Declaring variables */}}
{{ $msgID := .ReactionMessage.ID }}
{{ $bin := "bin:1251255316121653343" }}

{{/* Checks if the reaction is the bin emoji */}}
{{ if and (eq .Reaction.Emoji.APIName $bin) (eq .ReactionMessage.Author.ID .BotUser.ID) }}
	{{editMessage nil $msgID (complexMessageEdit "content" "**Message will be deleted in 5 seconds**") }}
	{{ deleteMessage nil $msgID 5 }}
{{ end }}
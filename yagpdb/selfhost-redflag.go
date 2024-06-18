{{/* Declares the variables */}}
{{ $msgID := .ReactionMessage.ID }}
{{ $redflag := "flag:1251303058542039202" }}
{{ $reactionadded := .Reaction.Emoji.APIName }}

{{/* Check if the reaction added is the red flag emoji */}}
{{ if ne $reactionadded $redflag }}
    {{ return }}
{{ end }}

{{/* Declaring the new fields */}}
{{ $redflags := (cslice (sdict 
	"name" "🚩 Repli.it Hosting Red Flags 🚩"
	"value" "While this may seem like a nice and free service, it has a lot more caveats than you might think, such as:\n- The machines are super underpowered.\n - This means your bot will lag a lot as it gets bigger.\n- You need to run a webserver alongside your bot to prevent it from being shut off. (I don't think this is an issue for ModMail).\n- Repl.it uses an ephemeral file system.\n - This means any file you saved via your bot will be overwritten when you next launch.\n- They use a shared IP for everything running on the service.\n - This one is important, if someone is running a user bot on their service and gets banned, everyone on that IP will be banned. **Including you.**"
	"inline" true
) (sdict 
	"name" "🚩 Heroku 🚩"
	"value" "- Bots are not what the platform is designed for.\n - Heroku is designed to provide web servers (like Django, Flask, etc). This is why they give you a domain name and open a port on their local emulator.\n- Heroku's environment is heavily containerized, making it significantly underpowered for a standard use case.\n- Heroku's environment is volatile.\n - In order to handle the insane amount of users trying to use it for their own applications, Heroku will dispose of your environment every time your application dies unless you pay.\n- Heroku has minimal system dependency control.\n - This is the reason why voice doesn't work natively on Heroku.\n- Heroku only offers a limited amount of time on their free program for your applications. If you exceed this limit, which you probably will, they'll shut down your application until your free credit resets."
	"inline" true
) )}}


{{/* Iterate over the embeds in the message */}}
{{ range .ReactionMessage.Embeds }}
    {{ $currentfieldnames := cslice }}
    {{ range .Fields }}
        {{ $currentfieldnames = $currentfieldnames.Append .Name }}
    {{ end }}
    {{ $embed := structToSdict . }}
    {{ $embed.Set "Fields" (cslice.AppendSlice $embed.Fields) }}
    {{ if ne $embed.Title "How do I self-host ModMail?" }}
        {{ return }}
    {{ end }}
    {{ range $redflags }}
        {{ if (in $currentfieldnames .name) }}
            {{ return }}
        {{ end }}
        {{ $embed.Set "Fields" ($embed.Fields.Append .)}}
    {{ end }}
    {{editMessage nil $msgID (complexMessageEdit "embed" $embed)}}
{{ end}}

{{ $buttons := cslice }}
{{ $deletebutton := cbutton "label" "Delete Response" "custom_id" "support-response-delete" "style" 4 "emoji" (sdict "id" "1251255316121653343") }}
{{ $bookmarkbutton := cbutton "label" "Bookmark Response" "custom_id" "support-response-bookmark" "style" 2 "emoji" (sdict "id" "1251243802207846566") }}
{{ $togglebutton := cbutton "label" "Show/Hide Extra Information" "custom_id" "support-response-toggle" "style" 1 "emoji" (sdict "id" "1258858981372330165") }}
{{ $premiumbutton := cbutton "label" "Buy Premium" "custom_id" "support-response-premium" "url" "https://modmail.xyz/premium" "style" "link" "emoji" (sdict "id" "1251273319110414429") }}
{{ $invitebutton := cbutton "label" "Invite ModMail" "custom_id" "support-response-invite" "url" "https://modmail.xyz/invite" "style" "link" "emoji" (sdict "id" "1251255870701047909") }}
{{ $commandsbutton := cbutton "label" "ModMail Commands (`=help`)" "custom_id" "support-response-commands" "url" "https://modmail.xyz/commands" "style" "link""emoji" (sdict "id" "1258858466081116293") }}
{{ $logsbutton := cbutton "label" "Example Logs" "custom_id" "support-response-logs" "url" "https://modmail.xyz/logs/d7586c153425000-10d1416086c01033-10d141608b802047" "style" "link" "emoji" (sdict "id" "1254891424965722122") }}
{{ idguidebutton := cbutton "label" "Discord ID Guide" "custom_id" "support-response-idguide" "url" "https://support.discord.com/hc/en-us/articles/206346498-Where-can-I-find-my-User-Server-Message-ID-" "style" "link" "emoji" (sdict "id" "1258860561869832246") }}

{{ $buttons = $buttons.Append $deletebutton $bookmarkbutton $togglebutton }}
{{ $message := complexMessage "buttons" $buttons }}
{{ sendMessage nil $message }}

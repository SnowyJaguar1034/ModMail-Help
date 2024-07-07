{{ $buttons := cslice }}
{{ $deletebutton := cbutton "label" "Delete Response" "custom_id" "support-response-delete" "style" 4 "emoji" (sdict "id" "1251255316121653343") }}
{{ $bookmarkbutton := cbutton "label" "Bookmark Response" "custom_id" "support-response-bookmark" "style" 2 "emoji" (sdict "id" "1251243802207846566") }}
{{ $togglebutton := cbutton "label" "Toggle Extra Information" "custom_id" "support-response-toggle" "style" 1 "emoji" (sdict "id" "1258858981372330165") }}
{{ $premiumbutton := cbutton "label" "Buy Premium" "custom_id" "support-response-premium" "url" "https://modmail.xyz/premium" "style" "link" "emoji" (sdict "id" "1251273319110414429") }}
{{ $invitebutton := cbutton "label" "Invite ModMail" "custom_id" "support-response-invite" "url" "https://modmail.xyz/invite" "style" "link" "emoji" (sdict "id" "1251255870701047909") }}
{{ $commandsbutton := cbutton "label" "ModMail Commands (`=help`)" "custom_id" "support-response-commands" "url" "https://modmail.xyz/commands" "style" "link""emoji" (sdict "id" "1258858466081116293") }}
{{ $logsbutton := cbutton "label" "Example Logs" "custom_id" "support-response-logs" "url" "https://modmail.xyz/logs/d7586c153425000-10d1416086c01033-10d141608b802047" "style" "link" "emoji" (sdict "id" "1254891424965722122") }}
{{ idguidebutton := cbutton "label" "Discord ID Guide" "custom_id" "support-response-idguide" "url" "https://support.discord.com/hc/en-us/articles/206346498-Where-can-I-find-my-User-Server-Message-ID-" "style" "link" "emoji" (sdict "id" "1258860561869832246") }}
{{ $githubbutton := cbutton "label" "ModMail GitHub" "custom_id" "support-response-github" "url" "https://github.com/chamburr/modmail" "style" "link" "emoji" (sdict "id" "...") }}
{{ $officialv3button := cbutton "label" "Official V3 Guide" "custom_id" "support-response-officialv3" "url" "https://github.com/chamburr/modmail#self-hosting" "style" "link" "emoji" (sdict "id" "...") }}
{{ $officialv2button := cbutton "label" "Official V2 Guide" "custom_id" "support-response-officialv3" "url" "https://github.com/chamburr/modmail/blob/v2.1.2/README.md#self-hosting" "style" "link" "emoji" (sdict "id" "...") }}
{{ $communityv2button := cbutton "label" "Community V2 Guide" "custom_id" "support-response-communityv2" "url" "https://gist.github.com/waterflamev8/cab61e680e2fb5ea6027cbf144732925" "style" "link" "emoji" (sdict "id" "...") }}

NEED TO ADD BUTTONS FOR THE SELFHOST GUIDES

{{ $buttons = $buttons.Append $deletebutton $bookmarkbutton $togglebutton }}
{{ $message := complexMessage "buttons" $buttons }}
{{ sendMessage nil $message }}

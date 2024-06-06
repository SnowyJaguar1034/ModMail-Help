Regex: only accepting d(?:irect message|m)s from friends|message (?:could not be deliver|not deliver|block)ed|don\'t share a server|no (?:shared|mutual) server|^clyde(?:[ \x2D]bot)?

Options:
only accepting direct messages from friends
only accepting dms from friends
message could not be delivered
message not delivered
message blocked
don't share a server
no shared server
no mutual server
clyde
clyde bot
clyde-bot
i'm blocked
i am blocked
im blocked
bot blocked me

your message could not be delivered. this is usially becuase you don't share a server with the recipient or the recipient is only accepting direct messages from friends. ypu can see the full list of reasons here

^(only accepting direct messages from friends|only accepting dms from friends|message could not be delivered|message not delivered|message blocked|don't share a server|no shared server|no mutual server|clyde|clyde bot|clyde-bot|i'm blocked|i am blocked|im blocked|bot blocked me)$

(?i)(only accepting d(irect )?m(essage)?s from friends|message ((could )?not( be)? delivered|blocked)|don't share a server|no (shared|mutual) server|clyde(\z|( |-)bot)|i('| a|)m blocked|bot blocked me)

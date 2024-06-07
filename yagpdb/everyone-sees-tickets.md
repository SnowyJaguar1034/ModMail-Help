```
Regex: wh(?:en someone makes a ticket everyone can see it|en someone makes a ticket everyone can see|y can everyone see (?:t(?:he (?:m(?:(?:odm)?ail m)?|ticket m)essage|icket)s|my (?:m(?:odm)?ail|tickets))|y can everyone see (?:my )?ticket)|everyone can see(?: ticket)?
```

Options:
ticket everyone can see
tickets everyone can see
ticket everyone sees
ticket everyone see's
tickets everyone sees
tickets everyone see's
everyone can see the mail message
everyone see the mail message
everyone sees the mail message
everyone see's the mail message
everyone can see the ticket message
everyone see the ticket message
everyone sees the ticket message
everyone see's the ticket message
everyone can see tickets
everyone see tickets
everyone sees tickets
everyone see's tickets
everyone see ticket
everyone sees ticket
everyone see's ticket
everyone can see my mail
everyone see my mail
everyone sees my mail
everyone see's my mail
everyone see my mod mail
everyone sees my mod mail
everyone see's my mod mail
everyone see my modmail
everyone sees my modmail
everyone see's my modmail
everyone see my mod-mail
everyone sees my mod-mail
everyone see's my mod-mail

```
(ticket|tickets|everyone) (can|see|sees|see's) (the )?(mail|ticket|tickets|my mail|mod mail|modmail|mod-mail) message$
```

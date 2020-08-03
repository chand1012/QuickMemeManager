# QuickMemeManager

Manages Patrons for [DiscordQuickMeme](https://github.com/chand1012/Discord-Quick-Meme). Here is how it will work when completed:

 - Will periodically scan the [DiscordQuickMeme](https://discord.gg/YNnp9uy) server for users
    - If a user is in one of the Patron roles, will check the database to see if user is in the correct table
        - If not in the table, send them a PM asking what server(s) they would like to boost.
    - If a user either leaves or is no longer in the Patron roles, will delete them from the table.

 - Will be use shards to make sure that if this becomes a big thing the bot can handle it.
 
 - That's really it.

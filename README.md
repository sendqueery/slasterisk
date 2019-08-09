# slasterisk

integration to allow asterisk to send notifications to slack

## to-do

*   add instructions for building `go-lame` for the appropriate architecture

## goal

whenever a voicemail is left for a given mailbox, send a notification to a slack user/channel along with the accompanying audio file

## config.json sections/values

(this goes in the config folder as `config.json` see example_config.json for a sample file)

### slackinfo

*   token: this is the api token for your slack application
*   channel_name: this is the channel to which you want slasterisk to send notifications, if applicable

### asteriskinfo

*   vm_dir: this is the absolute path on your asterisk server to where the voicemail folders are for each mailbox; I believe the default is `/var/spool/asterisk/voicemail/default/`, at least for FreePBX asterisk installs
*   extension: this is the extension for which you want slasterisk to trigger, if applicable

// This file is automatically generated. Do not modify it manually.

package main

import (
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
)

var manifest *model.Manifest

const manifestStr = `
{
  "id": "bitbucket",
  "name": "Bitbucket",
  "description": "Bitbucket plugin for Mattermost.",
  "homepage_url": "https://github.com/kosgrz/mattermost-plugin-bitbucket",
  "support_url": "https://github.com/kosgrz/mattermost-plugin-bitbucket/issues",
  "icon_path": "assets/icon.svg",
  "version": "1.1.0",
  "min_server_version": "5.25.0",
  "server": {
    "executables": {
      "linux-amd64": "server/dist/plugin-linux-amd64",
      "darwin-amd64": "server/dist/plugin-darwin-amd64",
      "windows-amd64": "server/dist/plugin-windows-amd64.exe"
    },
    "executable": ""
  },
  "webapp": {
    "bundle_path": "webapp/dist/main.js"
  },
  "settings_schema": {
    "header": "To set up the Bitbucket plugin, you need to register a Bitbucket OAuth consumer here https://bitbucket.org/YOURWORKSPACE/workspace/settings/oauth-consumers/new.",
    "footer": "* To report an issue, make a suggestion or a contribution, [check the repository](https://github.com/kosgrz/mattermost-plugin-bitbucket).",
    "settings": [
      {
        "key": "BitbucketOAuthClientID",
        "display_name": "Bitbucket OAuth Client ID (Key)",
        "type": "text",
        "help_text": "The client ID for the OAuth app registered with Bitbucket.",
        "placeholder": "",
        "default": null
      },
      {
        "key": "BitbucketOAuthClientSecret",
        "display_name": "Bitbucket OAuth Client Secret",
        "type": "text",
        "help_text": "The client secret for the OAuth app registered with Bitbucket.",
        "placeholder": "",
        "default": null
      },
      {
        "key": "WebhookSecret",
        "display_name": "Webhook Secret",
        "type": "generated",
        "help_text": "The webhook secret set in Bitbucket.",
        "placeholder": "",
        "default": null
      },
      {
        "key": "EncryptionKey",
        "display_name": "At Rest Encryption Key",
        "type": "generated",
        "help_text": "The AES encryption key used to encrypt stored access tokens.",
        "placeholder": "",
        "default": null
      },
      {
        "key": "BitbucketOrg",
        "display_name": "Bitbucket Organization",
        "type": "text",
        "help_text": "(Optional) Set to lock the plugin to a single Bitbucket organization.",
        "placeholder": "",
        "default": null
      },
      {
        "key": "EnablePrivateRepo",
        "display_name": "Enable Private Repositories",
        "type": "bool",
        "help_text": "(Optional) Allow the plugin to work with private repositories. Enabling private repositories will require existing users to reconnect their accounts to gain access to private repositories. A message will be automatically be sent to affected users next time they load Mattermost alerting them of this.",
        "placeholder": "",
        "default": null
      }
    ]
  }
}
`

func init() {
	manifest = model.ManifestFromJson(strings.NewReader(manifestStr))
}

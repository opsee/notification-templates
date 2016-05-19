package slack

import "github.com/hoisie/mustache"

var NewSignup = `{
  "text": ":moneybag: :rollin: NEW SIGNUP :pepe-trump: :moneybag:",
  "username": "CustomerBot",
  "icon_url": "https://s3-us-west-1.amazonaws.com/opsee-public-images/slack-avi-48-green.png",
  "attachments": [
    {
      "text": "{{user_name}} - {{user_email}}",
      "color": "#81C784"{{#referrer}},
      "fields": [
          {
              "title": "Referrer",
              "value": "{{referrer}}",
              "short": true
          }
          ]{{/referrer}}
    }
  ]
}
`

func init() {
	tmpl, err := mustache.ParseString(NewSignup)
	if err != nil {
		panic(err)
	}
	Templates["new-signup"] = tmpl
}

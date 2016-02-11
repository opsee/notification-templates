package slack
var BastionOnline = `{
  "text": "*BASTION ONLINE*",
  "username": "ErrorBot",
  "icon_url": "https://s3-us-west-1.amazonaws.com/opsee-public-images/slack-avi-48-green.png",
  "attachments": [
    {
      "text": "Customer: {{customer_id}}",
      "color": "#f44336",
      "fields": [
        {
          "title": "Bastion ID",
          "value": "{{bastion_id}}",
          "short": true
        },
        {
          "title": "Last Seen",
          "value": "{{last_seen}}",
          "short": true
        },
        {
          "title": "Current State",
          "value": "{{current_state}}",
          "short": true
        }
      ]
    }
  ]
}
`

package slack
var BastionOffline = `{
  "text": "*BASTION OFFLINE*",
  "username": "BastionTracker",
  "icon_url": "https://s3-us-west-1.amazonaws.com/opsee-public-images/slack-avi-48-red.png",
  "attachments": [
    {
      "text": "User: {{name}} ({{email}})",
      "color": "#f44336",
      "fields": [
        {
          "title": "Customer ID",
          "value": "{{customer_id}}",
          "short": true
        },
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

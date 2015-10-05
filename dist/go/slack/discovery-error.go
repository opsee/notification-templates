package slack
var DiscoveryError = `{
  "text": "Discovery Error!",
  "username": "CustomerBot",
  "icon_url": "https://s3-us-west-1.amazonaws.com/opsee-public-images/opsee-avi-48.jpg",
  "attachments": [
    {
      "text": "Customer ID: {{customer_id}}",
      "fields": [
        {
          "title": "User ID",
          "value": "{{user_id}}",
          "short": true
        },
        {
          "title": "AWS Region",
          "value": "{{aws_region}}",
          "short": true
        },
        {
          "title": "Instance Errors",
          "value": "{{instance_error_count}}",
          "short": true
        },
        {
          "title": "Group Errors",
          "value": "{{group_error_count}}",
          "short": true
        },
        {
          "title": "Last Error",
          "value": "{{last_error}}"
        }
      ]
    }
  ]
}
`

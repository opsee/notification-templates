package slack
var NewCustomer = `{
  "text": "New Opsee Customer!",
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
          "value": "{{region}}",
          "short": true
        },
        {
          "title": "EC2 Instances",
          "value": "{{instance_count}}",
          "short": true
        },
        {
          "title": "RDS Instances",
          "value": "{{db_instance_count}}",
          "short": true
        },
        {
          "title": "Security Groups",
          "value": "{{security_group_count}}",
          "short": true
        },
        {
          "title": "RDS Security Groups",
          "value": "{{db_security_group_count}}",
          "short": true
        },
        {
          "title": "Load Balancers",
          "value": "{{load_balancer_count}}",
          "short": true
        },
        {
          "title": "Autoscaling Groups",
          "value": "{{autoscaling_group_count}}",
          "short": true
        },
        {
          "title": "Opsee Created Checks",
          "value": "{{check_count}}",
          "short": true
        }
      ]
    }
  ]
}
`

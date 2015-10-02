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
          "value": "{{user_id}}"
        },
        {
          "title": "AWS Region",
          "value": "{{aws_region}}"
        },
        {
          "title": "EC2 Instances",
          "value": "{{instance_count}}"
        },
        {
          "title": "RDS Instances",
          "value": "{{db_instance_count}}"
        },
        {
          "title": "Security Groups",
          "value": "{{security_group_count}}"
        },
        {
          "title": "RDS Security Groups",
          "value": "{{db_security_group_count}}"
        },
        {
          "title": "Load Balancers",
          "value": "{{load_balancer_count}}"
        },
        {
          "title": "Autoscaling Groups",
          "value": "{{autoscaling_group_count}}"
        },
        {
          "title": "Opsee Created Checks",
          "valuefo": "{{checks_count}}"
        }
      ]
    }
  ]
}
`

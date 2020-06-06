package main

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceBool() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"condition": &schema.Schema{
				Type: schema.TypeBool,
				Required: true,
			},
			"name": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
			"error_message": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
		},

		Create:             resourceBoolCreate,
		Read:               resourceBoolRead,
		Update:             resourceBoolUpdate,
		Delete:             resourceBoolDelete,

	}
}


func resourceBoolCreate(d *schema.ResourceData,m interface{}) error {
	name := d.Get("name").(string)
	condition := d.Get("condition").(bool)
	if !condition {
		return fmt.Errorf("%s",d.Get("error_message").(string))
	}
	d.SetId(name)
	return resourceBoolRead(d,m)
}

func resourceBoolRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceBoolUpdate(d *schema.ResourceData, m interface{}) error {
	condition := d.Get("condition").(bool)
	if !condition {
		return fmt.Errorf("%s",d.Get("error_message").(string))
	}
	return resourceBoolRead(d, m)
}

func resourceBoolDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
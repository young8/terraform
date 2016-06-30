package aws

import (
        "log"

        "fmt"
        "github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/service/emr"
        "github.com/hashicorp/terraform/helper/schema"
)

func resourceAwsEMRTaskGroup() *schema.Resource {
        return &schema.Resource{
                Create: resourceAwsEMRTaskGroupCreate,
                Read:   resourceAwsEMRTaskGroupRead,
                Update: resourceAwsEMRTaskGroupUpdate,
                Delete: resourceAwsEMRTaskGroupDelete,
                Schema: map[string]*schema.Schema{
                        "cluster_id": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                                ForceNew: true,
                        },
                        "instance_type": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                        "instance_count": &schema.Schema{
                                Type:     schema.TypeInt,
                                Optional: true,
                                Default:  60,
                        },
                        "name": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                        },
                },
        }
}

func resourceAwsEMRTaskGroupCreate(d *schema.ResourceData, meta interface{}) error {
        conn := meta.(*AWSClient).emrconn

        clusterId := d.Get("cluster_id").(string)
        instanceType := d.Get("instance_type").(string)
        instanceCount := d.Get("instance_count").(int)
        groupName := d.Get("name").(string)

        log.Printf("[DEBUG] Creating EMR cluster task group")
        params := &emr.AddInstanceGroupsInput{
                InstanceGroups: []*emr.InstanceGroupConfig{
                        {
                                InstanceRole:  aws.String("TASK"),
                                InstanceCount: aws.Int64(int64(instanceCount)),
                                InstanceType:  aws.String(instanceType),
                                Name:          aws.String(groupName),
                        },
                },
                JobFlowId: aws.String(clusterId),
        }
        resp, err := conn.AddInstanceGroups(params)
        if err != nil {
                log.Printf("[ERROR] %s", err)
                return err
        }

        fmt.Println(resp)

        log.Printf("[DEBUG] Created EMR Cluster task group done...")
        d.SetId(*resp.InstanceGroupIds[0])

        return nil
}

func resourceAwsEMRTaskGroupRead(d *schema.ResourceData, meta interface{}) error {

        return nil
}

func resourceAwsEMRTaskGroupUpdate(d *schema.ResourceData, meta interface{}) error {
        conn := meta.(*AWSClient).emrconn

        log.Printf("[DEBUG] Modify EMR task group")
        instanceCount := d.Get("instance_count").(int)

        params := &emr.ModifyInstanceGroupsInput{
                InstanceGroups: []*emr.InstanceGroupModifyConfig{
                        {
                                InstanceGroupId: aws.String(d.Id()),
                                InstanceCount:   aws.Int64(int64(instanceCount)),
                        },
                },
        }
        respModify, errModify := conn.ModifyInstanceGroups(params)
        if errModify != nil {
                log.Printf("[ERROR] %s", errModify)
                return errModify
        }

        fmt.Println(respModify)

        log.Printf("[DEBUG] Modify EMR task group done...")

        return nil
}

func resourceAwsEMRTaskGroupDelete(d *schema.ResourceData, meta interface{}) error {

        return nil
}
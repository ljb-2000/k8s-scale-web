package controllers

import (
	"encoding/json"

	"time"

	"fmt"
	"os/exec"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type APIStruct struct {
	Metadata struct {
		Name string `json:"name"`
	} `json:"metadata"`
	Spec struct {
		CPUUtilization struct {
			TargetPercentage int `json:"targetPercentage"`
		} `json:"cpuUtilization"`
		MaxReplicas int `json:"maxReplicas"`
		MinReplicas int `json:"minReplicas"`
		ScaleRef    struct {
			APIVersion  string `json:"apiVersion"`
			Kind        string `json:"kind"`
			Name        string `json:"name"`
			Subresource string `json:"subresource"`
		} `json:"scaleRef"`
	} `json:"spec"`
}

type ProcessData struct {
	Namespace string
	Name      string
	Count     int
	Period    int
	Time      time.Time
}

var APIURL string

func init() {
	APIURL = beego.AppConfig.String("APIURL")
}

// APIControllerController operations for APIController
type APIControllerController struct {
	beego.Controller
}

// URLMapping ...
func (c *APIControllerController) URLMapping() {
	c.Mapping("ReplicationController", c.ReplicationController)
	c.Mapping("AutoScale", c.AutoScale)
	c.Mapping("RemoveScale", c.RemoveScale)
	c.Mapping("NewScale", c.NewScale)
	c.Mapping("NewTickScale", c.NewTickScale)
}

// @router /rc [get]
func (c *APIControllerController) ReplicationController() {
	http := httplib.Get(APIURL + "api/v1/replicationcontrollers")
	result := make(map[string]interface{})
	http.ToJSON(&result)
	c.Data["json"] = result
	c.ServeJSON()
}

// @router /as [get]
func (c *APIControllerController) AutoScale() {
	http := httplib.Get(APIURL + "apis/extensions/v1beta1/horizontalpodautoscalers")
	result := make(map[string]interface{})
	http.ToJSON(&result)
	c.Data["json"] = result
	c.ServeJSON()
}

// @router /as/remove/:namespace/:name [get]
func (c *APIControllerController) RemoveScale() {
	rc_namespace := c.Ctx.Input.Param(":namespace")
	rc_name := c.Ctx.Input.Param(":name")
	http := httplib.Delete(APIURL + "apis/extensions/v1beta1/namespaces/" + rc_namespace + "/horizontalpodautoscalers/" + rc_name)
	result := make(map[string]interface{})
	http.ToJSON(&result)
	// result["status"] = "Success"
	c.Data["json"] = result
	c.ServeJSON()
}

// @router /as/new [get]
func (c *APIControllerController) NewScale() {
	rc_namespace := c.GetString("namespace", "default")
	rc_name := c.GetString("name", "default")
	rc_cpu, _ := c.GetInt("cpu", 50)
	rc_min, _ := c.GetInt("min", 50)
	rc_max, _ := c.GetInt("max", 50)

	postData := new(APIStruct)
	postData.Metadata.Name = rc_name
	postData.Spec.ScaleRef.APIVersion = "v1"
	postData.Spec.ScaleRef.Kind = "ReplicationController"
	postData.Spec.ScaleRef.Name = rc_name
	postData.Spec.ScaleRef.Subresource = "scale"

	postData.Spec.MinReplicas = rc_min
	postData.Spec.MaxReplicas = rc_max
	postData.Spec.CPUUtilization.TargetPercentage = rc_cpu
	requestBody, _ := json.Marshal(postData)

	http := httplib.Post(APIURL + "apis/extensions/v1beta1/namespaces/" + rc_namespace + "/horizontalpodautoscalers")
	result := make(map[string]interface{})

	http.Body(string(requestBody))
	http.ToJSON(&result)
	c.Data["json"] = result
	c.ServeJSON()
}

//kubectl autoscale rc redis-master-rc --min=2 --max=5 --cpu-percent=10

// @router /as/newtick [get]
func (c *APIControllerController) NewTickScale() {
	rc_namespace := c.GetString("namespace", "default")
	rc_name := c.GetString("name", "default")
	rc_number, _ := c.GetInt("number", 50)
	rc_time_string := c.GetString("time", "")
	rc_period, _ := c.GetInt("period", 0)

	data := new(ProcessData)

	data.Namespace = rc_namespace
	data.Name = rc_name
	data.Count = rc_number
	data.Period = rc_period
	if data.Period == 1 {
		data.Time, _ = time.Parse("15:04:05", rc_time_string)
	} else {
		data.Time, _ = time.Parse("2006-01-02 15:04:05", rc_time_string)
	}

	CreateCron(data)
	result := make(map[string]interface{})
	result["status"] = "Success"
	c.Data["json"] = result
	c.ServeJSON()
}

func CreateCron(data *ProcessData) {
	cmdString := ""
	if data.Period == 1 {
		cmdString = fmt.Sprintf("echo '%d %d * * * root kubectl scale rc %s --replicas %d --namespace %s' >> /etc/crontab",
			data.Time.Minute(), data.Time.Hour(), data.Name, data.Count, data.Namespace)
	} else {
		cmdString = fmt.Sprintf("echo '%d %d %d %d %d root kubectl scale rc %s --replicas %d --namespace %s' >> /etc/crontab",
			data.Time.Minute(), data.Time.Hour(), data.Time.Day(), data.Time.Month(), data.Time.Year(), data.Name, data.Count, data.Namespace)
	}

	fmt.Println(cmdString)
	cmd := exec.Command("/bin/sh", "-c", cmdString)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(string(out))
}

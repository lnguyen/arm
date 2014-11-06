package consul

import (
	"fmt"

	consulapi "github.com/armon/consul-api"
)

//ConsulCollector collect where fingers are from consul
type ConsulCollector struct {
	client *consulapi.Client
}

//NewConsulCollector create new consul collector
func NewConsulCollector(consulAddress string) (*ConsulCollector, error) {
	config := consulapi.DefaultConfig()
	config.Address = consulAddress
	consulClient, err := consulapi.NewClient(config)
	if err != nil {
		return &ConsulCollector{}, err
	}
	return &ConsulCollector{client: consulClient}, nil
}

func (c *ConsulCollector) agentConfigData() map[string]interface{} {
	data, _ := c.client.Agent().Self()
	return data["Config"]
}

//DataCenter determine Data center name of collector
func (c *ConsulCollector) DataCenter() string {
	return c.agentConfigData()["Datacenter"].(string)
}

//ServiceByName Find services by Name
func (c *ConsulCollector) ServiceByName(name string) []string {
	var serviceList []string
	s, _, _ := c.client.Catalog().Service(name, "", &consulapi.QueryOptions{})
	for _, value := range s {
		fmt.Println(value.Node)
		fmt.Println(value.Address)

	}
	fmt.Println(s)
	return serviceList
}

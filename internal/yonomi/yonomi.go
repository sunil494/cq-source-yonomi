package yonomi

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

type Client struct {
}

type Option func(*Client)

type YonomiOutputStruct struct {
	Device struct {
		DisplayName        string
		ProductInformation struct {
			Description  string
			Manufacturer string
			Model        string
		}
		Traits []struct {
			State struct {
				Percentage struct {
					Reported struct {
						Value int
					}
				}
			}
		}
	}
}

type SchlageLockDataBlock struct {
	DisplayName  string
	Description  string
	Manufacturer string
	Model        string
	BatteryLife  int
}

func NewClient(opts ...Option) (*Client, error) {
	c := &Client{}
	return c, nil
}

func GetSchlageLockData(authorization string, deviceId string) (SchlageLockDataBlock, error) {
	client := graphql.NewClient("https://platform.yonomi.cloud/graphql")
	query := `
	query getMyDevices($id: ID!) {
		device(id: $id) {
		  id
		  productInformation {
			description
			manufacturer
			model
			serialNumber
		  }
		  displayName
		  traits {
			... on BatteryLevelDeviceTrait {
			  state {
				percentage {
				  reported {
					value
				  }
				}
			  }
			  name
			}
		  }
		  updatedAt
		}
	  }
	`
	request := graphql.NewRequest(query)
	request.Var("id", deviceId)
	request.Header.Set("Authorization", authorization)
	var response YonomiOutputStruct
	err := client.Run(context.Background(), request, &response)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
	output := SchlageLockDataBlock{
		DisplayName:  response.Device.DisplayName,
		Description:  response.Device.ProductInformation.Description,
		Manufacturer: response.Device.ProductInformation.Manufacturer,
		Model:        response.Device.ProductInformation.Model,
		BatteryLife:  response.Device.Traits[len(response.Device.Traits)-1].State.Percentage.Reported.Value,
	}
	return output, nil
}

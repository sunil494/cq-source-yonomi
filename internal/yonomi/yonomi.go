package yonomi

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

type Client struct {
}

type Option func(*Client)

type YonomiDevicesOutputStruct struct {
	Device struct {
		Id                 string
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

type YonomiDevicesSchlageTraitsStruct struct {
	Device struct {
		Traits []struct {
			State struct {
				IsJammed struct {
					Reported struct {
						Value bool
					}
				}
				IsLocked struct {
					Reported struct {
						Value bool
					}
				}
				PinCodeCredentials struct {
					Reported struct {
						Value struct {
							Edges []struct {
								Node struct {
									Name    string
									PinCode string
								}
							}
						}
					}
				}
				Percentage struct {
					Reported struct {
						Value int
					}
				}
			}
		}
	}
}

type DevicesDataBlock struct {
	Id           string
	DisplayName  string
	Description  string
	Manufacturer string
	Model        string
}

type SchlageTraitsDataBlock struct {
	BatteryLife        int
	IsJammed           bool
	IsLocked           bool
	PinCodeCredentials []struct {
		Node struct {
			Name    string
			PinCode string
		}
	}
}

func NewClient(opts ...Option) (*Client, error) {
	c := &Client{}
	return c, nil
}

func GetDevicesData(authorization string, deviceId string) (DevicesDataBlock, error) {
	client := graphql.NewClient("https://platform.yonomi.cloud/graphql")
	query := `
	query getMyDevices($id: ID!) {
		device(id: $id) {
		  id
		  productInformation {
			description
			manufacturer
			model
		  }
		  displayName
		}
	  }
	`
	request := graphql.NewRequest(query)
	request.Var("id", deviceId)
	request.Header.Set("Authorization", authorization)
	var response YonomiDevicesOutputStruct
	err := client.Run(context.Background(), request, &response)
	if err != nil {
		panic(err)
	}
	output := DevicesDataBlock{
		Id:           response.Device.Id,
		DisplayName:  response.Device.DisplayName,
		Description:  response.Device.ProductInformation.Description,
		Manufacturer: response.Device.ProductInformation.Manufacturer,
		Model:        response.Device.ProductInformation.Model,
	}
	return output, nil
}

func GetSchlageTraitsData(authorization string, deviceId string) (SchlageTraitsDataBlock, error) {
	client := graphql.NewClient("https://platform.yonomi.cloud/graphql")
	query := `
	query getMyDevices($id: ID!) {
		device(id: $id) {
			traits {
			  ... on BatteryLevelDeviceTrait {
				state {
				  percentage {
					reported {
					  value
					}
				  }
				}
			  }
			  ... on PinCodeCredentialDeviceTrait {
				state {
				  pinCodeCredentials {
					reported {
					  value {
						edges {
						  node {
							name
							pinCode
						  }
						}
					  }
					}
				  }
				}
			  }
			  ... on LockDeviceTrait {
				state {
				  isJammed {
					reported {
					  value
					}
				  }
				  isLocked {
					reported {
					  value
					}
				  }
				}
			  }
			}
		  }
  		}
	`
	request := graphql.NewRequest(query)
	request.Var("id", deviceId)
	request.Header.Set("Authorization", authorization)
	var response YonomiDevicesSchlageTraitsStruct
	err := client.Run(context.Background(), request, &response)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
	output := SchlageTraitsDataBlock{
		IsJammed:           response.Device.Traits[0].State.IsJammed.Reported.Value,
		IsLocked:           response.Device.Traits[0].State.IsLocked.Reported.Value,
		PinCodeCredentials: response.Device.Traits[1].State.PinCodeCredentials.Reported.Value.Edges,
		BatteryLife:        response.Device.Traits[3].State.Percentage.Reported.Value,
	}
	return output, nil
}

package tablestore

import (
	"os"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/golang/glog"
)

type TableClient struct {
	tableName string
	client    *tablestore.TableStoreClient
}

func InitClient() *TableClient {
	endpoint := os.Getenv("OTS_TEST_ENDPOINT")
	tableName := os.Getenv("TABLE_NAME")
	instanceName := os.Getenv("OTS_TEST_INSTANCENAME")
	accessKeyId := os.Getenv("OTS_TEST_KEYID")
	accessKeySecret := os.Getenv("OTS_TEST_SECRET")
	client := tablestore.NewClient(endpoint, instanceName, accessKeyId, accessKeySecret)
	return &TableClient{
		tableName: tableName,
		client:    client,
	}
}

func (tableClient *TableClient) Query(adcode, date string) (map[string]string, error) {
	getRowRequest := new(tablestore.GetRowRequest)
	criteria := new(tablestore.SingleRowQueryCriteria)
	putPk := &tablestore.PrimaryKey{}

	putPk.AddPrimaryKeyColumn("adcode", adcode)
	putPk.AddPrimaryKeyColumn("date", date)
	criteria.PrimaryKey = putPk

	getRowRequest.SingleRowQueryCriteria = criteria
	getRowRequest.SingleRowQueryCriteria.TableName = tableClient.tableName
	getRowRequest.SingleRowQueryCriteria.MaxVersion = 1
	getResp, err := tableClient.client.GetRow(getRowRequest)
	if err != nil {
		glog.Errorf("GetRow failed with error: %s", err.Error())
		return nil, err
	}
	weatherMap := make(map[string]string, 0)
	if getResp.PrimaryKey.PrimaryKeys != nil {
		for _, col := range getResp.PrimaryKey.PrimaryKeys {
			weatherMap[col.ColumnName] = col.Value.(string)
		}
	}
	if getResp.Columns != nil {
		for _, col := range getResp.Columns {
			// 过滤掉 id 信息
			if col.ColumnName == "id" {
				continue
			}
			weatherMap[col.ColumnName] = col.Value.(string)
		}
	}
	return weatherMap, nil
}

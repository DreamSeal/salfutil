package ossutil

import (
	"bytes"
	"casino_common/common/Error"
	"casino_common/common/log"
	ossCfg "salfutil/conf"
	"casino_common/proto/ddproto"
	"fmt"
	"salfutil/aliyun-oss-go-sdk/oss"
	"os"
	"strings"
)

var (
	endpoint   = ossCfg.EndPoint
	accessID   = ossCfg.AccessKey_ID
	accessKey  = ossCfg.SECRET
	bucketName = ossCfg.BucketName
)


//将回放数据存到oss
func SavePlaybackToOss(gameid,round,playbackNumber int32,data []byte) error{
	//data, err := proto.Marshal(p)
	if gameid < 1 || round < 1 || playbackNumber < 1 || data == nil || len(data) == 0 {
		log.E("保存回放数据异常gameid%v,round%v,playbackNumber%v,data%v",gameid,round,playbackNumber,data)
		return Error.NewError(-1,"存储数据不正确")
	}

	folderName := ""
	switch gameid {
	case int32(ddproto.CommonEnumGame_GID_PAOHUZI):
		folderName = "phz/"
	case int32(ddproto.CommonEnumGame_GID_PDK):
		folderName = "pdk/"
	case int32(ddproto.CommonEnumGame_GID_ZHUANZHUAN):
		folderName = "mjzhzh/"
	case int32(ddproto.CommonEnumGame_GID_ZHADAN):
		folderName = "zhadan/"
	case int32(ddproto.CommonEnumGame_GID_PHZ_SHAOYANGBOPI):
		folderName = "sybp/"
	}

	bucket, err := GetTestBucket(bucketName)
	if err != nil {
		return err
	}

	folderName = folderName + fmt.Sprintf("%v-%v",playbackNumber,round)
	log.T("回放存储路径%v",folderName)
	err = bucket.PutObject(folderName, bytes.NewReader(data))
	if err != nil {
		return err
	}
	return nil
}



// HandleError is the error handling method in the sample code
func HandleError(err error) {
	fmt.Println("occurred error:", err)
	os.Exit(-1)
}

// GetTestBucket creates the test bucket
func GetTestBucket(bucketName string) (*oss.Bucket, error) {
	// New client
	client, err := oss.New(endpoint, accessID, accessKey)
	if err != nil {
		return nil, err
	}

	isBucketExist, err := client.IsBucketExist(bucketName)
	if !isBucketExist {
		// Create bucket
		err = client.CreateBucket(bucketName)
		if err != nil {
			return nil, err
		}
	}

	// Get bucket
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}

// DeleteTestBucketAndLiveChannel 删除sample的channelname和bucket，该函数为了简化sample，让sample代码更明了
func DeleteTestBucketAndLiveChannel(bucketName string) error {
	// New Client
	client, err := oss.New(endpoint, accessID, accessKey)
	if err != nil {
		return err
	}

	// Get Bucket
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}

	marker := ""
	for {
		result, err := bucket.ListLiveChannel(oss.Marker(marker))
		if err != nil {
			log.E("bucket.ListLiveChannel(oss.Marker(marker)) err %v",err)
			//HandleError(err)
		}

		for _, channel := range result.LiveChannel {
			err := bucket.DeleteLiveChannel(channel.Name)
			if err != nil {
				log.E("bucket.DeleteLiveChannel(channel.Name) err %v",err)
				//HandleError(err)
			}
		}

		if result.IsTruncated {
			marker = result.NextMarker
		} else {
			break
		}
	}

	// Delete Bucket
	err = client.DeleteBucket(bucketName)
	if err != nil {
		return err
	}

	return nil
}

// DeleteTestBucketAndObject deletes the test bucket and its objects
func DeleteTestBucketAndObject(bucketName string) error {
	// New client
	client, err := oss.New(endpoint, accessID, accessKey)
	if err != nil {
		return err
	}

	// Get bucket
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}

	// Delete part
	keyMarker := oss.KeyMarker("")
	uploadIDMarker := oss.UploadIDMarker("")
	for {
		lmur, err := bucket.ListMultipartUploads(keyMarker, uploadIDMarker)
		if err != nil {
			return err
		}
		for _, upload := range lmur.Uploads {
			var imur = oss.InitiateMultipartUploadResult{Bucket: bucket.BucketName,
				Key: upload.Key, UploadID: upload.UploadID}
			err = bucket.AbortMultipartUpload(imur)
			if err != nil {
				return err
			}
		}
		keyMarker = oss.KeyMarker(lmur.NextKeyMarker)
		uploadIDMarker = oss.UploadIDMarker(lmur.NextUploadIDMarker)
		if !lmur.IsTruncated {
			break
		}
	}

	// Delete objects
	marker := oss.Marker("")
	for {
		lor, err := bucket.ListObjects(marker)
		if err != nil {
			return err
		}
		for _, object := range lor.Objects {
			err = bucket.DeleteObject(object.Key)
			if err != nil {
				return err
			}
		}
		marker = oss.Marker(lor.NextMarker)
		if !lor.IsTruncated {
			break
		}
	}

	// Delete bucket
	err = client.DeleteBucket(bucketName)
	if err != nil {
		return err
	}

	return nil
}

// Object defines pair of key and value
type Object struct {
	Key   string
	Value string
}

// CreateObjects creates some objects
func CreateObjects(bucket *oss.Bucket, objects []Object) error {
	for _, object := range objects {
		err := bucket.PutObject(object.Key, strings.NewReader(object.Value))
		if err != nil {
			return err
		}
	}
	return nil
}

// DeleteObjects deletes some objects.
func DeleteObjects(bucket *oss.Bucket, objects []Object) error {
	for _, object := range objects {
		err := bucket.DeleteObject(object.Key)
		if err != nil {
			return err
		}
	}
	return nil
}

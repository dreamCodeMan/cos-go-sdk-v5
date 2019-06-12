package cos

import (
	"encoding/xml"
	"fmt"
	"time"
)

// ListAllMyBucketsResult 获取bucket列表的结果
type ListAllMyBucketsResult struct {
	Owner struct {
		ID          string
		DisplayName string
	}

	Buckets struct {
		Bucket []struct {
			Name       string
			Location   string
			CreateDate string
		}
	}
}

// Error 错误消息
type Error struct {
	Code      string
	Message   string
	Resource  string
	RequestID string `xml:"RequestId"`
	TraceID   string `xml:"TaceId"`
}

// HTTPError http error struct
type HTTPError struct {
	Code    int
	Message string
}

// Error  error interface
func (he HTTPError) Error() string {
	return fmt.Sprintf("%d:%s", he.Code, he.Message)
}

// AccessControlPolicy acl return
type AccessControlPolicy struct {
	Owner struct {
		ID          string
		DisplayName string
	}
	AccessControlList struct {
		Grant []struct {
			Grantee struct {
				ID          string
				DisplayName string
			}
			Permission string
		}
	}
}

// ListBucketResult list bucket contents result
type ListBucketResult struct {
	Name         string
	EncodingType string `xml:"Encoding-Type"`
	Prefix       string
	Marker       string
	MaxKeys      int
	IsTruncated  bool
	NextMarker   string
	Contents     []struct {
		Key          string
		LastModified string
		ETag         string
		Size         int64
		Owner        struct {
			ID string
		}
		StorageClass string
	}
	CommonPrefixes []struct {
		Prefix string
	}
}

// ListMultipartUploadsResult list uploading task
type ListMultipartUploadsResult struct {
	Bucket             string
	EncodingType       string `xml:"Encoding-Type"`
	KeyMarker          string
	UploadIDMarker     string `xml:"UploadIdMarker"`
	NextKeyMarker      string
	NextUploadIDMarker string `xml:"NextUploadIdMarker"`
	MaxUploads         int
	IsTruncated        bool
	Prefix             string
	Delimiter          string
	Upload             []struct {
		Key          string
		UploadID     string
		StorageClass string
		Initiator    struct {
			UIN string
		}
		Owner struct {
			UID string
		}
		Initiated string
	}
	CommonPrefixes []struct {
		Prefix string
	}
}

// InitiateMultipartUploadResult init slice upload
type InitiateMultipartUploadResult struct {
	Bucket   string
	Key      string
	UploadID string `xml:"UploadId"`
}

// CompleteMultipartUploadResult compeleted slice upload
type CompleteMultipartUploadResult struct {
	Location string
	Bucket   string
	Key      string
	ETag     string
}

// ListObjectsResult ListObjects请求返回结果
type ListObjectsResult struct {
	XMLName        xml.Name           `xml:"ListBucketResult"`
	Prefix         string             `xml:"Prefix"`                // 本次查询结果的开始前缀
	Marker         string             `xml:"Marker"`                // 这次查询的起点
	NextMarker     string             `xml:"NextMarker"`            // 下次查询的起点
	MaxKeys        int                `xml:"MaxKeys"`               // 请求返回结果的最大数目
	IsTruncated    bool               `xml:"IsTruncated"`           // 是否所有的结果都已经返回
	Objects        []ObjectProperties `xml:"Contents"`              // Object类别
	CommonPrefixes []string           `xml:"CommonPrefixes>Prefix"` // 以delimiter结尾并有共同前缀的Object的集合
}

// ObjectProperties Objecct属性
type ObjectProperties struct {
	XMLName      xml.Name  `xml:"Contents"`
	Key          string    `xml:"Key"`          // Object的Key
	Size         int64     `xml:"Size"`         // Object的长度字节数
	ETag         string    `xml:"ETag"`         // 标示Object的内容
	Owner        Owner     `xml:"Owner"`        // 保存Object拥有者信息的容器
	LastModified time.Time `xml:"LastModified"` // Object最后修改时间
	StorageClass string    `xml:"StorageClass"` // Object的存储类型
}

// Owner Bucket/Object的owner
type Owner struct {
	XMLName     xml.Name `xml:"Owner"`
	ID          string   `xml:"ID"`          // 用户ID
	DisplayName string   `xml:"DisplayName"` // Owner名字
}

// SliceError slice upload err
type SliceError struct {
	Message string
}

// Error implements error
func (se SliceError) Error() string {
	return fmt.Sprintf("上传分片失败:%s", se.Message)
}

// ParamError slice upload err
type ParamError struct {
	Message string
}

// Error implements error
func (pe ParamError) Error() string {
	return fmt.Sprintf("参数错误:%s", pe.Message)
}

// FileError slice upload err
type FileError struct {
	Message string
}

// Error implements error
func (fe FileError) Error() string {
	return fmt.Sprintf("文件错误:%s", fe.Message)
}

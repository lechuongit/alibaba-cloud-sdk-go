package vod

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// VideoInGetVideoList is a nested struct in vod response
type VideoInGetVideoList struct {
	VideoId          string                  `json:"VideoId" xml:"VideoId"`
	Title            string                  `json:"Title" xml:"Title"`
	Tags             string                  `json:"Tags" xml:"Tags"`
	Status           string                  `json:"Status" xml:"Status"`
	Size             int                     `json:"Size" xml:"Size"`
	Duration         float64                 `json:"Duration" xml:"Duration"`
	Description      string                  `json:"Description" xml:"Description"`
	CreateTime       string                  `json:"CreateTime" xml:"CreateTime"`
	ModifyTime       string                  `json:"ModifyTime" xml:"ModifyTime"`
	ModificationTime string                  `json:"ModificationTime" xml:"ModificationTime"`
	CreationTime     string                  `json:"CreationTime" xml:"CreationTime"`
	CoverURL         string                  `json:"CoverURL" xml:"CoverURL"`
	CateId           int                     `json:"CateId" xml:"CateId"`
	CateName         string                  `json:"CateName" xml:"CateName"`
	StorageLocation  string                  `json:"StorageLocation" xml:"StorageLocation"`
	AppId            string                  `json:"AppId" xml:"AppId"`
	Snapshots        SnapshotsInGetVideoList `json:"Snapshots" xml:"Snapshots"`
}

package main

type StringSha struct {
	ObjectType 			string  `json:"docType"`
	StringSha			string  `json:"stringSha"`
} 

type FileSha struct {
	ObjectType 			string  `json:"docType"`
	FileSha				string  `json:"fileSha"`
} 

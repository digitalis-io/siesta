/* Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License. */

package siesta

import "testing"

var emptyFetchRequestBytes = []byte{0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
var singleFetchRequestBytes = []byte{0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x00, 0x03, 0xE8, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x01, 0x00, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x07, 0x5B, 0xCD, 0x15, 0x00, 0x00, 0x04, 0x00}
var multipleFetchRequestBytes = []byte{0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x00, 0x07, 0xD0, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x01, 0x00, 0x05, 0x6c, 0x6f, 0x67, 0x73, 0x31, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x3A, 0xDE, 0x68, 0xB1, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x96, 0x46, 0x19, 0xC7, 0x00, 0x00, 0x10, 0x00}

var emptyFetchResponseBytes = []byte{0x00, 0x00, 0x00, 0x00}
var singleFetchResponseBytes = []byte{0x00, 0x00, 0x00, 0x01, 0x00, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xE8, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xE8, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0xAA, 0xAA, 0xAA, 0xAA, 0x00, 0x00, 0x00, 0x04, 0xBB, 0xBB, 0xBB, 0xBB}

var invalidTopicsLengthFetchResponseBytes = []byte{0x00, 0x00, 0x00}
var invalidTopicFetchResponseBytes = []byte{0x00, 0x00, 0x00, 0x01, 0x00, 0x04, 0x6c, 0x6f, 0x67}
var invalidPartitionsLengthFetchBytes = []byte{0x00, 0x00, 0x00, 0x01, 0x00, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x00, 0x00}
var invalidPartitionFetchBytes = []byte{0x00, 0x00, 0x00, 0x01, 0x00, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00}
var invalidErrorCodeFetchBytes = []byte{0x00, 0x00, 0x00, 0x01, 0x00, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00}
var invalidHightwaterMarkOffsetFetchResponseBytes = []byte{0x00, 0x00, 0x00, 0x01, 0x00, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
var invalidMessageSetSizeFetchResponseBytes = []byte{0x00, 0x00, 0x00, 0x01, 0x00, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xE8, 0x00, 0x00, 0x00}

var invalidOffsetMessageSetBytes = []byte{0x00, 0x00, 0x00, 0x01, 0x00, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xE8, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xE8, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0xAA, 0xAA, 0xAA, 0xAA, 0x00, 0x00, 0x00, 0x04, 0xBB, 0xBB, 0xBB, 0xBB, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
var invalidMessageSizeMessageSetBytes = []byte{0x00, 0x00, 0x00, 0x01, 0x00, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xE8, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xE8, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0xAA, 0xAA, 0xAA, 0xAA, 0x00, 0x00, 0x00, 0x04, 0xBB, 0xBB, 0xBB, 0xBB, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xE8, 0x00, 0x00, 0x00}

var invalidCRCMessageBytes = []byte{0x00, 0x00, 0x00, 0x01, 0x00, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xE8, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xE8, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0xAA, 0xAA, 0xAA, 0xAA, 0x00, 0x00, 0x00, 0x04, 0xBB, 0xBB, 0xBB, 0xBB, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xE8, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x04}
var invalidMagicByteMessageBytes = []byte{0x00, 0x00, 0x00, 0x01, 0x00, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xE8, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xE8, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0xAA, 0xAA, 0xAA, 0xAA, 0x00, 0x00, 0x00, 0x04, 0xBB, 0xBB, 0xBB, 0xBB, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xE8, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x04, 0x00}
var invalidAttributesMessageBytes = []byte{0x00, 0x00, 0x00, 0x01, 0x00, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xE8, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xE8, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0xAA, 0xAA, 0xAA, 0xAA, 0x00, 0x00, 0x00, 0x04, 0xBB, 0xBB, 0xBB, 0xBB, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xE8, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x04, 0x00, 0x00}
var invalidKeyMessageBytes = []byte{0x00, 0x00, 0x00, 0x01, 0x00, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xE8, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xE8, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0xAA, 0xAA, 0xAA, 0xAA, 0x00, 0x00, 0x00, 0x04, 0xBB, 0xBB, 0xBB, 0xBB, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xE8, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00}
var invalidValueMessageBytes = []byte{0x00, 0x00, 0x00, 0x01, 0x00, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xE8, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xE8, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0xAA, 0xAA, 0xAA, 0xAA, 0x00, 0x00, 0x00, 0x04, 0xBB, 0xBB, 0xBB, 0xBB, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xE8, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0xAA, 0xAA, 0xAA, 0xAA, 0x00, 0x00, 0x00, 0x04, 0xBB, 0xBB}

func TestFetchRequest(t *testing.T) {
	emptyFetchRequest := new(FetchRequest)
	testRequest(t, emptyFetchRequest, emptyFetchRequestBytes)

	singleFetchRequest := new(FetchRequest)
	singleFetchRequest.MaxWaitTime = int32(1000)
	singleFetchRequest.MinBytes = int32(4)
	singleFetchRequest.AddFetch("logs", 1, 123456789, 1024)
	testRequest(t, singleFetchRequest, singleFetchRequestBytes)

	multipleFetchRequest := new(FetchRequest)
	multipleFetchRequest.MaxWaitTime = int32(2000)
	multipleFetchRequest.MinBytes = int32(8)
	multipleFetchRequest.AddFetch("logs1", 2, 987654321, 2048)
	multipleFetchRequest.AddFetch("logs1", 0, 11111111111, 4096)
	testRequest(t, multipleFetchRequest, multipleFetchRequestBytes)
}

func TestFetchResponse(t *testing.T) {
	emptyFetchResponse := new(FetchResponse)
	decode(t, emptyFetchResponse, emptyFetchResponseBytes)

	singleFetchResponse := new(FetchResponse)
	decode(t, singleFetchResponse, singleFetchResponseBytes)
	testGoodFetchResponse(t, singleFetchResponse)

	decodeErr(t, new(FetchResponse), invalidTopicsLengthFetchResponseBytes, NewDecodingError(ErrEOF, reasonInvalidBlocksLength))
	decodeErr(t, new(FetchResponse), invalidTopicFetchResponseBytes, NewDecodingError(ErrEOF, reasonInvalidBlockTopic))
	decodeErr(t, new(FetchResponse), invalidPartitionsLengthFetchBytes, NewDecodingError(ErrEOF, reasonInvalidFetchResponseDataLength))
	decodeErr(t, new(FetchResponse), invalidPartitionFetchBytes, NewDecodingError(ErrEOF, reasonInvalidFetchResponseDataPartition))
	decodeErr(t, new(FetchResponse), invalidErrorCodeFetchBytes, NewDecodingError(ErrEOF, reasonInvalidFetchResponseDataErrorCode))
	decodeErr(t, new(FetchResponse), invalidHightwaterMarkOffsetFetchResponseBytes, NewDecodingError(ErrEOF, reasonInvalidFetchResponseDataHighwaterMarkOffset))
	decodeErr(t, new(FetchResponse), invalidMessageSetSizeFetchResponseBytes, NewDecodingError(ErrEOF, reasonInvalidMessageSetLength))

	// partially cut fetch responses should be good as well and get as much data as possible
	cutFetchResponse := new(FetchResponse)
	decode(t, cutFetchResponse, invalidOffsetMessageSetBytes)
	testGoodFetchResponse(t, cutFetchResponse)

	cutFetchResponse = new(FetchResponse)
	decode(t, cutFetchResponse, invalidMessageSizeMessageSetBytes)
	testGoodFetchResponse(t, cutFetchResponse)

	cutFetchResponse = new(FetchResponse)
	decode(t, cutFetchResponse, invalidCRCMessageBytes)
	testGoodFetchResponse(t, cutFetchResponse)

	cutFetchResponse = new(FetchResponse)
	decode(t, cutFetchResponse, invalidMagicByteMessageBytes)
	testGoodFetchResponse(t, cutFetchResponse)

	cutFetchResponse = new(FetchResponse)
	decode(t, cutFetchResponse, invalidAttributesMessageBytes)
	testGoodFetchResponse(t, cutFetchResponse)

	cutFetchResponse = new(FetchResponse)
	decode(t, cutFetchResponse, invalidKeyMessageBytes)
	testGoodFetchResponse(t, cutFetchResponse)

	cutFetchResponse = new(FetchResponse)
	decode(t, cutFetchResponse, invalidValueMessageBytes)
	testGoodFetchResponse(t, cutFetchResponse)
}

func testGoodFetchResponse(t *testing.T, response *FetchResponse) {
	partitionData, exists := response.Blocks["logs"]
	assertFatal(t, exists, true)
	data, exists := partitionData[1]
	assertFatal(t, exists, true)
	assert(t, data.Error, ErrNoError)
	assert(t, data.HighwaterMarkOffset, int64(1000))
	assertFatal(t, len(data.Messages), 1)
	messageAndOffset := data.Messages[0]
	assert(t, messageAndOffset.Offset, int64(1000))
	messageData := messageAndOffset.Message
	assert(t, messageData.Crc, int32(1024))
	assert(t, messageData.MagicByte, int8(0))
	assert(t, messageData.Attributes, int8(0))
	assert(t, messageData.Key, []byte{0xAA, 0xAA, 0xAA, 0xAA})
	assert(t, messageData.Value, []byte{0xBB, 0xBB, 0xBB, 0xBB})

	messages, err := response.GetMessages()
	assertFatal(t, err, nil)
	assertFatal(t, len(messages), 1)
	message := messages[0]
	assert(t, message.Topic, "logs")
	assert(t, message.Partition, int32(1))
	assert(t, message.Offset, int64(1000))
	assert(t, message.Key, []byte{0xAA, 0xAA, 0xAA, 0xAA})
	assert(t, message.Value, []byte{0xBB, 0xBB, 0xBB, 0xBB})
}
